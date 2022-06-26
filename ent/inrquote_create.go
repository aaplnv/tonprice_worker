// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"main/ent/inrquote"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// INRQuoteCreate is the builder for creating a INRQuote entity.
type INRQuoteCreate struct {
	config
	mutation *INRQuoteMutation
	hooks    []Hook
}

// SetPrice sets the "price" field.
func (iqc *INRQuoteCreate) SetPrice(f float64) *INRQuoteCreate {
	iqc.mutation.SetPrice(f)
	return iqc
}

// SetTimestamp sets the "Timestamp" field.
func (iqc *INRQuoteCreate) SetTimestamp(t time.Time) *INRQuoteCreate {
	iqc.mutation.SetTimestamp(t)
	return iqc
}

// SetID sets the "id" field.
func (iqc *INRQuoteCreate) SetID(i int) *INRQuoteCreate {
	iqc.mutation.SetID(i)
	return iqc
}

// Mutation returns the INRQuoteMutation object of the builder.
func (iqc *INRQuoteCreate) Mutation() *INRQuoteMutation {
	return iqc.mutation
}

// Save creates the INRQuote in the database.
func (iqc *INRQuoteCreate) Save(ctx context.Context) (*INRQuote, error) {
	var (
		err  error
		node *INRQuote
	)
	if len(iqc.hooks) == 0 {
		if err = iqc.check(); err != nil {
			return nil, err
		}
		node, err = iqc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*INRQuoteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = iqc.check(); err != nil {
				return nil, err
			}
			iqc.mutation = mutation
			if node, err = iqc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(iqc.hooks) - 1; i >= 0; i-- {
			if iqc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = iqc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iqc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (iqc *INRQuoteCreate) SaveX(ctx context.Context) *INRQuote {
	v, err := iqc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (iqc *INRQuoteCreate) Exec(ctx context.Context) error {
	_, err := iqc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iqc *INRQuoteCreate) ExecX(ctx context.Context) {
	if err := iqc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iqc *INRQuoteCreate) check() error {
	if _, ok := iqc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`ent: missing required field "INRQuote.price"`)}
	}
	if _, ok := iqc.mutation.Timestamp(); !ok {
		return &ValidationError{Name: "Timestamp", err: errors.New(`ent: missing required field "INRQuote.Timestamp"`)}
	}
	return nil
}

func (iqc *INRQuoteCreate) sqlSave(ctx context.Context) (*INRQuote, error) {
	_node, _spec := iqc.createSpec()
	if err := sqlgraph.CreateNode(ctx, iqc.driver, _spec); err != nil {
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

func (iqc *INRQuoteCreate) createSpec() (*INRQuote, *sqlgraph.CreateSpec) {
	var (
		_node = &INRQuote{config: iqc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: inrquote.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: inrquote.FieldID,
			},
		}
	)
	if id, ok := iqc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := iqc.mutation.Price(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: inrquote.FieldPrice,
		})
		_node.Price = value
	}
	if value, ok := iqc.mutation.Timestamp(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: inrquote.FieldTimestamp,
		})
		_node.Timestamp = value
	}
	return _node, _spec
}

// INRQuoteCreateBulk is the builder for creating many INRQuote entities in bulk.
type INRQuoteCreateBulk struct {
	config
	builders []*INRQuoteCreate
}

// Save creates the INRQuote entities in the database.
func (iqcb *INRQuoteCreateBulk) Save(ctx context.Context) ([]*INRQuote, error) {
	specs := make([]*sqlgraph.CreateSpec, len(iqcb.builders))
	nodes := make([]*INRQuote, len(iqcb.builders))
	mutators := make([]Mutator, len(iqcb.builders))
	for i := range iqcb.builders {
		func(i int, root context.Context) {
			builder := iqcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*INRQuoteMutation)
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
					_, err = mutators[i+1].Mutate(root, iqcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, iqcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, iqcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (iqcb *INRQuoteCreateBulk) SaveX(ctx context.Context) []*INRQuote {
	v, err := iqcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (iqcb *INRQuoteCreateBulk) Exec(ctx context.Context) error {
	_, err := iqcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iqcb *INRQuoteCreateBulk) ExecX(ctx context.Context) {
	if err := iqcb.Exec(ctx); err != nil {
		panic(err)
	}
}