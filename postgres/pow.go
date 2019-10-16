package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/raedahgroup/dcrextdata/cache"
	"strconv"
	"strings"
	"time"

	"github.com/raedahgroup/dcrextdata/postgres/models"
	"github.com/raedahgroup/dcrextdata/pow"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func (pg *PgDb) PowTableName() string {
	return models.TableNames.PowData
}

func (pg *PgDb) LastPowEntryTime(source string) (time int64) {
	var rows *sql.Row

	if source == "" {
		rows = pg.db.QueryRow(lastPowEntryTime)
	} else {
		rows = pg.db.QueryRow(lastPowEntryTimeBySource, source)
	}

	err := rows.Scan(&time)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Errorf("Error in getting last PoW entry time: %s", err.Error())
		}
	}
	return
}

//
func (pg *PgDb) AddPowData(ctx context.Context, data []pow.PowData) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	added := 0
	for _, d := range data {
		powModel, err := responseToPowModel(d)
		if err != nil {
			return err
		}

		err = powModel.Insert(ctx, pg.db, boil.Infer())
		if err != nil {
			if !strings.Contains(err.Error(), "unique constraint") { // Ignore duplicate entries
				return err
			}
		}
		added++
	}
	if len(data) == 1 {
		log.Infof("Added %4d PoW   entry from %10s %s", added, data[0].Source, UnixTimeToString(data[0].Time))
	} else if len(data) > 1 {
		last := data[len(data)-1]
		log.Infof("Added %4d PoW entries from %10s %s to %s",
			added, last.Source, UnixTimeToString(data[0].Time), UnixTimeToString(last.Time))
	}

	return nil
}

func (pg *PgDb) AddPowDataFromSync(ctx context.Context, data interface{}) error {
	powModel, err := responseToPowModel(data.(pow.PowData))
	if err != nil {
		return err
	}

	err = powModel.Insert(ctx, pg.db, boil.Infer())
	if isUniqueConstraint(err) {
		return nil
	}

	return err
}

func responseToPowModel(data pow.PowData) (models.PowDatum, error) {
	return models.PowDatum{
		BTCPrice:     null.StringFrom(fmt.Sprint(data.BtcPrice)),
		CoinPrice:    null.StringFrom(fmt.Sprint(data.CoinPrice)),
		PoolHashrate: null.StringFrom(fmt.Sprintf("%.0f", data.PoolHashrate/pow.Thash)),
		Source:       data.Source,
		Time:         int(data.Time),
		Workers:      null.IntFrom(int(data.Workers)),
	}, nil
}

func (pg *PgDb) PowCount(ctx context.Context) (int64, error) {
	return models.PowData().Count(ctx, pg.db)
}

func (pg *PgDb) FetchPowData(ctx context.Context, offset, limit int) ([]pow.PowDataDto, int64, error) {
	powDatum, err := models.PowData(qm.Offset(offset), qm.Limit(limit), qm.OrderBy(fmt.Sprintf("%s DESC", models.PowDatumColumns.Time))).All(ctx, pg.db)
	if err != nil {
		return nil, 0, err
	}

	powCount, err := models.PowData().Count(ctx, pg.db)
	if err != nil {
		return nil, 0, err
	}

	var result []pow.PowDataDto
	for _, item := range powDatum {
		dto, err := pg.powDataModelToDto(item)
		if err != nil {
			return nil, 0, err
		}

		result = append(result, dto)
	}

	return result, powCount, nil
}

// FetchPowDataForSync returns PoW data for the sync operation
func (pg *PgDb) FetchPowDataForSync(ctx context.Context, date int64, skip, take int) ([]pow.PowData, int64, error) {
	powDatum, err := models.PowData(
		models.PowDatumWhere.Time.GT(int(date)),
		qm.Offset(skip), qm.Limit(take),
		qm.OrderBy(models.PowDatumColumns.Time)).All(ctx, pg.db)
	if err != nil {
		return nil, 0, err
	}

	powCount, err := models.PowData(models.PowDatumWhere.Time.GT(int(date))).Count(ctx, pg.db)
	if err != nil {
		return nil, 0, err
	}

	var result []pow.PowData
	for _, item := range powDatum {
		dto, err := pg.powDataModelToDomainObj(item)
		if err != nil {
			return nil, 0, err
		}

		result = append(result, dto)
	}

	return result, powCount, nil
}

func (pg *PgDb) FetchPowDataBySource(ctx context.Context, source string, offset, limit int) ([]pow.PowDataDto, int64, error) {
	powDatum, err := models.PowData(models.PowDatumWhere.Source.EQ(source), qm.Offset(offset), qm.Limit(limit), qm.OrderBy(fmt.Sprintf("%s DESC", models.PowDatumColumns.Time))).All(ctx, pg.db)
	if err != nil {
		return nil, 0, err
	}

	powCount, err := models.PowData(models.PowDatumWhere.Source.EQ(source)).Count(ctx, pg.db)
	if err != nil {
		return nil, 0, err
	}

	var result []pow.PowDataDto
	for _, item := range powDatum {
		dto, err := pg.powDataModelToDto(item)
		if err != nil {
			return nil, 0, err
		}
		result = append(result, dto)
	}

	return result, powCount, nil
}

func (pg *PgDb) GetPowDistinctDates(ctx context.Context, sources []string) ([]time.Time, error) {
	query := fmt.Sprintf("SELECT DISTINCT %s FROM %s WHERE %s IN ('%s') ORDER BY %s", models.PowDatumColumns.Time,
		models.TableNames.PowData,
		models.PowDatumColumns.Source, strings.Join(sources, "', '"), models.PowDatumColumns.Time)

	rows, err := pg.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var dates []time.Time

	for rows.Next() {
		var date int64
		err = rows.Scan(&date)
		if err != nil {
			return nil, err
		}
		dates = append(dates, time.Unix(date, 0).UTC())
	}
	return dates, nil
}

func (pg *PgDb) powDistinctDates(ctx context.Context, sources []string, startDate int64) ([]int64, error) {
	query := fmt.Sprintf("SELECT DISTINCT %s FROM %s WHERE %s IN ('%s') and %s > %d ORDER BY %s", models.PowDatumColumns.Time,
		models.TableNames.PowData,
		models.PowDatumColumns.Source, strings.Join(sources, "', '"),
		models.PowDatumColumns.Time, startDate, models.PowDatumColumns.Time)

	rows, err := pg.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var dates []int64

	for rows.Next() {
		var date int64
		err = rows.Scan(&date)
		if err != nil {
			return nil, err
		}
		dates = append(dates, date)
	}
	return dates, nil
}

func (pg *PgDb) FetchPowChartData(ctx context.Context, source string, dataType string) (records []pow.PowChartData, err error) {
	dataType = strings.ToLower(dataType)
	query := fmt.Sprintf("SELECT %s as date, %s as record FROM %s where %s = '%s' ORDER BY %s",
		models.PowDatumColumns.Time, dataType, models.TableNames.PowData, models.PowDatumColumns.Source, source, models.PowDatumColumns.Time)
	rows, err := pg.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var rec pow.PowChartData
		var unixDate int64
		err = rows.Scan(&unixDate, &rec.Record)
		if err != nil {
			return nil, err
		}

		rec.Date = time.Unix(unixDate, 0).UTC()
		records = append(records, rec)
	}

	return
}

func (pg *PgDb) FetchPowChartDatav(ctx context.Context, source string, dataType string) ([]pow.PowChartData, error) {
	powDatum, err := models.PowData(qm.Select(models.PowDatumColumns.Time, dataType),
		models.PowDatumWhere.Source.EQ(source), qm.OrderBy(models.PowDatumColumns.Time)).All(ctx, pg.db)
	if err != nil {
		return nil, err
	}

	var result []pow.PowChartData
	for _, item := range powDatum {
		var record string
		if dataType == models.PowDatumColumns.Workers {
			record = strconv.FormatInt(int64(item.Workers.Int), 10)
		} else if dataType == models.PowDatumColumns.PoolHashrate {
			record = item.PoolHashrate.String
		} else {
			return nil, fmt.Errorf("unsupported data type: %s", dataType)
		}
		powChartData := pow.PowChartData{
			Date:   time.Unix(int64(item.Time), 0).UTC(),
			Record: record,
		}
		result = append(result, powChartData)
	}

	return result, nil
}

func (pg *PgDb) powDataModelToDto(item *models.PowDatum) (dto pow.PowDataDto, err error) {
	poolHashRate, err := strconv.ParseFloat(item.PoolHashrate.String, 64)
	if err != nil {
		return dto, err
	}

	coinPrice, err := strconv.ParseFloat(item.CoinPrice.String, 64)
	if err != nil {
		return dto, err
	}

	bTCPrice, err := strconv.ParseFloat(item.BTCPrice.String, 64)
	if err != nil {
		return dto, err
	}

	return pow.PowDataDto{
		Time:           time.Unix(int64(item.Time), 0).UTC().Format(dateTemplate),
		PoolHashrateTh: fmt.Sprintf("%.0f", poolHashRate),
		Workers:        int64(item.Workers.Int),
		Source:         item.Source,
		CoinPrice:      coinPrice,
		BtcPrice:       bTCPrice,
	}, nil
}

func (pg *PgDb) powDataModelToDomainObj(item *models.PowDatum) (dto pow.PowData, err error) {
	poolHashRate, err := strconv.ParseFloat(item.PoolHashrate.String, 64)
	if err != nil {
		return dto, err
	}

	coinPrice, err := strconv.ParseFloat(item.CoinPrice.String, 64)
	if err != nil {
		return dto, err
	}

	bTCPrice, err := strconv.ParseFloat(item.BTCPrice.String, 64)
	if err != nil {
		return dto, err
	}

	return pow.PowData{
		Time:         int64(item.Time),
		PoolHashrate: poolHashRate,
		Workers:      int64(item.Workers.Int),
		Source:       item.Source,
		CoinPrice:    coinPrice,
		BtcPrice:     bTCPrice,
	}, nil
}

func (pg *PgDb) FetchPowSourceData(ctx context.Context) ([]pow.PowDataSource, error) {
	powDatum, err := models.PowData(qm.Select("source"), qm.GroupBy("source")).All(ctx, pg.db)
	if err != nil {
		return nil, err
	}

	var result []pow.PowDataSource
	for _, item := range powDatum {
		result = append(result, pow.PowDataSource{
			Source: item.Source,
		})
	}

	return result, nil
}

type powSet struct {
	time     map[string][]uint64
	workers  map[string][]uint64
	hashrate map[string][]uint64
}

func (pg *PgDb) fetchPowChart(ctx context.Context, charts *cache.ChartData) (interface{}, func(), error) {
	cancelFun := func() {}
	var powDataSet = powSet{
		time:     map[string][]uint64{},
		workers:  make(map[string][]uint64),
		hashrate: make(map[string][]uint64),
	}

	pools, err := pg.FetchPowSourceData(ctx)
	if err != nil {
		return nil, cancelFun, err
	}

	var poolSources []string
	for _, pool := range pools {
		poolSources = append(poolSources, pool.Source)
	}

	for _, pool := range poolSources {
		points, err := models.PowData(models.PowDatumWhere.Time.GT(int(charts.PowTime(pool)))).All(ctx, pg.db)
		if err != nil {
			return nil, cancelFun, fmt.Errorf("error in fetching records for %s: %s", pool, err.Error())
		}

		for _, record := range points {
			powDataSet.time[pool] = append(powDataSet.time[pool], uint64(record.Time))
			powDataSet.workers[pool] = append(powDataSet.workers[pool], uint64(record.Workers.Int))
			hashrate, _ := strconv.ParseInt(record.PoolHashrate.String, 10, 64)
			powDataSet.hashrate[pool] = append(powDataSet.hashrate[pool], uint64(hashrate))
		}
	}

	return powDataSet, cancelFun, nil
}

func appendPowChart(charts *cache.ChartData, data interface{}) error {
	powDataSet := data.(powSet)

	for pool, timeSet := range powDataSet.time {
		if charts.Pow.Time == nil {
			charts.Pow.Time = map[string]cache.ChartUints{}
		}
		charts.Pow.Time[pool] = append(charts.Pow.Time[pool], timeSet...)
	}

	for pool, workers := range powDataSet.workers {
		if charts.Pow.Workers == nil {
			charts.Pow.Workers = map[string]cache.ChartUints{}
		}
		charts.Pow.Workers[pool] = append(charts.Pow.Workers[pool], workers...)
	}

	for pool, hashrate := range powDataSet.hashrate {
		if charts.Pow.Hashrate == nil {
			charts.Pow.Hashrate = map[string]cache.ChartUints{}
		}
		charts.Pow.Hashrate[pool] = append(charts.Pow.Hashrate[pool], hashrate...)
	}

	return nil
}
