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
	"github.com/Southclaws/storyden/internal/ent/category"
	"github.com/Southclaws/storyden/internal/ent/post"
	"github.com/rs/xid"
)

// CategoryCreate is the builder for creating a Category entity.
type CategoryCreate struct {
	config
	mutation *CategoryMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (cc *CategoryCreate) SetCreatedAt(t time.Time) *CategoryCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CategoryCreate) SetNillableCreatedAt(t *time.Time) *CategoryCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CategoryCreate) SetUpdatedAt(t time.Time) *CategoryCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CategoryCreate) SetNillableUpdatedAt(t *time.Time) *CategoryCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetName sets the "name" field.
func (cc *CategoryCreate) SetName(s string) *CategoryCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetDescription sets the "description" field.
func (cc *CategoryCreate) SetDescription(s string) *CategoryCreate {
	cc.mutation.SetDescription(s)
	return cc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cc *CategoryCreate) SetNillableDescription(s *string) *CategoryCreate {
	if s != nil {
		cc.SetDescription(*s)
	}
	return cc
}

// SetColour sets the "colour" field.
func (cc *CategoryCreate) SetColour(s string) *CategoryCreate {
	cc.mutation.SetColour(s)
	return cc
}

// SetNillableColour sets the "colour" field if the given value is not nil.
func (cc *CategoryCreate) SetNillableColour(s *string) *CategoryCreate {
	if s != nil {
		cc.SetColour(*s)
	}
	return cc
}

// SetSort sets the "sort" field.
func (cc *CategoryCreate) SetSort(i int) *CategoryCreate {
	cc.mutation.SetSort(i)
	return cc
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (cc *CategoryCreate) SetNillableSort(i *int) *CategoryCreate {
	if i != nil {
		cc.SetSort(*i)
	}
	return cc
}

// SetAdmin sets the "admin" field.
func (cc *CategoryCreate) SetAdmin(b bool) *CategoryCreate {
	cc.mutation.SetAdmin(b)
	return cc
}

// SetNillableAdmin sets the "admin" field if the given value is not nil.
func (cc *CategoryCreate) SetNillableAdmin(b *bool) *CategoryCreate {
	if b != nil {
		cc.SetAdmin(*b)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *CategoryCreate) SetID(x xid.ID) *CategoryCreate {
	cc.mutation.SetID(x)
	return cc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cc *CategoryCreate) SetNillableID(x *xid.ID) *CategoryCreate {
	if x != nil {
		cc.SetID(*x)
	}
	return cc
}

// AddPostIDs adds the "posts" edge to the Post entity by IDs.
func (cc *CategoryCreate) AddPostIDs(ids ...xid.ID) *CategoryCreate {
	cc.mutation.AddPostIDs(ids...)
	return cc
}

// AddPosts adds the "posts" edges to the Post entity.
func (cc *CategoryCreate) AddPosts(p ...*Post) *CategoryCreate {
	ids := make([]xid.ID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cc.AddPostIDs(ids...)
}

// Mutation returns the CategoryMutation object of the builder.
func (cc *CategoryCreate) Mutation() *CategoryMutation {
	return cc.mutation
}

// Save creates the Category in the database.
func (cc *CategoryCreate) Save(ctx context.Context) (*Category, error) {
	cc.defaults()
	return withHooks[*Category, CategoryMutation](ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CategoryCreate) SaveX(ctx context.Context) *Category {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CategoryCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CategoryCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CategoryCreate) defaults() {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := category.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		v := category.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cc.mutation.Description(); !ok {
		v := category.DefaultDescription
		cc.mutation.SetDescription(v)
	}
	if _, ok := cc.mutation.Colour(); !ok {
		v := category.DefaultColour
		cc.mutation.SetColour(v)
	}
	if _, ok := cc.mutation.Sort(); !ok {
		v := category.DefaultSort
		cc.mutation.SetSort(v)
	}
	if _, ok := cc.mutation.Admin(); !ok {
		v := category.DefaultAdmin
		cc.mutation.SetAdmin(v)
	}
	if _, ok := cc.mutation.ID(); !ok {
		v := category.DefaultID()
		cc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CategoryCreate) check() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Category.created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Category.updated_at"`)}
	}
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Category.name"`)}
	}
	if _, ok := cc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Category.description"`)}
	}
	if _, ok := cc.mutation.Colour(); !ok {
		return &ValidationError{Name: "colour", err: errors.New(`ent: missing required field "Category.colour"`)}
	}
	if _, ok := cc.mutation.Sort(); !ok {
		return &ValidationError{Name: "sort", err: errors.New(`ent: missing required field "Category.sort"`)}
	}
	if _, ok := cc.mutation.Admin(); !ok {
		return &ValidationError{Name: "admin", err: errors.New(`ent: missing required field "Category.admin"`)}
	}
	if v, ok := cc.mutation.ID(); ok {
		if err := category.IDValidator(v.String()); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Category.id": %w`, err)}
		}
	}
	return nil
}

func (cc *CategoryCreate) sqlSave(ctx context.Context) (*Category, error) {
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

func (cc *CategoryCreate) createSpec() (*Category, *sqlgraph.CreateSpec) {
	var (
		_node = &Category{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(category.Table, sqlgraph.NewFieldSpec(category.FieldID, field.TypeString))
	)
	_spec.OnConflict = cc.conflict
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(category.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(category.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(category.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cc.mutation.Description(); ok {
		_spec.SetField(category.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := cc.mutation.Colour(); ok {
		_spec.SetField(category.FieldColour, field.TypeString, value)
		_node.Colour = value
	}
	if value, ok := cc.mutation.Sort(); ok {
		_spec.SetField(category.FieldSort, field.TypeInt, value)
		_node.Sort = value
	}
	if value, ok := cc.mutation.Admin(); ok {
		_spec.SetField(category.FieldAdmin, field.TypeBool, value)
		_node.Admin = value
	}
	if nodes := cc.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.PostsTable,
			Columns: []string{category.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: post.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Category.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CategoryUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (cc *CategoryCreate) OnConflict(opts ...sql.ConflictOption) *CategoryUpsertOne {
	cc.conflict = opts
	return &CategoryUpsertOne{
		create: cc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Category.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (cc *CategoryCreate) OnConflictColumns(columns ...string) *CategoryUpsertOne {
	cc.conflict = append(cc.conflict, sql.ConflictColumns(columns...))
	return &CategoryUpsertOne{
		create: cc,
	}
}

type (
	// CategoryUpsertOne is the builder for "upsert"-ing
	//  one Category node.
	CategoryUpsertOne struct {
		create *CategoryCreate
	}

	// CategoryUpsert is the "OnConflict" setter.
	CategoryUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *CategoryUpsert) SetUpdatedAt(v time.Time) *CategoryUpsert {
	u.Set(category.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CategoryUpsert) UpdateUpdatedAt() *CategoryUpsert {
	u.SetExcluded(category.FieldUpdatedAt)
	return u
}

// SetName sets the "name" field.
func (u *CategoryUpsert) SetName(v string) *CategoryUpsert {
	u.Set(category.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CategoryUpsert) UpdateName() *CategoryUpsert {
	u.SetExcluded(category.FieldName)
	return u
}

// SetDescription sets the "description" field.
func (u *CategoryUpsert) SetDescription(v string) *CategoryUpsert {
	u.Set(category.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *CategoryUpsert) UpdateDescription() *CategoryUpsert {
	u.SetExcluded(category.FieldDescription)
	return u
}

// SetColour sets the "colour" field.
func (u *CategoryUpsert) SetColour(v string) *CategoryUpsert {
	u.Set(category.FieldColour, v)
	return u
}

// UpdateColour sets the "colour" field to the value that was provided on create.
func (u *CategoryUpsert) UpdateColour() *CategoryUpsert {
	u.SetExcluded(category.FieldColour)
	return u
}

// SetSort sets the "sort" field.
func (u *CategoryUpsert) SetSort(v int) *CategoryUpsert {
	u.Set(category.FieldSort, v)
	return u
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *CategoryUpsert) UpdateSort() *CategoryUpsert {
	u.SetExcluded(category.FieldSort)
	return u
}

// AddSort adds v to the "sort" field.
func (u *CategoryUpsert) AddSort(v int) *CategoryUpsert {
	u.Add(category.FieldSort, v)
	return u
}

// SetAdmin sets the "admin" field.
func (u *CategoryUpsert) SetAdmin(v bool) *CategoryUpsert {
	u.Set(category.FieldAdmin, v)
	return u
}

// UpdateAdmin sets the "admin" field to the value that was provided on create.
func (u *CategoryUpsert) UpdateAdmin() *CategoryUpsert {
	u.SetExcluded(category.FieldAdmin)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Category.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(category.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *CategoryUpsertOne) UpdateNewValues() *CategoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(category.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(category.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Category.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *CategoryUpsertOne) Ignore() *CategoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CategoryUpsertOne) DoNothing() *CategoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CategoryCreate.OnConflict
// documentation for more info.
func (u *CategoryUpsertOne) Update(set func(*CategoryUpsert)) *CategoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CategoryUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CategoryUpsertOne) SetUpdatedAt(v time.Time) *CategoryUpsertOne {
	return u.Update(func(s *CategoryUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CategoryUpsertOne) UpdateUpdatedAt() *CategoryUpsertOne {
	return u.Update(func(s *CategoryUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetName sets the "name" field.
func (u *CategoryUpsertOne) SetName(v string) *CategoryUpsertOne {
	return u.Update(func(s *CategoryUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CategoryUpsertOne) UpdateName() *CategoryUpsertOne {
	return u.Update(func(s *CategoryUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *CategoryUpsertOne) SetDescription(v string) *CategoryUpsertOne {
	return u.Update(func(s *CategoryUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *CategoryUpsertOne) UpdateDescription() *CategoryUpsertOne {
	return u.Update(func(s *CategoryUpsert) {
		s.UpdateDescription()
	})
}

// SetColour sets the "colour" field.
func (u *CategoryUpsertOne) SetColour(v string) *CategoryUpsertOne {
	return u.Update(func(s *CategoryUpsert) {
		s.SetColour(v)
	})
}

// UpdateColour sets the "colour" field to the value that was provided on create.
func (u *CategoryUpsertOne) UpdateColour() *CategoryUpsertOne {
	return u.Update(func(s *CategoryUpsert) {
		s.UpdateColour()
	})
}

// SetSort sets the "sort" field.
func (u *CategoryUpsertOne) SetSort(v int) *CategoryUpsertOne {
	return u.Update(func(s *CategoryUpsert) {
		s.SetSort(v)
	})
}

// AddSort adds v to the "sort" field.
func (u *CategoryUpsertOne) AddSort(v int) *CategoryUpsertOne {
	return u.Update(func(s *CategoryUpsert) {
		s.AddSort(v)
	})
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *CategoryUpsertOne) UpdateSort() *CategoryUpsertOne {
	return u.Update(func(s *CategoryUpsert) {
		s.UpdateSort()
	})
}

// SetAdmin sets the "admin" field.
func (u *CategoryUpsertOne) SetAdmin(v bool) *CategoryUpsertOne {
	return u.Update(func(s *CategoryUpsert) {
		s.SetAdmin(v)
	})
}

// UpdateAdmin sets the "admin" field to the value that was provided on create.
func (u *CategoryUpsertOne) UpdateAdmin() *CategoryUpsertOne {
	return u.Update(func(s *CategoryUpsert) {
		s.UpdateAdmin()
	})
}

// Exec executes the query.
func (u *CategoryUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CategoryCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CategoryUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *CategoryUpsertOne) ID(ctx context.Context) (id xid.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: CategoryUpsertOne.ID is not supported by MySQL driver. Use CategoryUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CategoryUpsertOne) IDX(ctx context.Context) xid.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// CategoryCreateBulk is the builder for creating many Category entities in bulk.
type CategoryCreateBulk struct {
	config
	builders []*CategoryCreate
	conflict []sql.ConflictOption
}

// Save creates the Category entities in the database.
func (ccb *CategoryCreateBulk) Save(ctx context.Context) ([]*Category, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Category, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CategoryMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
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
func (ccb *CategoryCreateBulk) SaveX(ctx context.Context) []*Category {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CategoryCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CategoryCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Category.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CategoryUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (ccb *CategoryCreateBulk) OnConflict(opts ...sql.ConflictOption) *CategoryUpsertBulk {
	ccb.conflict = opts
	return &CategoryUpsertBulk{
		create: ccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Category.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ccb *CategoryCreateBulk) OnConflictColumns(columns ...string) *CategoryUpsertBulk {
	ccb.conflict = append(ccb.conflict, sql.ConflictColumns(columns...))
	return &CategoryUpsertBulk{
		create: ccb,
	}
}

// CategoryUpsertBulk is the builder for "upsert"-ing
// a bulk of Category nodes.
type CategoryUpsertBulk struct {
	create *CategoryCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Category.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(category.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *CategoryUpsertBulk) UpdateNewValues() *CategoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(category.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(category.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Category.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *CategoryUpsertBulk) Ignore() *CategoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CategoryUpsertBulk) DoNothing() *CategoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CategoryCreateBulk.OnConflict
// documentation for more info.
func (u *CategoryUpsertBulk) Update(set func(*CategoryUpsert)) *CategoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CategoryUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CategoryUpsertBulk) SetUpdatedAt(v time.Time) *CategoryUpsertBulk {
	return u.Update(func(s *CategoryUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CategoryUpsertBulk) UpdateUpdatedAt() *CategoryUpsertBulk {
	return u.Update(func(s *CategoryUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetName sets the "name" field.
func (u *CategoryUpsertBulk) SetName(v string) *CategoryUpsertBulk {
	return u.Update(func(s *CategoryUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CategoryUpsertBulk) UpdateName() *CategoryUpsertBulk {
	return u.Update(func(s *CategoryUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *CategoryUpsertBulk) SetDescription(v string) *CategoryUpsertBulk {
	return u.Update(func(s *CategoryUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *CategoryUpsertBulk) UpdateDescription() *CategoryUpsertBulk {
	return u.Update(func(s *CategoryUpsert) {
		s.UpdateDescription()
	})
}

// SetColour sets the "colour" field.
func (u *CategoryUpsertBulk) SetColour(v string) *CategoryUpsertBulk {
	return u.Update(func(s *CategoryUpsert) {
		s.SetColour(v)
	})
}

// UpdateColour sets the "colour" field to the value that was provided on create.
func (u *CategoryUpsertBulk) UpdateColour() *CategoryUpsertBulk {
	return u.Update(func(s *CategoryUpsert) {
		s.UpdateColour()
	})
}

// SetSort sets the "sort" field.
func (u *CategoryUpsertBulk) SetSort(v int) *CategoryUpsertBulk {
	return u.Update(func(s *CategoryUpsert) {
		s.SetSort(v)
	})
}

// AddSort adds v to the "sort" field.
func (u *CategoryUpsertBulk) AddSort(v int) *CategoryUpsertBulk {
	return u.Update(func(s *CategoryUpsert) {
		s.AddSort(v)
	})
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *CategoryUpsertBulk) UpdateSort() *CategoryUpsertBulk {
	return u.Update(func(s *CategoryUpsert) {
		s.UpdateSort()
	})
}

// SetAdmin sets the "admin" field.
func (u *CategoryUpsertBulk) SetAdmin(v bool) *CategoryUpsertBulk {
	return u.Update(func(s *CategoryUpsert) {
		s.SetAdmin(v)
	})
}

// UpdateAdmin sets the "admin" field to the value that was provided on create.
func (u *CategoryUpsertBulk) UpdateAdmin() *CategoryUpsertBulk {
	return u.Update(func(s *CategoryUpsert) {
		s.UpdateAdmin()
	})
}

// Exec executes the query.
func (u *CategoryUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the CategoryCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CategoryCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CategoryUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
