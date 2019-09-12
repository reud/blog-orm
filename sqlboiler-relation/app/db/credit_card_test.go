// Code generated by SQLBoiler 3.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package db

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

func testCreditCards(t *testing.T) {
	t.Parallel()

	query := CreditCards()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testCreditCardsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CreditCard{}
	if err = randomize.Struct(seed, o, creditCardDBTypes, true, creditCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CreditCard struct: %s", err)
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

	count, err := CreditCards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCreditCardsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CreditCard{}
	if err = randomize.Struct(seed, o, creditCardDBTypes, true, creditCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CreditCard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := CreditCards().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := CreditCards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCreditCardsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CreditCard{}
	if err = randomize.Struct(seed, o, creditCardDBTypes, true, creditCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CreditCard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CreditCardSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := CreditCards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCreditCardsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CreditCard{}
	if err = randomize.Struct(seed, o, creditCardDBTypes, true, creditCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CreditCard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := CreditCardExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if CreditCard exists: %s", err)
	}
	if !e {
		t.Errorf("Expected CreditCardExists to return true, but got false.")
	}
}

func testCreditCardsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CreditCard{}
	if err = randomize.Struct(seed, o, creditCardDBTypes, true, creditCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CreditCard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	creditCardFound, err := FindCreditCard(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if creditCardFound == nil {
		t.Error("want a record, got nil")
	}
}

func testCreditCardsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CreditCard{}
	if err = randomize.Struct(seed, o, creditCardDBTypes, true, creditCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CreditCard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = CreditCards().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testCreditCardsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CreditCard{}
	if err = randomize.Struct(seed, o, creditCardDBTypes, true, creditCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CreditCard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := CreditCards().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testCreditCardsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	creditCardOne := &CreditCard{}
	creditCardTwo := &CreditCard{}
	if err = randomize.Struct(seed, creditCardOne, creditCardDBTypes, false, creditCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CreditCard struct: %s", err)
	}
	if err = randomize.Struct(seed, creditCardTwo, creditCardDBTypes, false, creditCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CreditCard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = creditCardOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = creditCardTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := CreditCards().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testCreditCardsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	creditCardOne := &CreditCard{}
	creditCardTwo := &CreditCard{}
	if err = randomize.Struct(seed, creditCardOne, creditCardDBTypes, false, creditCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CreditCard struct: %s", err)
	}
	if err = randomize.Struct(seed, creditCardTwo, creditCardDBTypes, false, creditCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CreditCard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = creditCardOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = creditCardTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CreditCards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func creditCardBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *CreditCard) error {
	*o = CreditCard{}
	return nil
}

func creditCardAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *CreditCard) error {
	*o = CreditCard{}
	return nil
}

func creditCardAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *CreditCard) error {
	*o = CreditCard{}
	return nil
}

func creditCardBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *CreditCard) error {
	*o = CreditCard{}
	return nil
}

func creditCardAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *CreditCard) error {
	*o = CreditCard{}
	return nil
}

func creditCardBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *CreditCard) error {
	*o = CreditCard{}
	return nil
}

func creditCardAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *CreditCard) error {
	*o = CreditCard{}
	return nil
}

func creditCardBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *CreditCard) error {
	*o = CreditCard{}
	return nil
}

func creditCardAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *CreditCard) error {
	*o = CreditCard{}
	return nil
}

func testCreditCardsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &CreditCard{}
	o := &CreditCard{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, creditCardDBTypes, false); err != nil {
		t.Errorf("Unable to randomize CreditCard object: %s", err)
	}

	AddCreditCardHook(boil.BeforeInsertHook, creditCardBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	creditCardBeforeInsertHooks = []CreditCardHook{}

	AddCreditCardHook(boil.AfterInsertHook, creditCardAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	creditCardAfterInsertHooks = []CreditCardHook{}

	AddCreditCardHook(boil.AfterSelectHook, creditCardAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	creditCardAfterSelectHooks = []CreditCardHook{}

	AddCreditCardHook(boil.BeforeUpdateHook, creditCardBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	creditCardBeforeUpdateHooks = []CreditCardHook{}

	AddCreditCardHook(boil.AfterUpdateHook, creditCardAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	creditCardAfterUpdateHooks = []CreditCardHook{}

	AddCreditCardHook(boil.BeforeDeleteHook, creditCardBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	creditCardBeforeDeleteHooks = []CreditCardHook{}

	AddCreditCardHook(boil.AfterDeleteHook, creditCardAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	creditCardAfterDeleteHooks = []CreditCardHook{}

	AddCreditCardHook(boil.BeforeUpsertHook, creditCardBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	creditCardBeforeUpsertHooks = []CreditCardHook{}

	AddCreditCardHook(boil.AfterUpsertHook, creditCardAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	creditCardAfterUpsertHooks = []CreditCardHook{}
}

func testCreditCardsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CreditCard{}
	if err = randomize.Struct(seed, o, creditCardDBTypes, true, creditCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CreditCard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CreditCards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCreditCardsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CreditCard{}
	if err = randomize.Struct(seed, o, creditCardDBTypes, true); err != nil {
		t.Errorf("Unable to randomize CreditCard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(creditCardColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := CreditCards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCreditCardToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local CreditCard
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, creditCardDBTypes, true, creditCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CreditCard struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.UserID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.User().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := CreditCardSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*CreditCard)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.User = nil
	if err = local.L.LoadUser(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCreditCardToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a CreditCard
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, creditCardDBTypes, false, strmangle.SetComplement(creditCardPrimaryKeyColumns, creditCardColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.CreditCards[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.UserID, x.ID) {
			t.Error("foreign key was wrong value", a.UserID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UserID))
		reflect.Indirect(reflect.ValueOf(&a.UserID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.UserID, x.ID) {
			t.Error("foreign key was wrong value", a.UserID, x.ID)
		}
	}
}

func testCreditCardToOneRemoveOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a CreditCard
	var b User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, creditCardDBTypes, false, strmangle.SetComplement(creditCardPrimaryKeyColumns, creditCardColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetUser(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveUser(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.User().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.User != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.UserID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.CreditCards) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testCreditCardsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CreditCard{}
	if err = randomize.Struct(seed, o, creditCardDBTypes, true, creditCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CreditCard struct: %s", err)
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

func testCreditCardsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CreditCard{}
	if err = randomize.Struct(seed, o, creditCardDBTypes, true, creditCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CreditCard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CreditCardSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testCreditCardsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CreditCard{}
	if err = randomize.Struct(seed, o, creditCardDBTypes, true, creditCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CreditCard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := CreditCards().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	creditCardDBTypes = map[string]string{`ID`: `integer`, `Number`: `text`, `UserID`: `integer`}
	_                 = bytes.MinRead
)

func testCreditCardsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(creditCardPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(creditCardAllColumns) == len(creditCardPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &CreditCard{}
	if err = randomize.Struct(seed, o, creditCardDBTypes, true, creditCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CreditCard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CreditCards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, creditCardDBTypes, true, creditCardPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CreditCard struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testCreditCardsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(creditCardAllColumns) == len(creditCardPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &CreditCard{}
	if err = randomize.Struct(seed, o, creditCardDBTypes, true, creditCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CreditCard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CreditCards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, creditCardDBTypes, true, creditCardPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CreditCard struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(creditCardAllColumns, creditCardPrimaryKeyColumns) {
		fields = creditCardAllColumns
	} else {
		fields = strmangle.SetComplement(
			creditCardAllColumns,
			creditCardPrimaryKeyColumns,
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

	slice := CreditCardSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testCreditCardsUpsert(t *testing.T) {
	t.Parallel()

	if len(creditCardAllColumns) == len(creditCardPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := CreditCard{}
	if err = randomize.Struct(seed, &o, creditCardDBTypes, true); err != nil {
		t.Errorf("Unable to randomize CreditCard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert CreditCard: %s", err)
	}

	count, err := CreditCards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, creditCardDBTypes, false, creditCardPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CreditCard struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert CreditCard: %s", err)
	}

	count, err = CreditCards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}