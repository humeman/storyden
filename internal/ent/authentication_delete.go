// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Southclaws/storyden/internal/ent/authentication"
	"github.com/Southclaws/storyden/internal/ent/predicate"
)

// AuthenticationDelete is the builder for deleting a Authentication entity.
type AuthenticationDelete struct {
	config
	hooks    []Hook
	mutation *AuthenticationMutation
}

// Where appends a list predicates to the AuthenticationDelete builder.
func (ad *AuthenticationDelete) Where(ps ...predicate.Authentication) *AuthenticationDelete {
	ad.mutation.Where(ps...)
	return ad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ad *AuthenticationDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, AuthenticationMutation](ctx, ad.sqlExec, ad.mutation, ad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ad *AuthenticationDelete) ExecX(ctx context.Context) int {
	n, err := ad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ad *AuthenticationDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: authentication.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: authentication.FieldID,
			},
		},
	}
	if ps := ad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ad.mutation.done = true
	return affected, err
}

// AuthenticationDeleteOne is the builder for deleting a single Authentication entity.
type AuthenticationDeleteOne struct {
	ad *AuthenticationDelete
}

// Exec executes the deletion query.
func (ado *AuthenticationDeleteOne) Exec(ctx context.Context) error {
	n, err := ado.ad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{authentication.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ado *AuthenticationDeleteOne) ExecX(ctx context.Context) {
	ado.ad.ExecX(ctx)
}
