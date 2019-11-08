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

func testGithubs(t *testing.T) {
	t.Parallel()

	query := Githubs()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testGithubsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Github{}
	if err = randomize.Struct(seed, o, githubDBTypes, true, githubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Github struct: %s", err)
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

	count, err := Githubs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGithubsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Github{}
	if err = randomize.Struct(seed, o, githubDBTypes, true, githubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Github struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Githubs().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Githubs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGithubsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Github{}
	if err = randomize.Struct(seed, o, githubDBTypes, true, githubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Github struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := GithubSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Githubs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGithubsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Github{}
	if err = randomize.Struct(seed, o, githubDBTypes, true, githubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Github struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := GithubExists(ctx, tx, o.Date)
	if err != nil {
		t.Errorf("Unable to check if Github exists: %s", err)
	}
	if !e {
		t.Errorf("Expected GithubExists to return true, but got false.")
	}
}

func testGithubsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Github{}
	if err = randomize.Struct(seed, o, githubDBTypes, true, githubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Github struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	githubFound, err := FindGithub(ctx, tx, o.Date)
	if err != nil {
		t.Error(err)
	}

	if githubFound == nil {
		t.Error("want a record, got nil")
	}
}

func testGithubsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Github{}
	if err = randomize.Struct(seed, o, githubDBTypes, true, githubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Github struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Githubs().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testGithubsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Github{}
	if err = randomize.Struct(seed, o, githubDBTypes, true, githubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Github struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Githubs().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testGithubsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	githubOne := &Github{}
	githubTwo := &Github{}
	if err = randomize.Struct(seed, githubOne, githubDBTypes, false, githubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Github struct: %s", err)
	}
	if err = randomize.Struct(seed, githubTwo, githubDBTypes, false, githubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Github struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = githubOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = githubTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Githubs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testGithubsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	githubOne := &Github{}
	githubTwo := &Github{}
	if err = randomize.Struct(seed, githubOne, githubDBTypes, false, githubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Github struct: %s", err)
	}
	if err = randomize.Struct(seed, githubTwo, githubDBTypes, false, githubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Github struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = githubOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = githubTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Githubs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testGithubsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Github{}
	if err = randomize.Struct(seed, o, githubDBTypes, true, githubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Github struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Githubs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testGithubsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Github{}
	if err = randomize.Struct(seed, o, githubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Github struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(githubColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Githubs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testGithubsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Github{}
	if err = randomize.Struct(seed, o, githubDBTypes, true, githubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Github struct: %s", err)
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

func testGithubsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Github{}
	if err = randomize.Struct(seed, o, githubDBTypes, true, githubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Github struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := GithubSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testGithubsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Github{}
	if err = randomize.Struct(seed, o, githubDBTypes, true, githubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Github struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Githubs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	githubDBTypes = map[string]string{`Date`: `timestamp without time zone`, `Repository`: `character varying`, `Stars`: `integer`, `Folks`: `integer`}
	_             = bytes.MinRead
)

func testGithubsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(githubPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(githubAllColumns) == len(githubPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Github{}
	if err = randomize.Struct(seed, o, githubDBTypes, true, githubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Github struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Githubs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, githubDBTypes, true, githubPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Github struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testGithubsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(githubAllColumns) == len(githubPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Github{}
	if err = randomize.Struct(seed, o, githubDBTypes, true, githubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Github struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Githubs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, githubDBTypes, true, githubPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Github struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(githubAllColumns, githubPrimaryKeyColumns) {
		fields = githubAllColumns
	} else {
		fields = strmangle.SetComplement(
			githubAllColumns,
			githubPrimaryKeyColumns,
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

	slice := GithubSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testGithubsUpsert(t *testing.T) {
	t.Parallel()

	if len(githubAllColumns) == len(githubPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Github{}
	if err = randomize.Struct(seed, &o, githubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Github struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Github: %s", err)
	}

	count, err := Githubs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, githubDBTypes, false, githubPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Github struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Github: %s", err)
	}

	count, err = Githubs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
