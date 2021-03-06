// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"main/ent/inrquote"
	"main/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// INRQuoteDelete is the builder for deleting a INRQuote entity.
type INRQuoteDelete struct {
	config
	hooks    []Hook
	mutation *INRQuoteMutation
}

// Where appends a list predicates to the INRQuoteDelete builder.
func (iqd *INRQuoteDelete) Where(ps ...predicate.INRQuote) *INRQuoteDelete {
	iqd.mutation.Where(ps...)
	return iqd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (iqd *INRQuoteDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(iqd.hooks) == 0 {
		affected, err = iqd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*INRQuoteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			iqd.mutation = mutation
			affected, err = iqd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(iqd.hooks) - 1; i >= 0; i-- {
			if iqd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = iqd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iqd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (iqd *INRQuoteDelete) ExecX(ctx context.Context) int {
	n, err := iqd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (iqd *INRQuoteDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: inrquote.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: inrquote.FieldID,
			},
		},
	}
	if ps := iqd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, iqd.driver, _spec)
}

// INRQuoteDeleteOne is the builder for deleting a single INRQuote entity.
type INRQuoteDeleteOne struct {
	iqd *INRQuoteDelete
}

// Exec executes the deletion query.
func (iqdo *INRQuoteDeleteOne) Exec(ctx context.Context) error {
	n, err := iqdo.iqd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{inrquote.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (iqdo *INRQuoteDeleteOne) ExecX(ctx context.Context) {
	iqdo.iqd.ExecX(ctx)
}
