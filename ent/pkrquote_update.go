// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"main/ent/pkrquote"
	"main/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PKRQuoteUpdate is the builder for updating PKRQuote entities.
type PKRQuoteUpdate struct {
	config
	hooks    []Hook
	mutation *PKRQuoteMutation
}

// Where appends a list predicates to the PKRQuoteUpdate builder.
func (pqu *PKRQuoteUpdate) Where(ps ...predicate.PKRQuote) *PKRQuoteUpdate {
	pqu.mutation.Where(ps...)
	return pqu
}

// SetPrice sets the "price" field.
func (pqu *PKRQuoteUpdate) SetPrice(f float64) *PKRQuoteUpdate {
	pqu.mutation.ResetPrice()
	pqu.mutation.SetPrice(f)
	return pqu
}

// AddPrice adds f to the "price" field.
func (pqu *PKRQuoteUpdate) AddPrice(f float64) *PKRQuoteUpdate {
	pqu.mutation.AddPrice(f)
	return pqu
}

// SetTimestamp sets the "Timestamp" field.
func (pqu *PKRQuoteUpdate) SetTimestamp(t time.Time) *PKRQuoteUpdate {
	pqu.mutation.SetTimestamp(t)
	return pqu
}

// Mutation returns the PKRQuoteMutation object of the builder.
func (pqu *PKRQuoteUpdate) Mutation() *PKRQuoteMutation {
	return pqu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pqu *PKRQuoteUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pqu.hooks) == 0 {
		affected, err = pqu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PKRQuoteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pqu.mutation = mutation
			affected, err = pqu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pqu.hooks) - 1; i >= 0; i-- {
			if pqu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pqu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pqu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pqu *PKRQuoteUpdate) SaveX(ctx context.Context) int {
	affected, err := pqu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pqu *PKRQuoteUpdate) Exec(ctx context.Context) error {
	_, err := pqu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pqu *PKRQuoteUpdate) ExecX(ctx context.Context) {
	if err := pqu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pqu *PKRQuoteUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pkrquote.Table,
			Columns: pkrquote.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: pkrquote.FieldID,
			},
		},
	}
	if ps := pqu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pqu.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: pkrquote.FieldPrice,
		})
	}
	if value, ok := pqu.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: pkrquote.FieldPrice,
		})
	}
	if value, ok := pqu.mutation.Timestamp(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: pkrquote.FieldTimestamp,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pqu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pkrquote.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// PKRQuoteUpdateOne is the builder for updating a single PKRQuote entity.
type PKRQuoteUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PKRQuoteMutation
}

// SetPrice sets the "price" field.
func (pquo *PKRQuoteUpdateOne) SetPrice(f float64) *PKRQuoteUpdateOne {
	pquo.mutation.ResetPrice()
	pquo.mutation.SetPrice(f)
	return pquo
}

// AddPrice adds f to the "price" field.
func (pquo *PKRQuoteUpdateOne) AddPrice(f float64) *PKRQuoteUpdateOne {
	pquo.mutation.AddPrice(f)
	return pquo
}

// SetTimestamp sets the "Timestamp" field.
func (pquo *PKRQuoteUpdateOne) SetTimestamp(t time.Time) *PKRQuoteUpdateOne {
	pquo.mutation.SetTimestamp(t)
	return pquo
}

// Mutation returns the PKRQuoteMutation object of the builder.
func (pquo *PKRQuoteUpdateOne) Mutation() *PKRQuoteMutation {
	return pquo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pquo *PKRQuoteUpdateOne) Select(field string, fields ...string) *PKRQuoteUpdateOne {
	pquo.fields = append([]string{field}, fields...)
	return pquo
}

// Save executes the query and returns the updated PKRQuote entity.
func (pquo *PKRQuoteUpdateOne) Save(ctx context.Context) (*PKRQuote, error) {
	var (
		err  error
		node *PKRQuote
	)
	if len(pquo.hooks) == 0 {
		node, err = pquo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PKRQuoteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pquo.mutation = mutation
			node, err = pquo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pquo.hooks) - 1; i >= 0; i-- {
			if pquo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pquo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pquo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (pquo *PKRQuoteUpdateOne) SaveX(ctx context.Context) *PKRQuote {
	node, err := pquo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pquo *PKRQuoteUpdateOne) Exec(ctx context.Context) error {
	_, err := pquo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pquo *PKRQuoteUpdateOne) ExecX(ctx context.Context) {
	if err := pquo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pquo *PKRQuoteUpdateOne) sqlSave(ctx context.Context) (_node *PKRQuote, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pkrquote.Table,
			Columns: pkrquote.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: pkrquote.FieldID,
			},
		},
	}
	id, ok := pquo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PKRQuote.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pquo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pkrquote.FieldID)
		for _, f := range fields {
			if !pkrquote.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != pkrquote.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pquo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pquo.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: pkrquote.FieldPrice,
		})
	}
	if value, ok := pquo.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: pkrquote.FieldPrice,
		})
	}
	if value, ok := pquo.mutation.Timestamp(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: pkrquote.FieldTimestamp,
		})
	}
	_node = &PKRQuote{config: pquo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pquo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pkrquote.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
