// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"genomedb/ent/cds"
	"genomedb/ent/exon"
	"genomedb/ent/fiveprimeutr"
	"genomedb/ent/gene"
	"genomedb/ent/predicate"
	"genomedb/ent/threeprimeutr"
	"genomedb/ent/transcript"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TranscriptQuery is the builder for querying Transcript entities.
type TranscriptQuery struct {
	config
	limit             *int
	offset            *int
	unique            *bool
	order             []OrderFunc
	fields            []string
	predicates        []predicate.Transcript
	withGene          *GeneQuery
	withCds           *CdsQuery
	withExon          *ExonQuery
	withFivePrimeUtr  *FivePrimeUtrQuery
	withThreePrimeUtr *ThreePrimeUtrQuery
	withFKs           bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TranscriptQuery builder.
func (tq *TranscriptQuery) Where(ps ...predicate.Transcript) *TranscriptQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit adds a limit step to the query.
func (tq *TranscriptQuery) Limit(limit int) *TranscriptQuery {
	tq.limit = &limit
	return tq
}

// Offset adds an offset step to the query.
func (tq *TranscriptQuery) Offset(offset int) *TranscriptQuery {
	tq.offset = &offset
	return tq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tq *TranscriptQuery) Unique(unique bool) *TranscriptQuery {
	tq.unique = &unique
	return tq
}

// Order adds an order step to the query.
func (tq *TranscriptQuery) Order(o ...OrderFunc) *TranscriptQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// QueryGene chains the current query on the "gene" edge.
func (tq *TranscriptQuery) QueryGene() *GeneQuery {
	query := &GeneQuery{config: tq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(transcript.Table, transcript.FieldID, selector),
			sqlgraph.To(gene.Table, gene.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, transcript.GeneTable, transcript.GeneColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCds chains the current query on the "cds" edge.
func (tq *TranscriptQuery) QueryCds() *CdsQuery {
	query := &CdsQuery{config: tq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(transcript.Table, transcript.FieldID, selector),
			sqlgraph.To(cds.Table, cds.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, transcript.CdsTable, transcript.CdsColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryExon chains the current query on the "exon" edge.
func (tq *TranscriptQuery) QueryExon() *ExonQuery {
	query := &ExonQuery{config: tq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(transcript.Table, transcript.FieldID, selector),
			sqlgraph.To(exon.Table, exon.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, transcript.ExonTable, transcript.ExonColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFivePrimeUtr chains the current query on the "five_prime_utr" edge.
func (tq *TranscriptQuery) QueryFivePrimeUtr() *FivePrimeUtrQuery {
	query := &FivePrimeUtrQuery{config: tq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(transcript.Table, transcript.FieldID, selector),
			sqlgraph.To(fiveprimeutr.Table, fiveprimeutr.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, transcript.FivePrimeUtrTable, transcript.FivePrimeUtrColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryThreePrimeUtr chains the current query on the "three_prime_utr" edge.
func (tq *TranscriptQuery) QueryThreePrimeUtr() *ThreePrimeUtrQuery {
	query := &ThreePrimeUtrQuery{config: tq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(transcript.Table, transcript.FieldID, selector),
			sqlgraph.To(threeprimeutr.Table, threeprimeutr.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, transcript.ThreePrimeUtrTable, transcript.ThreePrimeUtrColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Transcript entity from the query.
// Returns a *NotFoundError when no Transcript was found.
func (tq *TranscriptQuery) First(ctx context.Context) (*Transcript, error) {
	nodes, err := tq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{transcript.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *TranscriptQuery) FirstX(ctx context.Context) *Transcript {
	node, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Transcript ID from the query.
// Returns a *NotFoundError when no Transcript ID was found.
func (tq *TranscriptQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = tq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{transcript.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tq *TranscriptQuery) FirstIDX(ctx context.Context) string {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Transcript entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Transcript entity is found.
// Returns a *NotFoundError when no Transcript entities are found.
func (tq *TranscriptQuery) Only(ctx context.Context) (*Transcript, error) {
	nodes, err := tq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{transcript.Label}
	default:
		return nil, &NotSingularError{transcript.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *TranscriptQuery) OnlyX(ctx context.Context) *Transcript {
	node, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Transcript ID in the query.
// Returns a *NotSingularError when more than one Transcript ID is found.
// Returns a *NotFoundError when no entities are found.
func (tq *TranscriptQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = tq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{transcript.Label}
	default:
		err = &NotSingularError{transcript.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tq *TranscriptQuery) OnlyIDX(ctx context.Context) string {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Transcripts.
func (tq *TranscriptQuery) All(ctx context.Context) ([]*Transcript, error) {
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return tq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (tq *TranscriptQuery) AllX(ctx context.Context) []*Transcript {
	nodes, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Transcript IDs.
func (tq *TranscriptQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := tq.Select(transcript.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *TranscriptQuery) IDsX(ctx context.Context) []string {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *TranscriptQuery) Count(ctx context.Context) (int, error) {
	if err := tq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return tq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (tq *TranscriptQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *TranscriptQuery) Exist(ctx context.Context) (bool, error) {
	if err := tq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return tq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (tq *TranscriptQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TranscriptQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *TranscriptQuery) Clone() *TranscriptQuery {
	if tq == nil {
		return nil
	}
	return &TranscriptQuery{
		config:            tq.config,
		limit:             tq.limit,
		offset:            tq.offset,
		order:             append([]OrderFunc{}, tq.order...),
		predicates:        append([]predicate.Transcript{}, tq.predicates...),
		withGene:          tq.withGene.Clone(),
		withCds:           tq.withCds.Clone(),
		withExon:          tq.withExon.Clone(),
		withFivePrimeUtr:  tq.withFivePrimeUtr.Clone(),
		withThreePrimeUtr: tq.withThreePrimeUtr.Clone(),
		// clone intermediate query.
		sql:    tq.sql.Clone(),
		path:   tq.path,
		unique: tq.unique,
	}
}

// WithGene tells the query-builder to eager-load the nodes that are connected to
// the "gene" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TranscriptQuery) WithGene(opts ...func(*GeneQuery)) *TranscriptQuery {
	query := &GeneQuery{config: tq.config}
	for _, opt := range opts {
		opt(query)
	}
	tq.withGene = query
	return tq
}

// WithCds tells the query-builder to eager-load the nodes that are connected to
// the "cds" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TranscriptQuery) WithCds(opts ...func(*CdsQuery)) *TranscriptQuery {
	query := &CdsQuery{config: tq.config}
	for _, opt := range opts {
		opt(query)
	}
	tq.withCds = query
	return tq
}

// WithExon tells the query-builder to eager-load the nodes that are connected to
// the "exon" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TranscriptQuery) WithExon(opts ...func(*ExonQuery)) *TranscriptQuery {
	query := &ExonQuery{config: tq.config}
	for _, opt := range opts {
		opt(query)
	}
	tq.withExon = query
	return tq
}

// WithFivePrimeUtr tells the query-builder to eager-load the nodes that are connected to
// the "five_prime_utr" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TranscriptQuery) WithFivePrimeUtr(opts ...func(*FivePrimeUtrQuery)) *TranscriptQuery {
	query := &FivePrimeUtrQuery{config: tq.config}
	for _, opt := range opts {
		opt(query)
	}
	tq.withFivePrimeUtr = query
	return tq
}

// WithThreePrimeUtr tells the query-builder to eager-load the nodes that are connected to
// the "three_prime_utr" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TranscriptQuery) WithThreePrimeUtr(opts ...func(*ThreePrimeUtrQuery)) *TranscriptQuery {
	query := &ThreePrimeUtrQuery{config: tq.config}
	for _, opt := range opts {
		opt(query)
	}
	tq.withThreePrimeUtr = query
	return tq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Strand string `json:"strand,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Transcript.Query().
//		GroupBy(transcript.FieldStrand).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tq *TranscriptQuery) GroupBy(field string, fields ...string) *TranscriptGroupBy {
	grbuild := &TranscriptGroupBy{config: tq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return tq.sqlQuery(ctx), nil
	}
	grbuild.label = transcript.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Strand string `json:"strand,omitempty"`
//	}
//
//	client.Transcript.Query().
//		Select(transcript.FieldStrand).
//		Scan(ctx, &v)
func (tq *TranscriptQuery) Select(fields ...string) *TranscriptSelect {
	tq.fields = append(tq.fields, fields...)
	selbuild := &TranscriptSelect{TranscriptQuery: tq}
	selbuild.label = transcript.Label
	selbuild.flds, selbuild.scan = &tq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a TranscriptSelect configured with the given aggregations.
func (tq *TranscriptQuery) Aggregate(fns ...AggregateFunc) *TranscriptSelect {
	return tq.Select().Aggregate(fns...)
}

func (tq *TranscriptQuery) prepareQuery(ctx context.Context) error {
	for _, f := range tq.fields {
		if !transcript.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tq.path != nil {
		prev, err := tq.path(ctx)
		if err != nil {
			return err
		}
		tq.sql = prev
	}
	return nil
}

func (tq *TranscriptQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Transcript, error) {
	var (
		nodes       = []*Transcript{}
		withFKs     = tq.withFKs
		_spec       = tq.querySpec()
		loadedTypes = [5]bool{
			tq.withGene != nil,
			tq.withCds != nil,
			tq.withExon != nil,
			tq.withFivePrimeUtr != nil,
			tq.withThreePrimeUtr != nil,
		}
	)
	if tq.withGene != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, transcript.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Transcript).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Transcript{config: tq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tq.withGene; query != nil {
		if err := tq.loadGene(ctx, query, nodes, nil,
			func(n *Transcript, e *Gene) { n.Edges.Gene = e }); err != nil {
			return nil, err
		}
	}
	if query := tq.withCds; query != nil {
		if err := tq.loadCds(ctx, query, nodes,
			func(n *Transcript) { n.Edges.Cds = []*Cds{} },
			func(n *Transcript, e *Cds) { n.Edges.Cds = append(n.Edges.Cds, e) }); err != nil {
			return nil, err
		}
	}
	if query := tq.withExon; query != nil {
		if err := tq.loadExon(ctx, query, nodes,
			func(n *Transcript) { n.Edges.Exon = []*Exon{} },
			func(n *Transcript, e *Exon) { n.Edges.Exon = append(n.Edges.Exon, e) }); err != nil {
			return nil, err
		}
	}
	if query := tq.withFivePrimeUtr; query != nil {
		if err := tq.loadFivePrimeUtr(ctx, query, nodes,
			func(n *Transcript) { n.Edges.FivePrimeUtr = []*FivePrimeUtr{} },
			func(n *Transcript, e *FivePrimeUtr) { n.Edges.FivePrimeUtr = append(n.Edges.FivePrimeUtr, e) }); err != nil {
			return nil, err
		}
	}
	if query := tq.withThreePrimeUtr; query != nil {
		if err := tq.loadThreePrimeUtr(ctx, query, nodes,
			func(n *Transcript) { n.Edges.ThreePrimeUtr = []*ThreePrimeUtr{} },
			func(n *Transcript, e *ThreePrimeUtr) { n.Edges.ThreePrimeUtr = append(n.Edges.ThreePrimeUtr, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tq *TranscriptQuery) loadGene(ctx context.Context, query *GeneQuery, nodes []*Transcript, init func(*Transcript), assign func(*Transcript, *Gene)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Transcript)
	for i := range nodes {
		if nodes[i].gene_transcripts == nil {
			continue
		}
		fk := *nodes[i].gene_transcripts
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(gene.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "gene_transcripts" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (tq *TranscriptQuery) loadCds(ctx context.Context, query *CdsQuery, nodes []*Transcript, init func(*Transcript), assign func(*Transcript, *Cds)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Transcript)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Cds(func(s *sql.Selector) {
		s.Where(sql.InValues(transcript.CdsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.transcript_cds
		if fk == nil {
			return fmt.Errorf(`foreign-key "transcript_cds" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "transcript_cds" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (tq *TranscriptQuery) loadExon(ctx context.Context, query *ExonQuery, nodes []*Transcript, init func(*Transcript), assign func(*Transcript, *Exon)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Transcript)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Exon(func(s *sql.Selector) {
		s.Where(sql.InValues(transcript.ExonColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.transcript_exon
		if fk == nil {
			return fmt.Errorf(`foreign-key "transcript_exon" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "transcript_exon" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (tq *TranscriptQuery) loadFivePrimeUtr(ctx context.Context, query *FivePrimeUtrQuery, nodes []*Transcript, init func(*Transcript), assign func(*Transcript, *FivePrimeUtr)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Transcript)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.InValues(transcript.FivePrimeUtrColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.transcript_five_prime_utr
		if fk == nil {
			return fmt.Errorf(`foreign-key "transcript_five_prime_utr" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "transcript_five_prime_utr" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (tq *TranscriptQuery) loadThreePrimeUtr(ctx context.Context, query *ThreePrimeUtrQuery, nodes []*Transcript, init func(*Transcript), assign func(*Transcript, *ThreePrimeUtr)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Transcript)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.ThreePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.InValues(transcript.ThreePrimeUtrColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.transcript_three_prime_utr
		if fk == nil {
			return fmt.Errorf(`foreign-key "transcript_three_prime_utr" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "transcript_three_prime_utr" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (tq *TranscriptQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tq.querySpec()
	_spec.Node.Columns = tq.fields
	if len(tq.fields) > 0 {
		_spec.Unique = tq.unique != nil && *tq.unique
	}
	return sqlgraph.CountNodes(ctx, tq.driver, _spec)
}

func (tq *TranscriptQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := tq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (tq *TranscriptQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   transcript.Table,
			Columns: transcript.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: transcript.FieldID,
			},
		},
		From:   tq.sql,
		Unique: true,
	}
	if unique := tq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := tq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, transcript.FieldID)
		for i := range fields {
			if fields[i] != transcript.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tq *TranscriptQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(transcript.Table)
	columns := tq.fields
	if len(columns) == 0 {
		columns = transcript.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tq.sql != nil {
		selector = tq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tq.unique != nil && *tq.unique {
		selector.Distinct()
	}
	for _, p := range tq.predicates {
		p(selector)
	}
	for _, p := range tq.order {
		p(selector)
	}
	if offset := tq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TranscriptGroupBy is the group-by builder for Transcript entities.
type TranscriptGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *TranscriptGroupBy) Aggregate(fns ...AggregateFunc) *TranscriptGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the group-by query and scans the result into the given value.
func (tgb *TranscriptGroupBy) Scan(ctx context.Context, v any) error {
	query, err := tgb.path(ctx)
	if err != nil {
		return err
	}
	tgb.sql = query
	return tgb.sqlScan(ctx, v)
}

func (tgb *TranscriptGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range tgb.fields {
		if !transcript.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := tgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tgb *TranscriptGroupBy) sqlQuery() *sql.Selector {
	selector := tgb.sql.Select()
	aggregation := make([]string, 0, len(tgb.fns))
	for _, fn := range tgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(tgb.fields)+len(tgb.fns))
		for _, f := range tgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(tgb.fields...)...)
}

// TranscriptSelect is the builder for selecting fields of Transcript entities.
type TranscriptSelect struct {
	*TranscriptQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ts *TranscriptSelect) Aggregate(fns ...AggregateFunc) *TranscriptSelect {
	ts.fns = append(ts.fns, fns...)
	return ts
}

// Scan applies the selector query and scans the result into the given value.
func (ts *TranscriptSelect) Scan(ctx context.Context, v any) error {
	if err := ts.prepareQuery(ctx); err != nil {
		return err
	}
	ts.sql = ts.TranscriptQuery.sqlQuery(ctx)
	return ts.sqlScan(ctx, v)
}

func (ts *TranscriptSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(ts.fns))
	for _, fn := range ts.fns {
		aggregation = append(aggregation, fn(ts.sql))
	}
	switch n := len(*ts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		ts.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		ts.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := ts.sql.Query()
	if err := ts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
