// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"main/ent/nzdquote"
	"main/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NZDQuoteUpdate is the builder for updating NZDQuote entities.
type NZDQuoteUpdate struct {
	config
	hooks    []Hook
	mutation *NZDQuoteMutation
}

// Where appends a list predicates to the NZDQuoteUpdate builder.
func (nqu *NZDQuoteUpdate) Where(ps ...predicate.NZDQuote) *NZDQuoteUpdate {
	nqu.mutation.Where(ps...)
	return nqu
}

// SetPrice sets the "price" field.
func (nqu *NZDQuoteUpdate) SetPrice(f float64) *NZDQuoteUpdate {
	nqu.mutation.ResetPrice()
	nqu.mutation.SetPrice(f)
	return nqu
}

// AddPrice adds f to the "price" field.
func (nqu *NZDQuoteUpdate) AddPrice(f float64) *NZDQuoteUpdate {
	nqu.mutation.AddPrice(f)
	return nqu
}

// SetTimestamp sets the "Timestamp" field.
func (nqu *NZDQuoteUpdate) SetTimestamp(t time.Time) *NZDQuoteUpdate {
	nqu.mutation.SetTimestamp(t)
	return nqu
}

// Mutation returns the NZDQuoteMutation object of the builder.
func (nqu *NZDQuoteUpdate) Mutation() *NZDQuoteMutation {
	return nqu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nqu *NZDQuoteUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(nqu.hooks) == 0 {
		affected, err = nqu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NZDQuoteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			nqu.mutation = mutation
			affected, err = nqu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(nqu.hooks) - 1; i >= 0; i-- {
			if nqu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nqu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nqu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (nqu *NZDQuoteUpdate) SaveX(ctx context.Context) int {
	affected, err := nqu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nqu *NZDQuoteUpdate) Exec(ctx context.Context) error {
	_, err := nqu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nqu *NZDQuoteUpdate) ExecX(ctx context.Context) {
	if err := nqu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nqu *NZDQuoteUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   nzdquote.Table,
			Columns: nzdquote.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: nzdquote.FieldID,
			},
		},
	}
	if ps := nqu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nqu.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: nzdquote.FieldPrice,
		})
	}
	if value, ok := nqu.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: nzdquote.FieldPrice,
		})
	}
	if value, ok := nqu.mutation.Timestamp(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: nzdquote.FieldTimestamp,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nqu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{nzdquote.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// NZDQuoteUpdateOne is the builder for updating a single NZDQuote entity.
type NZDQuoteUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NZDQuoteMutation
}

// SetPrice sets the "price" field.
func (nquo *NZDQuoteUpdateOne) SetPrice(f float64) *NZDQuoteUpdateOne {
	nquo.mutation.ResetPrice()
	nquo.mutation.SetPrice(f)
	return nquo
}

// AddPrice adds f to the "price" field.
func (nquo *NZDQuoteUpdateOne) AddPrice(f float64) *NZDQuoteUpdateOne {
	nquo.mutation.AddPrice(f)
	return nquo
}

// SetTimestamp sets the "Timestamp" field.
func (nquo *NZDQuoteUpdateOne) SetTimestamp(t time.Time) *NZDQuoteUpdateOne {
	nquo.mutation.SetTimestamp(t)
	return nquo
}

// Mutation returns the NZDQuoteMutation object of the builder.
func (nquo *NZDQuoteUpdateOne) Mutation() *NZDQuoteMutation {
	return nquo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nquo *NZDQuoteUpdateOne) Select(field string, fields ...string) *NZDQuoteUpdateOne {
	nquo.fields = append([]string{field}, fields...)
	return nquo
}

// Save executes the query and returns the updated NZDQuote entity.
func (nquo *NZDQuoteUpdateOne) Save(ctx context.Context) (*NZDQuote, error) {
	var (
		err  error
		node *NZDQuote
	)
	if len(nquo.hooks) == 0 {
		node, err = nquo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NZDQuoteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			nquo.mutation = mutation
			node, err = nquo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(nquo.hooks) - 1; i >= 0; i-- {
			if nquo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nquo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nquo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (nquo *NZDQuoteUpdateOne) SaveX(ctx context.Context) *NZDQuote {
	node, err := nquo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nquo *NZDQuoteUpdateOne) Exec(ctx context.Context) error {
	_, err := nquo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nquo *NZDQuoteUpdateOne) ExecX(ctx context.Context) {
	if err := nquo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nquo *NZDQuoteUpdateOne) sqlSave(ctx context.Context) (_node *NZDQuote, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   nzdquote.Table,
			Columns: nzdquote.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: nzdquote.FieldID,
			},
		},
	}
	id, ok := nquo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "NZDQuote.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nquo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, nzdquote.FieldID)
		for _, f := range fields {
			if !nzdquote.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != nzdquote.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nquo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nquo.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: nzdquote.FieldPrice,
		})
	}
	if value, ok := nquo.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: nzdquote.FieldPrice,
		})
	}
	if value, ok := nquo.mutation.Timestamp(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: nzdquote.FieldTimestamp,
		})
	}
	_node = &NZDQuote{config: nquo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nquo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{nzdquote.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
