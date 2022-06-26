// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"main/ent/predicate"
	"main/ent/usdquote"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// USDQuoteUpdate is the builder for updating USDQuote entities.
type USDQuoteUpdate struct {
	config
	hooks    []Hook
	mutation *USDQuoteMutation
}

// Where appends a list predicates to the USDQuoteUpdate builder.
func (uqu *USDQuoteUpdate) Where(ps ...predicate.USDQuote) *USDQuoteUpdate {
	uqu.mutation.Where(ps...)
	return uqu
}

// SetPrice sets the "price" field.
func (uqu *USDQuoteUpdate) SetPrice(f float64) *USDQuoteUpdate {
	uqu.mutation.ResetPrice()
	uqu.mutation.SetPrice(f)
	return uqu
}

// AddPrice adds f to the "price" field.
func (uqu *USDQuoteUpdate) AddPrice(f float64) *USDQuoteUpdate {
	uqu.mutation.AddPrice(f)
	return uqu
}

// SetTimestamp sets the "Timestamp" field.
func (uqu *USDQuoteUpdate) SetTimestamp(t time.Time) *USDQuoteUpdate {
	uqu.mutation.SetTimestamp(t)
	return uqu
}

// Mutation returns the USDQuoteMutation object of the builder.
func (uqu *USDQuoteUpdate) Mutation() *USDQuoteMutation {
	return uqu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uqu *USDQuoteUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uqu.hooks) == 0 {
		affected, err = uqu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*USDQuoteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uqu.mutation = mutation
			affected, err = uqu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uqu.hooks) - 1; i >= 0; i-- {
			if uqu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uqu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uqu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uqu *USDQuoteUpdate) SaveX(ctx context.Context) int {
	affected, err := uqu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uqu *USDQuoteUpdate) Exec(ctx context.Context) error {
	_, err := uqu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uqu *USDQuoteUpdate) ExecX(ctx context.Context) {
	if err := uqu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uqu *USDQuoteUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   usdquote.Table,
			Columns: usdquote.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usdquote.FieldID,
			},
		},
	}
	if ps := uqu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uqu.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: usdquote.FieldPrice,
		})
	}
	if value, ok := uqu.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: usdquote.FieldPrice,
		})
	}
	if value, ok := uqu.mutation.Timestamp(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usdquote.FieldTimestamp,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uqu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usdquote.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// USDQuoteUpdateOne is the builder for updating a single USDQuote entity.
type USDQuoteUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *USDQuoteMutation
}

// SetPrice sets the "price" field.
func (uquo *USDQuoteUpdateOne) SetPrice(f float64) *USDQuoteUpdateOne {
	uquo.mutation.ResetPrice()
	uquo.mutation.SetPrice(f)
	return uquo
}

// AddPrice adds f to the "price" field.
func (uquo *USDQuoteUpdateOne) AddPrice(f float64) *USDQuoteUpdateOne {
	uquo.mutation.AddPrice(f)
	return uquo
}

// SetTimestamp sets the "Timestamp" field.
func (uquo *USDQuoteUpdateOne) SetTimestamp(t time.Time) *USDQuoteUpdateOne {
	uquo.mutation.SetTimestamp(t)
	return uquo
}

// Mutation returns the USDQuoteMutation object of the builder.
func (uquo *USDQuoteUpdateOne) Mutation() *USDQuoteMutation {
	return uquo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uquo *USDQuoteUpdateOne) Select(field string, fields ...string) *USDQuoteUpdateOne {
	uquo.fields = append([]string{field}, fields...)
	return uquo
}

// Save executes the query and returns the updated USDQuote entity.
func (uquo *USDQuoteUpdateOne) Save(ctx context.Context) (*USDQuote, error) {
	var (
		err  error
		node *USDQuote
	)
	if len(uquo.hooks) == 0 {
		node, err = uquo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*USDQuoteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uquo.mutation = mutation
			node, err = uquo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uquo.hooks) - 1; i >= 0; i-- {
			if uquo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uquo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uquo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uquo *USDQuoteUpdateOne) SaveX(ctx context.Context) *USDQuote {
	node, err := uquo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uquo *USDQuoteUpdateOne) Exec(ctx context.Context) error {
	_, err := uquo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uquo *USDQuoteUpdateOne) ExecX(ctx context.Context) {
	if err := uquo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uquo *USDQuoteUpdateOne) sqlSave(ctx context.Context) (_node *USDQuote, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   usdquote.Table,
			Columns: usdquote.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usdquote.FieldID,
			},
		},
	}
	id, ok := uquo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "USDQuote.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uquo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usdquote.FieldID)
		for _, f := range fields {
			if !usdquote.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != usdquote.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uquo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uquo.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: usdquote.FieldPrice,
		})
	}
	if value, ok := uquo.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: usdquote.FieldPrice,
		})
	}
	if value, ok := uquo.mutation.Timestamp(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usdquote.FieldTimestamp,
		})
	}
	_node = &USDQuote{config: uquo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uquo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usdquote.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}