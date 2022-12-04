// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/artificialinc/contrib/entproto/internal/entprototest/ent/nobackref"
	"github.com/artificialinc/contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NoBackrefDelete is the builder for deleting a NoBackref entity.
type NoBackrefDelete struct {
	config
	hooks    []Hook
	mutation *NoBackrefMutation
}

// Where appends a list predicates to the NoBackrefDelete builder.
func (nbd *NoBackrefDelete) Where(ps ...predicate.NoBackref) *NoBackrefDelete {
	nbd.mutation.Where(ps...)
	return nbd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (nbd *NoBackrefDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(nbd.hooks) == 0 {
		affected, err = nbd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NoBackrefMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			nbd.mutation = mutation
			affected, err = nbd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(nbd.hooks) - 1; i >= 0; i-- {
			if nbd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nbd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nbd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (nbd *NoBackrefDelete) ExecX(ctx context.Context) int {
	n, err := nbd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (nbd *NoBackrefDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: nobackref.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: nobackref.FieldID,
			},
		},
	}
	if ps := nbd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, nbd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// NoBackrefDeleteOne is the builder for deleting a single NoBackref entity.
type NoBackrefDeleteOne struct {
	nbd *NoBackrefDelete
}

// Exec executes the deletion query.
func (nbdo *NoBackrefDeleteOne) Exec(ctx context.Context) error {
	n, err := nbdo.nbd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{nobackref.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (nbdo *NoBackrefDeleteOne) ExecX(ctx context.Context) {
	nbdo.nbd.ExecX(ctx)
}
