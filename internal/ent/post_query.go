// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Southclaws/storyden/internal/ent/account"
	"github.com/Southclaws/storyden/internal/ent/category"
	"github.com/Southclaws/storyden/internal/ent/post"
	"github.com/Southclaws/storyden/internal/ent/predicate"
	"github.com/Southclaws/storyden/internal/ent/react"
	"github.com/Southclaws/storyden/internal/ent/tag"
	"github.com/rs/xid"
)

// PostQuery is the builder for querying Post entities.
type PostQuery struct {
	config
	ctx          *QueryContext
	order        []OrderFunc
	inters       []Interceptor
	predicates   []predicate.Post
	withAuthor   *AccountQuery
	withCategory *CategoryQuery
	withTags     *TagQuery
	withRoot     *PostQuery
	withPosts    *PostQuery
	withReplyTo  *PostQuery
	withReplies  *PostQuery
	withReacts   *ReactQuery
	withFKs      bool
	modifiers    []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PostQuery builder.
func (pq *PostQuery) Where(ps ...predicate.Post) *PostQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit the number of records to be returned by this query.
func (pq *PostQuery) Limit(limit int) *PostQuery {
	pq.ctx.Limit = &limit
	return pq
}

// Offset to start from.
func (pq *PostQuery) Offset(offset int) *PostQuery {
	pq.ctx.Offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *PostQuery) Unique(unique bool) *PostQuery {
	pq.ctx.Unique = &unique
	return pq
}

// Order specifies how the records should be ordered.
func (pq *PostQuery) Order(o ...OrderFunc) *PostQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryAuthor chains the current query on the "author" edge.
func (pq *PostQuery) QueryAuthor() *AccountQuery {
	query := (&AccountClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(post.Table, post.FieldID, selector),
			sqlgraph.To(account.Table, account.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, post.AuthorTable, post.AuthorColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCategory chains the current query on the "category" edge.
func (pq *PostQuery) QueryCategory() *CategoryQuery {
	query := (&CategoryClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(post.Table, post.FieldID, selector),
			sqlgraph.To(category.Table, category.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, post.CategoryTable, post.CategoryColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTags chains the current query on the "tags" edge.
func (pq *PostQuery) QueryTags() *TagQuery {
	query := (&TagClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(post.Table, post.FieldID, selector),
			sqlgraph.To(tag.Table, tag.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, post.TagsTable, post.TagsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRoot chains the current query on the "root" edge.
func (pq *PostQuery) QueryRoot() *PostQuery {
	query := (&PostClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(post.Table, post.FieldID, selector),
			sqlgraph.To(post.Table, post.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, post.RootTable, post.RootColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPosts chains the current query on the "posts" edge.
func (pq *PostQuery) QueryPosts() *PostQuery {
	query := (&PostClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(post.Table, post.FieldID, selector),
			sqlgraph.To(post.Table, post.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, post.PostsTable, post.PostsColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryReplyTo chains the current query on the "replyTo" edge.
func (pq *PostQuery) QueryReplyTo() *PostQuery {
	query := (&PostClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(post.Table, post.FieldID, selector),
			sqlgraph.To(post.Table, post.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, post.ReplyToTable, post.ReplyToColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryReplies chains the current query on the "replies" edge.
func (pq *PostQuery) QueryReplies() *PostQuery {
	query := (&PostClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(post.Table, post.FieldID, selector),
			sqlgraph.To(post.Table, post.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, post.RepliesTable, post.RepliesColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryReacts chains the current query on the "reacts" edge.
func (pq *PostQuery) QueryReacts() *ReactQuery {
	query := (&ReactClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(post.Table, post.FieldID, selector),
			sqlgraph.To(react.Table, react.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, post.ReactsTable, post.ReactsColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Post entity from the query.
// Returns a *NotFoundError when no Post was found.
func (pq *PostQuery) First(ctx context.Context) (*Post, error) {
	nodes, err := pq.Limit(1).All(setContextOp(ctx, pq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{post.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PostQuery) FirstX(ctx context.Context) *Post {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Post ID from the query.
// Returns a *NotFoundError when no Post ID was found.
func (pq *PostQuery) FirstID(ctx context.Context) (id xid.ID, err error) {
	var ids []xid.ID
	if ids, err = pq.Limit(1).IDs(setContextOp(ctx, pq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{post.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PostQuery) FirstIDX(ctx context.Context) xid.ID {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Post entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Post entity is found.
// Returns a *NotFoundError when no Post entities are found.
func (pq *PostQuery) Only(ctx context.Context) (*Post, error) {
	nodes, err := pq.Limit(2).All(setContextOp(ctx, pq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{post.Label}
	default:
		return nil, &NotSingularError{post.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PostQuery) OnlyX(ctx context.Context) *Post {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Post ID in the query.
// Returns a *NotSingularError when more than one Post ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *PostQuery) OnlyID(ctx context.Context) (id xid.ID, err error) {
	var ids []xid.ID
	if ids, err = pq.Limit(2).IDs(setContextOp(ctx, pq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{post.Label}
	default:
		err = &NotSingularError{post.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PostQuery) OnlyIDX(ctx context.Context) xid.ID {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Posts.
func (pq *PostQuery) All(ctx context.Context) ([]*Post, error) {
	ctx = setContextOp(ctx, pq.ctx, "All")
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Post, *PostQuery]()
	return withInterceptors[[]*Post](ctx, pq, qr, pq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pq *PostQuery) AllX(ctx context.Context) []*Post {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Post IDs.
func (pq *PostQuery) IDs(ctx context.Context) (ids []xid.ID, err error) {
	if pq.ctx.Unique == nil && pq.path != nil {
		pq.Unique(true)
	}
	ctx = setContextOp(ctx, pq.ctx, "IDs")
	if err = pq.Select(post.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PostQuery) IDsX(ctx context.Context) []xid.ID {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PostQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pq.ctx, "Count")
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pq, querierCount[*PostQuery](), pq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PostQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PostQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pq.ctx, "Exist")
	switch _, err := pq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PostQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PostQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PostQuery) Clone() *PostQuery {
	if pq == nil {
		return nil
	}
	return &PostQuery{
		config:       pq.config,
		ctx:          pq.ctx.Clone(),
		order:        append([]OrderFunc{}, pq.order...),
		inters:       append([]Interceptor{}, pq.inters...),
		predicates:   append([]predicate.Post{}, pq.predicates...),
		withAuthor:   pq.withAuthor.Clone(),
		withCategory: pq.withCategory.Clone(),
		withTags:     pq.withTags.Clone(),
		withRoot:     pq.withRoot.Clone(),
		withPosts:    pq.withPosts.Clone(),
		withReplyTo:  pq.withReplyTo.Clone(),
		withReplies:  pq.withReplies.Clone(),
		withReacts:   pq.withReacts.Clone(),
		// clone intermediate query.
		sql:  pq.sql.Clone(),
		path: pq.path,
	}
}

// WithAuthor tells the query-builder to eager-load the nodes that are connected to
// the "author" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PostQuery) WithAuthor(opts ...func(*AccountQuery)) *PostQuery {
	query := (&AccountClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withAuthor = query
	return pq
}

// WithCategory tells the query-builder to eager-load the nodes that are connected to
// the "category" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PostQuery) WithCategory(opts ...func(*CategoryQuery)) *PostQuery {
	query := (&CategoryClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withCategory = query
	return pq
}

// WithTags tells the query-builder to eager-load the nodes that are connected to
// the "tags" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PostQuery) WithTags(opts ...func(*TagQuery)) *PostQuery {
	query := (&TagClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withTags = query
	return pq
}

// WithRoot tells the query-builder to eager-load the nodes that are connected to
// the "root" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PostQuery) WithRoot(opts ...func(*PostQuery)) *PostQuery {
	query := (&PostClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withRoot = query
	return pq
}

// WithPosts tells the query-builder to eager-load the nodes that are connected to
// the "posts" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PostQuery) WithPosts(opts ...func(*PostQuery)) *PostQuery {
	query := (&PostClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withPosts = query
	return pq
}

// WithReplyTo tells the query-builder to eager-load the nodes that are connected to
// the "replyTo" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PostQuery) WithReplyTo(opts ...func(*PostQuery)) *PostQuery {
	query := (&PostClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withReplyTo = query
	return pq
}

// WithReplies tells the query-builder to eager-load the nodes that are connected to
// the "replies" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PostQuery) WithReplies(opts ...func(*PostQuery)) *PostQuery {
	query := (&PostClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withReplies = query
	return pq
}

// WithReacts tells the query-builder to eager-load the nodes that are connected to
// the "reacts" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PostQuery) WithReacts(opts ...func(*ReactQuery)) *PostQuery {
	query := (&ReactClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withReacts = query
	return pq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Post.Query().
//		GroupBy(post.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pq *PostQuery) GroupBy(field string, fields ...string) *PostGroupBy {
	pq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PostGroupBy{build: pq}
	grbuild.flds = &pq.ctx.Fields
	grbuild.label = post.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Post.Query().
//		Select(post.FieldCreatedAt).
//		Scan(ctx, &v)
func (pq *PostQuery) Select(fields ...string) *PostSelect {
	pq.ctx.Fields = append(pq.ctx.Fields, fields...)
	sbuild := &PostSelect{PostQuery: pq}
	sbuild.label = post.Label
	sbuild.flds, sbuild.scan = &pq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PostSelect configured with the given aggregations.
func (pq *PostQuery) Aggregate(fns ...AggregateFunc) *PostSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *PostQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pq); err != nil {
				return err
			}
		}
	}
	for _, f := range pq.ctx.Fields {
		if !post.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *PostQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Post, error) {
	var (
		nodes       = []*Post{}
		withFKs     = pq.withFKs
		_spec       = pq.querySpec()
		loadedTypes = [8]bool{
			pq.withAuthor != nil,
			pq.withCategory != nil,
			pq.withTags != nil,
			pq.withRoot != nil,
			pq.withPosts != nil,
			pq.withReplyTo != nil,
			pq.withReplies != nil,
			pq.withReacts != nil,
		}
	)
	if pq.withAuthor != nil || pq.withCategory != nil || pq.withRoot != nil || pq.withReplyTo != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, post.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Post).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Post{config: pq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(pq.modifiers) > 0 {
		_spec.Modifiers = pq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pq.withAuthor; query != nil {
		if err := pq.loadAuthor(ctx, query, nodes, nil,
			func(n *Post, e *Account) { n.Edges.Author = e }); err != nil {
			return nil, err
		}
	}
	if query := pq.withCategory; query != nil {
		if err := pq.loadCategory(ctx, query, nodes, nil,
			func(n *Post, e *Category) { n.Edges.Category = e }); err != nil {
			return nil, err
		}
	}
	if query := pq.withTags; query != nil {
		if err := pq.loadTags(ctx, query, nodes,
			func(n *Post) { n.Edges.Tags = []*Tag{} },
			func(n *Post, e *Tag) { n.Edges.Tags = append(n.Edges.Tags, e) }); err != nil {
			return nil, err
		}
	}
	if query := pq.withRoot; query != nil {
		if err := pq.loadRoot(ctx, query, nodes, nil,
			func(n *Post, e *Post) { n.Edges.Root = e }); err != nil {
			return nil, err
		}
	}
	if query := pq.withPosts; query != nil {
		if err := pq.loadPosts(ctx, query, nodes,
			func(n *Post) { n.Edges.Posts = []*Post{} },
			func(n *Post, e *Post) { n.Edges.Posts = append(n.Edges.Posts, e) }); err != nil {
			return nil, err
		}
	}
	if query := pq.withReplyTo; query != nil {
		if err := pq.loadReplyTo(ctx, query, nodes, nil,
			func(n *Post, e *Post) { n.Edges.ReplyTo = e }); err != nil {
			return nil, err
		}
	}
	if query := pq.withReplies; query != nil {
		if err := pq.loadReplies(ctx, query, nodes,
			func(n *Post) { n.Edges.Replies = []*Post{} },
			func(n *Post, e *Post) { n.Edges.Replies = append(n.Edges.Replies, e) }); err != nil {
			return nil, err
		}
	}
	if query := pq.withReacts; query != nil {
		if err := pq.loadReacts(ctx, query, nodes,
			func(n *Post) { n.Edges.Reacts = []*React{} },
			func(n *Post, e *React) { n.Edges.Reacts = append(n.Edges.Reacts, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *PostQuery) loadAuthor(ctx context.Context, query *AccountQuery, nodes []*Post, init func(*Post), assign func(*Post, *Account)) error {
	ids := make([]xid.ID, 0, len(nodes))
	nodeids := make(map[xid.ID][]*Post)
	for i := range nodes {
		if nodes[i].account_posts == nil {
			continue
		}
		fk := *nodes[i].account_posts
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(account.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "account_posts" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (pq *PostQuery) loadCategory(ctx context.Context, query *CategoryQuery, nodes []*Post, init func(*Post), assign func(*Post, *Category)) error {
	ids := make([]xid.ID, 0, len(nodes))
	nodeids := make(map[xid.ID][]*Post)
	for i := range nodes {
		fk := nodes[i].CategoryID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(category.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "category_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (pq *PostQuery) loadTags(ctx context.Context, query *TagQuery, nodes []*Post, init func(*Post), assign func(*Post, *Tag)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[xid.ID]*Post)
	nids := make(map[xid.ID]map[*Post]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(post.TagsTable)
		s.Join(joinT).On(s.C(tag.FieldID), joinT.C(post.TagsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(post.TagsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(post.TagsPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(xid.ID)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := *values[0].(*xid.ID)
				inValue := *values[1].(*xid.ID)
				if nids[inValue] == nil {
					nids[inValue] = map[*Post]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Tag](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "tags" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (pq *PostQuery) loadRoot(ctx context.Context, query *PostQuery, nodes []*Post, init func(*Post), assign func(*Post, *Post)) error {
	ids := make([]xid.ID, 0, len(nodes))
	nodeids := make(map[xid.ID][]*Post)
	for i := range nodes {
		fk := nodes[i].RootPostID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(post.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "root_post_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (pq *PostQuery) loadPosts(ctx context.Context, query *PostQuery, nodes []*Post, init func(*Post), assign func(*Post, *Post)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[xid.ID]*Post)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Post(func(s *sql.Selector) {
		s.Where(sql.InValues(post.PostsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.RootPostID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "root_post_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (pq *PostQuery) loadReplyTo(ctx context.Context, query *PostQuery, nodes []*Post, init func(*Post), assign func(*Post, *Post)) error {
	ids := make([]xid.ID, 0, len(nodes))
	nodeids := make(map[xid.ID][]*Post)
	for i := range nodes {
		fk := nodes[i].ReplyToPostID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(post.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "reply_to_post_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (pq *PostQuery) loadReplies(ctx context.Context, query *PostQuery, nodes []*Post, init func(*Post), assign func(*Post, *Post)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[xid.ID]*Post)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Post(func(s *sql.Selector) {
		s.Where(sql.InValues(post.RepliesColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ReplyToPostID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "reply_to_post_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (pq *PostQuery) loadReacts(ctx context.Context, query *ReactQuery, nodes []*Post, init func(*Post), assign func(*Post, *React)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[xid.ID]*Post)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.React(func(s *sql.Selector) {
		s.Where(sql.InValues(post.ReactsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.post_reacts
		if fk == nil {
			return fmt.Errorf(`foreign-key "post_reacts" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "post_reacts" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (pq *PostQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	if len(pq.modifiers) > 0 {
		_spec.Modifiers = pq.modifiers
	}
	_spec.Node.Columns = pq.ctx.Fields
	if len(pq.ctx.Fields) > 0 {
		_spec.Unique = pq.ctx.Unique != nil && *pq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PostQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(post.Table, post.Columns, sqlgraph.NewFieldSpec(post.FieldID, field.TypeString))
	_spec.From = pq.sql
	if unique := pq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pq.path != nil {
		_spec.Unique = true
	}
	if fields := pq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, post.FieldID)
		for i := range fields {
			if fields[i] != post.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *PostQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(post.Table)
	columns := pq.ctx.Fields
	if len(columns) == 0 {
		columns = post.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.ctx.Unique != nil && *pq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range pq.modifiers {
		m(selector)
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (pq *PostQuery) Modify(modifiers ...func(s *sql.Selector)) *PostSelect {
	pq.modifiers = append(pq.modifiers, modifiers...)
	return pq.Select()
}

// PostGroupBy is the group-by builder for Post entities.
type PostGroupBy struct {
	selector
	build *PostQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PostGroupBy) Aggregate(fns ...AggregateFunc) *PostGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *PostGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pgb.build.ctx, "GroupBy")
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PostQuery, *PostGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *PostGroupBy) sqlScan(ctx context.Context, root *PostQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pgb.flds)+len(pgb.fns))
		for _, f := range *pgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PostSelect is the builder for selecting fields of Post entities.
type PostSelect struct {
	*PostQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *PostSelect) Aggregate(fns ...AggregateFunc) *PostSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PostSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ps.ctx, "Select")
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PostQuery, *PostSelect](ctx, ps.PostQuery, ps, ps.inters, v)
}

func (ps *PostSelect) sqlScan(ctx context.Context, root *PostQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ps.fns))
	for _, fn := range ps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ps *PostSelect) Modify(modifiers ...func(s *sql.Selector)) *PostSelect {
	ps.modifiers = append(ps.modifiers, modifiers...)
	return ps
}
