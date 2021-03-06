// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Cost is an object representing the database table.
type Cost struct {
	ID            int    `form:"id" boil:"id" json:"id" toml:"id" yaml:"id"`
	Active        bool   `form:"active" boil:"active" json:"active" toml:"active" yaml:"active"`
	Product       string `form:"product" boil:"product" json:"product" toml:"product" yaml:"product"`
	Cost          int16  `form:"cost" boil:"cost" json:"cost" toml:"cost" yaml:"cost"`
	Date          string `form:"date" boil:"date" json:"date" toml:"date" yaml:"date"`
	Time          string `form:"time" boil:"time" json:"time" toml:"time" yaml:"time"`
	PaymentObject int    `form:"payment_object" boil:"payment_object" json:"payment_object" toml:"payment_object" yaml:"payment_object"`

	R *costR `form:"-" boil:"-" json:"-" toml:"-" yaml:"-"`
	L costL  `form:"-" boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CostColumns = struct {
	ID            string
	Active        string
	Product       string
	Cost          string
	Date          string
	Time          string
	PaymentObject string
}{
	ID:            "id",
	Active:        "active",
	Product:       "product",
	Cost:          "cost",
	Date:          "date",
	Time:          "time",
	PaymentObject: "payment_object",
}

var CostTableColumns = struct {
	ID            string
	Active        string
	Product       string
	Cost          string
	Date          string
	Time          string
	PaymentObject string
}{
	ID:            "costs.id",
	Active:        "costs.active",
	Product:       "costs.product",
	Cost:          "costs.cost",
	Date:          "costs.date",
	Time:          "costs.time",
	PaymentObject: "costs.payment_object",
}

// Generated where

type whereHelperint16 struct{ field string }

func (w whereHelperint16) EQ(x int16) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint16) NEQ(x int16) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint16) LT(x int16) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint16) LTE(x int16) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint16) GT(x int16) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint16) GTE(x int16) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint16) IN(slice []int16) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint16) NIN(slice []int16) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var CostWhere = struct {
	ID            whereHelperint
	Active        whereHelperbool
	Product       whereHelperstring
	Cost          whereHelperint16
	Date          whereHelperstring
	Time          whereHelperstring
	PaymentObject whereHelperint
}{
	ID:            whereHelperint{field: "`costs`.`id`"},
	Active:        whereHelperbool{field: "`costs`.`active`"},
	Product:       whereHelperstring{field: "`costs`.`product`"},
	Cost:          whereHelperint16{field: "`costs`.`cost`"},
	Date:          whereHelperstring{field: "`costs`.`date`"},
	Time:          whereHelperstring{field: "`costs`.`time`"},
	PaymentObject: whereHelperint{field: "`costs`.`payment_object`"},
}

// CostRels is where relationship names are stored.
var CostRels = struct {
	CostPaymentObject string
}{
	CostPaymentObject: "CostPaymentObject",
}

// costR is where relationships are stored.
type costR struct {
	CostPaymentObject *PaymentObject `form:"CostPaymentObject" boil:"CostPaymentObject" json:"CostPaymentObject" toml:"CostPaymentObject" yaml:"CostPaymentObject"`
}

// NewStruct creates a new relationship struct
func (*costR) NewStruct() *costR {
	return &costR{}
}

// costL is where Load methods for each relationship are stored.
type costL struct{}

var (
	costAllColumns            = []string{"id", "active", "product", "cost", "date", "time", "payment_object"}
	costColumnsWithoutDefault = []string{"product", "cost", "date", "time", "payment_object"}
	costColumnsWithDefault    = []string{"id", "active"}
	costPrimaryKeyColumns     = []string{"id"}
)

type (
	// CostSlice is an alias for a slice of pointers to Cost.
	// This should almost always be used instead of []Cost.
	CostSlice []*Cost
	// CostHook is the signature for custom Cost hook methods
	CostHook func(context.Context, boil.ContextExecutor, *Cost) error

	costQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	costType                 = reflect.TypeOf(&Cost{})
	costMapping              = queries.MakeStructMapping(costType)
	costPrimaryKeyMapping, _ = queries.BindMapping(costType, costMapping, costPrimaryKeyColumns)
	costInsertCacheMut       sync.RWMutex
	costInsertCache          = make(map[string]insertCache)
	costUpdateCacheMut       sync.RWMutex
	costUpdateCache          = make(map[string]updateCache)
	costUpsertCacheMut       sync.RWMutex
	costUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var costBeforeInsertHooks []CostHook
var costBeforeUpdateHooks []CostHook
var costBeforeDeleteHooks []CostHook
var costBeforeUpsertHooks []CostHook

var costAfterInsertHooks []CostHook
var costAfterSelectHooks []CostHook
var costAfterUpdateHooks []CostHook
var costAfterDeleteHooks []CostHook
var costAfterUpsertHooks []CostHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Cost) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range costBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Cost) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range costBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Cost) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range costBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Cost) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range costBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Cost) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range costAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Cost) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range costAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Cost) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range costAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Cost) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range costAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Cost) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range costAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCostHook registers your hook function for all future operations.
func AddCostHook(hookPoint boil.HookPoint, costHook CostHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		costBeforeInsertHooks = append(costBeforeInsertHooks, costHook)
	case boil.BeforeUpdateHook:
		costBeforeUpdateHooks = append(costBeforeUpdateHooks, costHook)
	case boil.BeforeDeleteHook:
		costBeforeDeleteHooks = append(costBeforeDeleteHooks, costHook)
	case boil.BeforeUpsertHook:
		costBeforeUpsertHooks = append(costBeforeUpsertHooks, costHook)
	case boil.AfterInsertHook:
		costAfterInsertHooks = append(costAfterInsertHooks, costHook)
	case boil.AfterSelectHook:
		costAfterSelectHooks = append(costAfterSelectHooks, costHook)
	case boil.AfterUpdateHook:
		costAfterUpdateHooks = append(costAfterUpdateHooks, costHook)
	case boil.AfterDeleteHook:
		costAfterDeleteHooks = append(costAfterDeleteHooks, costHook)
	case boil.AfterUpsertHook:
		costAfterUpsertHooks = append(costAfterUpsertHooks, costHook)
	}
}

// One returns a single cost record from the query.
func (q costQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Cost, error) {
	o := &Cost{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for costs")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Cost records from the query.
func (q costQuery) All(ctx context.Context, exec boil.ContextExecutor) (CostSlice, error) {
	var o []*Cost

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Cost slice")
	}

	if len(costAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Cost records in the query.
func (q costQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count costs rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q costQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if costs exists")
	}

	return count > 0, nil
}

// CostPaymentObject pointed to by the foreign key.
func (o *Cost) CostPaymentObject(mods ...qm.QueryMod) paymentObjectQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.PaymentObject),
	}

	queryMods = append(queryMods, mods...)

	query := PaymentObjects(queryMods...)
	queries.SetFrom(query.Query, "`payment_objects`")

	return query
}

// LoadCostPaymentObject allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (costL) LoadCostPaymentObject(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCost interface{}, mods queries.Applicator) error {
	var slice []*Cost
	var object *Cost

	if singular {
		object = maybeCost.(*Cost)
	} else {
		slice = *maybeCost.(*[]*Cost)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &costR{}
		}
		args = append(args, object.PaymentObject)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &costR{}
			}

			for _, a := range args {
				if a == obj.PaymentObject {
					continue Outer
				}
			}

			args = append(args, obj.PaymentObject)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`payment_objects`),
		qm.WhereIn(`payment_objects.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load PaymentObject")
	}

	var resultSlice []*PaymentObject
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice PaymentObject")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for payment_objects")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for payment_objects")
	}

	if len(costAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.CostPaymentObject = foreign
		if foreign.R == nil {
			foreign.R = &paymentObjectR{}
		}
		foreign.R.Costs = append(foreign.R.Costs, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.PaymentObject == foreign.ID {
				local.R.CostPaymentObject = foreign
				if foreign.R == nil {
					foreign.R = &paymentObjectR{}
				}
				foreign.R.Costs = append(foreign.R.Costs, local)
				break
			}
		}
	}

	return nil
}

// SetCostPaymentObject of the cost to the related item.
// Sets o.R.CostPaymentObject to related.
// Adds o to related.R.Costs.
func (o *Cost) SetCostPaymentObject(ctx context.Context, exec boil.ContextExecutor, insert bool, related *PaymentObject) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `costs` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"payment_object"}),
		strmangle.WhereClause("`", "`", 0, costPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PaymentObject = related.ID
	if o.R == nil {
		o.R = &costR{
			CostPaymentObject: related,
		}
	} else {
		o.R.CostPaymentObject = related
	}

	if related.R == nil {
		related.R = &paymentObjectR{
			Costs: CostSlice{o},
		}
	} else {
		related.R.Costs = append(related.R.Costs, o)
	}

	return nil
}

// Costs retrieves all the records using an executor.
func Costs(mods ...qm.QueryMod) costQuery {
	mods = append(mods, qm.From("`costs`"))
	return costQuery{NewQuery(mods...)}
}

// FindCost retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCost(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Cost, error) {
	costObj := &Cost{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `costs` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, costObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from costs")
	}

	if err = costObj.doAfterSelectHooks(ctx, exec); err != nil {
		return costObj, err
	}

	return costObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Cost) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no costs provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(costColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	costInsertCacheMut.RLock()
	cache, cached := costInsertCache[key]
	costInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			costAllColumns,
			costColumnsWithDefault,
			costColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(costType, costMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(costType, costMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `costs` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `costs` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `costs` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, costPrimaryKeyColumns))
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into costs")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == costMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for costs")
	}

CacheNoHooks:
	if !cached {
		costInsertCacheMut.Lock()
		costInsertCache[key] = cache
		costInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Cost.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Cost) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	costUpdateCacheMut.RLock()
	cache, cached := costUpdateCache[key]
	costUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			costAllColumns,
			costPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update costs, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `costs` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, costPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(costType, costMapping, append(wl, costPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update costs row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for costs")
	}

	if !cached {
		costUpdateCacheMut.Lock()
		costUpdateCache[key] = cache
		costUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q costQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for costs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for costs")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CostSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), costPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `costs` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, costPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in cost slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all cost")
	}
	return rowsAff, nil
}

var mySQLCostUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Cost) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no costs provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(costColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCostUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
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
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	costUpsertCacheMut.RLock()
	cache, cached := costUpsertCache[key]
	costUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			costAllColumns,
			costColumnsWithDefault,
			costColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			costAllColumns,
			costPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert costs, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`costs`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `costs` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(costType, costMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(costType, costMapping, ret)
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for costs")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == costMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(costType, costMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for costs")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for costs")
	}

CacheNoHooks:
	if !cached {
		costUpsertCacheMut.Lock()
		costUpsertCache[key] = cache
		costUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Cost record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Cost) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Cost provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), costPrimaryKeyMapping)
	sql := "DELETE FROM `costs` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from costs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for costs")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q costQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no costQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from costs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for costs")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CostSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(costBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), costPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `costs` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, costPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cost slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for costs")
	}

	if len(costAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Cost) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCost(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CostSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CostSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), costPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `costs`.* FROM `costs` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, costPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CostSlice")
	}

	*o = slice

	return nil
}

// CostExists checks if the Cost row exists.
func CostExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `costs` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if costs exists")
	}

	return exists, nil
}
