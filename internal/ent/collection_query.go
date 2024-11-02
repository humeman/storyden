// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Southclaws/storyden/internal/ent/account"
	"github.com/Southclaws/storyden/internal/ent/asset"
	"github.com/Southclaws/storyden/internal/ent/collection"
	"github.com/Southclaws/storyden/internal/ent/collectionnode"
	"github.com/Southclaws/storyden/internal/ent/collectionpost"
	"github.com/Southclaws/storyden/internal/ent/node"
	"github.com/Southclaws/storyden/internal/ent/post"
	"github.com/Southclaws/storyden/internal/ent/predicate"
	"github.com/rs/xid"
)

// CollectionQuery is the builder for querying Collection entities.
type CollectionQuery struct {
	config
	ctx                 *QueryContext
	order               []collection.OrderOption
	inters              []Interceptor
	predicates          []predicate.Collection
	withOwner           *AccountQuery
	withCoverImage      *AssetQuery
	withPosts           *PostQuery
	withNodes           *NodeQuery
	withCollectionPosts *CollectionPostQuery
	withCollectionNodes *CollectionNodeQuery
	withFKs             bool
	modifiers           []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CollectionQuery builder.
func (cq *CollectionQuery) Where(ps ...predicate.Collection) *CollectionQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *CollectionQuery) Limit(limit int) *CollectionQuery {
	cq.ctx.Limit = &limit
	return cq
}

// Offset to start from.
func (cq *CollectionQuery) Offset(offset int) *CollectionQuery {
	cq.ctx.Offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *CollectionQuery) Unique(unique bool) *CollectionQuery {
	cq.ctx.Unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *CollectionQuery) Order(o ...collection.OrderOption) *CollectionQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryOwner chains the current query on the "owner" edge.
func (cq *CollectionQuery) QueryOwner() *AccountQuery {
	query := (&AccountClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(collection.Table, collection.FieldID, selector),
			sqlgraph.To(account.Table, account.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, collection.OwnerTable, collection.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCoverImage chains the current query on the "cover_image" edge.
func (cq *CollectionQuery) QueryCoverImage() *AssetQuery {
	query := (&AssetClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(collection.Table, collection.FieldID, selector),
			sqlgraph.To(asset.Table, asset.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, collection.CoverImageTable, collection.CoverImageColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPosts chains the current query on the "posts" edge.
func (cq *CollectionQuery) QueryPosts() *PostQuery {
	query := (&PostClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(collection.Table, collection.FieldID, selector),
			sqlgraph.To(post.Table, post.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, collection.PostsTable, collection.PostsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNodes chains the current query on the "nodes" edge.
func (cq *CollectionQuery) QueryNodes() *NodeQuery {
	query := (&NodeClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(collection.Table, collection.FieldID, selector),
			sqlgraph.To(node.Table, node.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, collection.NodesTable, collection.NodesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCollectionPosts chains the current query on the "collection_posts" edge.
func (cq *CollectionQuery) QueryCollectionPosts() *CollectionPostQuery {
	query := (&CollectionPostClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(collection.Table, collection.FieldID, selector),
			sqlgraph.To(collectionpost.Table, collectionpost.CollectionColumn),
			sqlgraph.Edge(sqlgraph.O2M, true, collection.CollectionPostsTable, collection.CollectionPostsColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCollectionNodes chains the current query on the "collection_nodes" edge.
func (cq *CollectionQuery) QueryCollectionNodes() *CollectionNodeQuery {
	query := (&CollectionNodeClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(collection.Table, collection.FieldID, selector),
			sqlgraph.To(collectionnode.Table, collectionnode.CollectionColumn),
			sqlgraph.Edge(sqlgraph.O2M, true, collection.CollectionNodesTable, collection.CollectionNodesColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Collection entity from the query.
// Returns a *NotFoundError when no Collection was found.
func (cq *CollectionQuery) First(ctx context.Context) (*Collection, error) {
	nodes, err := cq.Limit(1).All(setContextOp(ctx, cq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{collection.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *CollectionQuery) FirstX(ctx context.Context) *Collection {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Collection ID from the query.
// Returns a *NotFoundError when no Collection ID was found.
func (cq *CollectionQuery) FirstID(ctx context.Context) (id xid.ID, err error) {
	var ids []xid.ID
	if ids, err = cq.Limit(1).IDs(setContextOp(ctx, cq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{collection.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *CollectionQuery) FirstIDX(ctx context.Context) xid.ID {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Collection entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Collection entity is found.
// Returns a *NotFoundError when no Collection entities are found.
func (cq *CollectionQuery) Only(ctx context.Context) (*Collection, error) {
	nodes, err := cq.Limit(2).All(setContextOp(ctx, cq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{collection.Label}
	default:
		return nil, &NotSingularError{collection.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *CollectionQuery) OnlyX(ctx context.Context) *Collection {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Collection ID in the query.
// Returns a *NotSingularError when more than one Collection ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *CollectionQuery) OnlyID(ctx context.Context) (id xid.ID, err error) {
	var ids []xid.ID
	if ids, err = cq.Limit(2).IDs(setContextOp(ctx, cq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{collection.Label}
	default:
		err = &NotSingularError{collection.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *CollectionQuery) OnlyIDX(ctx context.Context) xid.ID {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Collections.
func (cq *CollectionQuery) All(ctx context.Context) ([]*Collection, error) {
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryAll)
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Collection, *CollectionQuery]()
	return withInterceptors[[]*Collection](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *CollectionQuery) AllX(ctx context.Context) []*Collection {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Collection IDs.
func (cq *CollectionQuery) IDs(ctx context.Context) (ids []xid.ID, err error) {
	if cq.ctx.Unique == nil && cq.path != nil {
		cq.Unique(true)
	}
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryIDs)
	if err = cq.Select(collection.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *CollectionQuery) IDsX(ctx context.Context) []xid.ID {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *CollectionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryCount)
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*CollectionQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *CollectionQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *CollectionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryExist)
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *CollectionQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CollectionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *CollectionQuery) Clone() *CollectionQuery {
	if cq == nil {
		return nil
	}
	return &CollectionQuery{
		config:              cq.config,
		ctx:                 cq.ctx.Clone(),
		order:               append([]collection.OrderOption{}, cq.order...),
		inters:              append([]Interceptor{}, cq.inters...),
		predicates:          append([]predicate.Collection{}, cq.predicates...),
		withOwner:           cq.withOwner.Clone(),
		withCoverImage:      cq.withCoverImage.Clone(),
		withPosts:           cq.withPosts.Clone(),
		withNodes:           cq.withNodes.Clone(),
		withCollectionPosts: cq.withCollectionPosts.Clone(),
		withCollectionNodes: cq.withCollectionNodes.Clone(),
		// clone intermediate query.
		sql:       cq.sql.Clone(),
		path:      cq.path,
		modifiers: append([]func(*sql.Selector){}, cq.modifiers...),
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CollectionQuery) WithOwner(opts ...func(*AccountQuery)) *CollectionQuery {
	query := (&AccountClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withOwner = query
	return cq
}

// WithCoverImage tells the query-builder to eager-load the nodes that are connected to
// the "cover_image" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CollectionQuery) WithCoverImage(opts ...func(*AssetQuery)) *CollectionQuery {
	query := (&AssetClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withCoverImage = query
	return cq
}

// WithPosts tells the query-builder to eager-load the nodes that are connected to
// the "posts" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CollectionQuery) WithPosts(opts ...func(*PostQuery)) *CollectionQuery {
	query := (&PostClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withPosts = query
	return cq
}

// WithNodes tells the query-builder to eager-load the nodes that are connected to
// the "nodes" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CollectionQuery) WithNodes(opts ...func(*NodeQuery)) *CollectionQuery {
	query := (&NodeClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withNodes = query
	return cq
}

// WithCollectionPosts tells the query-builder to eager-load the nodes that are connected to
// the "collection_posts" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CollectionQuery) WithCollectionPosts(opts ...func(*CollectionPostQuery)) *CollectionQuery {
	query := (&CollectionPostClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withCollectionPosts = query
	return cq
}

// WithCollectionNodes tells the query-builder to eager-load the nodes that are connected to
// the "collection_nodes" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CollectionQuery) WithCollectionNodes(opts ...func(*CollectionNodeQuery)) *CollectionQuery {
	query := (&CollectionNodeClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withCollectionNodes = query
	return cq
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
//	client.Collection.Query().
//		GroupBy(collection.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *CollectionQuery) GroupBy(field string, fields ...string) *CollectionGroupBy {
	cq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CollectionGroupBy{build: cq}
	grbuild.flds = &cq.ctx.Fields
	grbuild.label = collection.Label
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
//	client.Collection.Query().
//		Select(collection.FieldCreatedAt).
//		Scan(ctx, &v)
func (cq *CollectionQuery) Select(fields ...string) *CollectionSelect {
	cq.ctx.Fields = append(cq.ctx.Fields, fields...)
	sbuild := &CollectionSelect{CollectionQuery: cq}
	sbuild.label = collection.Label
	sbuild.flds, sbuild.scan = &cq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CollectionSelect configured with the given aggregations.
func (cq *CollectionQuery) Aggregate(fns ...AggregateFunc) *CollectionSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *CollectionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.ctx.Fields {
		if !collection.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *CollectionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Collection, error) {
	var (
		nodes       = []*Collection{}
		withFKs     = cq.withFKs
		_spec       = cq.querySpec()
		loadedTypes = [6]bool{
			cq.withOwner != nil,
			cq.withCoverImage != nil,
			cq.withPosts != nil,
			cq.withNodes != nil,
			cq.withCollectionPosts != nil,
			cq.withCollectionNodes != nil,
		}
	)
	if cq.withOwner != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, collection.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Collection).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Collection{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cq.withOwner; query != nil {
		if err := cq.loadOwner(ctx, query, nodes, nil,
			func(n *Collection, e *Account) { n.Edges.Owner = e }); err != nil {
			return nil, err
		}
	}
	if query := cq.withCoverImage; query != nil {
		if err := cq.loadCoverImage(ctx, query, nodes, nil,
			func(n *Collection, e *Asset) { n.Edges.CoverImage = e }); err != nil {
			return nil, err
		}
	}
	if query := cq.withPosts; query != nil {
		if err := cq.loadPosts(ctx, query, nodes,
			func(n *Collection) { n.Edges.Posts = []*Post{} },
			func(n *Collection, e *Post) { n.Edges.Posts = append(n.Edges.Posts, e) }); err != nil {
			return nil, err
		}
	}
	if query := cq.withNodes; query != nil {
		if err := cq.loadNodes(ctx, query, nodes,
			func(n *Collection) { n.Edges.Nodes = []*Node{} },
			func(n *Collection, e *Node) { n.Edges.Nodes = append(n.Edges.Nodes, e) }); err != nil {
			return nil, err
		}
	}
	if query := cq.withCollectionPosts; query != nil {
		if err := cq.loadCollectionPosts(ctx, query, nodes,
			func(n *Collection) { n.Edges.CollectionPosts = []*CollectionPost{} },
			func(n *Collection, e *CollectionPost) { n.Edges.CollectionPosts = append(n.Edges.CollectionPosts, e) }); err != nil {
			return nil, err
		}
	}
	if query := cq.withCollectionNodes; query != nil {
		if err := cq.loadCollectionNodes(ctx, query, nodes,
			func(n *Collection) { n.Edges.CollectionNodes = []*CollectionNode{} },
			func(n *Collection, e *CollectionNode) { n.Edges.CollectionNodes = append(n.Edges.CollectionNodes, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *CollectionQuery) loadOwner(ctx context.Context, query *AccountQuery, nodes []*Collection, init func(*Collection), assign func(*Collection, *Account)) error {
	ids := make([]xid.ID, 0, len(nodes))
	nodeids := make(map[xid.ID][]*Collection)
	for i := range nodes {
		if nodes[i].account_collections == nil {
			continue
		}
		fk := *nodes[i].account_collections
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
			return fmt.Errorf(`unexpected foreign-key "account_collections" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (cq *CollectionQuery) loadCoverImage(ctx context.Context, query *AssetQuery, nodes []*Collection, init func(*Collection), assign func(*Collection, *Asset)) error {
	ids := make([]xid.ID, 0, len(nodes))
	nodeids := make(map[xid.ID][]*Collection)
	for i := range nodes {
		if nodes[i].CoverAssetID == nil {
			continue
		}
		fk := *nodes[i].CoverAssetID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(asset.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "cover_asset_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (cq *CollectionQuery) loadPosts(ctx context.Context, query *PostQuery, nodes []*Collection, init func(*Collection), assign func(*Collection, *Post)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[xid.ID]*Collection)
	nids := make(map[xid.ID]map[*Collection]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(collection.PostsTable)
		s.Join(joinT).On(s.C(post.FieldID), joinT.C(collection.PostsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(collection.PostsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(collection.PostsPrimaryKey[0]))
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
					nids[inValue] = map[*Collection]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Post](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "posts" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (cq *CollectionQuery) loadNodes(ctx context.Context, query *NodeQuery, nodes []*Collection, init func(*Collection), assign func(*Collection, *Node)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[xid.ID]*Collection)
	nids := make(map[xid.ID]map[*Collection]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(collection.NodesTable)
		s.Join(joinT).On(s.C(node.FieldID), joinT.C(collection.NodesPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(collection.NodesPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(collection.NodesPrimaryKey[0]))
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
					nids[inValue] = map[*Collection]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Node](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "nodes" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (cq *CollectionQuery) loadCollectionPosts(ctx context.Context, query *CollectionPostQuery, nodes []*Collection, init func(*Collection), assign func(*Collection, *CollectionPost)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[xid.ID]*Collection)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(collectionpost.FieldCollectionID)
	}
	query.Where(predicate.CollectionPost(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(collection.CollectionPostsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.CollectionID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "collection_id" returned %v for node %v`, fk, n)
		}
		assign(node, n)
	}
	return nil
}
func (cq *CollectionQuery) loadCollectionNodes(ctx context.Context, query *CollectionNodeQuery, nodes []*Collection, init func(*Collection), assign func(*Collection, *CollectionNode)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[xid.ID]*Collection)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(collectionnode.FieldCollectionID)
	}
	query.Where(predicate.CollectionNode(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(collection.CollectionNodesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.CollectionID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "collection_id" returned %v for node %v`, fk, n)
		}
		assign(node, n)
	}
	return nil
}

func (cq *CollectionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	_spec.Node.Columns = cq.ctx.Fields
	if len(cq.ctx.Fields) > 0 {
		_spec.Unique = cq.ctx.Unique != nil && *cq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *CollectionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(collection.Table, collection.Columns, sqlgraph.NewFieldSpec(collection.FieldID, field.TypeString))
	_spec.From = cq.sql
	if unique := cq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cq.path != nil {
		_spec.Unique = true
	}
	if fields := cq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, collection.FieldID)
		for i := range fields {
			if fields[i] != collection.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if cq.withCoverImage != nil {
			_spec.Node.AddColumnOnce(collection.FieldCoverAssetID)
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *CollectionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(collection.Table)
	columns := cq.ctx.Fields
	if len(columns) == 0 {
		columns = collection.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.ctx.Unique != nil && *cq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range cq.modifiers {
		m(selector)
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cq *CollectionQuery) Modify(modifiers ...func(s *sql.Selector)) *CollectionSelect {
	cq.modifiers = append(cq.modifiers, modifiers...)
	return cq.Select()
}

// CollectionGroupBy is the group-by builder for Collection entities.
type CollectionGroupBy struct {
	selector
	build *CollectionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *CollectionGroupBy) Aggregate(fns ...AggregateFunc) *CollectionGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *CollectionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, ent.OpQueryGroupBy)
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CollectionQuery, *CollectionGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *CollectionGroupBy) sqlScan(ctx context.Context, root *CollectionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds)+len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CollectionSelect is the builder for selecting fields of Collection entities.
type CollectionSelect struct {
	*CollectionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *CollectionSelect) Aggregate(fns ...AggregateFunc) *CollectionSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *CollectionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, ent.OpQuerySelect)
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CollectionQuery, *CollectionSelect](ctx, cs.CollectionQuery, cs, cs.inters, v)
}

func (cs *CollectionSelect) sqlScan(ctx context.Context, root *CollectionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cs *CollectionSelect) Modify(modifiers ...func(s *sql.Selector)) *CollectionSelect {
	cs.modifiers = append(cs.modifiers, modifiers...)
	return cs
}
