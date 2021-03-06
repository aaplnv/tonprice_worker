// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"main/ent/tryquote"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TRYQuoteCreate is the builder for creating a TRYQuote entity.
type TRYQuoteCreate struct {
	config
	mutation *TRYQuoteMutation
	hooks    []Hook
}

// SetPrice sets the "price" field.
func (tqc *TRYQuoteCreate) SetPrice(f float64) *TRYQuoteCreate {
	tqc.mutation.SetPrice(f)
	return tqc
}

// SetTimestamp sets the "Timestamp" field.
func (tqc *TRYQuoteCreate) SetTimestamp(t time.Time) *TRYQuoteCreate {
	tqc.mutation.SetTimestamp(t)
	return tqc
}

// SetID sets the "id" field.
func (tqc *TRYQuoteCreate) SetID(i int) *TRYQuoteCreate {
	tqc.mutation.SetID(i)
	return tqc
}

// Mutation returns the TRYQuoteMutation object of the builder.
func (tqc *TRYQuoteCreate) Mutation() *TRYQuoteMutation {
	return tqc.mutation
}

// Save creates the TRYQuote in the database.
func (tqc *TRYQuoteCreate) Save(ctx context.Context) (*TRYQuote, error) {
	var (
		err  error
		node *TRYQuote
	)
	if len(tqc.hooks) == 0 {
		if err = tqc.check(); err != nil {
			return nil, err
		}
		node, err = tqc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TRYQuoteMutation)
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
func (tqc *TRYQuoteCreate) SaveX(ctx context.Context) *TRYQuote {
	v, err := tqc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tqc *TRYQuoteCreate) Exec(ctx context.Context) error {
	_, err := tqc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tqc *TRYQuoteCreate) ExecX(ctx context.Context) {
	if err := tqc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tqc *TRYQuoteCreate) check() error {
	if _, ok := tqc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`ent: missing required field "TRYQuote.price"`)}
	}
	if _, ok := tqc.mutation.Timestamp(); !ok {
		return &ValidationError{Name: "Timestamp", err: errors.New(`ent: missing required field "TRYQuote.Timestamp"`)}
	}
	return nil
}

func (tqc *TRYQuoteCreate) sqlSave(ctx context.Context) (*TRYQuote, error) {
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

func (tqc *TRYQuoteCreate) createSpec() (*TRYQuote, *sqlgraph.CreateSpec) {
	var (
		_node = &TRYQuote{config: tqc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: tryquote.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tryquote.FieldID,
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
			Column: tryquote.FieldPrice,
		})
		_node.Price = value
	}
	if value, ok := tqc.mutation.Timestamp(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tryquote.FieldTimestamp,
		})
		_node.Timestamp = value
	}
	return _node, _spec
}

// TRYQuoteCreateBulk is the builder for creating many TRYQuote entities in bulk.
type TRYQuoteCreateBulk struct {
	config
	builders []*TRYQuoteCreate
}

// Save creates the TRYQuote entities in the database.
func (tqcb *TRYQuoteCreateBulk) Save(ctx context.Context) ([]*TRYQuote, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tqcb.builders))
	nodes := make([]*TRYQuote, len(tqcb.builders))
	mutators := make([]Mutator, len(tqcb.builders))
	for i := range tqcb.builders {
		func(i int, root context.Context) {
			builder := tqcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TRYQuoteMutation)
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
func (tqcb *TRYQuoteCreateBulk) SaveX(ctx context.Context) []*TRYQuote {
	v, err := tqcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tqcb *TRYQuoteCreateBulk) Exec(ctx context.Context) error {
	_, err := tqcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tqcb *TRYQuoteCreateBulk) ExecX(ctx context.Context) {
	if err := tqcb.Exec(ctx); err != nil {
		panic(err)
	}
}
