// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"main/ent/twdquote"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TWDQuoteCreate is the builder for creating a TWDQuote entity.
type TWDQuoteCreate struct {
	config
	mutation *TWDQuoteMutation
	hooks    []Hook
}

// SetPrice sets the "price" field.
func (tqc *TWDQuoteCreate) SetPrice(f float64) *TWDQuoteCreate {
	tqc.mutation.SetPrice(f)
	return tqc
}

// SetTimestamp sets the "Timestamp" field.
func (tqc *TWDQuoteCreate) SetTimestamp(t time.Time) *TWDQuoteCreate {
	tqc.mutation.SetTimestamp(t)
	return tqc
}

// SetID sets the "id" field.
func (tqc *TWDQuoteCreate) SetID(i int) *TWDQuoteCreate {
	tqc.mutation.SetID(i)
	return tqc
}

// Mutation returns the TWDQuoteMutation object of the builder.
func (tqc *TWDQuoteCreate) Mutation() *TWDQuoteMutation {
	return tqc.mutation
}

// Save creates the TWDQuote in the database.
func (tqc *TWDQuoteCreate) Save(ctx context.Context) (*TWDQuote, error) {
	var (
		err  error
		node *TWDQuote
	)
	if len(tqc.hooks) == 0 {
		if err = tqc.check(); err != nil {
			return nil, err
		}
		node, err = tqc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TWDQuoteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tqc.check(); err != nil {
				return nil, err
			}
			tqc.mutation = mutation
			if node, err = tqc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tqc.hooks) - 1; i >= 0; i-- {
			if tqc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tqc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tqc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tqc *TWDQuoteCreate) SaveX(ctx context.Context) *TWDQuote {
	v, err := tqc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tqc *TWDQuoteCreate) Exec(ctx context.Context) error {
	_, err := tqc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tqc *TWDQuoteCreate) ExecX(ctx context.Context) {
	if err := tqc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tqc *TWDQuoteCreate) check() error {
	if _, ok := tqc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`ent: missing required field "TWDQuote.price"`)}
	}
	if _, ok := tqc.mutation.Timestamp(); !ok {
		return &ValidationError{Name: "Timestamp", err: errors.New(`ent: missing required field "TWDQuote.Timestamp"`)}
	}
	return nil
}

func (tqc *TWDQuoteCreate) sqlSave(ctx context.Context) (*TWDQuote, error) {
	_node, _spec := tqc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tqc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	return _node, nil
}

func (tqc *TWDQuoteCreate) createSpec() (*TWDQuote, *sqlgraph.CreateSpec) {
	var (
		_node = &TWDQuote{config: tqc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: twdquote.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: twdquote.FieldID,
			},
		}
	)
	if id, ok := tqc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tqc.mutation.Price(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: twdquote.FieldPrice,
		})
		_node.Price = value
	}
	if value, ok := tqc.mutation.Timestamp(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: twdquote.FieldTimestamp,
		})
		_node.Timestamp = value
	}
	return _node, _spec
}

// TWDQuoteCreateBulk is the builder for creating many TWDQuote entities in bulk.
type TWDQuoteCreateBulk struct {
	config
	builders []*TWDQuoteCreate
}

// Save creates the TWDQuote entities in the database.
func (tqcb *TWDQuoteCreateBulk) Save(ctx context.Context) ([]*TWDQuote, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tqcb.builders))
	nodes := make([]*TWDQuote, len(tqcb.builders))
	mutators := make([]Mutator, len(tqcb.builders))
	for i := range tqcb.builders {
		func(i int, root context.Context) {
			builder := tqcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TWDQuoteMutation)
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
					_, err = mutators[i+1].Mutate(root, tqcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tqcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tqcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tqcb *TWDQuoteCreateBulk) SaveX(ctx context.Context) []*TWDQuote {
	v, err := tqcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tqcb *TWDQuoteCreateBulk) Exec(ctx context.Context) error {
	_, err := tqcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tqcb *TWDQuoteCreateBulk) ExecX(ctx context.Context) {
	if err := tqcb.Exec(ctx); err != nil {
		panic(err)
	}
}
