// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package boilentity

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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Post is an object representing the database table.
type Post struct {
	ID        string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedBy string      `boil:"created_by" json:"created_by" toml:"created_by" yaml:"created_by"`
	Title     string      `boil:"title" json:"title" toml:"title" yaml:"title"`
	Content   null.String `boil:"content" json:"content,omitempty" toml:"content" yaml:"content,omitempty"`

	R *postR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L postL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PostColumns = struct {
	ID        string
	CreatedBy string
	Title     string
	Content   string
}{
	ID:        "id",
	CreatedBy: "created_by",
	Title:     "title",
	Content:   "content",
}

var PostTableColumns = struct {
	ID        string
	CreatedBy string
	Title     string
	Content   string
}{
	ID:        "post.id",
	CreatedBy: "post.created_by",
	Title:     "post.title",
	Content:   "post.content",
}

// Generated where

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelpernull_String) LIKE(x null.String) qm.QueryMod {
	return qm.Where(w.field+" LIKE ?", x)
}
func (w whereHelpernull_String) NLIKE(x null.String) qm.QueryMod {
	return qm.Where(w.field+" NOT LIKE ?", x)
}
func (w whereHelpernull_String) ILIKE(x null.String) qm.QueryMod {
	return qm.Where(w.field+" ILIKE ?", x)
}
func (w whereHelpernull_String) NILIKE(x null.String) qm.QueryMod {
	return qm.Where(w.field+" NOT ILIKE ?", x)
}
func (w whereHelpernull_String) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelpernull_String) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var PostWhere = struct {
	ID        whereHelperstring
	CreatedBy whereHelperstring
	Title     whereHelperstring
	Content   whereHelpernull_String
}{
	ID:        whereHelperstring{field: "\"post\".\"id\""},
	CreatedBy: whereHelperstring{field: "\"post\".\"created_by\""},
	Title:     whereHelperstring{field: "\"post\".\"title\""},
	Content:   whereHelpernull_String{field: "\"post\".\"content\""},
}

// PostRels is where relationship names are stored.
var PostRels = struct {
	CreatedByAccount string
}{
	CreatedByAccount: "CreatedByAccount",
}

// postR is where relationships are stored.
type postR struct {
	CreatedByAccount *Account `boil:"CreatedByAccount" json:"CreatedByAccount" toml:"CreatedByAccount" yaml:"CreatedByAccount"`
}

// NewStruct creates a new relationship struct
func (*postR) NewStruct() *postR {
	return &postR{}
}

func (r *postR) GetCreatedByAccount() *Account {
	if r == nil {
		return nil
	}
	return r.CreatedByAccount
}

// postL is where Load methods for each relationship are stored.
type postL struct{}

var (
	postAllColumns            = []string{"id", "created_by", "title", "content"}
	postColumnsWithoutDefault = []string{"created_by", "title"}
	postColumnsWithDefault    = []string{"id", "content"}
	postPrimaryKeyColumns     = []string{"id"}
	postGeneratedColumns      = []string{}
)

type (
	// PostSlice is an alias for a slice of pointers to Post.
	// This should almost always be used instead of []Post.
	PostSlice []*Post
	// PostHook is the signature for custom Post hook methods
	PostHook func(context.Context, boil.ContextExecutor, *Post) error

	postQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	postType                 = reflect.TypeOf(&Post{})
	postMapping              = queries.MakeStructMapping(postType)
	postPrimaryKeyMapping, _ = queries.BindMapping(postType, postMapping, postPrimaryKeyColumns)
	postInsertCacheMut       sync.RWMutex
	postInsertCache          = make(map[string]insertCache)
	postUpdateCacheMut       sync.RWMutex
	postUpdateCache          = make(map[string]updateCache)
	postUpsertCacheMut       sync.RWMutex
	postUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var postAfterSelectMu sync.Mutex
var postAfterSelectHooks []PostHook

var postBeforeInsertMu sync.Mutex
var postBeforeInsertHooks []PostHook
var postAfterInsertMu sync.Mutex
var postAfterInsertHooks []PostHook

var postBeforeUpdateMu sync.Mutex
var postBeforeUpdateHooks []PostHook
var postAfterUpdateMu sync.Mutex
var postAfterUpdateHooks []PostHook

var postBeforeDeleteMu sync.Mutex
var postBeforeDeleteHooks []PostHook
var postAfterDeleteMu sync.Mutex
var postAfterDeleteHooks []PostHook

var postBeforeUpsertMu sync.Mutex
var postBeforeUpsertHooks []PostHook
var postAfterUpsertMu sync.Mutex
var postAfterUpsertHooks []PostHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Post) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range postAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Post) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range postBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Post) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range postAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Post) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range postBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Post) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range postAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Post) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range postBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Post) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range postAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Post) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range postBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Post) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range postAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPostHook registers your hook function for all future operations.
func AddPostHook(hookPoint boil.HookPoint, postHook PostHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		postAfterSelectMu.Lock()
		postAfterSelectHooks = append(postAfterSelectHooks, postHook)
		postAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		postBeforeInsertMu.Lock()
		postBeforeInsertHooks = append(postBeforeInsertHooks, postHook)
		postBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		postAfterInsertMu.Lock()
		postAfterInsertHooks = append(postAfterInsertHooks, postHook)
		postAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		postBeforeUpdateMu.Lock()
		postBeforeUpdateHooks = append(postBeforeUpdateHooks, postHook)
		postBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		postAfterUpdateMu.Lock()
		postAfterUpdateHooks = append(postAfterUpdateHooks, postHook)
		postAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		postBeforeDeleteMu.Lock()
		postBeforeDeleteHooks = append(postBeforeDeleteHooks, postHook)
		postBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		postAfterDeleteMu.Lock()
		postAfterDeleteHooks = append(postAfterDeleteHooks, postHook)
		postAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		postBeforeUpsertMu.Lock()
		postBeforeUpsertHooks = append(postBeforeUpsertHooks, postHook)
		postBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		postAfterUpsertMu.Lock()
		postAfterUpsertHooks = append(postAfterUpsertHooks, postHook)
		postAfterUpsertMu.Unlock()
	}
}

// One returns a single post record from the query.
func (q postQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Post, error) {
	o := &Post{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "boilentity: failed to execute a one query for post")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Post records from the query.
func (q postQuery) All(ctx context.Context, exec boil.ContextExecutor) (PostSlice, error) {
	var o []*Post

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "boilentity: failed to assign all query results to Post slice")
	}

	if len(postAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Post records in the query.
func (q postQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "boilentity: failed to count post rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q postQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "boilentity: failed to check if post exists")
	}

	return count > 0, nil
}

// CreatedByAccount pointed to by the foreign key.
func (o *Post) CreatedByAccount(mods ...qm.QueryMod) accountQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.CreatedBy),
	}

	queryMods = append(queryMods, mods...)

	return Accounts(queryMods...)
}

// LoadCreatedByAccount allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (postL) LoadCreatedByAccount(ctx context.Context, e boil.ContextExecutor, singular bool, maybePost interface{}, mods queries.Applicator) error {
	var slice []*Post
	var object *Post

	if singular {
		var ok bool
		object, ok = maybePost.(*Post)
		if !ok {
			object = new(Post)
			ok = queries.SetFromEmbeddedStruct(&object, &maybePost)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybePost))
			}
		}
	} else {
		s, ok := maybePost.(*[]*Post)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybePost)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybePost))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &postR{}
		}
		args[object.CreatedBy] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &postR{}
			}

			args[obj.CreatedBy] = struct{}{}

		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`account`),
		qm.WhereIn(`account.id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Account")
	}

	var resultSlice []*Account
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Account")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for account")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for account")
	}

	if len(accountAfterSelectHooks) != 0 {
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
		object.R.CreatedByAccount = foreign
		if foreign.R == nil {
			foreign.R = &accountR{}
		}
		foreign.R.CreatedByPosts = append(foreign.R.CreatedByPosts, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.CreatedBy == foreign.ID {
				local.R.CreatedByAccount = foreign
				if foreign.R == nil {
					foreign.R = &accountR{}
				}
				foreign.R.CreatedByPosts = append(foreign.R.CreatedByPosts, local)
				break
			}
		}
	}

	return nil
}

// SetCreatedByAccount of the post to the related item.
// Sets o.R.CreatedByAccount to related.
// Adds o to related.R.CreatedByPosts.
func (o *Post) SetCreatedByAccount(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Account) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"post\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"created_by"}),
		strmangle.WhereClause("\"", "\"", 2, postPrimaryKeyColumns),
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

	o.CreatedBy = related.ID
	if o.R == nil {
		o.R = &postR{
			CreatedByAccount: related,
		}
	} else {
		o.R.CreatedByAccount = related
	}

	if related.R == nil {
		related.R = &accountR{
			CreatedByPosts: PostSlice{o},
		}
	} else {
		related.R.CreatedByPosts = append(related.R.CreatedByPosts, o)
	}

	return nil
}

// Posts retrieves all the records using an executor.
func Posts(mods ...qm.QueryMod) postQuery {
	mods = append(mods, qm.From("\"post\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"post\".*"})
	}

	return postQuery{q}
}

// FindPost retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPost(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Post, error) {
	postObj := &Post{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"post\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, postObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "boilentity: unable to select from post")
	}

	if err = postObj.doAfterSelectHooks(ctx, exec); err != nil {
		return postObj, err
	}

	return postObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Post) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("boilentity: no post provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(postColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	postInsertCacheMut.RLock()
	cache, cached := postInsertCache[key]
	postInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			postAllColumns,
			postColumnsWithDefault,
			postColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(postType, postMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(postType, postMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"post\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"post\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
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

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "boilentity: unable to insert into post")
	}

	if !cached {
		postInsertCacheMut.Lock()
		postInsertCache[key] = cache
		postInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Post.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Post) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	postUpdateCacheMut.RLock()
	cache, cached := postUpdateCache[key]
	postUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			postAllColumns,
			postPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("boilentity: unable to update post, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"post\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, postPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(postType, postMapping, append(wl, postPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "boilentity: unable to update post row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boilentity: failed to get rows affected by update for post")
	}

	if !cached {
		postUpdateCacheMut.Lock()
		postUpdateCache[key] = cache
		postUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q postQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "boilentity: unable to update all for post")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boilentity: unable to retrieve rows affected for post")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PostSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("boilentity: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), postPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"post\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, postPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boilentity: unable to update all in post slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boilentity: unable to retrieve rows affected all in update all post")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Post) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("boilentity: no post provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(postColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
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
	key := buf.String()
	strmangle.PutBuffer(buf)

	postUpsertCacheMut.RLock()
	cache, cached := postUpsertCache[key]
	postUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			postAllColumns,
			postColumnsWithDefault,
			postColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			postAllColumns,
			postPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("boilentity: unable to upsert post, could not build update column list")
		}

		ret := strmangle.SetComplement(postAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(postPrimaryKeyColumns) == 0 {
				return errors.New("boilentity: unable to upsert post, could not build conflict column list")
			}

			conflict = make([]string, len(postPrimaryKeyColumns))
			copy(conflict, postPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"post\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(postType, postMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(postType, postMapping, ret)
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
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "boilentity: unable to upsert post")
	}

	if !cached {
		postUpsertCacheMut.Lock()
		postUpsertCache[key] = cache
		postUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Post record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Post) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("boilentity: no Post provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), postPrimaryKeyMapping)
	sql := "DELETE FROM \"post\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boilentity: unable to delete from post")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boilentity: failed to get rows affected by delete for post")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q postQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("boilentity: no postQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "boilentity: unable to delete all from post")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boilentity: failed to get rows affected by deleteall for post")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PostSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(postBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), postPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"post\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, postPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boilentity: unable to delete all from post slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boilentity: failed to get rows affected by deleteall for post")
	}

	if len(postAfterDeleteHooks) != 0 {
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
func (o *Post) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPost(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PostSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PostSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), postPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"post\".* FROM \"post\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, postPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "boilentity: unable to reload all in PostSlice")
	}

	*o = slice

	return nil
}

// PostExists checks if the Post row exists.
func PostExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"post\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "boilentity: unable to check if post exists")
	}

	return exists, nil
}

// Exists checks if the Post row exists.
func (o *Post) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return PostExists(ctx, exec, o.ID)
}
