// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testHospitals(t *testing.T) {
	t.Parallel()

	query := Hospitals()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testHospitalsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Hospital{}
	if err = randomize.Struct(seed, o, hospitalDBTypes, true, hospitalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hospital struct: %s", err)
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

	count, err := Hospitals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testHospitalsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Hospital{}
	if err = randomize.Struct(seed, o, hospitalDBTypes, true, hospitalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hospital struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Hospitals().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Hospitals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testHospitalsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Hospital{}
	if err = randomize.Struct(seed, o, hospitalDBTypes, true, hospitalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hospital struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := HospitalSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Hospitals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testHospitalsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Hospital{}
	if err = randomize.Struct(seed, o, hospitalDBTypes, true, hospitalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hospital struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := HospitalExists(ctx, tx, o.HospitalID)
	if err != nil {
		t.Errorf("Unable to check if Hospital exists: %s", err)
	}
	if !e {
		t.Errorf("Expected HospitalExists to return true, but got false.")
	}
}

func testHospitalsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Hospital{}
	if err = randomize.Struct(seed, o, hospitalDBTypes, true, hospitalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hospital struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	hospitalFound, err := FindHospital(ctx, tx, o.HospitalID)
	if err != nil {
		t.Error(err)
	}

	if hospitalFound == nil {
		t.Error("want a record, got nil")
	}
}

func testHospitalsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Hospital{}
	if err = randomize.Struct(seed, o, hospitalDBTypes, true, hospitalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hospital struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Hospitals().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testHospitalsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Hospital{}
	if err = randomize.Struct(seed, o, hospitalDBTypes, true, hospitalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hospital struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Hospitals().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testHospitalsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	hospitalOne := &Hospital{}
	hospitalTwo := &Hospital{}
	if err = randomize.Struct(seed, hospitalOne, hospitalDBTypes, false, hospitalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hospital struct: %s", err)
	}
	if err = randomize.Struct(seed, hospitalTwo, hospitalDBTypes, false, hospitalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hospital struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = hospitalOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = hospitalTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Hospitals().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testHospitalsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	hospitalOne := &Hospital{}
	hospitalTwo := &Hospital{}
	if err = randomize.Struct(seed, hospitalOne, hospitalDBTypes, false, hospitalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hospital struct: %s", err)
	}
	if err = randomize.Struct(seed, hospitalTwo, hospitalDBTypes, false, hospitalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hospital struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = hospitalOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = hospitalTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Hospitals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func hospitalBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Hospital) error {
	*o = Hospital{}
	return nil
}

func hospitalAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Hospital) error {
	*o = Hospital{}
	return nil
}

func hospitalAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Hospital) error {
	*o = Hospital{}
	return nil
}

func hospitalBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Hospital) error {
	*o = Hospital{}
	return nil
}

func hospitalAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Hospital) error {
	*o = Hospital{}
	return nil
}

func hospitalBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Hospital) error {
	*o = Hospital{}
	return nil
}

func hospitalAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Hospital) error {
	*o = Hospital{}
	return nil
}

func hospitalBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Hospital) error {
	*o = Hospital{}
	return nil
}

func hospitalAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Hospital) error {
	*o = Hospital{}
	return nil
}

func testHospitalsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Hospital{}
	o := &Hospital{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, hospitalDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Hospital object: %s", err)
	}

	AddHospitalHook(boil.BeforeInsertHook, hospitalBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	hospitalBeforeInsertHooks = []HospitalHook{}

	AddHospitalHook(boil.AfterInsertHook, hospitalAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	hospitalAfterInsertHooks = []HospitalHook{}

	AddHospitalHook(boil.AfterSelectHook, hospitalAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	hospitalAfterSelectHooks = []HospitalHook{}

	AddHospitalHook(boil.BeforeUpdateHook, hospitalBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	hospitalBeforeUpdateHooks = []HospitalHook{}

	AddHospitalHook(boil.AfterUpdateHook, hospitalAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	hospitalAfterUpdateHooks = []HospitalHook{}

	AddHospitalHook(boil.BeforeDeleteHook, hospitalBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	hospitalBeforeDeleteHooks = []HospitalHook{}

	AddHospitalHook(boil.AfterDeleteHook, hospitalAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	hospitalAfterDeleteHooks = []HospitalHook{}

	AddHospitalHook(boil.BeforeUpsertHook, hospitalBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	hospitalBeforeUpsertHooks = []HospitalHook{}

	AddHospitalHook(boil.AfterUpsertHook, hospitalAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	hospitalAfterUpsertHooks = []HospitalHook{}
}

func testHospitalsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Hospital{}
	if err = randomize.Struct(seed, o, hospitalDBTypes, true, hospitalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hospital struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Hospitals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testHospitalsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Hospital{}
	if err = randomize.Struct(seed, o, hospitalDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Hospital struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(hospitalColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Hospitals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testHospitalsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Hospital{}
	if err = randomize.Struct(seed, o, hospitalDBTypes, true, hospitalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hospital struct: %s", err)
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

func testHospitalsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Hospital{}
	if err = randomize.Struct(seed, o, hospitalDBTypes, true, hospitalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hospital struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := HospitalSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testHospitalsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Hospital{}
	if err = randomize.Struct(seed, o, hospitalDBTypes, true, hospitalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hospital struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Hospitals().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	hospitalDBTypes = map[string]string{`HospitalID`: `int`, `Name`: `varchar`, `CreatedAt`: `datetime`, `Userid`: `varchar`, `Userpass`: `varchar`, `Mailaddress`: `varchar`}
	_               = bytes.MinRead
)

func testHospitalsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(hospitalPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(hospitalAllColumns) == len(hospitalPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Hospital{}
	if err = randomize.Struct(seed, o, hospitalDBTypes, true, hospitalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hospital struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Hospitals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, hospitalDBTypes, true, hospitalPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Hospital struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testHospitalsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(hospitalAllColumns) == len(hospitalPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Hospital{}
	if err = randomize.Struct(seed, o, hospitalDBTypes, true, hospitalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hospital struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Hospitals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, hospitalDBTypes, true, hospitalPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Hospital struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(hospitalAllColumns, hospitalPrimaryKeyColumns) {
		fields = hospitalAllColumns
	} else {
		fields = strmangle.SetComplement(
			hospitalAllColumns,
			hospitalPrimaryKeyColumns,
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

	slice := HospitalSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testHospitalsUpsert(t *testing.T) {
	t.Parallel()

	if len(hospitalAllColumns) == len(hospitalPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLHospitalUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Hospital{}
	if err = randomize.Struct(seed, &o, hospitalDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Hospital struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Hospital: %s", err)
	}

	count, err := Hospitals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, hospitalDBTypes, false, hospitalPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Hospital struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Hospital: %s", err)
	}

	count, err = Hospitals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
