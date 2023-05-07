// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/oplog-middleware/pkg/db/ent/oplog"
	"github.com/NpoolPlatform/oplog-middleware/pkg/db/ent/predicate"
)

// OpLogDelete is the builder for deleting a OpLog entity.
type OpLogDelete struct {
	config
	hooks    []Hook
	mutation *OpLogMutation
}

// Where appends a list predicates to the OpLogDelete builder.
func (old *OpLogDelete) Where(ps ...predicate.OpLog) *OpLogDelete {
	old.mutation.Where(ps...)
	return old
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (old *OpLogDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(old.hooks) == 0 {
		affected, err = old.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OpLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			old.mutation = mutation
			affected, err = old.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(old.hooks) - 1; i >= 0; i-- {
			if old.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = old.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, old.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (old *OpLogDelete) ExecX(ctx context.Context) int {
	n, err := old.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (old *OpLogDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: oplog.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: oplog.FieldID,
			},
		},
	}
	if ps := old.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, old.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// OpLogDeleteOne is the builder for deleting a single OpLog entity.
type OpLogDeleteOne struct {
	old *OpLogDelete
}

// Exec executes the deletion query.
func (oldo *OpLogDeleteOne) Exec(ctx context.Context) error {
	n, err := oldo.old.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{oplog.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (oldo *OpLogDeleteOne) ExecX(ctx context.Context) {
	oldo.old.ExecX(ctx)
}