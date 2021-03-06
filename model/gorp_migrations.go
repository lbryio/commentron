// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package model

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// GorpMigration is an object representing the database table.
type GorpMigration struct {
	ID        string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	AppliedAt null.Time `boil:"applied_at" json:"applied_at,omitempty" toml:"applied_at" yaml:"applied_at,omitempty"`

	R *gorpMigrationR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L gorpMigrationL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var GorpMigrationColumns = struct {
	ID        string
	AppliedAt string
}{
	ID:        "id",
	AppliedAt: "applied_at",
}

// Generated where

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Time) LT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Time) LTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Time) GT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Time) GTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var GorpMigrationWhere = struct {
	ID        whereHelperstring
	AppliedAt whereHelpernull_Time
}{
	ID:        whereHelperstring{field: "`gorp_migrations`.`id`"},
	AppliedAt: whereHelpernull_Time{field: "`gorp_migrations`.`applied_at`"},
}

// GorpMigrationRels is where relationship names are stored.
var GorpMigrationRels = struct {
}{}

// gorpMigrationR is where relationships are stored.
type gorpMigrationR struct {
}

// NewStruct creates a new relationship struct
func (*gorpMigrationR) NewStruct() *gorpMigrationR {
	return &gorpMigrationR{}
}

// gorpMigrationL is where Load methods for each relationship are stored.
type gorpMigrationL struct{}

var (
	gorpMigrationAllColumns            = []string{"id", "applied_at"}
	gorpMigrationColumnsWithoutDefault = []string{"id", "applied_at"}
	gorpMigrationColumnsWithDefault    = []string{}
	gorpMigrationPrimaryKeyColumns     = []string{"id"}
)

type (
	// GorpMigrationSlice is an alias for a slice of pointers to GorpMigration.
	// This should generally be used opposed to []GorpMigration.
	GorpMigrationSlice []*GorpMigration

	gorpMigrationQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	gorpMigrationType                 = reflect.TypeOf(&GorpMigration{})
	gorpMigrationMapping              = queries.MakeStructMapping(gorpMigrationType)
	gorpMigrationPrimaryKeyMapping, _ = queries.BindMapping(gorpMigrationType, gorpMigrationMapping, gorpMigrationPrimaryKeyColumns)
	gorpMigrationInsertCacheMut       sync.RWMutex
	gorpMigrationInsertCache          = make(map[string]insertCache)
	gorpMigrationUpdateCacheMut       sync.RWMutex
	gorpMigrationUpdateCache          = make(map[string]updateCache)
	gorpMigrationUpsertCacheMut       sync.RWMutex
	gorpMigrationUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single gorpMigration record from the query using the global executor.
func (q gorpMigrationQuery) OneG() (*GorpMigration, error) {
	return q.One(boil.GetDB())
}

// OneGP returns a single gorpMigration record from the query using the global executor, and panics on error.
func (q gorpMigrationQuery) OneGP() *GorpMigration {
	o, err := q.One(boil.GetDB())
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// OneP returns a single gorpMigration record from the query, and panics on error.
func (q gorpMigrationQuery) OneP(exec boil.Executor) *GorpMigration {
	o, err := q.One(exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single gorpMigration record from the query.
func (q gorpMigrationQuery) One(exec boil.Executor) (*GorpMigration, error) {
	o := &GorpMigration{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: failed to execute a one query for gorp_migrations")
	}

	return o, nil
}

// AllG returns all GorpMigration records from the query using the global executor.
func (q gorpMigrationQuery) AllG() (GorpMigrationSlice, error) {
	return q.All(boil.GetDB())
}

// AllGP returns all GorpMigration records from the query using the global executor, and panics on error.
func (q gorpMigrationQuery) AllGP() GorpMigrationSlice {
	o, err := q.All(boil.GetDB())
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// AllP returns all GorpMigration records from the query, and panics on error.
func (q gorpMigrationQuery) AllP(exec boil.Executor) GorpMigrationSlice {
	o, err := q.All(exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all GorpMigration records from the query.
func (q gorpMigrationQuery) All(exec boil.Executor) (GorpMigrationSlice, error) {
	var o []*GorpMigration

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "model: failed to assign all query results to GorpMigration slice")
	}

	return o, nil
}

// CountG returns the count of all GorpMigration records in the query, and panics on error.
func (q gorpMigrationQuery) CountG() (int64, error) {
	return q.Count(boil.GetDB())
}

// CountGP returns the count of all GorpMigration records in the query using the global executor, and panics on error.
func (q gorpMigrationQuery) CountGP() int64 {
	c, err := q.Count(boil.GetDB())
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// CountP returns the count of all GorpMigration records in the query, and panics on error.
func (q gorpMigrationQuery) CountP(exec boil.Executor) int64 {
	c, err := q.Count(exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all GorpMigration records in the query.
func (q gorpMigrationQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to count gorp_migrations rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q gorpMigrationQuery) ExistsG() (bool, error) {
	return q.Exists(boil.GetDB())
}

// ExistsGP checks if the row exists in the table using the global executor, and panics on error.
func (q gorpMigrationQuery) ExistsGP() bool {
	e, err := q.Exists(boil.GetDB())
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// ExistsP checks if the row exists in the table, and panics on error.
func (q gorpMigrationQuery) ExistsP(exec boil.Executor) bool {
	e, err := q.Exists(exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q gorpMigrationQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "model: failed to check if gorp_migrations exists")
	}

	return count > 0, nil
}

// GorpMigrations retrieves all the records using an executor.
func GorpMigrations(mods ...qm.QueryMod) gorpMigrationQuery {
	mods = append(mods, qm.From("`gorp_migrations`"))
	return gorpMigrationQuery{NewQuery(mods...)}
}

// FindGorpMigrationG retrieves a single record by ID.
func FindGorpMigrationG(iD string, selectCols ...string) (*GorpMigration, error) {
	return FindGorpMigration(boil.GetDB(), iD, selectCols...)
}

// FindGorpMigrationP retrieves a single record by ID with an executor, and panics on error.
func FindGorpMigrationP(exec boil.Executor, iD string, selectCols ...string) *GorpMigration {
	retobj, err := FindGorpMigration(exec, iD, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindGorpMigrationGP retrieves a single record by ID, and panics on error.
func FindGorpMigrationGP(iD string, selectCols ...string) *GorpMigration {
	retobj, err := FindGorpMigration(boil.GetDB(), iD, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindGorpMigration retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindGorpMigration(exec boil.Executor, iD string, selectCols ...string) (*GorpMigration, error) {
	gorpMigrationObj := &GorpMigration{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `gorp_migrations` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, gorpMigrationObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: unable to select from gorp_migrations")
	}

	return gorpMigrationObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *GorpMigration) InsertG(columns boil.Columns) error {
	return o.Insert(boil.GetDB(), columns)
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *GorpMigration) InsertP(exec boil.Executor, columns boil.Columns) {
	if err := o.Insert(exec, columns); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *GorpMigration) InsertGP(columns boil.Columns) {
	if err := o.Insert(boil.GetDB(), columns); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *GorpMigration) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("model: no gorp_migrations provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(gorpMigrationColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	gorpMigrationInsertCacheMut.RLock()
	cache, cached := gorpMigrationInsertCache[key]
	gorpMigrationInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			gorpMigrationAllColumns,
			gorpMigrationColumnsWithDefault,
			gorpMigrationColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(gorpMigrationType, gorpMigrationMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(gorpMigrationType, gorpMigrationMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `gorp_migrations` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `gorp_migrations` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `gorp_migrations` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, gorpMigrationPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	_, err = exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "model: unable to insert into gorp_migrations")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "model: unable to populate default values for gorp_migrations")
	}

CacheNoHooks:
	if !cached {
		gorpMigrationInsertCacheMut.Lock()
		gorpMigrationInsertCache[key] = cache
		gorpMigrationInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single GorpMigration record using the global executor.
// See Update for more documentation.
func (o *GorpMigration) UpdateG(columns boil.Columns) error {
	return o.Update(boil.GetDB(), columns)
}

// UpdateP uses an executor to update the GorpMigration, and panics on error.
// See Update for more documentation.
func (o *GorpMigration) UpdateP(exec boil.Executor, columns boil.Columns) {
	err := o.Update(exec, columns)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateGP a single GorpMigration record using the global executor. Panics on error.
// See Update for more documentation.
func (o *GorpMigration) UpdateGP(columns boil.Columns) {
	err := o.Update(boil.GetDB(), columns)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the GorpMigration.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *GorpMigration) Update(exec boil.Executor, columns boil.Columns) error {
	var err error
	key := makeCacheKey(columns, nil)
	gorpMigrationUpdateCacheMut.RLock()
	cache, cached := gorpMigrationUpdateCache[key]
	gorpMigrationUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			gorpMigrationAllColumns,
			gorpMigrationPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return errors.New("model: unable to update gorp_migrations, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `gorp_migrations` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, gorpMigrationPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(gorpMigrationType, gorpMigrationMapping, append(wl, gorpMigrationPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "model: unable to update gorp_migrations row")
	}

	if !cached {
		gorpMigrationUpdateCacheMut.Lock()
		gorpMigrationUpdateCache[key] = cache
		gorpMigrationUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q gorpMigrationQuery) UpdateAllP(exec boil.Executor, cols M) {
	err := q.UpdateAll(exec, cols)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllG updates all rows with the specified column values.
func (q gorpMigrationQuery) UpdateAllG(cols M) error {
	return q.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q gorpMigrationQuery) UpdateAll(exec boil.Executor, cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec(exec)
	if err != nil {
		return errors.Wrap(err, "model: unable to update all for gorp_migrations")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o GorpMigrationSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o GorpMigrationSlice) UpdateAllGP(cols M) {
	err := o.UpdateAll(boil.GetDB(), cols)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o GorpMigrationSlice) UpdateAllP(exec boil.Executor, cols M) {
	err := o.UpdateAll(exec, cols)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o GorpMigrationSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("model: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), gorpMigrationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `gorp_migrations` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, gorpMigrationPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "model: unable to update all in gorpMigration slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *GorpMigration) UpsertG(updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(boil.GetDB(), updateColumns, insertColumns)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *GorpMigration) UpsertGP(updateColumns, insertColumns boil.Columns) {
	if err := o.Upsert(boil.GetDB(), updateColumns, insertColumns); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *GorpMigration) UpsertP(exec boil.Executor, updateColumns, insertColumns boil.Columns) {
	if err := o.Upsert(exec, updateColumns, insertColumns); err != nil {
		panic(boil.WrapErr(err))
	}
}

var mySQLGorpMigrationUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *GorpMigration) Upsert(exec boil.Executor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("model: no gorp_migrations provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(gorpMigrationColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLGorpMigrationUniqueColumns, o)

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

	gorpMigrationUpsertCacheMut.RLock()
	cache, cached := gorpMigrationUpsertCache[key]
	gorpMigrationUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			gorpMigrationAllColumns,
			gorpMigrationColumnsWithDefault,
			gorpMigrationColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			gorpMigrationAllColumns,
			gorpMigrationPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("model: unable to upsert gorp_migrations, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "gorp_migrations", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `gorp_migrations` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(gorpMigrationType, gorpMigrationMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(gorpMigrationType, gorpMigrationMapping, ret)
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

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	_, err = exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "model: unable to upsert for gorp_migrations")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(gorpMigrationType, gorpMigrationMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "model: unable to retrieve unique values for gorp_migrations")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRow(cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "model: unable to populate default values for gorp_migrations")
	}

CacheNoHooks:
	if !cached {
		gorpMigrationUpsertCacheMut.Lock()
		gorpMigrationUpsertCache[key] = cache
		gorpMigrationUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single GorpMigration record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *GorpMigration) DeleteG() error {
	return o.Delete(boil.GetDB())
}

// DeleteP deletes a single GorpMigration record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *GorpMigration) DeleteP(exec boil.Executor) {
	err := o.Delete(exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteGP deletes a single GorpMigration record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *GorpMigration) DeleteGP() {
	err := o.Delete(boil.GetDB())
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single GorpMigration record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *GorpMigration) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("model: no GorpMigration provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), gorpMigrationPrimaryKeyMapping)
	sql := "DELETE FROM `gorp_migrations` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "model: unable to delete from gorp_migrations")
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q gorpMigrationQuery) DeleteAllP(exec boil.Executor) {
	err := q.DeleteAll(exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q gorpMigrationQuery) DeleteAll(exec boil.Executor) error {
	if q.Query == nil {
		return errors.New("model: no gorpMigrationQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec(exec)
	if err != nil {
		return errors.Wrap(err, "model: unable to delete all from gorp_migrations")
	}

	return nil
}

// DeleteAllG deletes all rows in the slice.
func (o GorpMigrationSlice) DeleteAllG() error {
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o GorpMigrationSlice) DeleteAllP(exec boil.Executor) {
	err := o.DeleteAll(exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o GorpMigrationSlice) DeleteAllGP() {
	err := o.DeleteAll(boil.GetDB())
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o GorpMigrationSlice) DeleteAll(exec boil.Executor) error {
	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), gorpMigrationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `gorp_migrations` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, gorpMigrationPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "model: unable to delete all from gorpMigration slice")
	}

	return nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *GorpMigration) ReloadG() error {
	if o == nil {
		return errors.New("model: no GorpMigration provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *GorpMigration) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadGP refetches the object from the database and panics on error.
func (o *GorpMigration) ReloadGP() {
	if err := o.Reload(boil.GetDB()); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *GorpMigration) Reload(exec boil.Executor) error {
	ret, err := FindGorpMigration(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *GorpMigrationSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("model: empty GorpMigrationSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *GorpMigrationSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *GorpMigrationSlice) ReloadAllGP() {
	if err := o.ReloadAll(boil.GetDB()); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *GorpMigrationSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := GorpMigrationSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), gorpMigrationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `gorp_migrations`.* FROM `gorp_migrations` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, gorpMigrationPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "model: unable to reload all in GorpMigrationSlice")
	}

	*o = slice

	return nil
}

// GorpMigrationExistsG checks if the GorpMigration row exists.
func GorpMigrationExistsG(iD string) (bool, error) {
	return GorpMigrationExists(boil.GetDB(), iD)
}

// GorpMigrationExistsP checks if the GorpMigration row exists. Panics on error.
func GorpMigrationExistsP(exec boil.Executor, iD string) bool {
	e, err := GorpMigrationExists(exec, iD)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// GorpMigrationExistsGP checks if the GorpMigration row exists. Panics on error.
func GorpMigrationExistsGP(iD string) bool {
	e, err := GorpMigrationExists(boil.GetDB(), iD)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// GorpMigrationExists checks if the GorpMigration row exists.
func GorpMigrationExists(exec boil.Executor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `gorp_migrations` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "model: unable to check if gorp_migrations exists")
	}

	return exists, nil
}
