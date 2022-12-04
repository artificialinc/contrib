// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"github.com/artificialinc/contrib/entproto/internal/todo/ent/multiwordschema"
	"github.com/artificialinc/contrib/entproto/internal/todo/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MultiWordSchemaQuery is the builder for querying MultiWordSchema entities.
type MultiWordSchemaQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.MultiWordSchema
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MultiWordSchemaQuery builder.
func (mwsq *MultiWordSchemaQuery) Where(ps ...predicate.MultiWordSchema) *MultiWordSchemaQuery {
	mwsq.predicates = append(mwsq.predicates, ps...)
	return mwsq
}

// Limit adds a limit step to the query.
func (mwsq *MultiWordSchemaQuery) Limit(limit int) *MultiWordSchemaQuery {
	mwsq.limit = &limit
	return mwsq
}

// Offset adds an offset step to the query.
func (mwsq *MultiWordSchemaQuery) Offset(offset int) *MultiWordSchemaQuery {
	mwsq.offset = &offset
	return mwsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mwsq *MultiWordSchemaQuery) Unique(unique bool) *MultiWordSchemaQuery {
	mwsq.unique = &unique
	return mwsq
}

// Order adds an order step to the query.
func (mwsq *MultiWordSchemaQuery) Order(o ...OrderFunc) *MultiWordSchemaQuery {
	mwsq.order = append(mwsq.order, o...)
	return mwsq
}

// First returns the first MultiWordSchema entity from the query.
// Returns a *NotFoundError when no MultiWordSchema was found.
func (mwsq *MultiWordSchemaQuery) First(ctx context.Context) (*MultiWordSchema, error) {
	nodes, err := mwsq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{multiwordschema.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mwsq *MultiWordSchemaQuery) FirstX(ctx context.Context) *MultiWordSchema {
	node, err := mwsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MultiWordSchema ID from the query.
// Returns a *NotFoundError when no MultiWordSchema ID was found.
func (mwsq *MultiWordSchemaQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mwsq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{multiwordschema.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mwsq *MultiWordSchemaQuery) FirstIDX(ctx context.Context) int {
	id, err := mwsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MultiWordSchema entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MultiWordSchema entity is found.
// Returns a *NotFoundError when no MultiWordSchema entities are found.
func (mwsq *MultiWordSchemaQuery) Only(ctx context.Context) (*MultiWordSchema, error) {
	nodes, err := mwsq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{multiwordschema.Label}
	default:
		return nil, &NotSingularError{multiwordschema.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mwsq *MultiWordSchemaQuery) OnlyX(ctx context.Context) *MultiWordSchema {
	node, err := mwsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MultiWordSchema ID in the query.
// Returns a *NotSingularError when more than one MultiWordSchema ID is found.
// Returns a *NotFoundError when no entities are found.
func (mwsq *MultiWordSchemaQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mwsq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{multiwordschema.Label}
	default:
		err = &NotSingularError{multiwordschema.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mwsq *MultiWordSchemaQuery) OnlyIDX(ctx context.Context) int {
	id, err := mwsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MultiWordSchemas.
func (mwsq *MultiWordSchemaQuery) All(ctx context.Context) ([]*MultiWordSchema, error) {
	if err := mwsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return mwsq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (mwsq *MultiWordSchemaQuery) AllX(ctx context.Context) []*MultiWordSchema {
	nodes, err := mwsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MultiWordSchema IDs.
func (mwsq *MultiWordSchemaQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := mwsq.Select(multiwordschema.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mwsq *MultiWordSchemaQuery) IDsX(ctx context.Context) []int {
	ids, err := mwsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mwsq *MultiWordSchemaQuery) Count(ctx context.Context) (int, error) {
	if err := mwsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return mwsq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (mwsq *MultiWordSchemaQuery) CountX(ctx context.Context) int {
	count, err := mwsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mwsq *MultiWordSchemaQuery) Exist(ctx context.Context) (bool, error) {
	if err := mwsq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return mwsq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (mwsq *MultiWordSchemaQuery) ExistX(ctx context.Context) bool {
	exist, err := mwsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MultiWordSchemaQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mwsq *MultiWordSchemaQuery) Clone() *MultiWordSchemaQuery {
	if mwsq == nil {
		return nil
	}
	return &MultiWordSchemaQuery{
		config:     mwsq.config,
		limit:      mwsq.limit,
		offset:     mwsq.offset,
		order:      append([]OrderFunc{}, mwsq.order...),
		predicates: append([]predicate.MultiWordSchema{}, mwsq.predicates...),
		// clone intermediate query.
		sql:    mwsq.sql.Clone(),
		path:   mwsq.path,
		unique: mwsq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Unit multiwordschema.Unit `json:"unit,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.MultiWordSchema.Query().
//		GroupBy(multiwordschema.FieldUnit).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (mwsq *MultiWordSchemaQuery) GroupBy(field string, fields ...string) *MultiWordSchemaGroupBy {
	grbuild := &MultiWordSchemaGroupBy{config: mwsq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := mwsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return mwsq.sqlQuery(ctx), nil
	}
	grbuild.label = multiwordschema.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Unit multiwordschema.Unit `json:"unit,omitempty"`
//	}
//
//	client.MultiWordSchema.Query().
//		Select(multiwordschema.FieldUnit).
//		Scan(ctx, &v)
func (mwsq *MultiWordSchemaQuery) Select(fields ...string) *MultiWordSchemaSelect {
	mwsq.fields = append(mwsq.fields, fields...)
	selbuild := &MultiWordSchemaSelect{MultiWordSchemaQuery: mwsq}
	selbuild.label = multiwordschema.Label
	selbuild.flds, selbuild.scan = &mwsq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a MultiWordSchemaSelect configured with the given aggregations.
func (mwsq *MultiWordSchemaQuery) Aggregate(fns ...AggregateFunc) *MultiWordSchemaSelect {
	return mwsq.Select().Aggregate(fns...)
}

func (mwsq *MultiWordSchemaQuery) prepareQuery(ctx context.Context) error {
	for _, f := range mwsq.fields {
		if !multiwordschema.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mwsq.path != nil {
		prev, err := mwsq.path(ctx)
		if err != nil {
			return err
		}
		mwsq.sql = prev
	}
	return nil
}

func (mwsq *MultiWordSchemaQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*MultiWordSchema, error) {
	var (
		nodes = []*MultiWordSchema{}
		_spec = mwsq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*MultiWordSchema).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &MultiWordSchema{config: mwsq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mwsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (mwsq *MultiWordSchemaQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mwsq.querySpec()
	_spec.Node.Columns = mwsq.fields
	if len(mwsq.fields) > 0 {
		_spec.Unique = mwsq.unique != nil && *mwsq.unique
	}
	return sqlgraph.CountNodes(ctx, mwsq.driver, _spec)
}

func (mwsq *MultiWordSchemaQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := mwsq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (mwsq *MultiWordSchemaQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   multiwordschema.Table,
			Columns: multiwordschema.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: multiwordschema.FieldID,
			},
		},
		From:   mwsq.sql,
		Unique: true,
	}
	if unique := mwsq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := mwsq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, multiwordschema.FieldID)
		for i := range fields {
			if fields[i] != multiwordschema.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mwsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mwsq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mwsq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mwsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mwsq *MultiWordSchemaQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mwsq.driver.Dialect())
	t1 := builder.Table(multiwordschema.Table)
	columns := mwsq.fields
	if len(columns) == 0 {
		columns = multiwordschema.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mwsq.sql != nil {
		selector = mwsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mwsq.unique != nil && *mwsq.unique {
		selector.Distinct()
	}
	for _, p := range mwsq.predicates {
		p(selector)
	}
	for _, p := range mwsq.order {
		p(selector)
	}
	if offset := mwsq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mwsq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MultiWordSchemaGroupBy is the group-by builder for MultiWordSchema entities.
type MultiWordSchemaGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mwsgb *MultiWordSchemaGroupBy) Aggregate(fns ...AggregateFunc) *MultiWordSchemaGroupBy {
	mwsgb.fns = append(mwsgb.fns, fns...)
	return mwsgb
}

// Scan applies the group-by query and scans the result into the given value.
func (mwsgb *MultiWordSchemaGroupBy) Scan(ctx context.Context, v any) error {
	query, err := mwsgb.path(ctx)
	if err != nil {
		return err
	}
	mwsgb.sql = query
	return mwsgb.sqlScan(ctx, v)
}

func (mwsgb *MultiWordSchemaGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range mwsgb.fields {
		if !multiwordschema.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := mwsgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mwsgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mwsgb *MultiWordSchemaGroupBy) sqlQuery() *sql.Selector {
	selector := mwsgb.sql.Select()
	aggregation := make([]string, 0, len(mwsgb.fns))
	for _, fn := range mwsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(mwsgb.fields)+len(mwsgb.fns))
		for _, f := range mwsgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(mwsgb.fields...)...)
}

// MultiWordSchemaSelect is the builder for selecting fields of MultiWordSchema entities.
type MultiWordSchemaSelect struct {
	*MultiWordSchemaQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (mwss *MultiWordSchemaSelect) Aggregate(fns ...AggregateFunc) *MultiWordSchemaSelect {
	mwss.fns = append(mwss.fns, fns...)
	return mwss
}

// Scan applies the selector query and scans the result into the given value.
func (mwss *MultiWordSchemaSelect) Scan(ctx context.Context, v any) error {
	if err := mwss.prepareQuery(ctx); err != nil {
		return err
	}
	mwss.sql = mwss.MultiWordSchemaQuery.sqlQuery(ctx)
	return mwss.sqlScan(ctx, v)
}

func (mwss *MultiWordSchemaSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(mwss.fns))
	for _, fn := range mwss.fns {
		aggregation = append(aggregation, fn(mwss.sql))
	}
	switch n := len(*mwss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		mwss.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		mwss.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := mwss.sql.Query()
	if err := mwss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
