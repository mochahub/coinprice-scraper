// Code generated by SQLBoiler 4.4.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testLastSyncs(t *testing.T) {
	t.Parallel()

	query := LastSyncs()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testLastSyncsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LastSync{}
	if err = randomize.Struct(seed, o, lastSyncDBTypes, true, lastSyncColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LastSyncTime struct: %s", err)
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

	count, err := LastSyncs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLastSyncsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LastSync{}
	if err = randomize.Struct(seed, o, lastSyncDBTypes, true, lastSyncColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LastSyncTime struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := LastSyncs().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := LastSyncs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLastSyncsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LastSync{}
	if err = randomize.Struct(seed, o, lastSyncDBTypes, true, lastSyncColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LastSyncTime struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := LastSyncSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := LastSyncs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLastSyncsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LastSync{}
	if err = randomize.Struct(seed, o, lastSyncDBTypes, true, lastSyncColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LastSyncTime struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := LastSyncExists(ctx, tx, o.BaseAsset, o.QuoteAsset, o.Exchange)
	if err != nil {
		t.Errorf("Unable to check if LastSyncTime exists: %s", err)
	}
	if !e {
		t.Errorf("Expected LastSyncExists to return true, but got false.")
	}
}

func testLastSyncsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LastSync{}
	if err = randomize.Struct(seed, o, lastSyncDBTypes, true, lastSyncColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LastSyncTime struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	lastSyncFound, err := FindLastSync(ctx, tx, o.BaseAsset, o.QuoteAsset, o.Exchange)
	if err != nil {
		t.Error(err)
	}

	if lastSyncFound == nil {
		t.Error("want a record, got nil")
	}
}

func testLastSyncsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LastSync{}
	if err = randomize.Struct(seed, o, lastSyncDBTypes, true, lastSyncColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LastSyncTime struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = LastSyncs().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testLastSyncsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LastSync{}
	if err = randomize.Struct(seed, o, lastSyncDBTypes, true, lastSyncColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LastSyncTime struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := LastSyncs().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testLastSyncsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	lastSyncOne := &LastSync{}
	lastSyncTwo := &LastSync{}
	if err = randomize.Struct(seed, lastSyncOne, lastSyncDBTypes, false, lastSyncColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LastSyncTime struct: %s", err)
	}
	if err = randomize.Struct(seed, lastSyncTwo, lastSyncDBTypes, false, lastSyncColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LastSyncTime struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = lastSyncOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = lastSyncTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := LastSyncs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testLastSyncsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	lastSyncOne := &LastSync{}
	lastSyncTwo := &LastSync{}
	if err = randomize.Struct(seed, lastSyncOne, lastSyncDBTypes, false, lastSyncColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LastSyncTime struct: %s", err)
	}
	if err = randomize.Struct(seed, lastSyncTwo, lastSyncDBTypes, false, lastSyncColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LastSyncTime struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = lastSyncOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = lastSyncTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := LastSyncs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func lastSyncBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *LastSync) error {
	*o = LastSync{}
	return nil
}

func lastSyncAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *LastSync) error {
	*o = LastSync{}
	return nil
}

func lastSyncAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *LastSync) error {
	*o = LastSync{}
	return nil
}

func lastSyncBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *LastSync) error {
	*o = LastSync{}
	return nil
}

func lastSyncAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *LastSync) error {
	*o = LastSync{}
	return nil
}

func lastSyncBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *LastSync) error {
	*o = LastSync{}
	return nil
}

func lastSyncAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *LastSync) error {
	*o = LastSync{}
	return nil
}

func lastSyncBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *LastSync) error {
	*o = LastSync{}
	return nil
}

func lastSyncAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *LastSync) error {
	*o = LastSync{}
	return nil
}

func testLastSyncsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &LastSync{}
	o := &LastSync{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, lastSyncDBTypes, false); err != nil {
		t.Errorf("Unable to randomize LastSyncTime object: %s", err)
	}

	AddLastSyncHook(boil.BeforeInsertHook, lastSyncBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	lastSyncBeforeInsertHooks = []LastSyncHook{}

	AddLastSyncHook(boil.AfterInsertHook, lastSyncAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	lastSyncAfterInsertHooks = []LastSyncHook{}

	AddLastSyncHook(boil.AfterSelectHook, lastSyncAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	lastSyncAfterSelectHooks = []LastSyncHook{}

	AddLastSyncHook(boil.BeforeUpdateHook, lastSyncBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	lastSyncBeforeUpdateHooks = []LastSyncHook{}

	AddLastSyncHook(boil.AfterUpdateHook, lastSyncAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	lastSyncAfterUpdateHooks = []LastSyncHook{}

	AddLastSyncHook(boil.BeforeDeleteHook, lastSyncBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	lastSyncBeforeDeleteHooks = []LastSyncHook{}

	AddLastSyncHook(boil.AfterDeleteHook, lastSyncAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	lastSyncAfterDeleteHooks = []LastSyncHook{}

	AddLastSyncHook(boil.BeforeUpsertHook, lastSyncBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	lastSyncBeforeUpsertHooks = []LastSyncHook{}

	AddLastSyncHook(boil.AfterUpsertHook, lastSyncAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	lastSyncAfterUpsertHooks = []LastSyncHook{}
}

func testLastSyncsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LastSync{}
	if err = randomize.Struct(seed, o, lastSyncDBTypes, true, lastSyncColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LastSyncTime struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := LastSyncs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testLastSyncsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LastSync{}
	if err = randomize.Struct(seed, o, lastSyncDBTypes, true); err != nil {
		t.Errorf("Unable to randomize LastSyncTime struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(lastSyncColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := LastSyncs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testLastSyncsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LastSync{}
	if err = randomize.Struct(seed, o, lastSyncDBTypes, true, lastSyncColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LastSyncTime struct: %s", err)
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

func testLastSyncsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LastSync{}
	if err = randomize.Struct(seed, o, lastSyncDBTypes, true, lastSyncColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LastSyncTime struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := LastSyncSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testLastSyncsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LastSync{}
	if err = randomize.Struct(seed, o, lastSyncDBTypes, true, lastSyncColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LastSyncTime struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := LastSyncs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	lastSyncDBTypes = map[string]string{`BaseAsset`: `text`, `QuoteAsset`: `text`, `Exchange`: `text`, `LastSyncTime`: `timestamp with time zone`}
	_               = bytes.MinRead
)

func testLastSyncsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(lastSyncPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(lastSyncAllColumns) == len(lastSyncPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &LastSync{}
	if err = randomize.Struct(seed, o, lastSyncDBTypes, true, lastSyncColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LastSyncTime struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := LastSyncs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, lastSyncDBTypes, true, lastSyncPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize LastSyncTime struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testLastSyncsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(lastSyncAllColumns) == len(lastSyncPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &LastSync{}
	if err = randomize.Struct(seed, o, lastSyncDBTypes, true, lastSyncColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LastSyncTime struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := LastSyncs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, lastSyncDBTypes, true, lastSyncPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize LastSyncTime struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(lastSyncAllColumns, lastSyncPrimaryKeyColumns) {
		fields = lastSyncAllColumns
	} else {
		fields = strmangle.SetComplement(
			lastSyncAllColumns,
			lastSyncPrimaryKeyColumns,
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

	slice := LastSyncSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testLastSyncsUpsert(t *testing.T) {
	t.Parallel()

	if len(lastSyncAllColumns) == len(lastSyncPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := LastSync{}
	if err = randomize.Struct(seed, &o, lastSyncDBTypes, true); err != nil {
		t.Errorf("Unable to randomize LastSyncTime struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert LastSyncTime: %s", err)
	}

	count, err := LastSyncs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, lastSyncDBTypes, false, lastSyncPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize LastSyncTime struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert LastSyncTime: %s", err)
	}

	count, err = LastSyncs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}