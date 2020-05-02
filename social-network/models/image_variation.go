// Code generated by SQLBoiler 4.0.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// ImageVariation is an object representing the database table.
type ImageVariation struct {
	ID      uint `boil:"id" json:"id" toml:"id" yaml:"id"`
	ImageID uint `boil:"image_id" json:"image_id" toml:"image_id" yaml:"image_id"`

	R *imageVariationR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L imageVariationL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ImageVariationColumns = struct {
	ID      string
	ImageID string
}{
	ID:      "id",
	ImageID: "image_id",
}

// Generated where

var ImageVariationWhere = struct {
	ID      whereHelperuint
	ImageID whereHelperuint
}{
	ID:      whereHelperuint{field: "`image_variation`.`id`"},
	ImageID: whereHelperuint{field: "`image_variation`.`image_id`"},
}

// ImageVariationRels is where relationship names are stored.
var ImageVariationRels = struct {
	Image string
}{
	Image: "Image",
}

// imageVariationR is where relationships are stored.
type imageVariationR struct {
	Image *Image
}

// NewStruct creates a new relationship struct
func (*imageVariationR) NewStruct() *imageVariationR {
	return &imageVariationR{}
}

// imageVariationL is where Load methods for each relationship are stored.
type imageVariationL struct{}

var (
	imageVariationAllColumns            = []string{"id", "image_id"}
	imageVariationColumnsWithoutDefault = []string{"image_id"}
	imageVariationColumnsWithDefault    = []string{"id"}
	imageVariationPrimaryKeyColumns     = []string{"id"}
)

type (
	// ImageVariationSlice is an alias for a slice of pointers to ImageVariation.
	// This should generally be used opposed to []ImageVariation.
	ImageVariationSlice []*ImageVariation
	// ImageVariationHook is the signature for custom ImageVariation hook methods
	ImageVariationHook func(context.Context, boil.ContextExecutor, *ImageVariation) error

	imageVariationQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	imageVariationType                 = reflect.TypeOf(&ImageVariation{})
	imageVariationMapping              = queries.MakeStructMapping(imageVariationType)
	imageVariationPrimaryKeyMapping, _ = queries.BindMapping(imageVariationType, imageVariationMapping, imageVariationPrimaryKeyColumns)
	imageVariationInsertCacheMut       sync.RWMutex
	imageVariationInsertCache          = make(map[string]insertCache)
	imageVariationUpdateCacheMut       sync.RWMutex
	imageVariationUpdateCache          = make(map[string]updateCache)
	imageVariationUpsertCacheMut       sync.RWMutex
	imageVariationUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var imageVariationBeforeInsertHooks []ImageVariationHook
var imageVariationBeforeUpdateHooks []ImageVariationHook
var imageVariationBeforeDeleteHooks []ImageVariationHook
var imageVariationBeforeUpsertHooks []ImageVariationHook

var imageVariationAfterInsertHooks []ImageVariationHook
var imageVariationAfterSelectHooks []ImageVariationHook
var imageVariationAfterUpdateHooks []ImageVariationHook
var imageVariationAfterDeleteHooks []ImageVariationHook
var imageVariationAfterUpsertHooks []ImageVariationHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ImageVariation) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range imageVariationBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ImageVariation) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range imageVariationBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ImageVariation) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range imageVariationBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ImageVariation) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range imageVariationBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ImageVariation) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range imageVariationAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ImageVariation) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range imageVariationAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ImageVariation) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range imageVariationAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ImageVariation) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range imageVariationAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ImageVariation) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range imageVariationAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddImageVariationHook registers your hook function for all future operations.
func AddImageVariationHook(hookPoint boil.HookPoint, imageVariationHook ImageVariationHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		imageVariationBeforeInsertHooks = append(imageVariationBeforeInsertHooks, imageVariationHook)
	case boil.BeforeUpdateHook:
		imageVariationBeforeUpdateHooks = append(imageVariationBeforeUpdateHooks, imageVariationHook)
	case boil.BeforeDeleteHook:
		imageVariationBeforeDeleteHooks = append(imageVariationBeforeDeleteHooks, imageVariationHook)
	case boil.BeforeUpsertHook:
		imageVariationBeforeUpsertHooks = append(imageVariationBeforeUpsertHooks, imageVariationHook)
	case boil.AfterInsertHook:
		imageVariationAfterInsertHooks = append(imageVariationAfterInsertHooks, imageVariationHook)
	case boil.AfterSelectHook:
		imageVariationAfterSelectHooks = append(imageVariationAfterSelectHooks, imageVariationHook)
	case boil.AfterUpdateHook:
		imageVariationAfterUpdateHooks = append(imageVariationAfterUpdateHooks, imageVariationHook)
	case boil.AfterDeleteHook:
		imageVariationAfterDeleteHooks = append(imageVariationAfterDeleteHooks, imageVariationHook)
	case boil.AfterUpsertHook:
		imageVariationAfterUpsertHooks = append(imageVariationAfterUpsertHooks, imageVariationHook)
	}
}

// One returns a single imageVariation record from the query.
func (q imageVariationQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ImageVariation, error) {
	o := &ImageVariation{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for image_variation")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ImageVariation records from the query.
func (q imageVariationQuery) All(ctx context.Context, exec boil.ContextExecutor) (ImageVariationSlice, error) {
	var o []*ImageVariation

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ImageVariation slice")
	}

	if len(imageVariationAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ImageVariation records in the query.
func (q imageVariationQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count image_variation rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q imageVariationQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if image_variation exists")
	}

	return count > 0, nil
}

// Image pointed to by the foreign key.
func (o *ImageVariation) Image(mods ...qm.QueryMod) imageQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.ImageID),
	}

	queryMods = append(queryMods, mods...)

	query := Images(queryMods...)
	queries.SetFrom(query.Query, "`image`")

	return query
}

// LoadImage allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (imageVariationL) LoadImage(ctx context.Context, e boil.ContextExecutor, singular bool, maybeImageVariation interface{}, mods queries.Applicator) error {
	var slice []*ImageVariation
	var object *ImageVariation

	if singular {
		object = maybeImageVariation.(*ImageVariation)
	} else {
		slice = *maybeImageVariation.(*[]*ImageVariation)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &imageVariationR{}
		}
		args = append(args, object.ImageID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &imageVariationR{}
			}

			for _, a := range args {
				if a == obj.ImageID {
					continue Outer
				}
			}

			args = append(args, obj.ImageID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`image`),
		qm.WhereIn(`image.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Image")
	}

	var resultSlice []*Image
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Image")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for image")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for image")
	}

	if len(imageVariationAfterSelectHooks) != 0 {
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
		object.R.Image = foreign
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ImageID == foreign.ID {
				local.R.Image = foreign
				break
			}
		}
	}

	return nil
}

// SetImage of the imageVariation to the related item.
// Sets o.R.Image to related.
// Adds o to related.R.ImageVariations.
func (o *ImageVariation) SetImage(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Image) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `image_variation` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"image_id"}),
		strmangle.WhereClause("`", "`", 0, imageVariationPrimaryKeyColumns),
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

	o.ImageID = related.ID
	if o.R == nil {
		o.R = &imageVariationR{
			Image: related,
		}
	} else {
		o.R.Image = related
	}

	if related.R == nil {
		related.R = &imageR{
			ImageVariations: ImageVariationSlice{o},
		}
	} else {
		related.R.ImageVariations = append(related.R.ImageVariations, o)
	}

	return nil
}

// ImageVariations retrieves all the records using an executor.
func ImageVariations(mods ...qm.QueryMod) imageVariationQuery {
	mods = append(mods, qm.From("`image_variation`"))
	return imageVariationQuery{NewQuery(mods...)}
}

// FindImageVariation retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindImageVariation(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*ImageVariation, error) {
	imageVariationObj := &ImageVariation{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `image_variation` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, imageVariationObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from image_variation")
	}

	return imageVariationObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ImageVariation) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no image_variation provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(imageVariationColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	imageVariationInsertCacheMut.RLock()
	cache, cached := imageVariationInsertCache[key]
	imageVariationInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			imageVariationAllColumns,
			imageVariationColumnsWithDefault,
			imageVariationColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(imageVariationType, imageVariationMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(imageVariationType, imageVariationMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `image_variation` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `image_variation` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `image_variation` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, imageVariationPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into image_variation")
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

	o.ID = uint(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == imageVariationMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for image_variation")
	}

CacheNoHooks:
	if !cached {
		imageVariationInsertCacheMut.Lock()
		imageVariationInsertCache[key] = cache
		imageVariationInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ImageVariation.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ImageVariation) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	imageVariationUpdateCacheMut.RLock()
	cache, cached := imageVariationUpdateCache[key]
	imageVariationUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			imageVariationAllColumns,
			imageVariationPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update image_variation, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `image_variation` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, imageVariationPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(imageVariationType, imageVariationMapping, append(wl, imageVariationPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update image_variation row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for image_variation")
	}

	if !cached {
		imageVariationUpdateCacheMut.Lock()
		imageVariationUpdateCache[key] = cache
		imageVariationUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q imageVariationQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for image_variation")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for image_variation")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ImageVariationSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), imageVariationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `image_variation` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, imageVariationPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in imageVariation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all imageVariation")
	}
	return rowsAff, nil
}

var mySQLImageVariationUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ImageVariation) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no image_variation provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(imageVariationColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLImageVariationUniqueColumns, o)

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

	imageVariationUpsertCacheMut.RLock()
	cache, cached := imageVariationUpsertCache[key]
	imageVariationUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			imageVariationAllColumns,
			imageVariationColumnsWithDefault,
			imageVariationColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			imageVariationAllColumns,
			imageVariationPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert image_variation, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "image_variation", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `image_variation` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(imageVariationType, imageVariationMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(imageVariationType, imageVariationMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for image_variation")
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

	o.ID = uint(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == imageVariationMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(imageVariationType, imageVariationMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for image_variation")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for image_variation")
	}

CacheNoHooks:
	if !cached {
		imageVariationUpsertCacheMut.Lock()
		imageVariationUpsertCache[key] = cache
		imageVariationUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ImageVariation record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ImageVariation) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ImageVariation provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), imageVariationPrimaryKeyMapping)
	sql := "DELETE FROM `image_variation` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from image_variation")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for image_variation")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q imageVariationQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no imageVariationQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from image_variation")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for image_variation")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ImageVariationSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(imageVariationBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), imageVariationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `image_variation` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, imageVariationPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from imageVariation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for image_variation")
	}

	if len(imageVariationAfterDeleteHooks) != 0 {
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
func (o *ImageVariation) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindImageVariation(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ImageVariationSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ImageVariationSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), imageVariationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `image_variation`.* FROM `image_variation` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, imageVariationPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ImageVariationSlice")
	}

	*o = slice

	return nil
}

// ImageVariationExists checks if the ImageVariation row exists.
func ImageVariationExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `image_variation` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if image_variation exists")
	}

	return exists, nil
}
