// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"main/ent/euroquote"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EUROQuoteCreate is the builder for creating a EUROQuote entity.
type EUROQuoteCreate struct {
	config
	mutation *EUROQuoteMutation
	hooks    []Hook
}

// SetPrice sets the "price" field.
func (eqc *EUROQuoteCreate) SetPrice(f float64) *EUROQuoteCreate {
	eqc.mutation.SetPrice(f)
	return eqc
}

// SetTimestamp sets the "Timestamp" field.
func (eqc *EUROQuoteCreate) SetTimestamp(t time.Time) *EUROQuoteCreate {
	eqc.mutation.SetTimestamp(t)
	return eqc
}

// SetID sets the "id" field.
func (eqc *EUROQuoteCreate) SetID(i int) *EUROQuoteCreate {
	eqc.mutation.SetID(i)
	return eqc
}

// Mutation returns the EUROQuoteMutation object of the builder.
func (eqc *EUROQuoteCreate) Mutation() *EUROQuoteMutation {
	return eqc.mutation
}

// Save creates the EUROQuote in the database.
func (eqc *EUROQuoteCreate) Save(ctx context.Context) (*EUROQuote, error) {
	var (
		err  error
		node *EUROQuote
	)
	if len(eqc.hooks) == 0 {
		if err = eqc.check(); err != nil {
			return nil, err
		}
		node, err = eqc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EUROQuoteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = eqc.check(); err != nil {
				return nil, err
			}
			eqc.mutation = mutation
			if node, err = eqc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(eqc.hooks) - 1; i >= 0; i-- {
			if eqc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = eqc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, eqc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (eqc *EUROQuoteCreate) SaveX(ctx context.Context) *EUROQuote {
	v, err := eqc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (eqc *EUROQuoteCreate) Exec(ctx context.Context) error {
	_, err := eqc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eqc *EUROQuoteCreate) ExecX(ctx context.Context) {
	if err := eqc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eqc *EUROQuoteCreate) check() error {
	if _, ok := eqc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`ent: missing required field "EUROQuote.price"`)}
	}
	if _, ok := eqc.mutation.Timestamp(); !ok {
		return &ValidationError{Name: "Timestamp", err: errors.New(`ent: missing required field "EUROQuote.Timestamp"`)}
	}
	return nil
}

func (eqc *EUROQuoteCreate) sqlSave(ctx context.Context) (*EUROQuote, error) {
	_node, _spec := eqc.createSpec()
	if err := sqlgraph.CreateNode(ctx, eqc.driver, _spec); err != nil {
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

func (eqc *EUROQuoteCreate) createSpec() (*EUROQuote, *sqlgraph.CreateSpec) {
	var (
		_node = &EUROQuote{config: eqc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: euroquote.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: euroquote.FieldID,
			},
		}
	)
	if id, ok := eqc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := eqc.mutation.Price(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: euroquote.FieldPrice,
		})
		_node.Price = value
	}
	if value, ok := eqc.mutation.Timestamp(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: euroquote.FieldTimestamp,
		})
		_node.Timestamp = value
	}
	return _node, _spec
}

// EUROQuoteCreateBulk is the builder for creating many EUROQuote entities in bulk.
type EUROQuoteCreateBulk struct {
	config
	builders []*EUROQuoteCreate
}

// Save creates the EUROQuote entities in the database.
func (eqcb *EUROQuoteCreateBulk) Save(ctx context.Context) ([]*EUROQuote, error) {
	specs := make([]*sqlgraph.CreateSpec, len(eqcb.builders))
	nodes := make([]*EUROQuote, len(eqcb.builders))
	mutators := make([]Mutator, len(eqcb.builders))
	for i := range eqcb.builders {
		func(i int, root context.Context) {
			builder := eqcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EUROQuoteMutation)
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
					_, err = mutators[i+1].Mutate(root, eqcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, eqcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, eqcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (eqcb *EUROQuoteCreateBulk) SaveX(ctx context.Context) []*EUROQuote {
	v, err := eqcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (eqcb *EUROQuoteCreateBulk) Exec(ctx context.Context) error {
	_, err := eqcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eqcb *EUROQuoteCreateBulk) ExecX(ctx context.Context) {
	if err := eqcb.Exec(ctx); err != nil {
		panic(err)
	}
}