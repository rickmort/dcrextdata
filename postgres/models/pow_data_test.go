// Code generated by SQLBoiler 3.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testPowData(t *testing.T) {
	t.Parallel()

	query := PowData()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testPowDataDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PowDatum{}
	if err = randomize.Struct(seed, o, powDatumDBTypes, true, powDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PowDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := PowData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPowDataQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PowDatum{}
	if err = randomize.Struct(seed, o, powDatumDBTypes, true, powDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PowDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := PowData().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := PowData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPowDataSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PowDatum{}
	if err = randomize.Struct(seed, o, powDatumDBTypes, true, powDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PowDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PowDatumSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := PowData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPowDataExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PowDatum{}
	if err = randomize.Struct(seed, o, powDatumDBTypes, true, powDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PowDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := PowDatumExists(ctx, tx, o.Time, o.Source)
	if err != nil {
		t.Errorf("Unable to check if PowDatum exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PowDatumExists to return true, but got false.")
	}
}

func testPowDataFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PowDatum{}
	if err = randomize.Struct(seed, o, powDatumDBTypes, true, powDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PowDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	powDatumFound, err := FindPowDatum(ctx, tx, o.Time, o.Source)
	if err != nil {
		t.Error(err)
	}

	if powDatumFound == nil {
		t.Error("want a record, got nil")
	}
}

func testPowDataBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PowDatum{}
	if err = randomize.Struct(seed, o, powDatumDBTypes, true, powDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PowDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = PowData().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testPowDataOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PowDatum{}
	if err = randomize.Struct(seed, o, powDatumDBTypes, true, powDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PowDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := PowData().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPowDataAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	powDatumOne := &PowDatum{}
	powDatumTwo := &PowDatum{}
	if err = randomize.Struct(seed, powDatumOne, powDatumDBTypes, false, powDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PowDatum struct: %s", err)
	}
	if err = randomize.Struct(seed, powDatumTwo, powDatumDBTypes, false, powDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PowDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = powDatumOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = powDatumTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := PowData().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPowDataCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	powDatumOne := &PowDatum{}
	powDatumTwo := &PowDatum{}
	if err = randomize.Struct(seed, powDatumOne, powDatumDBTypes, false, powDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PowDatum struct: %s", err)
	}
	if err = randomize.Struct(seed, powDatumTwo, powDatumDBTypes, false, powDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PowDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = powDatumOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = powDatumTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PowData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testPowDataInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PowDatum{}
	if err = randomize.Struct(seed, o, powDatumDBTypes, true, powDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PowDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PowData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPowDataInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PowDatum{}
	if err = randomize.Struct(seed, o, powDatumDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PowDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(powDatumColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := PowData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPowDataReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PowDatum{}
	if err = randomize.Struct(seed, o, powDatumDBTypes, true, powDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PowDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPowDataReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PowDatum{}
	if err = randomize.Struct(seed, o, powDatumDBTypes, true, powDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PowDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PowDatumSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPowDataSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PowDatum{}
	if err = randomize.Struct(seed, o, powDatumDBTypes, true, powDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PowDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := PowData().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	powDatumDBTypes = map[string]string{`PropagationTime`: `integer`, `PoolHashrate`: `character varying`, `Workers`: `integer`, `CoinPrice`: `character varying`, `BTCPrice`: `character varying`, `Source`: `character varying`}
	_               = bytes.MinRead
)

func testPowDataUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(powDatumPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(powDatumAllColumns) == len(powDatumPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &PowDatum{}
	if err = randomize.Struct(seed, o, powDatumDBTypes, true, powDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PowDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PowData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, powDatumDBTypes, true, powDatumPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PowDatum struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testPowDataSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(powDatumAllColumns) == len(powDatumPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &PowDatum{}
	if err = randomize.Struct(seed, o, powDatumDBTypes, true, powDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PowDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PowData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, powDatumDBTypes, true, powDatumPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PowDatum struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(powDatumAllColumns, powDatumPrimaryKeyColumns) {
		fields = powDatumAllColumns
	} else {
		fields = strmangle.SetComplement(
			powDatumAllColumns,
			powDatumPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := PowDatumSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testPowDataUpsert(t *testing.T) {
	t.Parallel()

	if len(powDatumAllColumns) == len(powDatumPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := PowDatum{}
	if err = randomize.Struct(seed, &o, powDatumDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PowDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert PowDatum: %s", err)
	}

	count, err := PowData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, powDatumDBTypes, false, powDatumPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PowDatum struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert PowDatum: %s", err)
	}

	count, err = PowData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
