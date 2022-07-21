// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"main/ent/nokquote"
	"main/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NOKQuoteDelete is the builder for deleting a NOKQuote entity.
type NOKQuoteDelete struct {
	config
	hooks    []Hook
	mutation *NOKQuoteMutation
}

// Where appends a list predicates to the NOKQuoteDelete builder.
func (nqd *NOKQuoteDelete) Where(ps ...predicate.NOKQuote) *NOKQuoteDelete {
	nqd.mutation.Where(ps...)
	return nqd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (nqd *NOKQuoteDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(nqd.hooks) == 0 {
		affected, err = nqd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NOKQuoteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			nqd.mutation = mutation
			affected, err = nqd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(nqd.hooks) - 1; i >= 0; i-- {
			if nqd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nqd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nqd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (nqd *NOKQuoteDelete) ExecX(ctx context.Context) int {
	n, err := nqd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (nqd *NOKQuoteDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: nokquote.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: nokquote.FieldID,
			},
		},
	}
	if ps := nqd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, nqd.driver, _spec)
}

// NOKQuoteDeleteOne is the builder for deleting a single NOKQuote entity.
type NOKQuoteDeleteOne struct {
	nqd *NOKQuoteDelete
}

// Exec executes the deletion query.
func (nqdo *NOKQuoteDeleteOne) Exec(ctx context.Context) error {
	n, err := nqdo.nqd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{nokquote.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (nqdo *NOKQuoteDeleteOne) ExecX(ctx context.Context) {
	nqdo.nqd.ExecX(ctx)
}
