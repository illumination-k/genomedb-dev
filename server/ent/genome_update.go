// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"genomedb/ent/genome"
	"genomedb/ent/locus"
	"genomedb/ent/predicate"
	"genomedb/ent/scaffold"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GenomeUpdate is the builder for updating Genome entities.
type GenomeUpdate struct {
	config
	hooks    []Hook
	mutation *GenomeMutation
}

// Where appends a list predicates to the GenomeUpdate builder.
func (gu *GenomeUpdate) Where(ps ...predicate.Genome) *GenomeUpdate {
	gu.mutation.Where(ps...)
	return gu
}

// SetCodonTable sets the "codon_table" field.
func (gu *GenomeUpdate) SetCodonTable(i int32) *GenomeUpdate {
	gu.mutation.ResetCodonTable()
	gu.mutation.SetCodonTable(i)
	return gu
}

// AddCodonTable adds i to the "codon_table" field.
func (gu *GenomeUpdate) AddCodonTable(i int32) *GenomeUpdate {
	gu.mutation.AddCodonTable(i)
	return gu
}

// AddLocuseIDs adds the "locuses" edge to the Locus entity by IDs.
func (gu *GenomeUpdate) AddLocuseIDs(ids ...string) *GenomeUpdate {
	gu.mutation.AddLocuseIDs(ids...)
	return gu
}

// AddLocuses adds the "locuses" edges to the Locus entity.
func (gu *GenomeUpdate) AddLocuses(l ...*Locus) *GenomeUpdate {
	ids := make([]string, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return gu.AddLocuseIDs(ids...)
}

// AddScaffoldIDs adds the "scaffolds" edge to the Scaffold entity by IDs.
func (gu *GenomeUpdate) AddScaffoldIDs(ids ...int) *GenomeUpdate {
	gu.mutation.AddScaffoldIDs(ids...)
	return gu
}

// AddScaffolds adds the "scaffolds" edges to the Scaffold entity.
func (gu *GenomeUpdate) AddScaffolds(s ...*Scaffold) *GenomeUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return gu.AddScaffoldIDs(ids...)
}

// Mutation returns the GenomeMutation object of the builder.
func (gu *GenomeUpdate) Mutation() *GenomeMutation {
	return gu.mutation
}

// ClearLocuses clears all "locuses" edges to the Locus entity.
func (gu *GenomeUpdate) ClearLocuses() *GenomeUpdate {
	gu.mutation.ClearLocuses()
	return gu
}

// RemoveLocuseIDs removes the "locuses" edge to Locus entities by IDs.
func (gu *GenomeUpdate) RemoveLocuseIDs(ids ...string) *GenomeUpdate {
	gu.mutation.RemoveLocuseIDs(ids...)
	return gu
}

// RemoveLocuses removes "locuses" edges to Locus entities.
func (gu *GenomeUpdate) RemoveLocuses(l ...*Locus) *GenomeUpdate {
	ids := make([]string, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return gu.RemoveLocuseIDs(ids...)
}

// ClearScaffolds clears all "scaffolds" edges to the Scaffold entity.
func (gu *GenomeUpdate) ClearScaffolds() *GenomeUpdate {
	gu.mutation.ClearScaffolds()
	return gu
}

// RemoveScaffoldIDs removes the "scaffolds" edge to Scaffold entities by IDs.
func (gu *GenomeUpdate) RemoveScaffoldIDs(ids ...int) *GenomeUpdate {
	gu.mutation.RemoveScaffoldIDs(ids...)
	return gu
}

// RemoveScaffolds removes "scaffolds" edges to Scaffold entities.
func (gu *GenomeUpdate) RemoveScaffolds(s ...*Scaffold) *GenomeUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return gu.RemoveScaffoldIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gu *GenomeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(gu.hooks) == 0 {
		if err = gu.check(); err != nil {
			return 0, err
		}
		affected, err = gu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GenomeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gu.check(); err != nil {
				return 0, err
			}
			gu.mutation = mutation
			affected, err = gu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gu.hooks) - 1; i >= 0; i-- {
			if gu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GenomeUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GenomeUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GenomeUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gu *GenomeUpdate) check() error {
	if v, ok := gu.mutation.CodonTable(); ok {
		if err := genome.CodonTableValidator(v); err != nil {
			return &ValidationError{Name: "codon_table", err: fmt.Errorf(`ent: validator failed for field "Genome.codon_table": %w`, err)}
		}
	}
	return nil
}

func (gu *GenomeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   genome.Table,
			Columns: genome.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: genome.FieldID,
			},
		},
	}
	if ps := gu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gu.mutation.CodonTable(); ok {
		_spec.SetField(genome.FieldCodonTable, field.TypeInt32, value)
	}
	if value, ok := gu.mutation.AddedCodonTable(); ok {
		_spec.AddField(genome.FieldCodonTable, field.TypeInt32, value)
	}
	if gu.mutation.LocusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   genome.LocusesTable,
			Columns: []string{genome.LocusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: locus.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.RemovedLocusesIDs(); len(nodes) > 0 && !gu.mutation.LocusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   genome.LocusesTable,
			Columns: []string{genome.LocusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: locus.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.LocusesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   genome.LocusesTable,
			Columns: []string{genome.LocusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: locus.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gu.mutation.ScaffoldsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   genome.ScaffoldsTable,
			Columns: []string{genome.ScaffoldsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: scaffold.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.RemovedScaffoldsIDs(); len(nodes) > 0 && !gu.mutation.ScaffoldsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   genome.ScaffoldsTable,
			Columns: []string{genome.ScaffoldsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: scaffold.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.ScaffoldsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   genome.ScaffoldsTable,
			Columns: []string{genome.ScaffoldsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: scaffold.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{genome.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// GenomeUpdateOne is the builder for updating a single Genome entity.
type GenomeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GenomeMutation
}

// SetCodonTable sets the "codon_table" field.
func (guo *GenomeUpdateOne) SetCodonTable(i int32) *GenomeUpdateOne {
	guo.mutation.ResetCodonTable()
	guo.mutation.SetCodonTable(i)
	return guo
}

// AddCodonTable adds i to the "codon_table" field.
func (guo *GenomeUpdateOne) AddCodonTable(i int32) *GenomeUpdateOne {
	guo.mutation.AddCodonTable(i)
	return guo
}

// AddLocuseIDs adds the "locuses" edge to the Locus entity by IDs.
func (guo *GenomeUpdateOne) AddLocuseIDs(ids ...string) *GenomeUpdateOne {
	guo.mutation.AddLocuseIDs(ids...)
	return guo
}

// AddLocuses adds the "locuses" edges to the Locus entity.
func (guo *GenomeUpdateOne) AddLocuses(l ...*Locus) *GenomeUpdateOne {
	ids := make([]string, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return guo.AddLocuseIDs(ids...)
}

// AddScaffoldIDs adds the "scaffolds" edge to the Scaffold entity by IDs.
func (guo *GenomeUpdateOne) AddScaffoldIDs(ids ...int) *GenomeUpdateOne {
	guo.mutation.AddScaffoldIDs(ids...)
	return guo
}

// AddScaffolds adds the "scaffolds" edges to the Scaffold entity.
func (guo *GenomeUpdateOne) AddScaffolds(s ...*Scaffold) *GenomeUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return guo.AddScaffoldIDs(ids...)
}

// Mutation returns the GenomeMutation object of the builder.
func (guo *GenomeUpdateOne) Mutation() *GenomeMutation {
	return guo.mutation
}

// ClearLocuses clears all "locuses" edges to the Locus entity.
func (guo *GenomeUpdateOne) ClearLocuses() *GenomeUpdateOne {
	guo.mutation.ClearLocuses()
	return guo
}

// RemoveLocuseIDs removes the "locuses" edge to Locus entities by IDs.
func (guo *GenomeUpdateOne) RemoveLocuseIDs(ids ...string) *GenomeUpdateOne {
	guo.mutation.RemoveLocuseIDs(ids...)
	return guo
}

// RemoveLocuses removes "locuses" edges to Locus entities.
func (guo *GenomeUpdateOne) RemoveLocuses(l ...*Locus) *GenomeUpdateOne {
	ids := make([]string, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return guo.RemoveLocuseIDs(ids...)
}

// ClearScaffolds clears all "scaffolds" edges to the Scaffold entity.
func (guo *GenomeUpdateOne) ClearScaffolds() *GenomeUpdateOne {
	guo.mutation.ClearScaffolds()
	return guo
}

// RemoveScaffoldIDs removes the "scaffolds" edge to Scaffold entities by IDs.
func (guo *GenomeUpdateOne) RemoveScaffoldIDs(ids ...int) *GenomeUpdateOne {
	guo.mutation.RemoveScaffoldIDs(ids...)
	return guo
}

// RemoveScaffolds removes "scaffolds" edges to Scaffold entities.
func (guo *GenomeUpdateOne) RemoveScaffolds(s ...*Scaffold) *GenomeUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return guo.RemoveScaffoldIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (guo *GenomeUpdateOne) Select(field string, fields ...string) *GenomeUpdateOne {
	guo.fields = append([]string{field}, fields...)
	return guo
}

// Save executes the query and returns the updated Genome entity.
func (guo *GenomeUpdateOne) Save(ctx context.Context) (*Genome, error) {
	var (
		err  error
		node *Genome
	)
	if len(guo.hooks) == 0 {
		if err = guo.check(); err != nil {
			return nil, err
		}
		node, err = guo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GenomeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = guo.check(); err != nil {
				return nil, err
			}
			guo.mutation = mutation
			node, err = guo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(guo.hooks) - 1; i >= 0; i-- {
			if guo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = guo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, guo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Genome)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from GenomeMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GenomeUpdateOne) SaveX(ctx context.Context) *Genome {
	node, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (guo *GenomeUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GenomeUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (guo *GenomeUpdateOne) check() error {
	if v, ok := guo.mutation.CodonTable(); ok {
		if err := genome.CodonTableValidator(v); err != nil {
			return &ValidationError{Name: "codon_table", err: fmt.Errorf(`ent: validator failed for field "Genome.codon_table": %w`, err)}
		}
	}
	return nil
}

func (guo *GenomeUpdateOne) sqlSave(ctx context.Context) (_node *Genome, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   genome.Table,
			Columns: genome.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: genome.FieldID,
			},
		},
	}
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Genome.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := guo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, genome.FieldID)
		for _, f := range fields {
			if !genome.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != genome.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := guo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := guo.mutation.CodonTable(); ok {
		_spec.SetField(genome.FieldCodonTable, field.TypeInt32, value)
	}
	if value, ok := guo.mutation.AddedCodonTable(); ok {
		_spec.AddField(genome.FieldCodonTable, field.TypeInt32, value)
	}
	if guo.mutation.LocusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   genome.LocusesTable,
			Columns: []string{genome.LocusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: locus.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.RemovedLocusesIDs(); len(nodes) > 0 && !guo.mutation.LocusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   genome.LocusesTable,
			Columns: []string{genome.LocusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: locus.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.LocusesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   genome.LocusesTable,
			Columns: []string{genome.LocusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: locus.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if guo.mutation.ScaffoldsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   genome.ScaffoldsTable,
			Columns: []string{genome.ScaffoldsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: scaffold.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.RemovedScaffoldsIDs(); len(nodes) > 0 && !guo.mutation.ScaffoldsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   genome.ScaffoldsTable,
			Columns: []string{genome.ScaffoldsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: scaffold.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.ScaffoldsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   genome.ScaffoldsTable,
			Columns: []string{genome.ScaffoldsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: scaffold.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Genome{config: guo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, guo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{genome.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
