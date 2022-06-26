// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"main/ent/plnquote"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PLNQuoteCreate is the builder for creating a PLNQuote entity.
type PLNQuoteCreate struct {
	config
	mutation *PLNQuoteMutation
	hooks    []Hook
}

// SetPrice sets the "price" field.
func (pqc *PLNQuoteCreate) SetPrice(f float64) *PLNQuoteCreate {
	pqc.mutation.SetPrice(f)
	return pqc
}

// SetTimestamp sets the "Timestamp" field.
func (pqc *PLNQuoteCreate) SetTimestamp(t time.Time) *PLNQuoteCreate {
	pqc.mutation.SetTimestamp(t)
	return pqc
}

// SetID sets the "id" field.
func (pqc *PLNQuoteCreate) SetID(i int) *PLNQuoteCreate {
	pqc.mutation.SetID(i)
	return pqc
}

// Mutation returns the PLNQuoteMutation object of the builder.
func (pqc *PLNQuoteCreate) Mutation() *PLNQuoteMutation {
	return pqc.mutation
}

// Save creates the PLNQuote in the database.
func (pqc *PLNQuoteCreate) Save(ctx context.Context) (*PLNQuote, error) {
	var (
		err  error
		node *PLNQuote
	)
	if len(pqc.hooks) == 0 {
		if err = pqc.check(); err != nil {
			return nil, err
		}
		node, err = pqc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PLNQuoteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pqc.check(); err != nil {
				return nil, err
			}
			pqc.mutation = mutation
			if node, err = pqc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pqc.hooks) - 1; i >= 0; i-- {
			if pqc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pqc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pqc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pqc *PLNQuoteCreate) SaveX(ctx context.Context) *PLNQuote {
	v, err := pqc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pqc *PLNQuoteCreate) Exec(ctx context.Context) error {
	_, err := pqc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pqc *PLNQuoteCreate) ExecX(ctx context.Context) {
	if err := pqc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pqc *PLNQuoteCreate) check() error {
	if _, ok := pqc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`ent: missing required field "PLNQuote.price"`)}
	}
	if _, ok := pqc.mutation.Timestamp(); !ok {
		return &ValidationError{Name: "Timestamp", err: errors.New(`ent: missing required field "PLNQuote.Timestamp"`)}
	}
	return nil
}

func (pqc *PLNQuoteCreate) sqlSave(ctx context.Context) (*PLNQuote, error) {
	_node, _spec := pqc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pqc.driver, _spec); err != nil {
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

func (pqc *PLNQuoteCreate) createSpec() (*PLNQuote, *sqlgraph.CreateSpec) {
	var (
		_node = &PLNQuote{config: pqc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: plnquote.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: plnquote.FieldID,
			},
		}
	)
	if id, ok := pqc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pqc.mutation.Price(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: plnquote.FieldPrice,
		})
		_node.Price = value
	}
	if value, ok := pqc.mutation.Timestamp(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: plnquote.FieldTimestamp,
		})
		_node.Timestamp = value
	}
	return _node, _spec
}

// PLNQuoteCreateBulk is the builder for creating many PLNQuote entities in bulk.
type PLNQuoteCreateBulk struct {
	config
	builders []*PLNQuoteCreate
}

// Save creates the PLNQuote entities in the database.
func (pqcb *PLNQuoteCreateBulk) Save(ctx context.Context) ([]*PLNQuote, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pqcb.builders))
	nodes := make([]*PLNQuote, len(pqcb.builders))
	mutators := make([]Mutator, len(pqcb.builders))
	for i := range pqcb.builders {
		func(i int, root context.Context) {
			builder := pqcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PLNQuoteMutation)
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
					_, err = mutators[i+1].Mutate(root, pqcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pqcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pqcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pqcb *PLNQuoteCreateBulk) SaveX(ctx context.Context) []*PLNQuote {
	v, err := pqcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pqcb *PLNQuoteCreateBulk) Exec(ctx context.Context) error {
	_, err := pqcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pqcb *PLNQuoteCreateBulk) ExecX(ctx context.Context) {
	if err := pqcb.Exec(ctx); err != nil {
		panic(err)
	}
}
