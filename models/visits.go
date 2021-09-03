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
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Visit is an object representing the database table.
type Visit struct {
	ID        int       `form:"id" boil:"id" json:"id" toml:"id" yaml:"id"`
	Active    null.Bool `form:"active" boil:"active" json:"active,omitempty" toml:"active" yaml:"active,omitempty"`
	StudentID int       `form:"student_id" boil:"student_id" json:"student_id" toml:"student_id" yaml:"student_id"`
	Date      string    `form:"date" boil:"date" json:"date" toml:"date" yaml:"date"`
	Time      string    `form:"time" boil:"time" json:"time" toml:"time" yaml:"time"`
	Type      string    `form:"type" boil:"type" json:"type" toml:"type" yaml:"type"`

	R *visitR `form:"-" boil:"-" json:"-" toml:"-" yaml:"-"`
	L visitL  `form:"-" boil:"-" json:"-" toml:"-" yaml:"-"`
}

var VisitColumns = struct {
	ID        string
	Active    string
	StudentID string
	Date      string
	Time      string
	Type      string
}{
	ID:        "id",
	Active:    "active",
	StudentID: "student_id",
	Date:      "date",
	Time:      "time",
	Type:      "type",
}

var VisitTableColumns = struct {
	ID        string
	Active    string
	StudentID string
	Date      string
	Time      string
	Type      string
}{
	ID:        "visits.id",
	Active:    "visits.active",
	StudentID: "visits.student_id",
	Date:      "visits.date",
	Time:      "visits.time",
	Type:      "visits.type",
}

// Generated where

var VisitWhere = struct {
	ID        whereHelperint
	Active    whereHelpernull_Bool
	StudentID whereHelperint
	Date      whereHelperstring
	Time      whereHelperstring
	Type      whereHelperstring
}{
	ID:        whereHelperint{field: "`visits`.`id`"},
	Active:    whereHelpernull_Bool{field: "`visits`.`active`"},
	StudentID: whereHelperint{field: "`visits`.`student_id`"},
	Date:      whereHelperstring{field: "`visits`.`date`"},
	Time:      whereHelperstring{field: "`visits`.`time`"},
	Type:      whereHelperstring{field: "`visits`.`type`"},
}

// VisitRels is where relationship names are stored.
var VisitRels = struct {
	Student string
}{
	Student: "Student",
}

// visitR is where relationships are stored.
type visitR struct {
	Student *Student `form:"Student" boil:"Student" json:"Student" toml:"Student" yaml:"Student"`
}

// NewStruct creates a new relationship struct
func (*visitR) NewStruct() *visitR {
	return &visitR{}
}

// visitL is where Load methods for each relationship are stored.
type visitL struct{}

var (
	visitAllColumns            = []string{"id", "active", "student_id", "date", "time", "type"}
	visitColumnsWithoutDefault = []string{"student_id", "date", "time", "type"}
	visitColumnsWithDefault    = []string{"id", "active"}
	visitPrimaryKeyColumns     = []string{"id"}
)

type (
	// VisitSlice is an alias for a slice of pointers to Visit.
	// This should almost always be used instead of []Visit.
	VisitSlice []*Visit
	// VisitHook is the signature for custom Visit hook methods
	VisitHook func(context.Context, boil.ContextExecutor, *Visit) error

	visitQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	visitType                 = reflect.TypeOf(&Visit{})
	visitMapping              = queries.MakeStructMapping(visitType)
	visitPrimaryKeyMapping, _ = queries.BindMapping(visitType, visitMapping, visitPrimaryKeyColumns)
	visitInsertCacheMut       sync.RWMutex
	visitInsertCache          = make(map[string]insertCache)
	visitUpdateCacheMut       sync.RWMutex
	visitUpdateCache          = make(map[string]updateCache)
	visitUpsertCacheMut       sync.RWMutex
	visitUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var visitBeforeInsertHooks []VisitHook
var visitBeforeUpdateHooks []VisitHook
var visitBeforeDeleteHooks []VisitHook
var visitBeforeUpsertHooks []VisitHook

var visitAfterInsertHooks []VisitHook
var visitAfterSelectHooks []VisitHook
var visitAfterUpdateHooks []VisitHook
var visitAfterDeleteHooks []VisitHook
var visitAfterUpsertHooks []VisitHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Visit) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range visitBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Visit) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range visitBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Visit) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range visitBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Visit) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range visitBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Visit) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range visitAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Visit) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range visitAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Visit) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range visitAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Visit) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range visitAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Visit) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range visitAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddVisitHook registers your hook function for all future operations.
func AddVisitHook(hookPoint boil.HookPoint, visitHook VisitHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		visitBeforeInsertHooks = append(visitBeforeInsertHooks, visitHook)
	case boil.BeforeUpdateHook:
		visitBeforeUpdateHooks = append(visitBeforeUpdateHooks, visitHook)
	case boil.BeforeDeleteHook:
		visitBeforeDeleteHooks = append(visitBeforeDeleteHooks, visitHook)
	case boil.BeforeUpsertHook:
		visitBeforeUpsertHooks = append(visitBeforeUpsertHooks, visitHook)
	case boil.AfterInsertHook:
		visitAfterInsertHooks = append(visitAfterInsertHooks, visitHook)
	case boil.AfterSelectHook:
		visitAfterSelectHooks = append(visitAfterSelectHooks, visitHook)
	case boil.AfterUpdateHook:
		visitAfterUpdateHooks = append(visitAfterUpdateHooks, visitHook)
	case boil.AfterDeleteHook:
		visitAfterDeleteHooks = append(visitAfterDeleteHooks, visitHook)
	case boil.AfterUpsertHook:
		visitAfterUpsertHooks = append(visitAfterUpsertHooks, visitHook)
	}
}

// One returns a single visit record from the query.
func (q visitQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Visit, error) {
	o := &Visit{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for visits")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Visit records from the query.
func (q visitQuery) All(ctx context.Context, exec boil.ContextExecutor) (VisitSlice, error) {
	var o []*Visit

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Visit slice")
	}

	if len(visitAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Visit records in the query.
func (q visitQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count visits rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q visitQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if visits exists")
	}

	return count > 0, nil
}

// Student pointed to by the foreign key.
func (o *Visit) Student(mods ...qm.QueryMod) studentQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.StudentID),
	}

	queryMods = append(queryMods, mods...)

	query := Students(queryMods...)
	queries.SetFrom(query.Query, "`students`")

	return query
}

// LoadStudent allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (visitL) LoadStudent(ctx context.Context, e boil.ContextExecutor, singular bool, maybeVisit interface{}, mods queries.Applicator) error {
	var slice []*Visit
	var object *Visit

	if singular {
		object = maybeVisit.(*Visit)
	} else {
		slice = *maybeVisit.(*[]*Visit)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &visitR{}
		}
		args = append(args, object.StudentID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &visitR{}
			}

			for _, a := range args {
				if a == obj.StudentID {
					continue Outer
				}
			}

			args = append(args, obj.StudentID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`students`),
		qm.WhereIn(`students.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Student")
	}

	var resultSlice []*Student
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Student")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for students")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for students")
	}

	if len(visitAfterSelectHooks) != 0 {
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
		object.R.Student = foreign
		if foreign.R == nil {
			foreign.R = &studentR{}
		}
		foreign.R.Visits = append(foreign.R.Visits, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.StudentID == foreign.ID {
				local.R.Student = foreign
				if foreign.R == nil {
					foreign.R = &studentR{}
				}
				foreign.R.Visits = append(foreign.R.Visits, local)
				break
			}
		}
	}

	return nil
}

// SetStudent of the visit to the related item.
// Sets o.R.Student to related.
// Adds o to related.R.Visits.
func (o *Visit) SetStudent(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Student) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `visits` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"student_id"}),
		strmangle.WhereClause("`", "`", 0, visitPrimaryKeyColumns),
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

	o.StudentID = related.ID
	if o.R == nil {
		o.R = &visitR{
			Student: related,
		}
	} else {
		o.R.Student = related
	}

	if related.R == nil {
		related.R = &studentR{
			Visits: VisitSlice{o},
		}
	} else {
		related.R.Visits = append(related.R.Visits, o)
	}

	return nil
}

// Visits retrieves all the records using an executor.
func Visits(mods ...qm.QueryMod) visitQuery {
	mods = append(mods, qm.From("`visits`"))
	return visitQuery{NewQuery(mods...)}
}

// FindVisit retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindVisit(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Visit, error) {
	visitObj := &Visit{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `visits` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, visitObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from visits")
	}

	if err = visitObj.doAfterSelectHooks(ctx, exec); err != nil {
		return visitObj, err
	}

	return visitObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Visit) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no visits provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(visitColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	visitInsertCacheMut.RLock()
	cache, cached := visitInsertCache[key]
	visitInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			visitAllColumns,
			visitColumnsWithDefault,
			visitColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(visitType, visitMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(visitType, visitMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `visits` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `visits` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `visits` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, visitPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into visits")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == visitMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for visits")
	}

CacheNoHooks:
	if !cached {
		visitInsertCacheMut.Lock()
		visitInsertCache[key] = cache
		visitInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Visit.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Visit) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	visitUpdateCacheMut.RLock()
	cache, cached := visitUpdateCache[key]
	visitUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			visitAllColumns,
			visitPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update visits, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `visits` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, visitPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(visitType, visitMapping, append(wl, visitPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update visits row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for visits")
	}

	if !cached {
		visitUpdateCacheMut.Lock()
		visitUpdateCache[key] = cache
		visitUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q visitQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for visits")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for visits")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o VisitSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), visitPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `visits` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, visitPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in visit slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all visit")
	}
	return rowsAff, nil
}

var mySQLVisitUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Visit) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no visits provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(visitColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLVisitUniqueColumns, o)

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

	visitUpsertCacheMut.RLock()
	cache, cached := visitUpsertCache[key]
	visitUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			visitAllColumns,
			visitColumnsWithDefault,
			visitColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			visitAllColumns,
			visitPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert visits, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`visits`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `visits` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(visitType, visitMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(visitType, visitMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for visits")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == visitMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(visitType, visitMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for visits")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for visits")
	}

CacheNoHooks:
	if !cached {
		visitUpsertCacheMut.Lock()
		visitUpsertCache[key] = cache
		visitUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Visit record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Visit) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Visit provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), visitPrimaryKeyMapping)
	sql := "DELETE FROM `visits` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from visits")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for visits")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q visitQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no visitQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from visits")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for visits")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o VisitSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(visitBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), visitPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `visits` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, visitPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from visit slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for visits")
	}

	if len(visitAfterDeleteHooks) != 0 {
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
func (o *Visit) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindVisit(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *VisitSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := VisitSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), visitPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `visits`.* FROM `visits` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, visitPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in VisitSlice")
	}

	*o = slice

	return nil
}

// VisitExists checks if the Visit row exists.
func VisitExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `visits` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if visits exists")
	}

	return exists, nil
}
