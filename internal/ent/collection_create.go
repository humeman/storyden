// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Southclaws/storyden/internal/ent/account"
	"github.com/Southclaws/storyden/internal/ent/asset"
	"github.com/Southclaws/storyden/internal/ent/collection"
	"github.com/Southclaws/storyden/internal/ent/node"
	"github.com/Southclaws/storyden/internal/ent/post"
	"github.com/rs/xid"
)

// CollectionCreate is the builder for creating a Collection entity.
type CollectionCreate struct {
	config
	mutation *CollectionMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (cc *CollectionCreate) SetCreatedAt(t time.Time) *CollectionCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CollectionCreate) SetNillableCreatedAt(t *time.Time) *CollectionCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CollectionCreate) SetUpdatedAt(t time.Time) *CollectionCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CollectionCreate) SetNillableUpdatedAt(t *time.Time) *CollectionCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetIndexedAt sets the "indexed_at" field.
func (cc *CollectionCreate) SetIndexedAt(t time.Time) *CollectionCreate {
	cc.mutation.SetIndexedAt(t)
	return cc
}

// SetNillableIndexedAt sets the "indexed_at" field if the given value is not nil.
func (cc *CollectionCreate) SetNillableIndexedAt(t *time.Time) *CollectionCreate {
	if t != nil {
		cc.SetIndexedAt(*t)
	}
	return cc
}

// SetName sets the "name" field.
func (cc *CollectionCreate) SetName(s string) *CollectionCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetSlug sets the "slug" field.
func (cc *CollectionCreate) SetSlug(s string) *CollectionCreate {
	cc.mutation.SetSlug(s)
	return cc
}

// SetDescription sets the "description" field.
func (cc *CollectionCreate) SetDescription(s string) *CollectionCreate {
	cc.mutation.SetDescription(s)
	return cc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cc *CollectionCreate) SetNillableDescription(s *string) *CollectionCreate {
	if s != nil {
		cc.SetDescription(*s)
	}
	return cc
}

// SetCoverAssetID sets the "cover_asset_id" field.
func (cc *CollectionCreate) SetCoverAssetID(x xid.ID) *CollectionCreate {
	cc.mutation.SetCoverAssetID(x)
	return cc
}

// SetNillableCoverAssetID sets the "cover_asset_id" field if the given value is not nil.
func (cc *CollectionCreate) SetNillableCoverAssetID(x *xid.ID) *CollectionCreate {
	if x != nil {
		cc.SetCoverAssetID(*x)
	}
	return cc
}

// SetVisibility sets the "visibility" field.
func (cc *CollectionCreate) SetVisibility(c collection.Visibility) *CollectionCreate {
	cc.mutation.SetVisibility(c)
	return cc
}

// SetNillableVisibility sets the "visibility" field if the given value is not nil.
func (cc *CollectionCreate) SetNillableVisibility(c *collection.Visibility) *CollectionCreate {
	if c != nil {
		cc.SetVisibility(*c)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *CollectionCreate) SetID(x xid.ID) *CollectionCreate {
	cc.mutation.SetID(x)
	return cc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cc *CollectionCreate) SetNillableID(x *xid.ID) *CollectionCreate {
	if x != nil {
		cc.SetID(*x)
	}
	return cc
}

// SetOwnerID sets the "owner" edge to the Account entity by ID.
func (cc *CollectionCreate) SetOwnerID(id xid.ID) *CollectionCreate {
	cc.mutation.SetOwnerID(id)
	return cc
}

// SetNillableOwnerID sets the "owner" edge to the Account entity by ID if the given value is not nil.
func (cc *CollectionCreate) SetNillableOwnerID(id *xid.ID) *CollectionCreate {
	if id != nil {
		cc = cc.SetOwnerID(*id)
	}
	return cc
}

// SetOwner sets the "owner" edge to the Account entity.
func (cc *CollectionCreate) SetOwner(a *Account) *CollectionCreate {
	return cc.SetOwnerID(a.ID)
}

// SetCoverImageID sets the "cover_image" edge to the Asset entity by ID.
func (cc *CollectionCreate) SetCoverImageID(id xid.ID) *CollectionCreate {
	cc.mutation.SetCoverImageID(id)
	return cc
}

// SetNillableCoverImageID sets the "cover_image" edge to the Asset entity by ID if the given value is not nil.
func (cc *CollectionCreate) SetNillableCoverImageID(id *xid.ID) *CollectionCreate {
	if id != nil {
		cc = cc.SetCoverImageID(*id)
	}
	return cc
}

// SetCoverImage sets the "cover_image" edge to the Asset entity.
func (cc *CollectionCreate) SetCoverImage(a *Asset) *CollectionCreate {
	return cc.SetCoverImageID(a.ID)
}

// AddPostIDs adds the "posts" edge to the Post entity by IDs.
func (cc *CollectionCreate) AddPostIDs(ids ...xid.ID) *CollectionCreate {
	cc.mutation.AddPostIDs(ids...)
	return cc
}

// AddPosts adds the "posts" edges to the Post entity.
func (cc *CollectionCreate) AddPosts(p ...*Post) *CollectionCreate {
	ids := make([]xid.ID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cc.AddPostIDs(ids...)
}

// AddNodeIDs adds the "nodes" edge to the Node entity by IDs.
func (cc *CollectionCreate) AddNodeIDs(ids ...xid.ID) *CollectionCreate {
	cc.mutation.AddNodeIDs(ids...)
	return cc
}

// AddNodes adds the "nodes" edges to the Node entity.
func (cc *CollectionCreate) AddNodes(n ...*Node) *CollectionCreate {
	ids := make([]xid.ID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return cc.AddNodeIDs(ids...)
}

// Mutation returns the CollectionMutation object of the builder.
func (cc *CollectionCreate) Mutation() *CollectionMutation {
	return cc.mutation
}

// Save creates the Collection in the database.
func (cc *CollectionCreate) Save(ctx context.Context) (*Collection, error) {
	cc.defaults()
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CollectionCreate) SaveX(ctx context.Context) *Collection {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CollectionCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CollectionCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CollectionCreate) defaults() {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := collection.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		v := collection.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cc.mutation.Visibility(); !ok {
		v := collection.DefaultVisibility
		cc.mutation.SetVisibility(v)
	}
	if _, ok := cc.mutation.ID(); !ok {
		v := collection.DefaultID()
		cc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CollectionCreate) check() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Collection.created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Collection.updated_at"`)}
	}
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Collection.name"`)}
	}
	if _, ok := cc.mutation.Slug(); !ok {
		return &ValidationError{Name: "slug", err: errors.New(`ent: missing required field "Collection.slug"`)}
	}
	if _, ok := cc.mutation.Visibility(); !ok {
		return &ValidationError{Name: "visibility", err: errors.New(`ent: missing required field "Collection.visibility"`)}
	}
	if v, ok := cc.mutation.Visibility(); ok {
		if err := collection.VisibilityValidator(v); err != nil {
			return &ValidationError{Name: "visibility", err: fmt.Errorf(`ent: validator failed for field "Collection.visibility": %w`, err)}
		}
	}
	if v, ok := cc.mutation.ID(); ok {
		if err := collection.IDValidator(v.String()); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Collection.id": %w`, err)}
		}
	}
	return nil
}

func (cc *CollectionCreate) sqlSave(ctx context.Context) (*Collection, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*xid.ID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CollectionCreate) createSpec() (*Collection, *sqlgraph.CreateSpec) {
	var (
		_node = &Collection{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(collection.Table, sqlgraph.NewFieldSpec(collection.FieldID, field.TypeString))
	)
	_spec.OnConflict = cc.conflict
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(collection.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(collection.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.IndexedAt(); ok {
		_spec.SetField(collection.FieldIndexedAt, field.TypeTime, value)
		_node.IndexedAt = &value
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(collection.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cc.mutation.Slug(); ok {
		_spec.SetField(collection.FieldSlug, field.TypeString, value)
		_node.Slug = value
	}
	if value, ok := cc.mutation.Description(); ok {
		_spec.SetField(collection.FieldDescription, field.TypeString, value)
		_node.Description = &value
	}
	if value, ok := cc.mutation.Visibility(); ok {
		_spec.SetField(collection.FieldVisibility, field.TypeEnum, value)
		_node.Visibility = value
	}
	if nodes := cc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   collection.OwnerTable,
			Columns: []string{collection.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.account_collections = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.CoverImageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   collection.CoverImageTable,
			Columns: []string{collection.CoverImageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(asset.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CoverAssetID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   collection.PostsTable,
			Columns: collection.PostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &CollectionPostCreate{config: cc.config, mutation: newCollectionPostMutation(cc.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.NodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   collection.NodesTable,
			Columns: collection.NodesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &CollectionNodeCreate{config: cc.config, mutation: newCollectionNodeMutation(cc.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Collection.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CollectionUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (cc *CollectionCreate) OnConflict(opts ...sql.ConflictOption) *CollectionUpsertOne {
	cc.conflict = opts
	return &CollectionUpsertOne{
		create: cc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Collection.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (cc *CollectionCreate) OnConflictColumns(columns ...string) *CollectionUpsertOne {
	cc.conflict = append(cc.conflict, sql.ConflictColumns(columns...))
	return &CollectionUpsertOne{
		create: cc,
	}
}

type (
	// CollectionUpsertOne is the builder for "upsert"-ing
	//  one Collection node.
	CollectionUpsertOne struct {
		create *CollectionCreate
	}

	// CollectionUpsert is the "OnConflict" setter.
	CollectionUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *CollectionUpsert) SetUpdatedAt(v time.Time) *CollectionUpsert {
	u.Set(collection.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CollectionUpsert) UpdateUpdatedAt() *CollectionUpsert {
	u.SetExcluded(collection.FieldUpdatedAt)
	return u
}

// SetIndexedAt sets the "indexed_at" field.
func (u *CollectionUpsert) SetIndexedAt(v time.Time) *CollectionUpsert {
	u.Set(collection.FieldIndexedAt, v)
	return u
}

// UpdateIndexedAt sets the "indexed_at" field to the value that was provided on create.
func (u *CollectionUpsert) UpdateIndexedAt() *CollectionUpsert {
	u.SetExcluded(collection.FieldIndexedAt)
	return u
}

// ClearIndexedAt clears the value of the "indexed_at" field.
func (u *CollectionUpsert) ClearIndexedAt() *CollectionUpsert {
	u.SetNull(collection.FieldIndexedAt)
	return u
}

// SetName sets the "name" field.
func (u *CollectionUpsert) SetName(v string) *CollectionUpsert {
	u.Set(collection.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CollectionUpsert) UpdateName() *CollectionUpsert {
	u.SetExcluded(collection.FieldName)
	return u
}

// SetSlug sets the "slug" field.
func (u *CollectionUpsert) SetSlug(v string) *CollectionUpsert {
	u.Set(collection.FieldSlug, v)
	return u
}

// UpdateSlug sets the "slug" field to the value that was provided on create.
func (u *CollectionUpsert) UpdateSlug() *CollectionUpsert {
	u.SetExcluded(collection.FieldSlug)
	return u
}

// SetDescription sets the "description" field.
func (u *CollectionUpsert) SetDescription(v string) *CollectionUpsert {
	u.Set(collection.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *CollectionUpsert) UpdateDescription() *CollectionUpsert {
	u.SetExcluded(collection.FieldDescription)
	return u
}

// ClearDescription clears the value of the "description" field.
func (u *CollectionUpsert) ClearDescription() *CollectionUpsert {
	u.SetNull(collection.FieldDescription)
	return u
}

// SetCoverAssetID sets the "cover_asset_id" field.
func (u *CollectionUpsert) SetCoverAssetID(v xid.ID) *CollectionUpsert {
	u.Set(collection.FieldCoverAssetID, v)
	return u
}

// UpdateCoverAssetID sets the "cover_asset_id" field to the value that was provided on create.
func (u *CollectionUpsert) UpdateCoverAssetID() *CollectionUpsert {
	u.SetExcluded(collection.FieldCoverAssetID)
	return u
}

// ClearCoverAssetID clears the value of the "cover_asset_id" field.
func (u *CollectionUpsert) ClearCoverAssetID() *CollectionUpsert {
	u.SetNull(collection.FieldCoverAssetID)
	return u
}

// SetVisibility sets the "visibility" field.
func (u *CollectionUpsert) SetVisibility(v collection.Visibility) *CollectionUpsert {
	u.Set(collection.FieldVisibility, v)
	return u
}

// UpdateVisibility sets the "visibility" field to the value that was provided on create.
func (u *CollectionUpsert) UpdateVisibility() *CollectionUpsert {
	u.SetExcluded(collection.FieldVisibility)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Collection.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(collection.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *CollectionUpsertOne) UpdateNewValues() *CollectionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(collection.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(collection.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Collection.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *CollectionUpsertOne) Ignore() *CollectionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CollectionUpsertOne) DoNothing() *CollectionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CollectionCreate.OnConflict
// documentation for more info.
func (u *CollectionUpsertOne) Update(set func(*CollectionUpsert)) *CollectionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CollectionUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CollectionUpsertOne) SetUpdatedAt(v time.Time) *CollectionUpsertOne {
	return u.Update(func(s *CollectionUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CollectionUpsertOne) UpdateUpdatedAt() *CollectionUpsertOne {
	return u.Update(func(s *CollectionUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetIndexedAt sets the "indexed_at" field.
func (u *CollectionUpsertOne) SetIndexedAt(v time.Time) *CollectionUpsertOne {
	return u.Update(func(s *CollectionUpsert) {
		s.SetIndexedAt(v)
	})
}

// UpdateIndexedAt sets the "indexed_at" field to the value that was provided on create.
func (u *CollectionUpsertOne) UpdateIndexedAt() *CollectionUpsertOne {
	return u.Update(func(s *CollectionUpsert) {
		s.UpdateIndexedAt()
	})
}

// ClearIndexedAt clears the value of the "indexed_at" field.
func (u *CollectionUpsertOne) ClearIndexedAt() *CollectionUpsertOne {
	return u.Update(func(s *CollectionUpsert) {
		s.ClearIndexedAt()
	})
}

// SetName sets the "name" field.
func (u *CollectionUpsertOne) SetName(v string) *CollectionUpsertOne {
	return u.Update(func(s *CollectionUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CollectionUpsertOne) UpdateName() *CollectionUpsertOne {
	return u.Update(func(s *CollectionUpsert) {
		s.UpdateName()
	})
}

// SetSlug sets the "slug" field.
func (u *CollectionUpsertOne) SetSlug(v string) *CollectionUpsertOne {
	return u.Update(func(s *CollectionUpsert) {
		s.SetSlug(v)
	})
}

// UpdateSlug sets the "slug" field to the value that was provided on create.
func (u *CollectionUpsertOne) UpdateSlug() *CollectionUpsertOne {
	return u.Update(func(s *CollectionUpsert) {
		s.UpdateSlug()
	})
}

// SetDescription sets the "description" field.
func (u *CollectionUpsertOne) SetDescription(v string) *CollectionUpsertOne {
	return u.Update(func(s *CollectionUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *CollectionUpsertOne) UpdateDescription() *CollectionUpsertOne {
	return u.Update(func(s *CollectionUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *CollectionUpsertOne) ClearDescription() *CollectionUpsertOne {
	return u.Update(func(s *CollectionUpsert) {
		s.ClearDescription()
	})
}

// SetCoverAssetID sets the "cover_asset_id" field.
func (u *CollectionUpsertOne) SetCoverAssetID(v xid.ID) *CollectionUpsertOne {
	return u.Update(func(s *CollectionUpsert) {
		s.SetCoverAssetID(v)
	})
}

// UpdateCoverAssetID sets the "cover_asset_id" field to the value that was provided on create.
func (u *CollectionUpsertOne) UpdateCoverAssetID() *CollectionUpsertOne {
	return u.Update(func(s *CollectionUpsert) {
		s.UpdateCoverAssetID()
	})
}

// ClearCoverAssetID clears the value of the "cover_asset_id" field.
func (u *CollectionUpsertOne) ClearCoverAssetID() *CollectionUpsertOne {
	return u.Update(func(s *CollectionUpsert) {
		s.ClearCoverAssetID()
	})
}

// SetVisibility sets the "visibility" field.
func (u *CollectionUpsertOne) SetVisibility(v collection.Visibility) *CollectionUpsertOne {
	return u.Update(func(s *CollectionUpsert) {
		s.SetVisibility(v)
	})
}

// UpdateVisibility sets the "visibility" field to the value that was provided on create.
func (u *CollectionUpsertOne) UpdateVisibility() *CollectionUpsertOne {
	return u.Update(func(s *CollectionUpsert) {
		s.UpdateVisibility()
	})
}

// Exec executes the query.
func (u *CollectionUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CollectionCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CollectionUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *CollectionUpsertOne) ID(ctx context.Context) (id xid.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: CollectionUpsertOne.ID is not supported by MySQL driver. Use CollectionUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CollectionUpsertOne) IDX(ctx context.Context) xid.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// CollectionCreateBulk is the builder for creating many Collection entities in bulk.
type CollectionCreateBulk struct {
	config
	err      error
	builders []*CollectionCreate
	conflict []sql.ConflictOption
}

// Save creates the Collection entities in the database.
func (ccb *CollectionCreateBulk) Save(ctx context.Context) ([]*Collection, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Collection, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CollectionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CollectionCreateBulk) SaveX(ctx context.Context) []*Collection {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CollectionCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CollectionCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Collection.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CollectionUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (ccb *CollectionCreateBulk) OnConflict(opts ...sql.ConflictOption) *CollectionUpsertBulk {
	ccb.conflict = opts
	return &CollectionUpsertBulk{
		create: ccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Collection.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ccb *CollectionCreateBulk) OnConflictColumns(columns ...string) *CollectionUpsertBulk {
	ccb.conflict = append(ccb.conflict, sql.ConflictColumns(columns...))
	return &CollectionUpsertBulk{
		create: ccb,
	}
}

// CollectionUpsertBulk is the builder for "upsert"-ing
// a bulk of Collection nodes.
type CollectionUpsertBulk struct {
	create *CollectionCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Collection.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(collection.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *CollectionUpsertBulk) UpdateNewValues() *CollectionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(collection.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(collection.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Collection.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *CollectionUpsertBulk) Ignore() *CollectionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CollectionUpsertBulk) DoNothing() *CollectionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CollectionCreateBulk.OnConflict
// documentation for more info.
func (u *CollectionUpsertBulk) Update(set func(*CollectionUpsert)) *CollectionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CollectionUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CollectionUpsertBulk) SetUpdatedAt(v time.Time) *CollectionUpsertBulk {
	return u.Update(func(s *CollectionUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CollectionUpsertBulk) UpdateUpdatedAt() *CollectionUpsertBulk {
	return u.Update(func(s *CollectionUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetIndexedAt sets the "indexed_at" field.
func (u *CollectionUpsertBulk) SetIndexedAt(v time.Time) *CollectionUpsertBulk {
	return u.Update(func(s *CollectionUpsert) {
		s.SetIndexedAt(v)
	})
}

// UpdateIndexedAt sets the "indexed_at" field to the value that was provided on create.
func (u *CollectionUpsertBulk) UpdateIndexedAt() *CollectionUpsertBulk {
	return u.Update(func(s *CollectionUpsert) {
		s.UpdateIndexedAt()
	})
}

// ClearIndexedAt clears the value of the "indexed_at" field.
func (u *CollectionUpsertBulk) ClearIndexedAt() *CollectionUpsertBulk {
	return u.Update(func(s *CollectionUpsert) {
		s.ClearIndexedAt()
	})
}

// SetName sets the "name" field.
func (u *CollectionUpsertBulk) SetName(v string) *CollectionUpsertBulk {
	return u.Update(func(s *CollectionUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CollectionUpsertBulk) UpdateName() *CollectionUpsertBulk {
	return u.Update(func(s *CollectionUpsert) {
		s.UpdateName()
	})
}

// SetSlug sets the "slug" field.
func (u *CollectionUpsertBulk) SetSlug(v string) *CollectionUpsertBulk {
	return u.Update(func(s *CollectionUpsert) {
		s.SetSlug(v)
	})
}

// UpdateSlug sets the "slug" field to the value that was provided on create.
func (u *CollectionUpsertBulk) UpdateSlug() *CollectionUpsertBulk {
	return u.Update(func(s *CollectionUpsert) {
		s.UpdateSlug()
	})
}

// SetDescription sets the "description" field.
func (u *CollectionUpsertBulk) SetDescription(v string) *CollectionUpsertBulk {
	return u.Update(func(s *CollectionUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *CollectionUpsertBulk) UpdateDescription() *CollectionUpsertBulk {
	return u.Update(func(s *CollectionUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *CollectionUpsertBulk) ClearDescription() *CollectionUpsertBulk {
	return u.Update(func(s *CollectionUpsert) {
		s.ClearDescription()
	})
}

// SetCoverAssetID sets the "cover_asset_id" field.
func (u *CollectionUpsertBulk) SetCoverAssetID(v xid.ID) *CollectionUpsertBulk {
	return u.Update(func(s *CollectionUpsert) {
		s.SetCoverAssetID(v)
	})
}

// UpdateCoverAssetID sets the "cover_asset_id" field to the value that was provided on create.
func (u *CollectionUpsertBulk) UpdateCoverAssetID() *CollectionUpsertBulk {
	return u.Update(func(s *CollectionUpsert) {
		s.UpdateCoverAssetID()
	})
}

// ClearCoverAssetID clears the value of the "cover_asset_id" field.
func (u *CollectionUpsertBulk) ClearCoverAssetID() *CollectionUpsertBulk {
	return u.Update(func(s *CollectionUpsert) {
		s.ClearCoverAssetID()
	})
}

// SetVisibility sets the "visibility" field.
func (u *CollectionUpsertBulk) SetVisibility(v collection.Visibility) *CollectionUpsertBulk {
	return u.Update(func(s *CollectionUpsert) {
		s.SetVisibility(v)
	})
}

// UpdateVisibility sets the "visibility" field to the value that was provided on create.
func (u *CollectionUpsertBulk) UpdateVisibility() *CollectionUpsertBulk {
	return u.Update(func(s *CollectionUpsert) {
		s.UpdateVisibility()
	})
}

// Exec executes the query.
func (u *CollectionUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the CollectionCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CollectionCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CollectionUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
