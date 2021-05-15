// Code generated by SQLBoiler 4.4.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package model

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

func testSessions(t *testing.T) {
	t.Parallel()

	query := Sessions()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testSessionsSoftDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Session{}
	if err = randomize.Struct(seed, o, sessionDBTypes, true, sessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx, false); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Sessions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSessionsQuerySoftDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Session{}
	if err = randomize.Struct(seed, o, sessionDBTypes, true, sessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Sessions().DeleteAll(ctx, tx, false); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Sessions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSessionsSliceSoftDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Session{}
	if err = randomize.Struct(seed, o, sessionDBTypes, true, sessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SessionSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx, false); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Sessions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSessionsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Session{}
	if err = randomize.Struct(seed, o, sessionDBTypes, true, sessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx, true); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Sessions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSessionsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Session{}
	if err = randomize.Struct(seed, o, sessionDBTypes, true, sessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Sessions().DeleteAll(ctx, tx, true); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Sessions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSessionsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Session{}
	if err = randomize.Struct(seed, o, sessionDBTypes, true, sessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SessionSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx, true); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Sessions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSessionsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Session{}
	if err = randomize.Struct(seed, o, sessionDBTypes, true, sessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := SessionExists(ctx, tx, o.Sessionid)
	if err != nil {
		t.Errorf("Unable to check if Session exists: %s", err)
	}
	if !e {
		t.Errorf("Expected SessionExists to return true, but got false.")
	}
}

func testSessionsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Session{}
	if err = randomize.Struct(seed, o, sessionDBTypes, true, sessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	sessionFound, err := FindSession(ctx, tx, o.Sessionid)
	if err != nil {
		t.Error(err)
	}

	if sessionFound == nil {
		t.Error("want a record, got nil")
	}
}

func testSessionsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Session{}
	if err = randomize.Struct(seed, o, sessionDBTypes, true, sessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Sessions().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testSessionsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Session{}
	if err = randomize.Struct(seed, o, sessionDBTypes, true, sessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Sessions().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testSessionsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	sessionOne := &Session{}
	sessionTwo := &Session{}
	if err = randomize.Struct(seed, sessionOne, sessionDBTypes, false, sessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
	}
	if err = randomize.Struct(seed, sessionTwo, sessionDBTypes, false, sessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = sessionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = sessionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Sessions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testSessionsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	sessionOne := &Session{}
	sessionTwo := &Session{}
	if err = randomize.Struct(seed, sessionOne, sessionDBTypes, false, sessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
	}
	if err = randomize.Struct(seed, sessionTwo, sessionDBTypes, false, sessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = sessionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = sessionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Sessions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func sessionBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Session) error {
	*o = Session{}
	return nil
}

func sessionAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Session) error {
	*o = Session{}
	return nil
}

func sessionAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Session) error {
	*o = Session{}
	return nil
}

func sessionBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Session) error {
	*o = Session{}
	return nil
}

func sessionAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Session) error {
	*o = Session{}
	return nil
}

func sessionBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Session) error {
	*o = Session{}
	return nil
}

func sessionAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Session) error {
	*o = Session{}
	return nil
}

func sessionBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Session) error {
	*o = Session{}
	return nil
}

func sessionAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Session) error {
	*o = Session{}
	return nil
}

func testSessionsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Session{}
	o := &Session{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, sessionDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Session object: %s", err)
	}

	AddSessionHook(boil.BeforeInsertHook, sessionBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	sessionBeforeInsertHooks = []SessionHook{}

	AddSessionHook(boil.AfterInsertHook, sessionAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	sessionAfterInsertHooks = []SessionHook{}

	AddSessionHook(boil.AfterSelectHook, sessionAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	sessionAfterSelectHooks = []SessionHook{}

	AddSessionHook(boil.BeforeUpdateHook, sessionBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	sessionBeforeUpdateHooks = []SessionHook{}

	AddSessionHook(boil.AfterUpdateHook, sessionAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	sessionAfterUpdateHooks = []SessionHook{}

	AddSessionHook(boil.BeforeDeleteHook, sessionBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	sessionBeforeDeleteHooks = []SessionHook{}

	AddSessionHook(boil.AfterDeleteHook, sessionAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	sessionAfterDeleteHooks = []SessionHook{}

	AddSessionHook(boil.BeforeUpsertHook, sessionBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	sessionBeforeUpsertHooks = []SessionHook{}

	AddSessionHook(boil.AfterUpsertHook, sessionAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	sessionAfterUpsertHooks = []SessionHook{}
}

func testSessionsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Session{}
	if err = randomize.Struct(seed, o, sessionDBTypes, true, sessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Sessions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSessionsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Session{}
	if err = randomize.Struct(seed, o, sessionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(sessionColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Sessions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSessionsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Session{}
	if err = randomize.Struct(seed, o, sessionDBTypes, true, sessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
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

func testSessionsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Session{}
	if err = randomize.Struct(seed, o, sessionDBTypes, true, sessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SessionSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testSessionsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Session{}
	if err = randomize.Struct(seed, o, sessionDBTypes, true, sessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Sessions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	sessionDBTypes = map[string]string{`Sessionid`: `text`, `Expiry`: `timestamp with time zone`, `UserID`: `text`, `CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`, `DeletedAt`: `timestamp with time zone`}
	_              = bytes.MinRead
)

func testSessionsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(sessionPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(sessionAllColumns) == len(sessionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Session{}
	if err = randomize.Struct(seed, o, sessionDBTypes, true, sessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Sessions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, sessionDBTypes, true, sessionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testSessionsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(sessionAllColumns) == len(sessionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Session{}
	if err = randomize.Struct(seed, o, sessionDBTypes, true, sessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Sessions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, sessionDBTypes, true, sessionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(sessionAllColumns, sessionPrimaryKeyColumns) {
		fields = sessionAllColumns
	} else {
		fields = strmangle.SetComplement(
			sessionAllColumns,
			sessionPrimaryKeyColumns,
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

	slice := SessionSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testSessionsUpsert(t *testing.T) {
	t.Parallel()

	if len(sessionAllColumns) == len(sessionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Session{}
	if err = randomize.Struct(seed, &o, sessionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Session: %s", err)
	}

	count, err := Sessions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, sessionDBTypes, false, sessionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Session: %s", err)
	}

	count, err = Sessions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
