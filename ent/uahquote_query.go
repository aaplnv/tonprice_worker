// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"main/ent/predicate"
	"main/ent/uahquote"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UAHQuoteQuery is the builder for querying UAHQuote entities.
type UAHQuoteQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.UAHQuote
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UAHQuoteQuery builder.
func (uqq *UAHQuoteQuery) Where(ps ...predicate.UAHQuote) *UAHQuoteQuery {
	uqq.predicates = append(uqq.predicates, ps...)
	return uqq
}

// Limit adds a limit step to the query.
func (uqq *UAHQuoteQuery) Limit(limit int) *UAHQuoteQuery {
	uqq.limit = &limit
	return uqq
}

// Offset adds an offset step to the query.
func (uqq *UAHQuoteQuery) Offset(offset int) *UAHQuoteQuery {
	uqq.offset = &offset
	return uqq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (uqq *UAHQuoteQuery) Unique(unique bool) *UAHQuoteQuery {
	uqq.unique = &unique
	return uqq
}

// Order adds an order step to the query.
func (uqq *UAHQuoteQuery) Order(o ...OrderFunc) *UAHQuoteQuery {
	uqq.order = append(uqq.order, o...)
	return uqq
}

// First returns the first UAHQuote entity from the query.
// Returns a *NotFoundError when no UAHQuote was found.
func (uqq *UAHQuoteQuery) First(ctx context.Context) (*UAHQuote, error) {
	nodes, err := uqq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{uahquote.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uqq *UAHQuoteQuery) FirstX(ctx context.Context) *UAHQuote {
	node, err := uqq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UAHQuote ID from the query.
// Returns a *NotFoundError when no UAHQuote ID was found.
func (uqq *UAHQuoteQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uqq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{uahquote.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (uqq *UAHQuoteQuery) FirstIDX(ctx context.Context) int {
	id, err := uqq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UAHQuote entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UAHQuote entity is found.
// Returns a *NotFoundError when no UAHQuote entities are found.
func (uqq *UAHQuoteQuery) Only(ctx context.Context) (*UAHQuote, error) {
	nodes, err := uqq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{uahquote.Label}
	default:
		return nil, &NotSingularError{uahquote.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uqq *UAHQuoteQuery) OnlyX(ctx context.Context) *UAHQuote {
	node, err := uqq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UAHQuote ID in the query.
// Returns a *NotSingularError when more than one UAHQuote ID is found.
// Returns a *NotFoundError when no entities are found.
func (uqq *UAHQuoteQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uqq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{uahquote.Label}
	default:
		err = &NotSingularError{uahquote.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (uqq *UAHQuoteQuery) OnlyIDX(ctx context.Context) int {
	id, err := uqq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UAHQuotes.
func (uqq *UAHQuoteQuery) All(ctx context.Context) ([]*UAHQuote, error) {
	if err := uqq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return uqq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (uqq *UAHQuoteQuery) AllX(ctx context.Context) []*UAHQuote {
	nodes, err := uqq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UAHQuote IDs.
func (uqq *UAHQuoteQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := uqq.Select(uahquote.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uqq *UAHQuoteQuery) IDsX(ctx context.Context) []int {
	ids, err := uqq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uqq *UAHQuoteQuery) Count(ctx context.Context) (int, error) {
	if err := uqq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return uqq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (uqq *UAHQuoteQuery) CountX(ctx context.Context) int {
	count, err := uqq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uqq *UAHQuoteQuery) Exist(ctx context.Context) (bool, error) {
	if err := uqq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return uqq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (uqq *UAHQuoteQuery) ExistX(ctx context.Context) bool {
	exist, err := uqq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UAHQuoteQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uqq *UAHQuoteQuery) Clone() *UAHQuoteQuery {
	if uqq == nil {
		return nil
	}
	return &UAHQuoteQuery{
		config:     uqq.config,
		limit:      uqq.limit,
		offset:     uqq.offset,
		order:      append([]OrderFunc{}, uqq.order...),
		predicates: append([]predicate.UAHQuote{}, uqq.predicates...),
		// clone intermediate query.
		sql:    uqq.sql.Clone(),
		path:   uqq.path,
		unique: uqq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Price float64 `json:"price,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UAHQuote.Query().
//		GroupBy(uahquote.FieldPrice).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (uqq *UAHQuoteQuery) GroupBy(field string, fields ...string) *UAHQuoteGroupBy {
	group := &UAHQuoteGroupBy{config: uqq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := uqq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return uqq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Price float64 `json:"price,omitempty"`
//	}
//
//	client.UAHQuote.Query().
//		Select(uahquote.FieldPrice).
//		Scan(ctx, &v)
//
func (uqq *UAHQuoteQuery) Select(fields ...string) *UAHQuoteSelect {
	uqq.fields = append(uqq.fields, fields...)
	return &UAHQuoteSelect{UAHQuoteQuery: uqq}
}

func (uqq *UAHQuoteQuery) prepareQuery(ctx context.Context) error {
	for _, f := range uqq.fields {
		if !uahquote.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if uqq.path != nil {
		prev, err := uqq.path(ctx)
		if err != nil {
			return err
		}
		uqq.sql = prev
	}
	return nil
}

func (uqq *UAHQuoteQuery) sqlAll(ctx context.Context) ([]*UAHQuote, error) {
	var (
		nodes = []*UAHQuote{}
		_spec = uqq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &UAHQuote{config: uqq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, uqq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (uqq *UAHQuoteQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uqq.querySpec()
	_spec.Node.Columns = uqq.fields
	if len(uqq.fields) > 0 {
		_spec.Unique = uqq.unique != nil && *uqq.unique
	}
	return sqlgraph.CountNodes(ctx, uqq.driver, _spec)
}

func (uqq *UAHQuoteQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := uqq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (uqq *UAHQuoteQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   uahquote.Table,
			Columns: uahquote.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: uahquote.FieldID,
			},
		},
		From:   uqq.sql,
		Unique: true,
	}
	if unique := uqq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := uqq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, uahquote.FieldID)
		for i := range fields {
			if fields[i] != uahquote.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := uqq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uqq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uqq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uqq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uqq *UAHQuoteQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(uqq.driver.Dialect())
	t1 := builder.Table(uahquote.Table)
	columns := uqq.fields
	if len(columns) == 0 {
		columns = uahquote.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if uqq.sql != nil {
		selector = uqq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if uqq.unique != nil && *uqq.unique {
		selector.Distinct()
	}
	for _, p := range uqq.predicates {
		p(selector)
	}
	for _, p := range uqq.order {
		p(selector)
	}
	if offset := uqq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uqq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UAHQuoteGroupBy is the group-by builder for UAHQuote entities.
type UAHQuoteGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (uqgb *UAHQuoteGroupBy) Aggregate(fns ...AggregateFunc) *UAHQuoteGroupBy {
	uqgb.fns = append(uqgb.fns, fns...)
	return uqgb
}

// Scan applies the group-by query and scans the result into the given value.
func (uqgb *UAHQuoteGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := uqgb.path(ctx)
	if err != nil {
		return err
	}
	uqgb.sql = query
	return uqgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (uqgb *UAHQuoteGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := uqgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (uqgb *UAHQuoteGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(uqgb.fields) > 1 {
		return nil, errors.New("ent: UAHQuoteGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := uqgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (uqgb *UAHQuoteGroupBy) StringsX(ctx context.Context) []string {
	v, err := uqgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (uqgb *UAHQuoteGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = uqgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uahquote.Label}
	default:
		err = fmt.Errorf("ent: UAHQuoteGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (uqgb *UAHQuoteGroupBy) StringX(ctx context.Context) string {
	v, err := uqgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (uqgb *UAHQuoteGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(uqgb.fields) > 1 {
		return nil, errors.New("ent: UAHQuoteGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := uqgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (uqgb *UAHQuoteGroupBy) IntsX(ctx context.Context) []int {
	v, err := uqgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (uqgb *UAHQuoteGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = uqgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uahquote.Label}
	default:
		err = fmt.Errorf("ent: UAHQuoteGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (uqgb *UAHQuoteGroupBy) IntX(ctx context.Context) int {
	v, err := uqgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (uqgb *UAHQuoteGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(uqgb.fields) > 1 {
		return nil, errors.New("ent: UAHQuoteGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := uqgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (uqgb *UAHQuoteGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := uqgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (uqgb *UAHQuoteGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = uqgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uahquote.Label}
	default:
		err = fmt.Errorf("ent: UAHQuoteGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (uqgb *UAHQuoteGroupBy) Float64X(ctx context.Context) float64 {
	v, err := uqgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (uqgb *UAHQuoteGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(uqgb.fields) > 1 {
		return nil, errors.New("ent: UAHQuoteGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := uqgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (uqgb *UAHQuoteGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := uqgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (uqgb *UAHQuoteGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = uqgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uahquote.Label}
	default:
		err = fmt.Errorf("ent: UAHQuoteGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (uqgb *UAHQuoteGroupBy) BoolX(ctx context.Context) bool {
	v, err := uqgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uqgb *UAHQuoteGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range uqgb.fields {
		if !uahquote.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := uqgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uqgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (uqgb *UAHQuoteGroupBy) sqlQuery() *sql.Selector {
	selector := uqgb.sql.Select()
	aggregation := make([]string, 0, len(uqgb.fns))
	for _, fn := range uqgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(uqgb.fields)+len(uqgb.fns))
		for _, f := range uqgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(uqgb.fields...)...)
}

// UAHQuoteSelect is the builder for selecting fields of UAHQuote entities.
type UAHQuoteSelect struct {
	*UAHQuoteQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (uqs *UAHQuoteSelect) Scan(ctx context.Context, v interface{}) error {
	if err := uqs.prepareQuery(ctx); err != nil {
		return err
	}
	uqs.sql = uqs.UAHQuoteQuery.sqlQuery(ctx)
	return uqs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (uqs *UAHQuoteSelect) ScanX(ctx context.Context, v interface{}) {
	if err := uqs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (uqs *UAHQuoteSelect) Strings(ctx context.Context) ([]string, error) {
	if len(uqs.fields) > 1 {
		return nil, errors.New("ent: UAHQuoteSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := uqs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (uqs *UAHQuoteSelect) StringsX(ctx context.Context) []string {
	v, err := uqs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (uqs *UAHQuoteSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = uqs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uahquote.Label}
	default:
		err = fmt.Errorf("ent: UAHQuoteSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (uqs *UAHQuoteSelect) StringX(ctx context.Context) string {
	v, err := uqs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (uqs *UAHQuoteSelect) Ints(ctx context.Context) ([]int, error) {
	if len(uqs.fields) > 1 {
		return nil, errors.New("ent: UAHQuoteSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := uqs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (uqs *UAHQuoteSelect) IntsX(ctx context.Context) []int {
	v, err := uqs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (uqs *UAHQuoteSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = uqs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uahquote.Label}
	default:
		err = fmt.Errorf("ent: UAHQuoteSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (uqs *UAHQuoteSelect) IntX(ctx context.Context) int {
	v, err := uqs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (uqs *UAHQuoteSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(uqs.fields) > 1 {
		return nil, errors.New("ent: UAHQuoteSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := uqs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (uqs *UAHQuoteSelect) Float64sX(ctx context.Context) []float64 {
	v, err := uqs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (uqs *UAHQuoteSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = uqs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uahquote.Label}
	default:
		err = fmt.Errorf("ent: UAHQuoteSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (uqs *UAHQuoteSelect) Float64X(ctx context.Context) float64 {
	v, err := uqs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (uqs *UAHQuoteSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(uqs.fields) > 1 {
		return nil, errors.New("ent: UAHQuoteSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := uqs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (uqs *UAHQuoteSelect) BoolsX(ctx context.Context) []bool {
	v, err := uqs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (uqs *UAHQuoteSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = uqs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uahquote.Label}
	default:
		err = fmt.Errorf("ent: UAHQuoteSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (uqs *UAHQuoteSelect) BoolX(ctx context.Context) bool {
	v, err := uqs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uqs *UAHQuoteSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := uqs.sql.Query()
	if err := uqs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
