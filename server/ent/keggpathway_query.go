// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"genomedb/ent/keggpathway"
	"genomedb/ent/predicate"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// KeggPathwayQuery is the builder for querying KeggPathway entities.
type KeggPathwayQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.KeggPathway
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the KeggPathwayQuery builder.
func (kpq *KeggPathwayQuery) Where(ps ...predicate.KeggPathway) *KeggPathwayQuery {
	kpq.predicates = append(kpq.predicates, ps...)
	return kpq
}

// Limit adds a limit step to the query.
func (kpq *KeggPathwayQuery) Limit(limit int) *KeggPathwayQuery {
	kpq.limit = &limit
	return kpq
}

// Offset adds an offset step to the query.
func (kpq *KeggPathwayQuery) Offset(offset int) *KeggPathwayQuery {
	kpq.offset = &offset
	return kpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (kpq *KeggPathwayQuery) Unique(unique bool) *KeggPathwayQuery {
	kpq.unique = &unique
	return kpq
}

// Order adds an order step to the query.
func (kpq *KeggPathwayQuery) Order(o ...OrderFunc) *KeggPathwayQuery {
	kpq.order = append(kpq.order, o...)
	return kpq
}

// First returns the first KeggPathway entity from the query.
// Returns a *NotFoundError when no KeggPathway was found.
func (kpq *KeggPathwayQuery) First(ctx context.Context) (*KeggPathway, error) {
	nodes, err := kpq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{keggpathway.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (kpq *KeggPathwayQuery) FirstX(ctx context.Context) *KeggPathway {
	node, err := kpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first KeggPathway ID from the query.
// Returns a *NotFoundError when no KeggPathway ID was found.
func (kpq *KeggPathwayQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = kpq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{keggpathway.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (kpq *KeggPathwayQuery) FirstIDX(ctx context.Context) string {
	id, err := kpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single KeggPathway entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one KeggPathway entity is found.
// Returns a *NotFoundError when no KeggPathway entities are found.
func (kpq *KeggPathwayQuery) Only(ctx context.Context) (*KeggPathway, error) {
	nodes, err := kpq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{keggpathway.Label}
	default:
		return nil, &NotSingularError{keggpathway.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (kpq *KeggPathwayQuery) OnlyX(ctx context.Context) *KeggPathway {
	node, err := kpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only KeggPathway ID in the query.
// Returns a *NotSingularError when more than one KeggPathway ID is found.
// Returns a *NotFoundError when no entities are found.
func (kpq *KeggPathwayQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = kpq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{keggpathway.Label}
	default:
		err = &NotSingularError{keggpathway.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (kpq *KeggPathwayQuery) OnlyIDX(ctx context.Context) string {
	id, err := kpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of KeggPathways.
func (kpq *KeggPathwayQuery) All(ctx context.Context) ([]*KeggPathway, error) {
	if err := kpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return kpq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (kpq *KeggPathwayQuery) AllX(ctx context.Context) []*KeggPathway {
	nodes, err := kpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of KeggPathway IDs.
func (kpq *KeggPathwayQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := kpq.Select(keggpathway.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (kpq *KeggPathwayQuery) IDsX(ctx context.Context) []string {
	ids, err := kpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (kpq *KeggPathwayQuery) Count(ctx context.Context) (int, error) {
	if err := kpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return kpq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (kpq *KeggPathwayQuery) CountX(ctx context.Context) int {
	count, err := kpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (kpq *KeggPathwayQuery) Exist(ctx context.Context) (bool, error) {
	if err := kpq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return kpq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (kpq *KeggPathwayQuery) ExistX(ctx context.Context) bool {
	exist, err := kpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the KeggPathwayQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (kpq *KeggPathwayQuery) Clone() *KeggPathwayQuery {
	if kpq == nil {
		return nil
	}
	return &KeggPathwayQuery{
		config:     kpq.config,
		limit:      kpq.limit,
		offset:     kpq.offset,
		order:      append([]OrderFunc{}, kpq.order...),
		predicates: append([]predicate.KeggPathway{}, kpq.predicates...),
		// clone intermediate query.
		sql:    kpq.sql.Clone(),
		path:   kpq.path,
		unique: kpq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (kpq *KeggPathwayQuery) GroupBy(field string, fields ...string) *KeggPathwayGroupBy {
	grbuild := &KeggPathwayGroupBy{config: kpq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := kpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return kpq.sqlQuery(ctx), nil
	}
	grbuild.label = keggpathway.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (kpq *KeggPathwayQuery) Select(fields ...string) *KeggPathwaySelect {
	kpq.fields = append(kpq.fields, fields...)
	selbuild := &KeggPathwaySelect{KeggPathwayQuery: kpq}
	selbuild.label = keggpathway.Label
	selbuild.flds, selbuild.scan = &kpq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a KeggPathwaySelect configured with the given aggregations.
func (kpq *KeggPathwayQuery) Aggregate(fns ...AggregateFunc) *KeggPathwaySelect {
	return kpq.Select().Aggregate(fns...)
}

func (kpq *KeggPathwayQuery) prepareQuery(ctx context.Context) error {
	for _, f := range kpq.fields {
		if !keggpathway.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if kpq.path != nil {
		prev, err := kpq.path(ctx)
		if err != nil {
			return err
		}
		kpq.sql = prev
	}
	return nil
}

func (kpq *KeggPathwayQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*KeggPathway, error) {
	var (
		nodes = []*KeggPathway{}
		_spec = kpq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*KeggPathway).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &KeggPathway{config: kpq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, kpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (kpq *KeggPathwayQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := kpq.querySpec()
	_spec.Node.Columns = kpq.fields
	if len(kpq.fields) > 0 {
		_spec.Unique = kpq.unique != nil && *kpq.unique
	}
	return sqlgraph.CountNodes(ctx, kpq.driver, _spec)
}

func (kpq *KeggPathwayQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := kpq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (kpq *KeggPathwayQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   keggpathway.Table,
			Columns: keggpathway.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: keggpathway.FieldID,
			},
		},
		From:   kpq.sql,
		Unique: true,
	}
	if unique := kpq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := kpq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, keggpathway.FieldID)
		for i := range fields {
			if fields[i] != keggpathway.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := kpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := kpq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := kpq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := kpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (kpq *KeggPathwayQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(kpq.driver.Dialect())
	t1 := builder.Table(keggpathway.Table)
	columns := kpq.fields
	if len(columns) == 0 {
		columns = keggpathway.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if kpq.sql != nil {
		selector = kpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if kpq.unique != nil && *kpq.unique {
		selector.Distinct()
	}
	for _, p := range kpq.predicates {
		p(selector)
	}
	for _, p := range kpq.order {
		p(selector)
	}
	if offset := kpq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := kpq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// KeggPathwayGroupBy is the group-by builder for KeggPathway entities.
type KeggPathwayGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (kpgb *KeggPathwayGroupBy) Aggregate(fns ...AggregateFunc) *KeggPathwayGroupBy {
	kpgb.fns = append(kpgb.fns, fns...)
	return kpgb
}

// Scan applies the group-by query and scans the result into the given value.
func (kpgb *KeggPathwayGroupBy) Scan(ctx context.Context, v any) error {
	query, err := kpgb.path(ctx)
	if err != nil {
		return err
	}
	kpgb.sql = query
	return kpgb.sqlScan(ctx, v)
}

func (kpgb *KeggPathwayGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range kpgb.fields {
		if !keggpathway.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := kpgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := kpgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (kpgb *KeggPathwayGroupBy) sqlQuery() *sql.Selector {
	selector := kpgb.sql.Select()
	aggregation := make([]string, 0, len(kpgb.fns))
	for _, fn := range kpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(kpgb.fields)+len(kpgb.fns))
		for _, f := range kpgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(kpgb.fields...)...)
}

// KeggPathwaySelect is the builder for selecting fields of KeggPathway entities.
type KeggPathwaySelect struct {
	*KeggPathwayQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (kps *KeggPathwaySelect) Aggregate(fns ...AggregateFunc) *KeggPathwaySelect {
	kps.fns = append(kps.fns, fns...)
	return kps
}

// Scan applies the selector query and scans the result into the given value.
func (kps *KeggPathwaySelect) Scan(ctx context.Context, v any) error {
	if err := kps.prepareQuery(ctx); err != nil {
		return err
	}
	kps.sql = kps.KeggPathwayQuery.sqlQuery(ctx)
	return kps.sqlScan(ctx, v)
}

func (kps *KeggPathwaySelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(kps.fns))
	for _, fn := range kps.fns {
		aggregation = append(aggregation, fn(kps.sql))
	}
	switch n := len(*kps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		kps.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		kps.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := kps.sql.Query()
	if err := kps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}