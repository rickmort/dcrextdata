package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/raedahgroup/dcrextdata/datasync"
	"github.com/raedahgroup/dcrextdata/postgres/models"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func (pg *PgDb) SaveSyncHistory(ctx context.Context, history datasync.History) error {
	prevHistory, err := models.SyncHistories(
		models.SyncHistoryWhere.TableName.EQ(history.Table),
		models.SyncHistoryWhere.Source.EQ(history.Source)).One(ctx, pg.db)

	if err == nil {
		_, prevErr := prevHistory.Delete(ctx, pg.db)
		if prevErr != nil {
			log.Warnf("error in deleting previous %s sync history for %s, %s", history.Table, history.Source, prevErr.Error())
		}
	}
	historyModel := models.SyncHistory{
		Date: history.Date, Source: history.Source, TableName: history.Table,
	}
	if err := historyModel.Insert(ctx, pg.db, boil.Infer()); err != nil {
		return err
	}
	return nil
}

func (pg *PgDb) FetchSyncHistory(ctx context.Context, tableName string, source string) (datasync.History, error) {
	history, err := models.SyncHistories(
		models.SyncHistoryWhere.TableName.EQ(tableName),
		models.SyncHistoryWhere.Source.EQ(source),
		qm.OrderBy(fmt.Sprintf("%s desc", models.SyncHistoryColumns.Date))).One(ctx, pg.db)
	if err != nil {
		if err != sql.ErrNoRows {
			return datasync.History{}, err
		}

		history = &models.SyncHistory{
			ID:        0,
			TableName: tableName,
			Source:    source,
			Date:      time.Date(2019, 2, 1, 0, 0, 0, 0, nil),
		}
	}
	return datasync.History{Date: history.Date, Table: history.TableName, Source: history.Source}, nil
}

func (pg *PgDb) TableNames() []string {
	return []string{
		models.TableNames.Vote,
		models.TableNames.Block,
		models.TableNames.Mempool,
		models.TableNames.Exchange,
		models.TableNames.ExchangeTick,
		models.TableNames.VSP,
		models.TableNames.VSPTick,
		models.TableNames.PowData,
		models.TableNames.SyncHistory,
	}
}
