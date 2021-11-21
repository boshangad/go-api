// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/boshangad/v1/ent/appoption"
	"github.com/boshangad/v1/ent/predicate"
)

// AppOptionDelete is the builder for deleting a AppOption entity.
type AppOptionDelete struct {
	config
	hooks    []Hook
	mutation *AppOptionMutation
}

// Where appends a list predicates to the AppOptionDelete builder.
func (aod *AppOptionDelete) Where(ps ...predicate.AppOption) *AppOptionDelete {
	aod.mutation.Where(ps...)
	return aod
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (aod *AppOptionDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(aod.hooks) == 0 {
		affected, err = aod.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppOptionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aod.mutation = mutation
			affected, err = aod.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(aod.hooks) - 1; i >= 0; i-- {
			if aod.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = aod.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aod.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (aod *AppOptionDelete) ExecX(ctx context.Context) int {
	n, err := aod.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (aod *AppOptionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: appoption.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: appoption.FieldID,
			},
		},
	}
	if ps := aod.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, aod.driver, _spec)
}

// AppOptionDeleteOne is the builder for deleting a single AppOption entity.
type AppOptionDeleteOne struct {
	aod *AppOptionDelete
}

// Exec executes the deletion query.
func (aodo *AppOptionDeleteOne) Exec(ctx context.Context) error {
	n, err := aodo.aod.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{appoption.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (aodo *AppOptionDeleteOne) ExecX(ctx context.Context) {
	aodo.aod.ExecX(ctx)
}
