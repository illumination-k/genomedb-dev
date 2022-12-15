// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"genomedb/ent/cds"
	"genomedb/ent/exon"
	"genomedb/ent/fiveprimeutr"
	"genomedb/ent/gene"
	"genomedb/ent/threeprimeutr"
	"genomedb/ent/transcript"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TranscriptCreate is the builder for creating a Transcript entity.
type TranscriptCreate struct {
	config
	mutation *TranscriptMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetStrand sets the "strand" field.
func (tc *TranscriptCreate) SetStrand(s string) *TranscriptCreate {
	tc.mutation.SetStrand(s)
	return tc
}

// SetType sets the "type" field.
func (tc *TranscriptCreate) SetType(s string) *TranscriptCreate {
	tc.mutation.SetType(s)
	return tc
}

// SetGenomeSeq sets the "genome_seq" field.
func (tc *TranscriptCreate) SetGenomeSeq(s string) *TranscriptCreate {
	tc.mutation.SetGenomeSeq(s)
	return tc
}

// SetTranscriptSeq sets the "transcript_seq" field.
func (tc *TranscriptCreate) SetTranscriptSeq(s string) *TranscriptCreate {
	tc.mutation.SetTranscriptSeq(s)
	return tc
}

// SetCdsSeq sets the "cds_seq" field.
func (tc *TranscriptCreate) SetCdsSeq(s string) *TranscriptCreate {
	tc.mutation.SetCdsSeq(s)
	return tc
}

// SetProteinSeq sets the "protein_seq" field.
func (tc *TranscriptCreate) SetProteinSeq(s string) *TranscriptCreate {
	tc.mutation.SetProteinSeq(s)
	return tc
}

// SetID sets the "id" field.
func (tc *TranscriptCreate) SetID(s string) *TranscriptCreate {
	tc.mutation.SetID(s)
	return tc
}

// SetGeneID sets the "gene" edge to the Gene entity by ID.
func (tc *TranscriptCreate) SetGeneID(id string) *TranscriptCreate {
	tc.mutation.SetGeneID(id)
	return tc
}

// SetNillableGeneID sets the "gene" edge to the Gene entity by ID if the given value is not nil.
func (tc *TranscriptCreate) SetNillableGeneID(id *string) *TranscriptCreate {
	if id != nil {
		tc = tc.SetGeneID(*id)
	}
	return tc
}

// SetGene sets the "gene" edge to the Gene entity.
func (tc *TranscriptCreate) SetGene(g *Gene) *TranscriptCreate {
	return tc.SetGeneID(g.ID)
}

// AddCdIDs adds the "cds" edge to the Cds entity by IDs.
func (tc *TranscriptCreate) AddCdIDs(ids ...int) *TranscriptCreate {
	tc.mutation.AddCdIDs(ids...)
	return tc
}

// AddCds adds the "cds" edges to the Cds entity.
func (tc *TranscriptCreate) AddCds(c ...*Cds) *TranscriptCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tc.AddCdIDs(ids...)
}

// AddExonIDs adds the "exon" edge to the Exon entity by IDs.
func (tc *TranscriptCreate) AddExonIDs(ids ...int) *TranscriptCreate {
	tc.mutation.AddExonIDs(ids...)
	return tc
}

// AddExon adds the "exon" edges to the Exon entity.
func (tc *TranscriptCreate) AddExon(e ...*Exon) *TranscriptCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return tc.AddExonIDs(ids...)
}

// AddFivePrimeUtrIDs adds the "five_prime_utr" edge to the FivePrimeUtr entity by IDs.
func (tc *TranscriptCreate) AddFivePrimeUtrIDs(ids ...int) *TranscriptCreate {
	tc.mutation.AddFivePrimeUtrIDs(ids...)
	return tc
}

// AddFivePrimeUtr adds the "five_prime_utr" edges to the FivePrimeUtr entity.
func (tc *TranscriptCreate) AddFivePrimeUtr(f ...*FivePrimeUtr) *TranscriptCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return tc.AddFivePrimeUtrIDs(ids...)
}

// AddThreePrimeUtrIDs adds the "three_prime_utr" edge to the ThreePrimeUtr entity by IDs.
func (tc *TranscriptCreate) AddThreePrimeUtrIDs(ids ...int) *TranscriptCreate {
	tc.mutation.AddThreePrimeUtrIDs(ids...)
	return tc
}

// AddThreePrimeUtr adds the "three_prime_utr" edges to the ThreePrimeUtr entity.
func (tc *TranscriptCreate) AddThreePrimeUtr(t ...*ThreePrimeUtr) *TranscriptCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tc.AddThreePrimeUtrIDs(ids...)
}

// Mutation returns the TranscriptMutation object of the builder.
func (tc *TranscriptCreate) Mutation() *TranscriptMutation {
	return tc.mutation
}

// Save creates the Transcript in the database.
func (tc *TranscriptCreate) Save(ctx context.Context) (*Transcript, error) {
	var (
		err  error
		node *Transcript
	)
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TranscriptMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			if node, err = tc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			if tc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Transcript)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TranscriptMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TranscriptCreate) SaveX(ctx context.Context) *Transcript {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TranscriptCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TranscriptCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TranscriptCreate) check() error {
	if _, ok := tc.mutation.Strand(); !ok {
		return &ValidationError{Name: "strand", err: errors.New(`ent: missing required field "Transcript.strand"`)}
	}
	if _, ok := tc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Transcript.type"`)}
	}
	if _, ok := tc.mutation.GenomeSeq(); !ok {
		return &ValidationError{Name: "genome_seq", err: errors.New(`ent: missing required field "Transcript.genome_seq"`)}
	}
	if _, ok := tc.mutation.TranscriptSeq(); !ok {
		return &ValidationError{Name: "transcript_seq", err: errors.New(`ent: missing required field "Transcript.transcript_seq"`)}
	}
	if _, ok := tc.mutation.CdsSeq(); !ok {
		return &ValidationError{Name: "cds_seq", err: errors.New(`ent: missing required field "Transcript.cds_seq"`)}
	}
	if _, ok := tc.mutation.ProteinSeq(); !ok {
		return &ValidationError{Name: "protein_seq", err: errors.New(`ent: missing required field "Transcript.protein_seq"`)}
	}
	return nil
}

func (tc *TranscriptCreate) sqlSave(ctx context.Context) (*Transcript, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Transcript.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (tc *TranscriptCreate) createSpec() (*Transcript, *sqlgraph.CreateSpec) {
	var (
		_node = &Transcript{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: transcript.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: transcript.FieldID,
			},
		}
	)
	_spec.OnConflict = tc.conflict
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tc.mutation.Strand(); ok {
		_spec.SetField(transcript.FieldStrand, field.TypeString, value)
		_node.Strand = value
	}
	if value, ok := tc.mutation.GetType(); ok {
		_spec.SetField(transcript.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := tc.mutation.GenomeSeq(); ok {
		_spec.SetField(transcript.FieldGenomeSeq, field.TypeString, value)
		_node.GenomeSeq = value
	}
	if value, ok := tc.mutation.TranscriptSeq(); ok {
		_spec.SetField(transcript.FieldTranscriptSeq, field.TypeString, value)
		_node.TranscriptSeq = value
	}
	if value, ok := tc.mutation.CdsSeq(); ok {
		_spec.SetField(transcript.FieldCdsSeq, field.TypeString, value)
		_node.CdsSeq = value
	}
	if value, ok := tc.mutation.ProteinSeq(); ok {
		_spec.SetField(transcript.FieldProteinSeq, field.TypeString, value)
		_node.ProteinSeq = value
	}
	if nodes := tc.mutation.GeneIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transcript.GeneTable,
			Columns: []string{transcript.GeneColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: gene.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.gene_transcripts = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.CdsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   transcript.CdsTable,
			Columns: []string{transcript.CdsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: cds.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.ExonIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   transcript.ExonTable,
			Columns: []string{transcript.ExonColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: exon.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.FivePrimeUtrIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   transcript.FivePrimeUtrTable,
			Columns: []string{transcript.FivePrimeUtrColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: fiveprimeutr.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.ThreePrimeUtrIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   transcript.ThreePrimeUtrTable,
			Columns: []string{transcript.ThreePrimeUtrColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: threeprimeutr.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Transcript.Create().
//		SetStrand(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TranscriptUpsert) {
//			SetStrand(v+v).
//		}).
//		Exec(ctx)
func (tc *TranscriptCreate) OnConflict(opts ...sql.ConflictOption) *TranscriptUpsertOne {
	tc.conflict = opts
	return &TranscriptUpsertOne{
		create: tc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Transcript.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tc *TranscriptCreate) OnConflictColumns(columns ...string) *TranscriptUpsertOne {
	tc.conflict = append(tc.conflict, sql.ConflictColumns(columns...))
	return &TranscriptUpsertOne{
		create: tc,
	}
}

type (
	// TranscriptUpsertOne is the builder for "upsert"-ing
	//  one Transcript node.
	TranscriptUpsertOne struct {
		create *TranscriptCreate
	}

	// TranscriptUpsert is the "OnConflict" setter.
	TranscriptUpsert struct {
		*sql.UpdateSet
	}
)

// SetStrand sets the "strand" field.
func (u *TranscriptUpsert) SetStrand(v string) *TranscriptUpsert {
	u.Set(transcript.FieldStrand, v)
	return u
}

// UpdateStrand sets the "strand" field to the value that was provided on create.
func (u *TranscriptUpsert) UpdateStrand() *TranscriptUpsert {
	u.SetExcluded(transcript.FieldStrand)
	return u
}

// SetType sets the "type" field.
func (u *TranscriptUpsert) SetType(v string) *TranscriptUpsert {
	u.Set(transcript.FieldType, v)
	return u
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *TranscriptUpsert) UpdateType() *TranscriptUpsert {
	u.SetExcluded(transcript.FieldType)
	return u
}

// SetGenomeSeq sets the "genome_seq" field.
func (u *TranscriptUpsert) SetGenomeSeq(v string) *TranscriptUpsert {
	u.Set(transcript.FieldGenomeSeq, v)
	return u
}

// UpdateGenomeSeq sets the "genome_seq" field to the value that was provided on create.
func (u *TranscriptUpsert) UpdateGenomeSeq() *TranscriptUpsert {
	u.SetExcluded(transcript.FieldGenomeSeq)
	return u
}

// SetTranscriptSeq sets the "transcript_seq" field.
func (u *TranscriptUpsert) SetTranscriptSeq(v string) *TranscriptUpsert {
	u.Set(transcript.FieldTranscriptSeq, v)
	return u
}

// UpdateTranscriptSeq sets the "transcript_seq" field to the value that was provided on create.
func (u *TranscriptUpsert) UpdateTranscriptSeq() *TranscriptUpsert {
	u.SetExcluded(transcript.FieldTranscriptSeq)
	return u
}

// SetCdsSeq sets the "cds_seq" field.
func (u *TranscriptUpsert) SetCdsSeq(v string) *TranscriptUpsert {
	u.Set(transcript.FieldCdsSeq, v)
	return u
}

// UpdateCdsSeq sets the "cds_seq" field to the value that was provided on create.
func (u *TranscriptUpsert) UpdateCdsSeq() *TranscriptUpsert {
	u.SetExcluded(transcript.FieldCdsSeq)
	return u
}

// SetProteinSeq sets the "protein_seq" field.
func (u *TranscriptUpsert) SetProteinSeq(v string) *TranscriptUpsert {
	u.Set(transcript.FieldProteinSeq, v)
	return u
}

// UpdateProteinSeq sets the "protein_seq" field to the value that was provided on create.
func (u *TranscriptUpsert) UpdateProteinSeq() *TranscriptUpsert {
	u.SetExcluded(transcript.FieldProteinSeq)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Transcript.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(transcript.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *TranscriptUpsertOne) UpdateNewValues() *TranscriptUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(transcript.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Transcript.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *TranscriptUpsertOne) Ignore() *TranscriptUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TranscriptUpsertOne) DoNothing() *TranscriptUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TranscriptCreate.OnConflict
// documentation for more info.
func (u *TranscriptUpsertOne) Update(set func(*TranscriptUpsert)) *TranscriptUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TranscriptUpsert{UpdateSet: update})
	}))
	return u
}

// SetStrand sets the "strand" field.
func (u *TranscriptUpsertOne) SetStrand(v string) *TranscriptUpsertOne {
	return u.Update(func(s *TranscriptUpsert) {
		s.SetStrand(v)
	})
}

// UpdateStrand sets the "strand" field to the value that was provided on create.
func (u *TranscriptUpsertOne) UpdateStrand() *TranscriptUpsertOne {
	return u.Update(func(s *TranscriptUpsert) {
		s.UpdateStrand()
	})
}

// SetType sets the "type" field.
func (u *TranscriptUpsertOne) SetType(v string) *TranscriptUpsertOne {
	return u.Update(func(s *TranscriptUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *TranscriptUpsertOne) UpdateType() *TranscriptUpsertOne {
	return u.Update(func(s *TranscriptUpsert) {
		s.UpdateType()
	})
}

// SetGenomeSeq sets the "genome_seq" field.
func (u *TranscriptUpsertOne) SetGenomeSeq(v string) *TranscriptUpsertOne {
	return u.Update(func(s *TranscriptUpsert) {
		s.SetGenomeSeq(v)
	})
}

// UpdateGenomeSeq sets the "genome_seq" field to the value that was provided on create.
func (u *TranscriptUpsertOne) UpdateGenomeSeq() *TranscriptUpsertOne {
	return u.Update(func(s *TranscriptUpsert) {
		s.UpdateGenomeSeq()
	})
}

// SetTranscriptSeq sets the "transcript_seq" field.
func (u *TranscriptUpsertOne) SetTranscriptSeq(v string) *TranscriptUpsertOne {
	return u.Update(func(s *TranscriptUpsert) {
		s.SetTranscriptSeq(v)
	})
}

// UpdateTranscriptSeq sets the "transcript_seq" field to the value that was provided on create.
func (u *TranscriptUpsertOne) UpdateTranscriptSeq() *TranscriptUpsertOne {
	return u.Update(func(s *TranscriptUpsert) {
		s.UpdateTranscriptSeq()
	})
}

// SetCdsSeq sets the "cds_seq" field.
func (u *TranscriptUpsertOne) SetCdsSeq(v string) *TranscriptUpsertOne {
	return u.Update(func(s *TranscriptUpsert) {
		s.SetCdsSeq(v)
	})
}

// UpdateCdsSeq sets the "cds_seq" field to the value that was provided on create.
func (u *TranscriptUpsertOne) UpdateCdsSeq() *TranscriptUpsertOne {
	return u.Update(func(s *TranscriptUpsert) {
		s.UpdateCdsSeq()
	})
}

// SetProteinSeq sets the "protein_seq" field.
func (u *TranscriptUpsertOne) SetProteinSeq(v string) *TranscriptUpsertOne {
	return u.Update(func(s *TranscriptUpsert) {
		s.SetProteinSeq(v)
	})
}

// UpdateProteinSeq sets the "protein_seq" field to the value that was provided on create.
func (u *TranscriptUpsertOne) UpdateProteinSeq() *TranscriptUpsertOne {
	return u.Update(func(s *TranscriptUpsert) {
		s.UpdateProteinSeq()
	})
}

// Exec executes the query.
func (u *TranscriptUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TranscriptCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TranscriptUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TranscriptUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: TranscriptUpsertOne.ID is not supported by MySQL driver. Use TranscriptUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TranscriptUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TranscriptCreateBulk is the builder for creating many Transcript entities in bulk.
type TranscriptCreateBulk struct {
	config
	builders []*TranscriptCreate
	conflict []sql.ConflictOption
}

// Save creates the Transcript entities in the database.
func (tcb *TranscriptCreateBulk) Save(ctx context.Context) ([]*Transcript, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Transcript, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TranscriptMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = tcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TranscriptCreateBulk) SaveX(ctx context.Context) []*Transcript {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TranscriptCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TranscriptCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Transcript.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TranscriptUpsert) {
//			SetStrand(v+v).
//		}).
//		Exec(ctx)
func (tcb *TranscriptCreateBulk) OnConflict(opts ...sql.ConflictOption) *TranscriptUpsertBulk {
	tcb.conflict = opts
	return &TranscriptUpsertBulk{
		create: tcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Transcript.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tcb *TranscriptCreateBulk) OnConflictColumns(columns ...string) *TranscriptUpsertBulk {
	tcb.conflict = append(tcb.conflict, sql.ConflictColumns(columns...))
	return &TranscriptUpsertBulk{
		create: tcb,
	}
}

// TranscriptUpsertBulk is the builder for "upsert"-ing
// a bulk of Transcript nodes.
type TranscriptUpsertBulk struct {
	create *TranscriptCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Transcript.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(transcript.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *TranscriptUpsertBulk) UpdateNewValues() *TranscriptUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(transcript.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Transcript.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *TranscriptUpsertBulk) Ignore() *TranscriptUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TranscriptUpsertBulk) DoNothing() *TranscriptUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TranscriptCreateBulk.OnConflict
// documentation for more info.
func (u *TranscriptUpsertBulk) Update(set func(*TranscriptUpsert)) *TranscriptUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TranscriptUpsert{UpdateSet: update})
	}))
	return u
}

// SetStrand sets the "strand" field.
func (u *TranscriptUpsertBulk) SetStrand(v string) *TranscriptUpsertBulk {
	return u.Update(func(s *TranscriptUpsert) {
		s.SetStrand(v)
	})
}

// UpdateStrand sets the "strand" field to the value that was provided on create.
func (u *TranscriptUpsertBulk) UpdateStrand() *TranscriptUpsertBulk {
	return u.Update(func(s *TranscriptUpsert) {
		s.UpdateStrand()
	})
}

// SetType sets the "type" field.
func (u *TranscriptUpsertBulk) SetType(v string) *TranscriptUpsertBulk {
	return u.Update(func(s *TranscriptUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *TranscriptUpsertBulk) UpdateType() *TranscriptUpsertBulk {
	return u.Update(func(s *TranscriptUpsert) {
		s.UpdateType()
	})
}

// SetGenomeSeq sets the "genome_seq" field.
func (u *TranscriptUpsertBulk) SetGenomeSeq(v string) *TranscriptUpsertBulk {
	return u.Update(func(s *TranscriptUpsert) {
		s.SetGenomeSeq(v)
	})
}

// UpdateGenomeSeq sets the "genome_seq" field to the value that was provided on create.
func (u *TranscriptUpsertBulk) UpdateGenomeSeq() *TranscriptUpsertBulk {
	return u.Update(func(s *TranscriptUpsert) {
		s.UpdateGenomeSeq()
	})
}

// SetTranscriptSeq sets the "transcript_seq" field.
func (u *TranscriptUpsertBulk) SetTranscriptSeq(v string) *TranscriptUpsertBulk {
	return u.Update(func(s *TranscriptUpsert) {
		s.SetTranscriptSeq(v)
	})
}

// UpdateTranscriptSeq sets the "transcript_seq" field to the value that was provided on create.
func (u *TranscriptUpsertBulk) UpdateTranscriptSeq() *TranscriptUpsertBulk {
	return u.Update(func(s *TranscriptUpsert) {
		s.UpdateTranscriptSeq()
	})
}

// SetCdsSeq sets the "cds_seq" field.
func (u *TranscriptUpsertBulk) SetCdsSeq(v string) *TranscriptUpsertBulk {
	return u.Update(func(s *TranscriptUpsert) {
		s.SetCdsSeq(v)
	})
}

// UpdateCdsSeq sets the "cds_seq" field to the value that was provided on create.
func (u *TranscriptUpsertBulk) UpdateCdsSeq() *TranscriptUpsertBulk {
	return u.Update(func(s *TranscriptUpsert) {
		s.UpdateCdsSeq()
	})
}

// SetProteinSeq sets the "protein_seq" field.
func (u *TranscriptUpsertBulk) SetProteinSeq(v string) *TranscriptUpsertBulk {
	return u.Update(func(s *TranscriptUpsert) {
		s.SetProteinSeq(v)
	})
}

// UpdateProteinSeq sets the "protein_seq" field to the value that was provided on create.
func (u *TranscriptUpsertBulk) UpdateProteinSeq() *TranscriptUpsertBulk {
	return u.Update(func(s *TranscriptUpsert) {
		s.UpdateProteinSeq()
	})
}

// Exec executes the query.
func (u *TranscriptUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TranscriptCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TranscriptCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TranscriptUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
