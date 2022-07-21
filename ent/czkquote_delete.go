// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"main/ent/czkquote"
	"main/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CZKQuoteDelete is the builder for deleting a CZKQuote entity.
type CZKQuoteDelete struct {
	config
	hooks    []Hook
	mutation *CZKQuoteMutation
}

// Where appends a list predicates to the CZKQuoteDelete builder.
func (cqd *CZKQuoteDelete) Where(ps ...predicate.CZKQuote) *CZKQuoteDelete {
	cqd.mutation.Where(ps...)
	return cqd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cqd *CZKQuoteDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cqd.hooks) == 0 {
		affected, err = cqd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CZKQuoteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cqd.mutation = mutation
			affected, err = cqd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cqd.hooks) - 1; i >= 0; i-- {
			if cqd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cqd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cqd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (cqd *CZKQuoteDelete) ExecX(ctx context.Context) int {
	n, err := cqd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cqd *CZKQuoteDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: czkquote.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: czkquote.FieldID,
			},
		},
	}
	if ps := cqd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, cqd.driver, _spec)
}

// CZKQuoteDeleteOne is the builder for deleting a single CZKQuote entity.
type CZKQuoteDeleteOne struct {
	cqd *CZKQuoteDelete
}

// Exec executes the deletion query.
func (cqdo *CZKQuoteDeleteOne) Exec(ctx context.Context) error {
	n, err := cqdo.cqd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{czkquote.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cqdo *CZKQuoteDeleteOne) ExecX(ctx context.Context) {
	cqdo.cqd.ExecX(ctx)
}
