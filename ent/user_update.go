// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"main/ent/predicate"
	"main/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetTelegramId sets the "TelegramId" field.
func (uu *UserUpdate) SetTelegramId(i int64) *UserUpdate {
	uu.mutation.ResetTelegramId()
	uu.mutation.SetTelegramId(i)
	return uu
}

// AddTelegramId adds i to the "TelegramId" field.
func (uu *UserUpdate) AddTelegramId(i int64) *UserUpdate {
	uu.mutation.AddTelegramId(i)
	return uu
}

// SetActiveStable sets the "ActiveStable" field.
func (uu *UserUpdate) SetActiveStable(s string) *UserUpdate {
	uu.mutation.SetActiveStable(s)
	return uu
}

// SetNillableActiveStable sets the "ActiveStable" field if the given value is not nil.
func (uu *UserUpdate) SetNillableActiveStable(s *string) *UserUpdate {
	if s != nil {
		uu.SetActiveStable(*s)
	}
	return uu
}

// ClearActiveStable clears the value of the "ActiveStable" field.
func (uu *UserUpdate) ClearActiveStable() *UserUpdate {
	uu.mutation.ClearActiveStable()
	return uu
}

// SetAllStables sets the "AllStables" field.
func (uu *UserUpdate) SetAllStables(s string) *UserUpdate {
	uu.mutation.SetAllStables(s)
	return uu
}

// SetNillableAllStables sets the "AllStables" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAllStables(s *string) *UserUpdate {
	if s != nil {
		uu.SetAllStables(*s)
	}
	return uu
}

// ClearAllStables clears the value of the "AllStables" field.
func (uu *UserUpdate) ClearAllStables() *UserUpdate {
	uu.mutation.ClearAllStables()
	return uu
}

// SetRegTime sets the "RegTime" field.
func (uu *UserUpdate) SetRegTime(t time.Time) *UserUpdate {
	uu.mutation.SetRegTime(t)
	return uu
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uu.hooks) == 0 {
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			if uu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.TelegramId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: user.FieldTelegramId,
		})
	}
	if value, ok := uu.mutation.AddedTelegramId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: user.FieldTelegramId,
		})
	}
	if value, ok := uu.mutation.ActiveStable(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldActiveStable,
		})
	}
	if uu.mutation.ActiveStableCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldActiveStable,
		})
	}
	if value, ok := uu.mutation.AllStables(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldAllStables,
		})
	}
	if uu.mutation.AllStablesCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldAllStables,
		})
	}
	if value, ok := uu.mutation.RegTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldRegTime,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetTelegramId sets the "TelegramId" field.
func (uuo *UserUpdateOne) SetTelegramId(i int64) *UserUpdateOne {
	uuo.mutation.ResetTelegramId()
	uuo.mutation.SetTelegramId(i)
	return uuo
}

// AddTelegramId adds i to the "TelegramId" field.
func (uuo *UserUpdateOne) AddTelegramId(i int64) *UserUpdateOne {
	uuo.mutation.AddTelegramId(i)
	return uuo
}

// SetActiveStable sets the "ActiveStable" field.
func (uuo *UserUpdateOne) SetActiveStable(s string) *UserUpdateOne {
	uuo.mutation.SetActiveStable(s)
	return uuo
}

// SetNillableActiveStable sets the "ActiveStable" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableActiveStable(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetActiveStable(*s)
	}
	return uuo
}

// ClearActiveStable clears the value of the "ActiveStable" field.
func (uuo *UserUpdateOne) ClearActiveStable() *UserUpdateOne {
	uuo.mutation.ClearActiveStable()
	return uuo
}

// SetAllStables sets the "AllStables" field.
func (uuo *UserUpdateOne) SetAllStables(s string) *UserUpdateOne {
	uuo.mutation.SetAllStables(s)
	return uuo
}

// SetNillableAllStables sets the "AllStables" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAllStables(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetAllStables(*s)
	}
	return uuo
}

// ClearAllStables clears the value of the "AllStables" field.
func (uuo *UserUpdateOne) ClearAllStables() *UserUpdateOne {
	uuo.mutation.ClearAllStables()
	return uuo
}

// SetRegTime sets the "RegTime" field.
func (uuo *UserUpdateOne) SetRegTime(t time.Time) *UserUpdateOne {
	uuo.mutation.SetRegTime(t)
	return uuo
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	if len(uuo.hooks) == 0 {
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			if uuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.TelegramId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: user.FieldTelegramId,
		})
	}
	if value, ok := uuo.mutation.AddedTelegramId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: user.FieldTelegramId,
		})
	}
	if value, ok := uuo.mutation.ActiveStable(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldActiveStable,
		})
	}
	if uuo.mutation.ActiveStableCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldActiveStable,
		})
	}
	if value, ok := uuo.mutation.AllStables(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldAllStables,
		})
	}
	if uuo.mutation.AllStablesCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldAllStables,
		})
	}
	if value, ok := uuo.mutation.RegTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldRegTime,
		})
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
