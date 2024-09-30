// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Southclaws/storyden/internal/ent/account"
	"github.com/Southclaws/storyden/internal/ent/event"
	"github.com/Southclaws/storyden/internal/ent/eventparticipant"
	"github.com/Southclaws/storyden/internal/ent/predicate"
	"github.com/rs/xid"
)

// EventParticipantUpdate is the builder for updating EventParticipant entities.
type EventParticipantUpdate struct {
	config
	hooks     []Hook
	mutation  *EventParticipantMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the EventParticipantUpdate builder.
func (epu *EventParticipantUpdate) Where(ps ...predicate.EventParticipant) *EventParticipantUpdate {
	epu.mutation.Where(ps...)
	return epu
}

// SetRole sets the "role" field.
func (epu *EventParticipantUpdate) SetRole(s string) *EventParticipantUpdate {
	epu.mutation.SetRole(s)
	return epu
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (epu *EventParticipantUpdate) SetNillableRole(s *string) *EventParticipantUpdate {
	if s != nil {
		epu.SetRole(*s)
	}
	return epu
}

// SetStatus sets the "status" field.
func (epu *EventParticipantUpdate) SetStatus(s string) *EventParticipantUpdate {
	epu.mutation.SetStatus(s)
	return epu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (epu *EventParticipantUpdate) SetNillableStatus(s *string) *EventParticipantUpdate {
	if s != nil {
		epu.SetStatus(*s)
	}
	return epu
}

// SetAccountID sets the "account_id" field.
func (epu *EventParticipantUpdate) SetAccountID(x xid.ID) *EventParticipantUpdate {
	epu.mutation.SetAccountID(x)
	return epu
}

// SetNillableAccountID sets the "account_id" field if the given value is not nil.
func (epu *EventParticipantUpdate) SetNillableAccountID(x *xid.ID) *EventParticipantUpdate {
	if x != nil {
		epu.SetAccountID(*x)
	}
	return epu
}

// SetEventID sets the "event_id" field.
func (epu *EventParticipantUpdate) SetEventID(x xid.ID) *EventParticipantUpdate {
	epu.mutation.SetEventID(x)
	return epu
}

// SetNillableEventID sets the "event_id" field if the given value is not nil.
func (epu *EventParticipantUpdate) SetNillableEventID(x *xid.ID) *EventParticipantUpdate {
	if x != nil {
		epu.SetEventID(*x)
	}
	return epu
}

// SetAccount sets the "account" edge to the Account entity.
func (epu *EventParticipantUpdate) SetAccount(a *Account) *EventParticipantUpdate {
	return epu.SetAccountID(a.ID)
}

// SetEvent sets the "event" edge to the Event entity.
func (epu *EventParticipantUpdate) SetEvent(e *Event) *EventParticipantUpdate {
	return epu.SetEventID(e.ID)
}

// Mutation returns the EventParticipantMutation object of the builder.
func (epu *EventParticipantUpdate) Mutation() *EventParticipantMutation {
	return epu.mutation
}

// ClearAccount clears the "account" edge to the Account entity.
func (epu *EventParticipantUpdate) ClearAccount() *EventParticipantUpdate {
	epu.mutation.ClearAccount()
	return epu
}

// ClearEvent clears the "event" edge to the Event entity.
func (epu *EventParticipantUpdate) ClearEvent() *EventParticipantUpdate {
	epu.mutation.ClearEvent()
	return epu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (epu *EventParticipantUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, epu.sqlSave, epu.mutation, epu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (epu *EventParticipantUpdate) SaveX(ctx context.Context) int {
	affected, err := epu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (epu *EventParticipantUpdate) Exec(ctx context.Context) error {
	_, err := epu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (epu *EventParticipantUpdate) ExecX(ctx context.Context) {
	if err := epu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (epu *EventParticipantUpdate) check() error {
	if epu.mutation.AccountCleared() && len(epu.mutation.AccountIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "EventParticipant.account"`)
	}
	if epu.mutation.EventCleared() && len(epu.mutation.EventIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "EventParticipant.event"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (epu *EventParticipantUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *EventParticipantUpdate {
	epu.modifiers = append(epu.modifiers, modifiers...)
	return epu
}

func (epu *EventParticipantUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := epu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(eventparticipant.Table, eventparticipant.Columns, sqlgraph.NewFieldSpec(eventparticipant.FieldID, field.TypeString))
	if ps := epu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := epu.mutation.Role(); ok {
		_spec.SetField(eventparticipant.FieldRole, field.TypeString, value)
	}
	if value, ok := epu.mutation.Status(); ok {
		_spec.SetField(eventparticipant.FieldStatus, field.TypeString, value)
	}
	if epu.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   eventparticipant.AccountTable,
			Columns: []string{eventparticipant.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := epu.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   eventparticipant.AccountTable,
			Columns: []string{eventparticipant.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if epu.mutation.EventCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   eventparticipant.EventTable,
			Columns: []string{eventparticipant.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := epu.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   eventparticipant.EventTable,
			Columns: []string{eventparticipant.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(epu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, epu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{eventparticipant.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	epu.mutation.done = true
	return n, nil
}

// EventParticipantUpdateOne is the builder for updating a single EventParticipant entity.
type EventParticipantUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *EventParticipantMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetRole sets the "role" field.
func (epuo *EventParticipantUpdateOne) SetRole(s string) *EventParticipantUpdateOne {
	epuo.mutation.SetRole(s)
	return epuo
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (epuo *EventParticipantUpdateOne) SetNillableRole(s *string) *EventParticipantUpdateOne {
	if s != nil {
		epuo.SetRole(*s)
	}
	return epuo
}

// SetStatus sets the "status" field.
func (epuo *EventParticipantUpdateOne) SetStatus(s string) *EventParticipantUpdateOne {
	epuo.mutation.SetStatus(s)
	return epuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (epuo *EventParticipantUpdateOne) SetNillableStatus(s *string) *EventParticipantUpdateOne {
	if s != nil {
		epuo.SetStatus(*s)
	}
	return epuo
}

// SetAccountID sets the "account_id" field.
func (epuo *EventParticipantUpdateOne) SetAccountID(x xid.ID) *EventParticipantUpdateOne {
	epuo.mutation.SetAccountID(x)
	return epuo
}

// SetNillableAccountID sets the "account_id" field if the given value is not nil.
func (epuo *EventParticipantUpdateOne) SetNillableAccountID(x *xid.ID) *EventParticipantUpdateOne {
	if x != nil {
		epuo.SetAccountID(*x)
	}
	return epuo
}

// SetEventID sets the "event_id" field.
func (epuo *EventParticipantUpdateOne) SetEventID(x xid.ID) *EventParticipantUpdateOne {
	epuo.mutation.SetEventID(x)
	return epuo
}

// SetNillableEventID sets the "event_id" field if the given value is not nil.
func (epuo *EventParticipantUpdateOne) SetNillableEventID(x *xid.ID) *EventParticipantUpdateOne {
	if x != nil {
		epuo.SetEventID(*x)
	}
	return epuo
}

// SetAccount sets the "account" edge to the Account entity.
func (epuo *EventParticipantUpdateOne) SetAccount(a *Account) *EventParticipantUpdateOne {
	return epuo.SetAccountID(a.ID)
}

// SetEvent sets the "event" edge to the Event entity.
func (epuo *EventParticipantUpdateOne) SetEvent(e *Event) *EventParticipantUpdateOne {
	return epuo.SetEventID(e.ID)
}

// Mutation returns the EventParticipantMutation object of the builder.
func (epuo *EventParticipantUpdateOne) Mutation() *EventParticipantMutation {
	return epuo.mutation
}

// ClearAccount clears the "account" edge to the Account entity.
func (epuo *EventParticipantUpdateOne) ClearAccount() *EventParticipantUpdateOne {
	epuo.mutation.ClearAccount()
	return epuo
}

// ClearEvent clears the "event" edge to the Event entity.
func (epuo *EventParticipantUpdateOne) ClearEvent() *EventParticipantUpdateOne {
	epuo.mutation.ClearEvent()
	return epuo
}

// Where appends a list predicates to the EventParticipantUpdate builder.
func (epuo *EventParticipantUpdateOne) Where(ps ...predicate.EventParticipant) *EventParticipantUpdateOne {
	epuo.mutation.Where(ps...)
	return epuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (epuo *EventParticipantUpdateOne) Select(field string, fields ...string) *EventParticipantUpdateOne {
	epuo.fields = append([]string{field}, fields...)
	return epuo
}

// Save executes the query and returns the updated EventParticipant entity.
func (epuo *EventParticipantUpdateOne) Save(ctx context.Context) (*EventParticipant, error) {
	return withHooks(ctx, epuo.sqlSave, epuo.mutation, epuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (epuo *EventParticipantUpdateOne) SaveX(ctx context.Context) *EventParticipant {
	node, err := epuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (epuo *EventParticipantUpdateOne) Exec(ctx context.Context) error {
	_, err := epuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (epuo *EventParticipantUpdateOne) ExecX(ctx context.Context) {
	if err := epuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (epuo *EventParticipantUpdateOne) check() error {
	if epuo.mutation.AccountCleared() && len(epuo.mutation.AccountIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "EventParticipant.account"`)
	}
	if epuo.mutation.EventCleared() && len(epuo.mutation.EventIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "EventParticipant.event"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (epuo *EventParticipantUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *EventParticipantUpdateOne {
	epuo.modifiers = append(epuo.modifiers, modifiers...)
	return epuo
}

func (epuo *EventParticipantUpdateOne) sqlSave(ctx context.Context) (_node *EventParticipant, err error) {
	if err := epuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(eventparticipant.Table, eventparticipant.Columns, sqlgraph.NewFieldSpec(eventparticipant.FieldID, field.TypeString))
	id, ok := epuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "EventParticipant.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := epuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, eventparticipant.FieldID)
		for _, f := range fields {
			if !eventparticipant.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != eventparticipant.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := epuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := epuo.mutation.Role(); ok {
		_spec.SetField(eventparticipant.FieldRole, field.TypeString, value)
	}
	if value, ok := epuo.mutation.Status(); ok {
		_spec.SetField(eventparticipant.FieldStatus, field.TypeString, value)
	}
	if epuo.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   eventparticipant.AccountTable,
			Columns: []string{eventparticipant.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := epuo.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   eventparticipant.AccountTable,
			Columns: []string{eventparticipant.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if epuo.mutation.EventCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   eventparticipant.EventTable,
			Columns: []string{eventparticipant.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := epuo.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   eventparticipant.EventTable,
			Columns: []string{eventparticipant.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(epuo.modifiers...)
	_node = &EventParticipant{config: epuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, epuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{eventparticipant.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	epuo.mutation.done = true
	return _node, nil
}
