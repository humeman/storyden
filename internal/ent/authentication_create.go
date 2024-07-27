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
	"github.com/Southclaws/storyden/internal/ent/authentication"
	"github.com/Southclaws/storyden/internal/ent/email"
	"github.com/rs/xid"
)

// AuthenticationCreate is the builder for creating a Authentication entity.
type AuthenticationCreate struct {
	config
	mutation *AuthenticationMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (ac *AuthenticationCreate) SetCreatedAt(t time.Time) *AuthenticationCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AuthenticationCreate) SetNillableCreatedAt(t *time.Time) *AuthenticationCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetService sets the "service" field.
func (ac *AuthenticationCreate) SetService(s string) *AuthenticationCreate {
	ac.mutation.SetService(s)
	return ac
}

// SetIdentifier sets the "identifier" field.
func (ac *AuthenticationCreate) SetIdentifier(s string) *AuthenticationCreate {
	ac.mutation.SetIdentifier(s)
	return ac
}

// SetToken sets the "token" field.
func (ac *AuthenticationCreate) SetToken(s string) *AuthenticationCreate {
	ac.mutation.SetToken(s)
	return ac
}

// SetName sets the "name" field.
func (ac *AuthenticationCreate) SetName(s string) *AuthenticationCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ac *AuthenticationCreate) SetNillableName(s *string) *AuthenticationCreate {
	if s != nil {
		ac.SetName(*s)
	}
	return ac
}

// SetMetadata sets the "metadata" field.
func (ac *AuthenticationCreate) SetMetadata(m map[string]interface{}) *AuthenticationCreate {
	ac.mutation.SetMetadata(m)
	return ac
}

// SetID sets the "id" field.
func (ac *AuthenticationCreate) SetID(x xid.ID) *AuthenticationCreate {
	ac.mutation.SetID(x)
	return ac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ac *AuthenticationCreate) SetNillableID(x *xid.ID) *AuthenticationCreate {
	if x != nil {
		ac.SetID(*x)
	}
	return ac
}

// SetAccountID sets the "account" edge to the Account entity by ID.
func (ac *AuthenticationCreate) SetAccountID(id xid.ID) *AuthenticationCreate {
	ac.mutation.SetAccountID(id)
	return ac
}

// SetAccount sets the "account" edge to the Account entity.
func (ac *AuthenticationCreate) SetAccount(a *Account) *AuthenticationCreate {
	return ac.SetAccountID(a.ID)
}

// AddEmailAddresIDs adds the "email_address" edge to the Email entity by IDs.
func (ac *AuthenticationCreate) AddEmailAddresIDs(ids ...xid.ID) *AuthenticationCreate {
	ac.mutation.AddEmailAddresIDs(ids...)
	return ac
}

// AddEmailAddress adds the "email_address" edges to the Email entity.
func (ac *AuthenticationCreate) AddEmailAddress(e ...*Email) *AuthenticationCreate {
	ids := make([]xid.ID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ac.AddEmailAddresIDs(ids...)
}

// Mutation returns the AuthenticationMutation object of the builder.
func (ac *AuthenticationCreate) Mutation() *AuthenticationMutation {
	return ac.mutation
}

// Save creates the Authentication in the database.
func (ac *AuthenticationCreate) Save(ctx context.Context) (*Authentication, error) {
	ac.defaults()
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AuthenticationCreate) SaveX(ctx context.Context) *Authentication {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AuthenticationCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AuthenticationCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AuthenticationCreate) defaults() {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := authentication.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.ID(); !ok {
		v := authentication.DefaultID()
		ac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AuthenticationCreate) check() error {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Authentication.created_at"`)}
	}
	if _, ok := ac.mutation.Service(); !ok {
		return &ValidationError{Name: "service", err: errors.New(`ent: missing required field "Authentication.service"`)}
	}
	if v, ok := ac.mutation.Service(); ok {
		if err := authentication.ServiceValidator(v); err != nil {
			return &ValidationError{Name: "service", err: fmt.Errorf(`ent: validator failed for field "Authentication.service": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Identifier(); !ok {
		return &ValidationError{Name: "identifier", err: errors.New(`ent: missing required field "Authentication.identifier"`)}
	}
	if _, ok := ac.mutation.Token(); !ok {
		return &ValidationError{Name: "token", err: errors.New(`ent: missing required field "Authentication.token"`)}
	}
	if v, ok := ac.mutation.Token(); ok {
		if err := authentication.TokenValidator(v); err != nil {
			return &ValidationError{Name: "token", err: fmt.Errorf(`ent: validator failed for field "Authentication.token": %w`, err)}
		}
	}
	if v, ok := ac.mutation.ID(); ok {
		if err := authentication.IDValidator(v.String()); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Authentication.id": %w`, err)}
		}
	}
	if _, ok := ac.mutation.AccountID(); !ok {
		return &ValidationError{Name: "account", err: errors.New(`ent: missing required edge "Authentication.account"`)}
	}
	return nil
}

func (ac *AuthenticationCreate) sqlSave(ctx context.Context) (*Authentication, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
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
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AuthenticationCreate) createSpec() (*Authentication, *sqlgraph.CreateSpec) {
	var (
		_node = &Authentication{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(authentication.Table, sqlgraph.NewFieldSpec(authentication.FieldID, field.TypeString))
	)
	_spec.OnConflict = ac.conflict
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(authentication.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.Service(); ok {
		_spec.SetField(authentication.FieldService, field.TypeString, value)
		_node.Service = value
	}
	if value, ok := ac.mutation.Identifier(); ok {
		_spec.SetField(authentication.FieldIdentifier, field.TypeString, value)
		_node.Identifier = value
	}
	if value, ok := ac.mutation.Token(); ok {
		_spec.SetField(authentication.FieldToken, field.TypeString, value)
		_node.Token = value
	}
	if value, ok := ac.mutation.Name(); ok {
		_spec.SetField(authentication.FieldName, field.TypeString, value)
		_node.Name = &value
	}
	if value, ok := ac.mutation.Metadata(); ok {
		_spec.SetField(authentication.FieldMetadata, field.TypeJSON, value)
		_node.Metadata = value
	}
	if nodes := ac.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   authentication.AccountTable,
			Columns: []string{authentication.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.account_authentication = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.EmailAddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   authentication.EmailAddressTable,
			Columns: []string{authentication.EmailAddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(email.FieldID, field.TypeString),
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
//	client.Authentication.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AuthenticationUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (ac *AuthenticationCreate) OnConflict(opts ...sql.ConflictOption) *AuthenticationUpsertOne {
	ac.conflict = opts
	return &AuthenticationUpsertOne{
		create: ac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Authentication.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ac *AuthenticationCreate) OnConflictColumns(columns ...string) *AuthenticationUpsertOne {
	ac.conflict = append(ac.conflict, sql.ConflictColumns(columns...))
	return &AuthenticationUpsertOne{
		create: ac,
	}
}

type (
	// AuthenticationUpsertOne is the builder for "upsert"-ing
	//  one Authentication node.
	AuthenticationUpsertOne struct {
		create *AuthenticationCreate
	}

	// AuthenticationUpsert is the "OnConflict" setter.
	AuthenticationUpsert struct {
		*sql.UpdateSet
	}
)

// SetService sets the "service" field.
func (u *AuthenticationUpsert) SetService(v string) *AuthenticationUpsert {
	u.Set(authentication.FieldService, v)
	return u
}

// UpdateService sets the "service" field to the value that was provided on create.
func (u *AuthenticationUpsert) UpdateService() *AuthenticationUpsert {
	u.SetExcluded(authentication.FieldService)
	return u
}

// SetIdentifier sets the "identifier" field.
func (u *AuthenticationUpsert) SetIdentifier(v string) *AuthenticationUpsert {
	u.Set(authentication.FieldIdentifier, v)
	return u
}

// UpdateIdentifier sets the "identifier" field to the value that was provided on create.
func (u *AuthenticationUpsert) UpdateIdentifier() *AuthenticationUpsert {
	u.SetExcluded(authentication.FieldIdentifier)
	return u
}

// SetToken sets the "token" field.
func (u *AuthenticationUpsert) SetToken(v string) *AuthenticationUpsert {
	u.Set(authentication.FieldToken, v)
	return u
}

// UpdateToken sets the "token" field to the value that was provided on create.
func (u *AuthenticationUpsert) UpdateToken() *AuthenticationUpsert {
	u.SetExcluded(authentication.FieldToken)
	return u
}

// SetName sets the "name" field.
func (u *AuthenticationUpsert) SetName(v string) *AuthenticationUpsert {
	u.Set(authentication.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AuthenticationUpsert) UpdateName() *AuthenticationUpsert {
	u.SetExcluded(authentication.FieldName)
	return u
}

// ClearName clears the value of the "name" field.
func (u *AuthenticationUpsert) ClearName() *AuthenticationUpsert {
	u.SetNull(authentication.FieldName)
	return u
}

// SetMetadata sets the "metadata" field.
func (u *AuthenticationUpsert) SetMetadata(v map[string]interface{}) *AuthenticationUpsert {
	u.Set(authentication.FieldMetadata, v)
	return u
}

// UpdateMetadata sets the "metadata" field to the value that was provided on create.
func (u *AuthenticationUpsert) UpdateMetadata() *AuthenticationUpsert {
	u.SetExcluded(authentication.FieldMetadata)
	return u
}

// ClearMetadata clears the value of the "metadata" field.
func (u *AuthenticationUpsert) ClearMetadata() *AuthenticationUpsert {
	u.SetNull(authentication.FieldMetadata)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Authentication.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(authentication.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AuthenticationUpsertOne) UpdateNewValues() *AuthenticationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(authentication.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(authentication.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Authentication.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *AuthenticationUpsertOne) Ignore() *AuthenticationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AuthenticationUpsertOne) DoNothing() *AuthenticationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AuthenticationCreate.OnConflict
// documentation for more info.
func (u *AuthenticationUpsertOne) Update(set func(*AuthenticationUpsert)) *AuthenticationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AuthenticationUpsert{UpdateSet: update})
	}))
	return u
}

// SetService sets the "service" field.
func (u *AuthenticationUpsertOne) SetService(v string) *AuthenticationUpsertOne {
	return u.Update(func(s *AuthenticationUpsert) {
		s.SetService(v)
	})
}

// UpdateService sets the "service" field to the value that was provided on create.
func (u *AuthenticationUpsertOne) UpdateService() *AuthenticationUpsertOne {
	return u.Update(func(s *AuthenticationUpsert) {
		s.UpdateService()
	})
}

// SetIdentifier sets the "identifier" field.
func (u *AuthenticationUpsertOne) SetIdentifier(v string) *AuthenticationUpsertOne {
	return u.Update(func(s *AuthenticationUpsert) {
		s.SetIdentifier(v)
	})
}

// UpdateIdentifier sets the "identifier" field to the value that was provided on create.
func (u *AuthenticationUpsertOne) UpdateIdentifier() *AuthenticationUpsertOne {
	return u.Update(func(s *AuthenticationUpsert) {
		s.UpdateIdentifier()
	})
}

// SetToken sets the "token" field.
func (u *AuthenticationUpsertOne) SetToken(v string) *AuthenticationUpsertOne {
	return u.Update(func(s *AuthenticationUpsert) {
		s.SetToken(v)
	})
}

// UpdateToken sets the "token" field to the value that was provided on create.
func (u *AuthenticationUpsertOne) UpdateToken() *AuthenticationUpsertOne {
	return u.Update(func(s *AuthenticationUpsert) {
		s.UpdateToken()
	})
}

// SetName sets the "name" field.
func (u *AuthenticationUpsertOne) SetName(v string) *AuthenticationUpsertOne {
	return u.Update(func(s *AuthenticationUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AuthenticationUpsertOne) UpdateName() *AuthenticationUpsertOne {
	return u.Update(func(s *AuthenticationUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *AuthenticationUpsertOne) ClearName() *AuthenticationUpsertOne {
	return u.Update(func(s *AuthenticationUpsert) {
		s.ClearName()
	})
}

// SetMetadata sets the "metadata" field.
func (u *AuthenticationUpsertOne) SetMetadata(v map[string]interface{}) *AuthenticationUpsertOne {
	return u.Update(func(s *AuthenticationUpsert) {
		s.SetMetadata(v)
	})
}

// UpdateMetadata sets the "metadata" field to the value that was provided on create.
func (u *AuthenticationUpsertOne) UpdateMetadata() *AuthenticationUpsertOne {
	return u.Update(func(s *AuthenticationUpsert) {
		s.UpdateMetadata()
	})
}

// ClearMetadata clears the value of the "metadata" field.
func (u *AuthenticationUpsertOne) ClearMetadata() *AuthenticationUpsertOne {
	return u.Update(func(s *AuthenticationUpsert) {
		s.ClearMetadata()
	})
}

// Exec executes the query.
func (u *AuthenticationUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AuthenticationCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AuthenticationUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AuthenticationUpsertOne) ID(ctx context.Context) (id xid.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: AuthenticationUpsertOne.ID is not supported by MySQL driver. Use AuthenticationUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AuthenticationUpsertOne) IDX(ctx context.Context) xid.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AuthenticationCreateBulk is the builder for creating many Authentication entities in bulk.
type AuthenticationCreateBulk struct {
	config
	err      error
	builders []*AuthenticationCreate
	conflict []sql.ConflictOption
}

// Save creates the Authentication entities in the database.
func (acb *AuthenticationCreateBulk) Save(ctx context.Context) ([]*Authentication, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Authentication, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AuthenticationMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = acb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AuthenticationCreateBulk) SaveX(ctx context.Context) []*Authentication {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AuthenticationCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AuthenticationCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Authentication.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AuthenticationUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (acb *AuthenticationCreateBulk) OnConflict(opts ...sql.ConflictOption) *AuthenticationUpsertBulk {
	acb.conflict = opts
	return &AuthenticationUpsertBulk{
		create: acb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Authentication.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (acb *AuthenticationCreateBulk) OnConflictColumns(columns ...string) *AuthenticationUpsertBulk {
	acb.conflict = append(acb.conflict, sql.ConflictColumns(columns...))
	return &AuthenticationUpsertBulk{
		create: acb,
	}
}

// AuthenticationUpsertBulk is the builder for "upsert"-ing
// a bulk of Authentication nodes.
type AuthenticationUpsertBulk struct {
	create *AuthenticationCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Authentication.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(authentication.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AuthenticationUpsertBulk) UpdateNewValues() *AuthenticationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(authentication.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(authentication.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Authentication.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *AuthenticationUpsertBulk) Ignore() *AuthenticationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AuthenticationUpsertBulk) DoNothing() *AuthenticationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AuthenticationCreateBulk.OnConflict
// documentation for more info.
func (u *AuthenticationUpsertBulk) Update(set func(*AuthenticationUpsert)) *AuthenticationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AuthenticationUpsert{UpdateSet: update})
	}))
	return u
}

// SetService sets the "service" field.
func (u *AuthenticationUpsertBulk) SetService(v string) *AuthenticationUpsertBulk {
	return u.Update(func(s *AuthenticationUpsert) {
		s.SetService(v)
	})
}

// UpdateService sets the "service" field to the value that was provided on create.
func (u *AuthenticationUpsertBulk) UpdateService() *AuthenticationUpsertBulk {
	return u.Update(func(s *AuthenticationUpsert) {
		s.UpdateService()
	})
}

// SetIdentifier sets the "identifier" field.
func (u *AuthenticationUpsertBulk) SetIdentifier(v string) *AuthenticationUpsertBulk {
	return u.Update(func(s *AuthenticationUpsert) {
		s.SetIdentifier(v)
	})
}

// UpdateIdentifier sets the "identifier" field to the value that was provided on create.
func (u *AuthenticationUpsertBulk) UpdateIdentifier() *AuthenticationUpsertBulk {
	return u.Update(func(s *AuthenticationUpsert) {
		s.UpdateIdentifier()
	})
}

// SetToken sets the "token" field.
func (u *AuthenticationUpsertBulk) SetToken(v string) *AuthenticationUpsertBulk {
	return u.Update(func(s *AuthenticationUpsert) {
		s.SetToken(v)
	})
}

// UpdateToken sets the "token" field to the value that was provided on create.
func (u *AuthenticationUpsertBulk) UpdateToken() *AuthenticationUpsertBulk {
	return u.Update(func(s *AuthenticationUpsert) {
		s.UpdateToken()
	})
}

// SetName sets the "name" field.
func (u *AuthenticationUpsertBulk) SetName(v string) *AuthenticationUpsertBulk {
	return u.Update(func(s *AuthenticationUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AuthenticationUpsertBulk) UpdateName() *AuthenticationUpsertBulk {
	return u.Update(func(s *AuthenticationUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *AuthenticationUpsertBulk) ClearName() *AuthenticationUpsertBulk {
	return u.Update(func(s *AuthenticationUpsert) {
		s.ClearName()
	})
}

// SetMetadata sets the "metadata" field.
func (u *AuthenticationUpsertBulk) SetMetadata(v map[string]interface{}) *AuthenticationUpsertBulk {
	return u.Update(func(s *AuthenticationUpsert) {
		s.SetMetadata(v)
	})
}

// UpdateMetadata sets the "metadata" field to the value that was provided on create.
func (u *AuthenticationUpsertBulk) UpdateMetadata() *AuthenticationUpsertBulk {
	return u.Update(func(s *AuthenticationUpsert) {
		s.UpdateMetadata()
	})
}

// ClearMetadata clears the value of the "metadata" field.
func (u *AuthenticationUpsertBulk) ClearMetadata() *AuthenticationUpsertBulk {
	return u.Update(func(s *AuthenticationUpsert) {
		s.ClearMetadata()
	})
}

// Exec executes the query.
func (u *AuthenticationUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AuthenticationCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AuthenticationCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AuthenticationUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
