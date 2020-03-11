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
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// NetworkSnapshot is an object representing the database table.
type NetworkSnapshot struct {
	Timestamp           int64  `boil:"timestamp" json:"timestamp" toml:"timestamp" yaml:"timestamp"`
	Height              int64  `boil:"height" json:"height" toml:"height" yaml:"height"`
	NodeCount           int    `boil:"node_count" json:"node_count" toml:"node_count" yaml:"node_count"`
	ReachableNodes      int    `boil:"reachable_nodes" json:"reachable_nodes" toml:"reachable_nodes" yaml:"reachable_nodes"`
	OldestNodeTimestamp int64  `boil:"oldest_node_timestamp" json:"oldest_node_timestamp" toml:"oldest_node_timestamp" yaml:"oldest_node_timestamp"`
	OldestNode          string `boil:"oldest_node" json:"oldest_node" toml:"oldest_node" yaml:"oldest_node"`
	Latency             int    `boil:"latency" json:"latency" toml:"latency" yaml:"latency"`

	R *networkSnapshotR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L networkSnapshotL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var NetworkSnapshotColumns = struct {
	Timestamp           string
	Height              string
	NodeCount           string
	ReachableNodes      string
	OldestNodeTimestamp string
	OldestNode          string
	Latency             string
}{
	Timestamp:           "timestamp",
	Height:              "height",
	NodeCount:           "node_count",
	ReachableNodes:      "reachable_nodes",
	OldestNodeTimestamp: "oldest_node_timestamp",
	OldestNode:          "oldest_node",
	Latency:             "latency",
}

// Generated where

var NetworkSnapshotWhere = struct {
	Timestamp           whereHelperint64
	Height              whereHelperint64
	NodeCount           whereHelperint
	ReachableNodes      whereHelperint
	OldestNodeTimestamp whereHelperint64
	OldestNode          whereHelperstring
	Latency             whereHelperint
}{
	Timestamp:           whereHelperint64{field: "\"network_snapshot\".\"timestamp\""},
	Height:              whereHelperint64{field: "\"network_snapshot\".\"height\""},
	NodeCount:           whereHelperint{field: "\"network_snapshot\".\"node_count\""},
	ReachableNodes:      whereHelperint{field: "\"network_snapshot\".\"reachable_nodes\""},
	OldestNodeTimestamp: whereHelperint64{field: "\"network_snapshot\".\"oldest_node_timestamp\""},
	OldestNode:          whereHelperstring{field: "\"network_snapshot\".\"oldest_node\""},
	Latency:             whereHelperint{field: "\"network_snapshot\".\"latency\""},
}

// NetworkSnapshotRels is where relationship names are stored.
var NetworkSnapshotRels = struct {
}{}

// networkSnapshotR is where relationships are stored.
type networkSnapshotR struct {
}

// NewStruct creates a new relationship struct
func (*networkSnapshotR) NewStruct() *networkSnapshotR {
	return &networkSnapshotR{}
}

// networkSnapshotL is where Load methods for each relationship are stored.
type networkSnapshotL struct{}

var (
	networkSnapshotAllColumns            = []string{"timestamp", "height", "node_count", "reachable_nodes", "oldest_node_timestamp", "oldest_node", "latency"}
	networkSnapshotColumnsWithoutDefault = []string{"timestamp", "height", "node_count", "reachable_nodes"}
	networkSnapshotColumnsWithDefault    = []string{"oldest_node_timestamp", "oldest_node", "latency"}
	networkSnapshotPrimaryKeyColumns     = []string{"timestamp"}
)

type (
	// NetworkSnapshotSlice is an alias for a slice of pointers to NetworkSnapshot.
	// This should generally be used opposed to []NetworkSnapshot.
	NetworkSnapshotSlice []*NetworkSnapshot

	networkSnapshotQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	networkSnapshotType                 = reflect.TypeOf(&NetworkSnapshot{})
	networkSnapshotMapping              = queries.MakeStructMapping(networkSnapshotType)
	networkSnapshotPrimaryKeyMapping, _ = queries.BindMapping(networkSnapshotType, networkSnapshotMapping, networkSnapshotPrimaryKeyColumns)
	networkSnapshotInsertCacheMut       sync.RWMutex
	networkSnapshotInsertCache          = make(map[string]insertCache)
	networkSnapshotUpdateCacheMut       sync.RWMutex
	networkSnapshotUpdateCache          = make(map[string]updateCache)
	networkSnapshotUpsertCacheMut       sync.RWMutex
	networkSnapshotUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single networkSnapshot record from the query.
func (q networkSnapshotQuery) One(ctx context.Context, exec boil.ContextExecutor) (*NetworkSnapshot, error) {
	o := &NetworkSnapshot{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for network_snapshot")
	}

	return o, nil
}

// All returns all NetworkSnapshot records from the query.
func (q networkSnapshotQuery) All(ctx context.Context, exec boil.ContextExecutor) (NetworkSnapshotSlice, error) {
	var o []*NetworkSnapshot

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to NetworkSnapshot slice")
	}

	return o, nil
}

// Count returns the count of all NetworkSnapshot records in the query.
func (q networkSnapshotQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count network_snapshot rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q networkSnapshotQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if network_snapshot exists")
	}

	return count > 0, nil
}

// NetworkSnapshots retrieves all the records using an executor.
func NetworkSnapshots(mods ...qm.QueryMod) networkSnapshotQuery {
	mods = append(mods, qm.From("\"network_snapshot\""))
	return networkSnapshotQuery{NewQuery(mods...)}
}

// FindNetworkSnapshot retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindNetworkSnapshot(ctx context.Context, exec boil.ContextExecutor, timestamp int64, selectCols ...string) (*NetworkSnapshot, error) {
	networkSnapshotObj := &NetworkSnapshot{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"network_snapshot\" where \"timestamp\"=$1", sel,
	)

	q := queries.Raw(query, timestamp)

	err := q.Bind(ctx, exec, networkSnapshotObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from network_snapshot")
	}

	return networkSnapshotObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *NetworkSnapshot) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no network_snapshot provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(networkSnapshotColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	networkSnapshotInsertCacheMut.RLock()
	cache, cached := networkSnapshotInsertCache[key]
	networkSnapshotInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			networkSnapshotAllColumns,
			networkSnapshotColumnsWithDefault,
			networkSnapshotColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(networkSnapshotType, networkSnapshotMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(networkSnapshotType, networkSnapshotMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"network_snapshot\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"network_snapshot\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into network_snapshot")
	}

	if !cached {
		networkSnapshotInsertCacheMut.Lock()
		networkSnapshotInsertCache[key] = cache
		networkSnapshotInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the NetworkSnapshot.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *NetworkSnapshot) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	networkSnapshotUpdateCacheMut.RLock()
	cache, cached := networkSnapshotUpdateCache[key]
	networkSnapshotUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			networkSnapshotAllColumns,
			networkSnapshotPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("models: unable to update network_snapshot, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"network_snapshot\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, networkSnapshotPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(networkSnapshotType, networkSnapshotMapping, append(wl, networkSnapshotPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update network_snapshot row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for network_snapshot")
	}

	if !cached {
		networkSnapshotUpdateCacheMut.Lock()
		networkSnapshotUpdateCache[key] = cache
		networkSnapshotUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q networkSnapshotQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for network_snapshot")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for network_snapshot")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o NetworkSnapshotSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), networkSnapshotPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"network_snapshot\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, networkSnapshotPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in networkSnapshot slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all networkSnapshot")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *NetworkSnapshot) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no network_snapshot provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(networkSnapshotColumnsWithDefault, o)

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

	networkSnapshotUpsertCacheMut.RLock()
	cache, cached := networkSnapshotUpsertCache[key]
	networkSnapshotUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			networkSnapshotAllColumns,
			networkSnapshotColumnsWithDefault,
			networkSnapshotColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			networkSnapshotAllColumns,
			networkSnapshotPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert network_snapshot, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(networkSnapshotPrimaryKeyColumns))
			copy(conflict, networkSnapshotPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"network_snapshot\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(networkSnapshotType, networkSnapshotMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(networkSnapshotType, networkSnapshotMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert network_snapshot")
	}

	if !cached {
		networkSnapshotUpsertCacheMut.Lock()
		networkSnapshotUpsertCache[key] = cache
		networkSnapshotUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single NetworkSnapshot record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *NetworkSnapshot) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no NetworkSnapshot provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), networkSnapshotPrimaryKeyMapping)
	sql := "DELETE FROM \"network_snapshot\" WHERE \"timestamp\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from network_snapshot")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for network_snapshot")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q networkSnapshotQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no networkSnapshotQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from network_snapshot")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for network_snapshot")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o NetworkSnapshotSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), networkSnapshotPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"network_snapshot\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, networkSnapshotPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from networkSnapshot slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for network_snapshot")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *NetworkSnapshot) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindNetworkSnapshot(ctx, exec, o.Timestamp)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *NetworkSnapshotSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := NetworkSnapshotSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), networkSnapshotPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"network_snapshot\".* FROM \"network_snapshot\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, networkSnapshotPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in NetworkSnapshotSlice")
	}

	*o = slice

	return nil
}

// NetworkSnapshotExists checks if the NetworkSnapshot row exists.
func NetworkSnapshotExists(ctx context.Context, exec boil.ContextExecutor, timestamp int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"network_snapshot\" where \"timestamp\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, timestamp)
	}
	row := exec.QueryRowContext(ctx, sql, timestamp)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if network_snapshot exists")
	}

	return exists, nil
}
