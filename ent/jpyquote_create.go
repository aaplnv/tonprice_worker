// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"main/ent/jpyquote"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// JPYQuoteCreate is the builder for creating a JPYQuote entity.
type JPYQuoteCreate struct {
	config
	mutation *JPYQuoteMutation
	hooks    []Hook
}

// SetPrice sets the "price" field.
func (jqc *JPYQuoteCreate) SetPrice(f float64) *JPYQuoteCreate {
	jqc.mutation.SetPrice(f)
	return jqc
}

// SetTimestamp sets the "Timestamp" field.
func (jqc *JPYQuoteCreate) SetTimestamp(t time.Time) *JPYQuoteCreate {
	jqc.mutation.SetTimestamp(t)
	return jqc
}

// SetID sets the "id" field.
func (jqc *JPYQuoteCreate) SetID(i int) *JPYQuoteCreate {
	jqc.mutation.SetID(i)
	return jqc
}

// Mutation returns the JPYQuoteMutation object of the builder.
func (jqc *JPYQuoteCreate) Mutation() *JPYQuoteMutation {
	return jqc.mutation
}

// Save creates the JPYQuote in the database.
func (jqc *JPYQuoteCreate) Save(ctx context.Context) (*JPYQuote, error) {
	var (
		err  error
		node *JPYQuote
	)
	if len(jqc.hooks) == 0 {
		if err = jqc.check(); err != nil {
			return nil, err
		}
		node, err = jqc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*JPYQuoteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = jqc.check(); err != nil {
				return nil, err
			}
			jqc.mutation = mutation
			if node, err = jqc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(jqc.hooks) - 1; i >= 0; i-- {
			if jqc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = jqc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, jqc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (jqc *JPYQuoteCreate) SaveX(ctx context.Context) *JPYQuote {
	v, err := jqc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (jqc *JPYQuoteCreate) Exec(ctx context.Context) error {
	_, err := jqc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jqc *JPYQuoteCreate) ExecX(ctx context.Context) {
	if err := jqc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (jqc *JPYQuoteCreate) check() error {
	if _, ok := jqc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`ent: missing required field "JPYQuote.price"`)}
	}
	if _, ok := jqc.mutation.Timestamp(); !ok {
		return &ValidationError{Name: "Timestamp", err: errors.New(`ent: missing required field "JPYQuote.Timestamp"`)}
	}
	return nil
}

func (jqc *JPYQuoteCreate) sqlSave(ctx context.Context) (*JPYQuote, error) {
	_node, _spec := jqc.createSpec()
	if err := sqlgraph.CreateNode(ctx, jqc.driver, _spec); err != nil {
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

func (jqc *JPYQuoteCreate) createSpec() (*JPYQuote, *sqlgraph.CreateSpec) {
	var (
		_node = &JPYQuote{config: jqc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: jpyquote.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: jpyquote.FieldID,
			},
		}
	)
	if id, ok := jqc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := jqc.mutation.Price(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: jpyquote.FieldPrice,
		})
		_node.Price = value
	}
	if value, ok := jqc.mutation.Timestamp(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: jpyquote.FieldTimestamp,
		})
		_node.Timestamp = value
	}
	return _node, _spec
}

// JPYQuoteCreateBulk is the builder for creating many JPYQuote entities in bulk.
type JPYQuoteCreateBulk struct {
	config
	builders []*JPYQuoteCreate
}

// Save creates the JPYQuote entities in the database.
func (jqcb *JPYQuoteCreateBulk) Save(ctx context.Context) ([]*JPYQuote, error) {
	specs := make([]*sqlgraph.CreateSpec, len(jqcb.builders))
	nodes := make([]*JPYQuote, len(jqcb.builders))
	mutators := make([]Mutator, len(jqcb.builders))
	for i := range jqcb.builders {
		func(i int, root context.Context) {
			builder := jqcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*JPYQuoteMutation)
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
					_, err = mutators[i+1].Mutate(root, jqcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, jqcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, jqcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (jqcb *JPYQuoteCreateBulk) SaveX(ctx context.Context) []*JPYQuote {
	v, err := jqcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (jqcb *JPYQuoteCreateBulk) Exec(ctx context.Context) error {
	_, err := jqcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jqcb *JPYQuoteCreateBulk) ExecX(ctx context.Context) {
	if err := jqcb.Exec(ctx); err != nil {
		panic(err)
	}
}
