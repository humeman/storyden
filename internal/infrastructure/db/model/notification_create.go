// Code generated by ent, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Southclaws/storyden/internal/infrastructure/db/model/notification"
	"github.com/Southclaws/storyden/internal/infrastructure/db/model/subscription"
	"github.com/rs/xid"
)

// NotificationCreate is the builder for creating a Notification entity.
type NotificationCreate struct {
	config
	mutation *NotificationMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (nc *NotificationCreate) SetCreatedAt(t time.Time) *NotificationCreate {
	nc.mutation.SetCreatedAt(t)
	return nc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (nc *NotificationCreate) SetNillableCreatedAt(t *time.Time) *NotificationCreate {
	if t != nil {
		nc.SetCreatedAt(*t)
	}
	return nc
}

// SetTitle sets the "title" field.
func (nc *NotificationCreate) SetTitle(s string) *NotificationCreate {
	nc.mutation.SetTitle(s)
	return nc
}

// SetDescription sets the "description" field.
func (nc *NotificationCreate) SetDescription(s string) *NotificationCreate {
	nc.mutation.SetDescription(s)
	return nc
}

// SetLink sets the "link" field.
func (nc *NotificationCreate) SetLink(s string) *NotificationCreate {
	nc.mutation.SetLink(s)
	return nc
}

// SetRead sets the "read" field.
func (nc *NotificationCreate) SetRead(b bool) *NotificationCreate {
	nc.mutation.SetRead(b)
	return nc
}

// SetID sets the "id" field.
func (nc *NotificationCreate) SetID(x xid.ID) *NotificationCreate {
	nc.mutation.SetID(x)
	return nc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (nc *NotificationCreate) SetNillableID(x *xid.ID) *NotificationCreate {
	if x != nil {
		nc.SetID(*x)
	}
	return nc
}

// SetSubscriptionID sets the "subscription" edge to the Subscription entity by ID.
func (nc *NotificationCreate) SetSubscriptionID(id xid.ID) *NotificationCreate {
	nc.mutation.SetSubscriptionID(id)
	return nc
}

// SetNillableSubscriptionID sets the "subscription" edge to the Subscription entity by ID if the given value is not nil.
func (nc *NotificationCreate) SetNillableSubscriptionID(id *xid.ID) *NotificationCreate {
	if id != nil {
		nc = nc.SetSubscriptionID(*id)
	}
	return nc
}

// SetSubscription sets the "subscription" edge to the Subscription entity.
func (nc *NotificationCreate) SetSubscription(s *Subscription) *NotificationCreate {
	return nc.SetSubscriptionID(s.ID)
}

// Mutation returns the NotificationMutation object of the builder.
func (nc *NotificationCreate) Mutation() *NotificationMutation {
	return nc.mutation
}

// Save creates the Notification in the database.
func (nc *NotificationCreate) Save(ctx context.Context) (*Notification, error) {
	var (
		err  error
		node *Notification
	)
	nc.defaults()
	if len(nc.hooks) == 0 {
		if err = nc.check(); err != nil {
			return nil, err
		}
		node, err = nc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NotificationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = nc.check(); err != nil {
				return nil, err
			}
			nc.mutation = mutation
			if node, err = nc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(nc.hooks) - 1; i >= 0; i-- {
			if nc.hooks[i] == nil {
				return nil, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = nc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, nc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Notification)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from NotificationMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (nc *NotificationCreate) SaveX(ctx context.Context) *Notification {
	v, err := nc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nc *NotificationCreate) Exec(ctx context.Context) error {
	_, err := nc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nc *NotificationCreate) ExecX(ctx context.Context) {
	if err := nc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nc *NotificationCreate) defaults() {
	if _, ok := nc.mutation.CreatedAt(); !ok {
		v := notification.DefaultCreatedAt()
		nc.mutation.SetCreatedAt(v)
	}
	if _, ok := nc.mutation.ID(); !ok {
		v := notification.DefaultID()
		nc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nc *NotificationCreate) check() error {
	if _, ok := nc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`model: missing required field "Notification.created_at"`)}
	}
	if _, ok := nc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`model: missing required field "Notification.title"`)}
	}
	if _, ok := nc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`model: missing required field "Notification.description"`)}
	}
	if _, ok := nc.mutation.Link(); !ok {
		return &ValidationError{Name: "link", err: errors.New(`model: missing required field "Notification.link"`)}
	}
	if _, ok := nc.mutation.Read(); !ok {
		return &ValidationError{Name: "read", err: errors.New(`model: missing required field "Notification.read"`)}
	}
	if v, ok := nc.mutation.ID(); ok {
		if err := notification.IDValidator(v.String()); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`model: validator failed for field "Notification.id": %w`, err)}
		}
	}
	return nil
}

func (nc *NotificationCreate) sqlSave(ctx context.Context) (*Notification, error) {
	_node, _spec := nc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nc.driver, _spec); err != nil {
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
	return _node, nil
}

func (nc *NotificationCreate) createSpec() (*Notification, *sqlgraph.CreateSpec) {
	var (
		_node = &Notification{config: nc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: notification.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: notification.FieldID,
			},
		}
	)
	_spec.OnConflict = nc.conflict
	if id, ok := nc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := nc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: notification.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := nc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: notification.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := nc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: notification.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := nc.mutation.Link(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: notification.FieldLink,
		})
		_node.Link = value
	}
	if value, ok := nc.mutation.Read(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: notification.FieldRead,
		})
		_node.Read = value
	}
	if nodes := nc.mutation.SubscriptionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   notification.SubscriptionTable,
			Columns: []string{notification.SubscriptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: subscription.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.notification_subscription = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Notification.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.NotificationUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (nc *NotificationCreate) OnConflict(opts ...sql.ConflictOption) *NotificationUpsertOne {
	nc.conflict = opts
	return &NotificationUpsertOne{
		create: nc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Notification.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (nc *NotificationCreate) OnConflictColumns(columns ...string) *NotificationUpsertOne {
	nc.conflict = append(nc.conflict, sql.ConflictColumns(columns...))
	return &NotificationUpsertOne{
		create: nc,
	}
}

type (
	// NotificationUpsertOne is the builder for "upsert"-ing
	//  one Notification node.
	NotificationUpsertOne struct {
		create *NotificationCreate
	}

	// NotificationUpsert is the "OnConflict" setter.
	NotificationUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *NotificationUpsert) SetCreatedAt(v time.Time) *NotificationUpsert {
	u.Set(notification.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *NotificationUpsert) UpdateCreatedAt() *NotificationUpsert {
	u.SetExcluded(notification.FieldCreatedAt)
	return u
}

// SetTitle sets the "title" field.
func (u *NotificationUpsert) SetTitle(v string) *NotificationUpsert {
	u.Set(notification.FieldTitle, v)
	return u
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *NotificationUpsert) UpdateTitle() *NotificationUpsert {
	u.SetExcluded(notification.FieldTitle)
	return u
}

// SetDescription sets the "description" field.
func (u *NotificationUpsert) SetDescription(v string) *NotificationUpsert {
	u.Set(notification.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *NotificationUpsert) UpdateDescription() *NotificationUpsert {
	u.SetExcluded(notification.FieldDescription)
	return u
}

// SetLink sets the "link" field.
func (u *NotificationUpsert) SetLink(v string) *NotificationUpsert {
	u.Set(notification.FieldLink, v)
	return u
}

// UpdateLink sets the "link" field to the value that was provided on create.
func (u *NotificationUpsert) UpdateLink() *NotificationUpsert {
	u.SetExcluded(notification.FieldLink)
	return u
}

// SetRead sets the "read" field.
func (u *NotificationUpsert) SetRead(v bool) *NotificationUpsert {
	u.Set(notification.FieldRead, v)
	return u
}

// UpdateRead sets the "read" field to the value that was provided on create.
func (u *NotificationUpsert) UpdateRead() *NotificationUpsert {
	u.SetExcluded(notification.FieldRead)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Notification.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(notification.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *NotificationUpsertOne) UpdateNewValues() *NotificationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(notification.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(notification.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Notification.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *NotificationUpsertOne) Ignore() *NotificationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *NotificationUpsertOne) DoNothing() *NotificationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the NotificationCreate.OnConflict
// documentation for more info.
func (u *NotificationUpsertOne) Update(set func(*NotificationUpsert)) *NotificationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&NotificationUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *NotificationUpsertOne) SetCreatedAt(v time.Time) *NotificationUpsertOne {
	return u.Update(func(s *NotificationUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *NotificationUpsertOne) UpdateCreatedAt() *NotificationUpsertOne {
	return u.Update(func(s *NotificationUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetTitle sets the "title" field.
func (u *NotificationUpsertOne) SetTitle(v string) *NotificationUpsertOne {
	return u.Update(func(s *NotificationUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *NotificationUpsertOne) UpdateTitle() *NotificationUpsertOne {
	return u.Update(func(s *NotificationUpsert) {
		s.UpdateTitle()
	})
}

// SetDescription sets the "description" field.
func (u *NotificationUpsertOne) SetDescription(v string) *NotificationUpsertOne {
	return u.Update(func(s *NotificationUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *NotificationUpsertOne) UpdateDescription() *NotificationUpsertOne {
	return u.Update(func(s *NotificationUpsert) {
		s.UpdateDescription()
	})
}

// SetLink sets the "link" field.
func (u *NotificationUpsertOne) SetLink(v string) *NotificationUpsertOne {
	return u.Update(func(s *NotificationUpsert) {
		s.SetLink(v)
	})
}

// UpdateLink sets the "link" field to the value that was provided on create.
func (u *NotificationUpsertOne) UpdateLink() *NotificationUpsertOne {
	return u.Update(func(s *NotificationUpsert) {
		s.UpdateLink()
	})
}

// SetRead sets the "read" field.
func (u *NotificationUpsertOne) SetRead(v bool) *NotificationUpsertOne {
	return u.Update(func(s *NotificationUpsert) {
		s.SetRead(v)
	})
}

// UpdateRead sets the "read" field to the value that was provided on create.
func (u *NotificationUpsertOne) UpdateRead() *NotificationUpsertOne {
	return u.Update(func(s *NotificationUpsert) {
		s.UpdateRead()
	})
}

// Exec executes the query.
func (u *NotificationUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for NotificationCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *NotificationUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *NotificationUpsertOne) ID(ctx context.Context) (id xid.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("model: NotificationUpsertOne.ID is not supported by MySQL driver. Use NotificationUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *NotificationUpsertOne) IDX(ctx context.Context) xid.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// NotificationCreateBulk is the builder for creating many Notification entities in bulk.
type NotificationCreateBulk struct {
	config
	builders []*NotificationCreate
	conflict []sql.ConflictOption
}

// Save creates the Notification entities in the database.
func (ncb *NotificationCreateBulk) Save(ctx context.Context) ([]*Notification, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ncb.builders))
	nodes := make([]*Notification, len(ncb.builders))
	mutators := make([]Mutator, len(ncb.builders))
	for i := range ncb.builders {
		func(i int, root context.Context) {
			builder := ncb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NotificationMutation)
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
					_, err = mutators[i+1].Mutate(root, ncb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ncb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ncb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ncb *NotificationCreateBulk) SaveX(ctx context.Context) []*Notification {
	v, err := ncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ncb *NotificationCreateBulk) Exec(ctx context.Context) error {
	_, err := ncb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ncb *NotificationCreateBulk) ExecX(ctx context.Context) {
	if err := ncb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Notification.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.NotificationUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (ncb *NotificationCreateBulk) OnConflict(opts ...sql.ConflictOption) *NotificationUpsertBulk {
	ncb.conflict = opts
	return &NotificationUpsertBulk{
		create: ncb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Notification.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ncb *NotificationCreateBulk) OnConflictColumns(columns ...string) *NotificationUpsertBulk {
	ncb.conflict = append(ncb.conflict, sql.ConflictColumns(columns...))
	return &NotificationUpsertBulk{
		create: ncb,
	}
}

// NotificationUpsertBulk is the builder for "upsert"-ing
// a bulk of Notification nodes.
type NotificationUpsertBulk struct {
	create *NotificationCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Notification.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(notification.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *NotificationUpsertBulk) UpdateNewValues() *NotificationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(notification.FieldID)
				return
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(notification.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Notification.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *NotificationUpsertBulk) Ignore() *NotificationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *NotificationUpsertBulk) DoNothing() *NotificationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the NotificationCreateBulk.OnConflict
// documentation for more info.
func (u *NotificationUpsertBulk) Update(set func(*NotificationUpsert)) *NotificationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&NotificationUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *NotificationUpsertBulk) SetCreatedAt(v time.Time) *NotificationUpsertBulk {
	return u.Update(func(s *NotificationUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *NotificationUpsertBulk) UpdateCreatedAt() *NotificationUpsertBulk {
	return u.Update(func(s *NotificationUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetTitle sets the "title" field.
func (u *NotificationUpsertBulk) SetTitle(v string) *NotificationUpsertBulk {
	return u.Update(func(s *NotificationUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *NotificationUpsertBulk) UpdateTitle() *NotificationUpsertBulk {
	return u.Update(func(s *NotificationUpsert) {
		s.UpdateTitle()
	})
}

// SetDescription sets the "description" field.
func (u *NotificationUpsertBulk) SetDescription(v string) *NotificationUpsertBulk {
	return u.Update(func(s *NotificationUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *NotificationUpsertBulk) UpdateDescription() *NotificationUpsertBulk {
	return u.Update(func(s *NotificationUpsert) {
		s.UpdateDescription()
	})
}

// SetLink sets the "link" field.
func (u *NotificationUpsertBulk) SetLink(v string) *NotificationUpsertBulk {
	return u.Update(func(s *NotificationUpsert) {
		s.SetLink(v)
	})
}

// UpdateLink sets the "link" field to the value that was provided on create.
func (u *NotificationUpsertBulk) UpdateLink() *NotificationUpsertBulk {
	return u.Update(func(s *NotificationUpsert) {
		s.UpdateLink()
	})
}

// SetRead sets the "read" field.
func (u *NotificationUpsertBulk) SetRead(v bool) *NotificationUpsertBulk {
	return u.Update(func(s *NotificationUpsert) {
		s.SetRead(v)
	})
}

// UpdateRead sets the "read" field to the value that was provided on create.
func (u *NotificationUpsertBulk) UpdateRead() *NotificationUpsertBulk {
	return u.Update(func(s *NotificationUpsert) {
		s.UpdateRead()
	})
}

// Exec executes the query.
func (u *NotificationUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("model: OnConflict was set for builder %d. Set it on the NotificationCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for NotificationCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *NotificationUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
