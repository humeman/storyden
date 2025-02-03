// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Southclaws/storyden/internal/ent/property"
	"github.com/Southclaws/storyden/internal/ent/propertyschema"
	"github.com/Southclaws/storyden/internal/ent/propertyschemafield"
	"github.com/rs/xid"
)

// PropertySchemaFieldCreate is the builder for creating a PropertySchemaField entity.
type PropertySchemaFieldCreate struct {
	config
	mutation *PropertySchemaFieldMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (psfc *PropertySchemaFieldCreate) SetName(s string) *PropertySchemaFieldCreate {
	psfc.mutation.SetName(s)
	return psfc
}

// SetType sets the "type" field.
func (psfc *PropertySchemaFieldCreate) SetType(s string) *PropertySchemaFieldCreate {
	psfc.mutation.SetType(s)
	return psfc
}

// SetSort sets the "sort" field.
func (psfc *PropertySchemaFieldCreate) SetSort(s string) *PropertySchemaFieldCreate {
	psfc.mutation.SetSort(s)
	return psfc
}

// SetSchemaID sets the "schema_id" field.
func (psfc *PropertySchemaFieldCreate) SetSchemaID(x xid.ID) *PropertySchemaFieldCreate {
	psfc.mutation.SetSchemaID(x)
	return psfc
}

// SetID sets the "id" field.
func (psfc *PropertySchemaFieldCreate) SetID(x xid.ID) *PropertySchemaFieldCreate {
	psfc.mutation.SetID(x)
	return psfc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (psfc *PropertySchemaFieldCreate) SetNillableID(x *xid.ID) *PropertySchemaFieldCreate {
	if x != nil {
		psfc.SetID(*x)
	}
	return psfc
}

// SetSchema sets the "schema" edge to the PropertySchema entity.
func (psfc *PropertySchemaFieldCreate) SetSchema(p *PropertySchema) *PropertySchemaFieldCreate {
	return psfc.SetSchemaID(p.ID)
}

// AddPropertyIDs adds the "properties" edge to the Property entity by IDs.
func (psfc *PropertySchemaFieldCreate) AddPropertyIDs(ids ...xid.ID) *PropertySchemaFieldCreate {
	psfc.mutation.AddPropertyIDs(ids...)
	return psfc
}

// AddProperties adds the "properties" edges to the Property entity.
func (psfc *PropertySchemaFieldCreate) AddProperties(p ...*Property) *PropertySchemaFieldCreate {
	ids := make([]xid.ID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return psfc.AddPropertyIDs(ids...)
}

// Mutation returns the PropertySchemaFieldMutation object of the builder.
func (psfc *PropertySchemaFieldCreate) Mutation() *PropertySchemaFieldMutation {
	return psfc.mutation
}

// Save creates the PropertySchemaField in the database.
func (psfc *PropertySchemaFieldCreate) Save(ctx context.Context) (*PropertySchemaField, error) {
	psfc.defaults()
	return withHooks(ctx, psfc.sqlSave, psfc.mutation, psfc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (psfc *PropertySchemaFieldCreate) SaveX(ctx context.Context) *PropertySchemaField {
	v, err := psfc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (psfc *PropertySchemaFieldCreate) Exec(ctx context.Context) error {
	_, err := psfc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psfc *PropertySchemaFieldCreate) ExecX(ctx context.Context) {
	if err := psfc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (psfc *PropertySchemaFieldCreate) defaults() {
	if _, ok := psfc.mutation.ID(); !ok {
		v := propertyschemafield.DefaultID()
		psfc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (psfc *PropertySchemaFieldCreate) check() error {
	if _, ok := psfc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "PropertySchemaField.name"`)}
	}
	if _, ok := psfc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "PropertySchemaField.type"`)}
	}
	if _, ok := psfc.mutation.Sort(); !ok {
		return &ValidationError{Name: "sort", err: errors.New(`ent: missing required field "PropertySchemaField.sort"`)}
	}
	if _, ok := psfc.mutation.SchemaID(); !ok {
		return &ValidationError{Name: "schema_id", err: errors.New(`ent: missing required field "PropertySchemaField.schema_id"`)}
	}
	if v, ok := psfc.mutation.ID(); ok {
		if err := propertyschemafield.IDValidator(v.String()); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "PropertySchemaField.id": %w`, err)}
		}
	}
	if len(psfc.mutation.SchemaIDs()) == 0 {
		return &ValidationError{Name: "schema", err: errors.New(`ent: missing required edge "PropertySchemaField.schema"`)}
	}
	return nil
}

func (psfc *PropertySchemaFieldCreate) sqlSave(ctx context.Context) (*PropertySchemaField, error) {
	if err := psfc.check(); err != nil {
		return nil, err
	}
	_node, _spec := psfc.createSpec()
	if err := sqlgraph.CreateNode(ctx, psfc.driver, _spec); err != nil {
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
	psfc.mutation.id = &_node.ID
	psfc.mutation.done = true
	return _node, nil
}

func (psfc *PropertySchemaFieldCreate) createSpec() (*PropertySchemaField, *sqlgraph.CreateSpec) {
	var (
		_node = &PropertySchemaField{config: psfc.config}
		_spec = sqlgraph.NewCreateSpec(propertyschemafield.Table, sqlgraph.NewFieldSpec(propertyschemafield.FieldID, field.TypeString))
	)
	_spec.OnConflict = psfc.conflict
	if id, ok := psfc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := psfc.mutation.Name(); ok {
		_spec.SetField(propertyschemafield.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := psfc.mutation.GetType(); ok {
		_spec.SetField(propertyschemafield.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := psfc.mutation.Sort(); ok {
		_spec.SetField(propertyschemafield.FieldSort, field.TypeString, value)
		_node.Sort = value
	}
	if nodes := psfc.mutation.SchemaIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   propertyschemafield.SchemaTable,
			Columns: []string{propertyschemafield.SchemaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(propertyschema.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.SchemaID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := psfc.mutation.PropertiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   propertyschemafield.PropertiesTable,
			Columns: []string{propertyschemafield.PropertiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(property.FieldID, field.TypeString),
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
//	client.PropertySchemaField.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PropertySchemaFieldUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (psfc *PropertySchemaFieldCreate) OnConflict(opts ...sql.ConflictOption) *PropertySchemaFieldUpsertOne {
	psfc.conflict = opts
	return &PropertySchemaFieldUpsertOne{
		create: psfc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.PropertySchemaField.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (psfc *PropertySchemaFieldCreate) OnConflictColumns(columns ...string) *PropertySchemaFieldUpsertOne {
	psfc.conflict = append(psfc.conflict, sql.ConflictColumns(columns...))
	return &PropertySchemaFieldUpsertOne{
		create: psfc,
	}
}

type (
	// PropertySchemaFieldUpsertOne is the builder for "upsert"-ing
	//  one PropertySchemaField node.
	PropertySchemaFieldUpsertOne struct {
		create *PropertySchemaFieldCreate
	}

	// PropertySchemaFieldUpsert is the "OnConflict" setter.
	PropertySchemaFieldUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *PropertySchemaFieldUpsert) SetName(v string) *PropertySchemaFieldUpsert {
	u.Set(propertyschemafield.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PropertySchemaFieldUpsert) UpdateName() *PropertySchemaFieldUpsert {
	u.SetExcluded(propertyschemafield.FieldName)
	return u
}

// SetType sets the "type" field.
func (u *PropertySchemaFieldUpsert) SetType(v string) *PropertySchemaFieldUpsert {
	u.Set(propertyschemafield.FieldType, v)
	return u
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *PropertySchemaFieldUpsert) UpdateType() *PropertySchemaFieldUpsert {
	u.SetExcluded(propertyschemafield.FieldType)
	return u
}

// SetSort sets the "sort" field.
func (u *PropertySchemaFieldUpsert) SetSort(v string) *PropertySchemaFieldUpsert {
	u.Set(propertyschemafield.FieldSort, v)
	return u
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *PropertySchemaFieldUpsert) UpdateSort() *PropertySchemaFieldUpsert {
	u.SetExcluded(propertyschemafield.FieldSort)
	return u
}

// SetSchemaID sets the "schema_id" field.
func (u *PropertySchemaFieldUpsert) SetSchemaID(v xid.ID) *PropertySchemaFieldUpsert {
	u.Set(propertyschemafield.FieldSchemaID, v)
	return u
}

// UpdateSchemaID sets the "schema_id" field to the value that was provided on create.
func (u *PropertySchemaFieldUpsert) UpdateSchemaID() *PropertySchemaFieldUpsert {
	u.SetExcluded(propertyschemafield.FieldSchemaID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.PropertySchemaField.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(propertyschemafield.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *PropertySchemaFieldUpsertOne) UpdateNewValues() *PropertySchemaFieldUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(propertyschemafield.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.PropertySchemaField.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *PropertySchemaFieldUpsertOne) Ignore() *PropertySchemaFieldUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PropertySchemaFieldUpsertOne) DoNothing() *PropertySchemaFieldUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PropertySchemaFieldCreate.OnConflict
// documentation for more info.
func (u *PropertySchemaFieldUpsertOne) Update(set func(*PropertySchemaFieldUpsert)) *PropertySchemaFieldUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PropertySchemaFieldUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *PropertySchemaFieldUpsertOne) SetName(v string) *PropertySchemaFieldUpsertOne {
	return u.Update(func(s *PropertySchemaFieldUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PropertySchemaFieldUpsertOne) UpdateName() *PropertySchemaFieldUpsertOne {
	return u.Update(func(s *PropertySchemaFieldUpsert) {
		s.UpdateName()
	})
}

// SetType sets the "type" field.
func (u *PropertySchemaFieldUpsertOne) SetType(v string) *PropertySchemaFieldUpsertOne {
	return u.Update(func(s *PropertySchemaFieldUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *PropertySchemaFieldUpsertOne) UpdateType() *PropertySchemaFieldUpsertOne {
	return u.Update(func(s *PropertySchemaFieldUpsert) {
		s.UpdateType()
	})
}

// SetSort sets the "sort" field.
func (u *PropertySchemaFieldUpsertOne) SetSort(v string) *PropertySchemaFieldUpsertOne {
	return u.Update(func(s *PropertySchemaFieldUpsert) {
		s.SetSort(v)
	})
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *PropertySchemaFieldUpsertOne) UpdateSort() *PropertySchemaFieldUpsertOne {
	return u.Update(func(s *PropertySchemaFieldUpsert) {
		s.UpdateSort()
	})
}

// SetSchemaID sets the "schema_id" field.
func (u *PropertySchemaFieldUpsertOne) SetSchemaID(v xid.ID) *PropertySchemaFieldUpsertOne {
	return u.Update(func(s *PropertySchemaFieldUpsert) {
		s.SetSchemaID(v)
	})
}

// UpdateSchemaID sets the "schema_id" field to the value that was provided on create.
func (u *PropertySchemaFieldUpsertOne) UpdateSchemaID() *PropertySchemaFieldUpsertOne {
	return u.Update(func(s *PropertySchemaFieldUpsert) {
		s.UpdateSchemaID()
	})
}

// Exec executes the query.
func (u *PropertySchemaFieldUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PropertySchemaFieldCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PropertySchemaFieldUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *PropertySchemaFieldUpsertOne) ID(ctx context.Context) (id xid.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: PropertySchemaFieldUpsertOne.ID is not supported by MySQL driver. Use PropertySchemaFieldUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *PropertySchemaFieldUpsertOne) IDX(ctx context.Context) xid.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// PropertySchemaFieldCreateBulk is the builder for creating many PropertySchemaField entities in bulk.
type PropertySchemaFieldCreateBulk struct {
	config
	err      error
	builders []*PropertySchemaFieldCreate
	conflict []sql.ConflictOption
}

// Save creates the PropertySchemaField entities in the database.
func (psfcb *PropertySchemaFieldCreateBulk) Save(ctx context.Context) ([]*PropertySchemaField, error) {
	if psfcb.err != nil {
		return nil, psfcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(psfcb.builders))
	nodes := make([]*PropertySchemaField, len(psfcb.builders))
	mutators := make([]Mutator, len(psfcb.builders))
	for i := range psfcb.builders {
		func(i int, root context.Context) {
			builder := psfcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PropertySchemaFieldMutation)
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
					_, err = mutators[i+1].Mutate(root, psfcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = psfcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, psfcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, psfcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (psfcb *PropertySchemaFieldCreateBulk) SaveX(ctx context.Context) []*PropertySchemaField {
	v, err := psfcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (psfcb *PropertySchemaFieldCreateBulk) Exec(ctx context.Context) error {
	_, err := psfcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psfcb *PropertySchemaFieldCreateBulk) ExecX(ctx context.Context) {
	if err := psfcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.PropertySchemaField.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PropertySchemaFieldUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (psfcb *PropertySchemaFieldCreateBulk) OnConflict(opts ...sql.ConflictOption) *PropertySchemaFieldUpsertBulk {
	psfcb.conflict = opts
	return &PropertySchemaFieldUpsertBulk{
		create: psfcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.PropertySchemaField.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (psfcb *PropertySchemaFieldCreateBulk) OnConflictColumns(columns ...string) *PropertySchemaFieldUpsertBulk {
	psfcb.conflict = append(psfcb.conflict, sql.ConflictColumns(columns...))
	return &PropertySchemaFieldUpsertBulk{
		create: psfcb,
	}
}

// PropertySchemaFieldUpsertBulk is the builder for "upsert"-ing
// a bulk of PropertySchemaField nodes.
type PropertySchemaFieldUpsertBulk struct {
	create *PropertySchemaFieldCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.PropertySchemaField.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(propertyschemafield.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *PropertySchemaFieldUpsertBulk) UpdateNewValues() *PropertySchemaFieldUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(propertyschemafield.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.PropertySchemaField.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *PropertySchemaFieldUpsertBulk) Ignore() *PropertySchemaFieldUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PropertySchemaFieldUpsertBulk) DoNothing() *PropertySchemaFieldUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PropertySchemaFieldCreateBulk.OnConflict
// documentation for more info.
func (u *PropertySchemaFieldUpsertBulk) Update(set func(*PropertySchemaFieldUpsert)) *PropertySchemaFieldUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PropertySchemaFieldUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *PropertySchemaFieldUpsertBulk) SetName(v string) *PropertySchemaFieldUpsertBulk {
	return u.Update(func(s *PropertySchemaFieldUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PropertySchemaFieldUpsertBulk) UpdateName() *PropertySchemaFieldUpsertBulk {
	return u.Update(func(s *PropertySchemaFieldUpsert) {
		s.UpdateName()
	})
}

// SetType sets the "type" field.
func (u *PropertySchemaFieldUpsertBulk) SetType(v string) *PropertySchemaFieldUpsertBulk {
	return u.Update(func(s *PropertySchemaFieldUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *PropertySchemaFieldUpsertBulk) UpdateType() *PropertySchemaFieldUpsertBulk {
	return u.Update(func(s *PropertySchemaFieldUpsert) {
		s.UpdateType()
	})
}

// SetSort sets the "sort" field.
func (u *PropertySchemaFieldUpsertBulk) SetSort(v string) *PropertySchemaFieldUpsertBulk {
	return u.Update(func(s *PropertySchemaFieldUpsert) {
		s.SetSort(v)
	})
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *PropertySchemaFieldUpsertBulk) UpdateSort() *PropertySchemaFieldUpsertBulk {
	return u.Update(func(s *PropertySchemaFieldUpsert) {
		s.UpdateSort()
	})
}

// SetSchemaID sets the "schema_id" field.
func (u *PropertySchemaFieldUpsertBulk) SetSchemaID(v xid.ID) *PropertySchemaFieldUpsertBulk {
	return u.Update(func(s *PropertySchemaFieldUpsert) {
		s.SetSchemaID(v)
	})
}

// UpdateSchemaID sets the "schema_id" field to the value that was provided on create.
func (u *PropertySchemaFieldUpsertBulk) UpdateSchemaID() *PropertySchemaFieldUpsertBulk {
	return u.Update(func(s *PropertySchemaFieldUpsert) {
		s.UpdateSchemaID()
	})
}

// Exec executes the query.
func (u *PropertySchemaFieldUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the PropertySchemaFieldCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PropertySchemaFieldCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PropertySchemaFieldUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
