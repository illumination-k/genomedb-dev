// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"genomedb/ent/nomeclature"
	"genomedb/ent/predicate"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NomeclatureQuery is the builder for querying Nomeclature entities.
type NomeclatureQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Nomeclature
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the NomeclatureQuery builder.
func (nq *NomeclatureQuery) Where(ps ...predicate.Nomeclature) *NomeclatureQuery {
	nq.predicates = append(nq.predicates, ps...)
	return nq
}

// Limit adds a limit step to the query.
func (nq *NomeclatureQuery) Limit(limit int) *NomeclatureQuery {
	nq.limit = &limit
	return nq
}

// Offset adds an offset step to the query.
func (nq *NomeclatureQuery) Offset(offset int) *NomeclatureQuery {
	nq.offset = &offset
	return nq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (nq *NomeclatureQuery) Unique(unique bool) *NomeclatureQuery {
	nq.unique = &unique
	return nq
}

// Order adds an order step to the query.
func (nq *NomeclatureQuery) Order(o ...OrderFunc) *NomeclatureQuery {
	nq.order = append(nq.order, o...)
	return nq
}

// First returns the first Nomeclature entity from the query.
// Returns a *NotFoundError when no Nomeclature was found.
func (nq *NomeclatureQuery) First(ctx context.Context) (*Nomeclature, error) {
	nodes, err := nq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{nomeclature.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (nq *NomeclatureQuery) FirstX(ctx context.Context) *Nomeclature {
	node, err := nq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Nomeclature ID from the query.
// Returns a *NotFoundError when no Nomeclature ID was found.
func (nq *NomeclatureQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = nq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{nomeclature.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (nq *NomeclatureQuery) FirstIDX(ctx context.Context) int {
	id, err := nq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Nomeclature entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Nomeclature entity is found.
// Returns a *NotFoundError when no Nomeclature entities are found.
func (nq *NomeclatureQuery) Only(ctx context.Context) (*Nomeclature, error) {
	nodes, err := nq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{nomeclature.Label}
	default:
		return nil, &NotSingularError{nomeclature.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (nq *NomeclatureQuery) OnlyX(ctx context.Context) *Nomeclature {
	node, err := nq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Nomeclature ID in the query.
// Returns a *NotSingularError when more than one Nomeclature ID is found.
// Returns a *NotFoundError when no entities are found.
func (nq *NomeclatureQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = nq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{nomeclature.Label}
	default:
		err = &NotSingularError{nomeclature.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (nq *NomeclatureQuery) OnlyIDX(ctx context.Context) int {
	id, err := nq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Nomeclatures.
func (nq *NomeclatureQuery) All(ctx context.Context) ([]*Nomeclature, error) {
	if err := nq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return nq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (nq *NomeclatureQuery) AllX(ctx context.Context) []*Nomeclature {
	nodes, err := nq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Nomeclature IDs.
func (nq *NomeclatureQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := nq.Select(nomeclature.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (nq *NomeclatureQuery) IDsX(ctx context.Context) []int {
	ids, err := nq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (nq *NomeclatureQuery) Count(ctx context.Context) (int, error) {
	if err := nq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return nq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (nq *NomeclatureQuery) CountX(ctx context.Context) int {
	count, err := nq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (nq *NomeclatureQuery) Exist(ctx context.Context) (bool, error) {
	if err := nq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return nq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (nq *NomeclatureQuery) ExistX(ctx context.Context) bool {
	exist, err := nq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NomeclatureQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (nq *NomeclatureQuery) Clone() *NomeclatureQuery {
	if nq == nil {
		return nil
	}
	return &NomeclatureQuery{
		config:     nq.config,
		limit:      nq.limit,
		offset:     nq.offset,
		order:      append([]OrderFunc{}, nq.order...),
		predicates: append([]predicate.Nomeclature{}, nq.predicates...),
		// clone intermediate query.
		sql:    nq.sql.Clone(),
		path:   nq.path,
		unique: nq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Nomeclature.Query().
//		GroupBy(nomeclature.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (nq *NomeclatureQuery) GroupBy(field string, fields ...string) *NomeclatureGroupBy {
	grbuild := &NomeclatureGroupBy{config: nq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return nq.sqlQuery(ctx), nil
	}
	grbuild.label = nomeclature.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Nomeclature.Query().
//		Select(nomeclature.FieldName).
//		Scan(ctx, &v)
func (nq *NomeclatureQuery) Select(fields ...string) *NomeclatureSelect {
	nq.fields = append(nq.fields, fields...)
	selbuild := &NomeclatureSelect{NomeclatureQuery: nq}
	selbuild.label = nomeclature.Label
	selbuild.flds, selbuild.scan = &nq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a NomeclatureSelect configured with the given aggregations.
func (nq *NomeclatureQuery) Aggregate(fns ...AggregateFunc) *NomeclatureSelect {
	return nq.Select().Aggregate(fns...)
}

func (nq *NomeclatureQuery) prepareQuery(ctx context.Context) error {
	for _, f := range nq.fields {
		if !nomeclature.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if nq.path != nil {
		prev, err := nq.path(ctx)
		if err != nil {
			return err
		}
		nq.sql = prev
	}
	return nil
}

func (nq *NomeclatureQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Nomeclature, error) {
	var (
		nodes = []*Nomeclature{}
		_spec = nq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Nomeclature).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Nomeclature{config: nq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, nq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (nq *NomeclatureQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := nq.querySpec()
	_spec.Node.Columns = nq.fields
	if len(nq.fields) > 0 {
		_spec.Unique = nq.unique != nil && *nq.unique
	}
	return sqlgraph.CountNodes(ctx, nq.driver, _spec)
}

func (nq *NomeclatureQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := nq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (nq *NomeclatureQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   nomeclature.Table,
			Columns: nomeclature.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: nomeclature.FieldID,
			},
		},
		From:   nq.sql,
		Unique: true,
	}
	if unique := nq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := nq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, nomeclature.FieldID)
		for i := range fields {
			if fields[i] != nomeclature.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := nq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := nq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := nq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := nq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (nq *NomeclatureQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(nq.driver.Dialect())
	t1 := builder.Table(nomeclature.Table)
	columns := nq.fields
	if len(columns) == 0 {
		columns = nomeclature.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if nq.sql != nil {
		selector = nq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if nq.unique != nil && *nq.unique {
		selector.Distinct()
	}
	for _, p := range nq.predicates {
		p(selector)
	}
	for _, p := range nq.order {
		p(selector)
	}
	if offset := nq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := nq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// NomeclatureGroupBy is the group-by builder for Nomeclature entities.
type NomeclatureGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ngb *NomeclatureGroupBy) Aggregate(fns ...AggregateFunc) *NomeclatureGroupBy {
	ngb.fns = append(ngb.fns, fns...)
	return ngb
}

// Scan applies the group-by query and scans the result into the given value.
func (ngb *NomeclatureGroupBy) Scan(ctx context.Context, v any) error {
	query, err := ngb.path(ctx)
	if err != nil {
		return err
	}
	ngb.sql = query
	return ngb.sqlScan(ctx, v)
}

func (ngb *NomeclatureGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range ngb.fields {
		if !nomeclature.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ngb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ngb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ngb *NomeclatureGroupBy) sqlQuery() *sql.Selector {
	selector := ngb.sql.Select()
	aggregation := make([]string, 0, len(ngb.fns))
	for _, fn := range ngb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ngb.fields)+len(ngb.fns))
		for _, f := range ngb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ngb.fields...)...)
}

// NomeclatureSelect is the builder for selecting fields of Nomeclature entities.
type NomeclatureSelect struct {
	*NomeclatureQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ns *NomeclatureSelect) Aggregate(fns ...AggregateFunc) *NomeclatureSelect {
	ns.fns = append(ns.fns, fns...)
	return ns
}

// Scan applies the selector query and scans the result into the given value.
func (ns *NomeclatureSelect) Scan(ctx context.Context, v any) error {
	if err := ns.prepareQuery(ctx); err != nil {
		return err
	}
	ns.sql = ns.NomeclatureQuery.sqlQuery(ctx)
	return ns.sqlScan(ctx, v)
}

func (ns *NomeclatureSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(ns.fns))
	for _, fn := range ns.fns {
		aggregation = append(aggregation, fn(ns.sql))
	}
	switch n := len(*ns.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		ns.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		ns.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := ns.sql.Query()
	if err := ns.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}