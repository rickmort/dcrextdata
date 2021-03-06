// Code generated by SQLBoiler 3.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Mempool is an object representing the database table.
type Mempool struct {
	Time                 time.Time    `boil:"time" json:"time" toml:"time" yaml:"time"`
	FirstSeenTime        null.Time    `boil:"first_seen_time" json:"first_seen_time,omitempty" toml:"first_seen_time" yaml:"first_seen_time,omitempty"`
	NumberOfTransactions null.Int     `boil:"number_of_transactions" json:"number_of_transactions,omitempty" toml:"number_of_transactions" yaml:"number_of_transactions,omitempty"`
	Voters               null.Int     `boil:"voters" json:"voters,omitempty" toml:"voters" yaml:"voters,omitempty"`
	Tickets              null.Int     `boil:"tickets" json:"tickets,omitempty" toml:"tickets" yaml:"tickets,omitempty"`
	Revocations          null.Int     `boil:"revocations" json:"revocations,omitempty" toml:"revocations" yaml:"revocations,omitempty"`
	Size                 null.Int     `boil:"size" json:"size,omitempty" toml:"size" yaml:"size,omitempty"`
	TotalFee             null.Float64 `boil:"total_fee" json:"total_fee,omitempty" toml:"total_fee" yaml:"total_fee,omitempty"`
	Total                null.Float64 `boil:"total" json:"total,omitempty" toml:"total" yaml:"total,omitempty"`

	R *mempoolR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L mempoolL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var MempoolColumns = struct {
	Time                 string
	FirstSeenTime        string
	NumberOfTransactions string
	Voters               string
	Tickets              string
	Revocations          string
	Size                 string
	TotalFee             string
	Total                string
}{
	Time:                 "time",
	FirstSeenTime:        "first_seen_time",
	NumberOfTransactions: "number_of_transactions",
	Voters:               "voters",
	Tickets:              "tickets",
	Revocations:          "revocations",
	Size:                 "size",
	TotalFee:             "total_fee",
	Total:                "total",
}

// Generated where

type whereHelpernull_Int struct{ field string }

func (w whereHelpernull_Int) EQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int) NEQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Int) LT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int) LTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int) GT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int) GTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpernull_Float64 struct{ field string }

func (w whereHelpernull_Float64) EQ(x null.Float64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Float64) NEQ(x null.Float64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Float64) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Float64) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Float64) LT(x null.Float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Float64) LTE(x null.Float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Float64) GT(x null.Float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Float64) GTE(x null.Float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var MempoolWhere = struct {
	Time                 whereHelpertime_Time
	FirstSeenTime        whereHelpernull_Time
	NumberOfTransactions whereHelpernull_Int
	Voters               whereHelpernull_Int
	Tickets              whereHelpernull_Int
	Revocations          whereHelpernull_Int
	Size                 whereHelpernull_Int
	TotalFee             whereHelpernull_Float64
	Total                whereHelpernull_Float64
}{
	Time:                 whereHelpertime_Time{field: "\"mempool\".\"time\""},
	FirstSeenTime:        whereHelpernull_Time{field: "\"mempool\".\"first_seen_time\""},
	NumberOfTransactions: whereHelpernull_Int{field: "\"mempool\".\"number_of_transactions\""},
	Voters:               whereHelpernull_Int{field: "\"mempool\".\"voters\""},
	Tickets:              whereHelpernull_Int{field: "\"mempool\".\"tickets\""},
	Revocations:          whereHelpernull_Int{field: "\"mempool\".\"revocations\""},
	Size:                 whereHelpernull_Int{field: "\"mempool\".\"size\""},
	TotalFee:             whereHelpernull_Float64{field: "\"mempool\".\"total_fee\""},
	Total:                whereHelpernull_Float64{field: "\"mempool\".\"total\""},
}

// MempoolRels is where relationship names are stored.
var MempoolRels = struct {
}{}

// mempoolR is where relationships are stored.
type mempoolR struct {
}

// NewStruct creates a new relationship struct
func (*mempoolR) NewStruct() *mempoolR {
	return &mempoolR{}
}

// mempoolL is where Load methods for each relationship are stored.
type mempoolL struct{}

var (
	mempoolAllColumns            = []string{"time", "first_seen_time", "number_of_transactions", "voters", "tickets", "revocations", "size", "total_fee", "total"}
	mempoolColumnsWithoutDefault = []string{"time", "first_seen_time", "number_of_transactions", "voters", "tickets", "revocations", "size", "total_fee", "total"}
	mempoolColumnsWithDefault    = []string{}
	mempoolPrimaryKeyColumns     = []string{"time"}
)

type (
	// MempoolSlice is an alias for a slice of pointers to Mempool.
	// This should generally be used opposed to []Mempool.
	MempoolSlice []*Mempool

	mempoolQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	mempoolType                 = reflect.TypeOf(&Mempool{})
	mempoolMapping              = queries.MakeStructMapping(mempoolType)
	mempoolPrimaryKeyMapping, _ = queries.BindMapping(mempoolType, mempoolMapping, mempoolPrimaryKeyColumns)
	mempoolInsertCacheMut       sync.RWMutex
	mempoolInsertCache          = make(map[string]insertCache)
	mempoolUpdateCacheMut       sync.RWMutex
	mempoolUpdateCache          = make(map[string]updateCache)
	mempoolUpsertCacheMut       sync.RWMutex
	mempoolUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single mempool record from the query.
func (q mempoolQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Mempool, error) {
	o := &Mempool{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for mempool")
	}

	return o, nil
}

// All returns all Mempool records from the query.
func (q mempoolQuery) All(ctx context.Context, exec boil.ContextExecutor) (MempoolSlice, error) {
	var o []*Mempool

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Mempool slice")
	}

	return o, nil
}

// Count returns the count of all Mempool records in the query.
func (q mempoolQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count mempool rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q mempoolQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if mempool exists")
	}

	return count > 0, nil
}

// Mempools retrieves all the records using an executor.
func Mempools(mods ...qm.QueryMod) mempoolQuery {
	mods = append(mods, qm.From("\"mempool\""))
	return mempoolQuery{NewQuery(mods...)}
}

// FindMempool retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMempool(ctx context.Context, exec boil.ContextExecutor, time time.Time, selectCols ...string) (*Mempool, error) {
	mempoolObj := &Mempool{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"mempool\" where \"time\"=$1", sel,
	)

	q := queries.Raw(query, time)

	err := q.Bind(ctx, exec, mempoolObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from mempool")
	}

	return mempoolObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Mempool) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no mempool provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(mempoolColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	mempoolInsertCacheMut.RLock()
	cache, cached := mempoolInsertCache[key]
	mempoolInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			mempoolAllColumns,
			mempoolColumnsWithDefault,
			mempoolColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(mempoolType, mempoolMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(mempoolType, mempoolMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"mempool\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"mempool\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into mempool")
	}

	if !cached {
		mempoolInsertCacheMut.Lock()
		mempoolInsertCache[key] = cache
		mempoolInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Mempool.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Mempool) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	mempoolUpdateCacheMut.RLock()
	cache, cached := mempoolUpdateCache[key]
	mempoolUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			mempoolAllColumns,
			mempoolPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("models: unable to update mempool, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"mempool\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, mempoolPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(mempoolType, mempoolMapping, append(wl, mempoolPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update mempool row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for mempool")
	}

	if !cached {
		mempoolUpdateCacheMut.Lock()
		mempoolUpdateCache[key] = cache
		mempoolUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q mempoolQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for mempool")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for mempool")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o MempoolSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), mempoolPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"mempool\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, mempoolPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in mempool slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all mempool")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Mempool) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no mempool provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(mempoolColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	mempoolUpsertCacheMut.RLock()
	cache, cached := mempoolUpsertCache[key]
	mempoolUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			mempoolAllColumns,
			mempoolColumnsWithDefault,
			mempoolColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			mempoolAllColumns,
			mempoolPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert mempool, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(mempoolPrimaryKeyColumns))
			copy(conflict, mempoolPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"mempool\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(mempoolType, mempoolMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(mempoolType, mempoolMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert mempool")
	}

	if !cached {
		mempoolUpsertCacheMut.Lock()
		mempoolUpsertCache[key] = cache
		mempoolUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Mempool record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Mempool) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Mempool provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), mempoolPrimaryKeyMapping)
	sql := "DELETE FROM \"mempool\" WHERE \"time\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from mempool")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for mempool")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q mempoolQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no mempoolQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from mempool")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for mempool")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o MempoolSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), mempoolPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"mempool\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, mempoolPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from mempool slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for mempool")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Mempool) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindMempool(ctx, exec, o.Time)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MempoolSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := MempoolSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), mempoolPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"mempool\".* FROM \"mempool\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, mempoolPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in MempoolSlice")
	}

	*o = slice

	return nil
}

// MempoolExists checks if the Mempool row exists.
func MempoolExists(ctx context.Context, exec boil.ContextExecutor, time time.Time) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"mempool\" where \"time\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, time)
	}
	row := exec.QueryRowContext(ctx, sql, time)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if mempool exists")
	}

	return exists, nil
}
