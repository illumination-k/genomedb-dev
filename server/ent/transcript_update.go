// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"genomedb/bio/gffio"
	"genomedb/ent/domainannotation"
	"genomedb/ent/goterm"
	"genomedb/ent/locus"
	"genomedb/ent/predicate"
	"genomedb/ent/transcript"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
)

// TranscriptUpdate is the builder for updating Transcript entities.
type TranscriptUpdate struct {
	config
	hooks    []Hook
	mutation *TranscriptMutation
}

// Where appends a list predicates to the TranscriptUpdate builder.
func (tu *TranscriptUpdate) Where(ps ...predicate.Transcript) *TranscriptUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetSeqname sets the "seqname" field.
func (tu *TranscriptUpdate) SetSeqname(s string) *TranscriptUpdate {
	tu.mutation.SetSeqname(s)
	return tu
}

// SetStrand sets the "strand" field.
func (tu *TranscriptUpdate) SetStrand(s string) *TranscriptUpdate {
	tu.mutation.SetStrand(s)
	return tu
}

// SetType sets the "type" field.
func (tu *TranscriptUpdate) SetType(s string) *TranscriptUpdate {
	tu.mutation.SetType(s)
	return tu
}

// SetSource sets the "source" field.
func (tu *TranscriptUpdate) SetSource(s string) *TranscriptUpdate {
	tu.mutation.SetSource(s)
	return tu
}

// SetStart sets the "start" field.
func (tu *TranscriptUpdate) SetStart(i int32) *TranscriptUpdate {
	tu.mutation.ResetStart()
	tu.mutation.SetStart(i)
	return tu
}

// AddStart adds i to the "start" field.
func (tu *TranscriptUpdate) AddStart(i int32) *TranscriptUpdate {
	tu.mutation.AddStart(i)
	return tu
}

// SetEnd sets the "end" field.
func (tu *TranscriptUpdate) SetEnd(i int32) *TranscriptUpdate {
	tu.mutation.ResetEnd()
	tu.mutation.SetEnd(i)
	return tu
}

// AddEnd adds i to the "end" field.
func (tu *TranscriptUpdate) AddEnd(i int32) *TranscriptUpdate {
	tu.mutation.AddEnd(i)
	return tu
}

// SetExon sets the "exon" field.
func (tu *TranscriptUpdate) SetExon(gr []gffio.GffRecord) *TranscriptUpdate {
	tu.mutation.SetExon(gr)
	return tu
}

// AppendExon appends gr to the "exon" field.
func (tu *TranscriptUpdate) AppendExon(gr []gffio.GffRecord) *TranscriptUpdate {
	tu.mutation.AppendExon(gr)
	return tu
}

// SetFivePrimeUtr sets the "five_prime_utr" field.
func (tu *TranscriptUpdate) SetFivePrimeUtr(gr []gffio.GffRecord) *TranscriptUpdate {
	tu.mutation.SetFivePrimeUtr(gr)
	return tu
}

// AppendFivePrimeUtr appends gr to the "five_prime_utr" field.
func (tu *TranscriptUpdate) AppendFivePrimeUtr(gr []gffio.GffRecord) *TranscriptUpdate {
	tu.mutation.AppendFivePrimeUtr(gr)
	return tu
}

// SetThreePrimeUtr sets the "three_prime_utr" field.
func (tu *TranscriptUpdate) SetThreePrimeUtr(gr []gffio.GffRecord) *TranscriptUpdate {
	tu.mutation.SetThreePrimeUtr(gr)
	return tu
}

// AppendThreePrimeUtr appends gr to the "three_prime_utr" field.
func (tu *TranscriptUpdate) AppendThreePrimeUtr(gr []gffio.GffRecord) *TranscriptUpdate {
	tu.mutation.AppendThreePrimeUtr(gr)
	return tu
}

// SetCds sets the "cds" field.
func (tu *TranscriptUpdate) SetCds(gr []gffio.GffRecord) *TranscriptUpdate {
	tu.mutation.SetCds(gr)
	return tu
}

// AppendCds appends gr to the "cds" field.
func (tu *TranscriptUpdate) AppendCds(gr []gffio.GffRecord) *TranscriptUpdate {
	tu.mutation.AppendCds(gr)
	return tu
}

// SetGenomicSequence sets the "genomic_sequence" field.
func (tu *TranscriptUpdate) SetGenomicSequence(s string) *TranscriptUpdate {
	tu.mutation.SetGenomicSequence(s)
	return tu
}

// SetExonSequence sets the "exon_sequence" field.
func (tu *TranscriptUpdate) SetExonSequence(s string) *TranscriptUpdate {
	tu.mutation.SetExonSequence(s)
	return tu
}

// SetCdsSequence sets the "cds_sequence" field.
func (tu *TranscriptUpdate) SetCdsSequence(s string) *TranscriptUpdate {
	tu.mutation.SetCdsSequence(s)
	return tu
}

// SetProteinSequence sets the "protein_sequence" field.
func (tu *TranscriptUpdate) SetProteinSequence(s string) *TranscriptUpdate {
	tu.mutation.SetProteinSequence(s)
	return tu
}

// SetLocusID sets the "locus" edge to the Locus entity by ID.
func (tu *TranscriptUpdate) SetLocusID(id string) *TranscriptUpdate {
	tu.mutation.SetLocusID(id)
	return tu
}

// SetNillableLocusID sets the "locus" edge to the Locus entity by ID if the given value is not nil.
func (tu *TranscriptUpdate) SetNillableLocusID(id *string) *TranscriptUpdate {
	if id != nil {
		tu = tu.SetLocusID(*id)
	}
	return tu
}

// SetLocus sets the "locus" edge to the Locus entity.
func (tu *TranscriptUpdate) SetLocus(l *Locus) *TranscriptUpdate {
	return tu.SetLocusID(l.ID)
}

// AddGotermIDs adds the "goterms" edge to the GoTerm entity by IDs.
func (tu *TranscriptUpdate) AddGotermIDs(ids ...string) *TranscriptUpdate {
	tu.mutation.AddGotermIDs(ids...)
	return tu
}

// AddGoterms adds the "goterms" edges to the GoTerm entity.
func (tu *TranscriptUpdate) AddGoterms(g ...*GoTerm) *TranscriptUpdate {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return tu.AddGotermIDs(ids...)
}

// AddDomainIDs adds the "domains" edge to the DomainAnnotation entity by IDs.
func (tu *TranscriptUpdate) AddDomainIDs(ids ...string) *TranscriptUpdate {
	tu.mutation.AddDomainIDs(ids...)
	return tu
}

// AddDomains adds the "domains" edges to the DomainAnnotation entity.
func (tu *TranscriptUpdate) AddDomains(d ...*DomainAnnotation) *TranscriptUpdate {
	ids := make([]string, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return tu.AddDomainIDs(ids...)
}

// Mutation returns the TranscriptMutation object of the builder.
func (tu *TranscriptUpdate) Mutation() *TranscriptMutation {
	return tu.mutation
}

// ClearLocus clears the "locus" edge to the Locus entity.
func (tu *TranscriptUpdate) ClearLocus() *TranscriptUpdate {
	tu.mutation.ClearLocus()
	return tu
}

// ClearGoterms clears all "goterms" edges to the GoTerm entity.
func (tu *TranscriptUpdate) ClearGoterms() *TranscriptUpdate {
	tu.mutation.ClearGoterms()
	return tu
}

// RemoveGotermIDs removes the "goterms" edge to GoTerm entities by IDs.
func (tu *TranscriptUpdate) RemoveGotermIDs(ids ...string) *TranscriptUpdate {
	tu.mutation.RemoveGotermIDs(ids...)
	return tu
}

// RemoveGoterms removes "goterms" edges to GoTerm entities.
func (tu *TranscriptUpdate) RemoveGoterms(g ...*GoTerm) *TranscriptUpdate {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return tu.RemoveGotermIDs(ids...)
}

// ClearDomains clears all "domains" edges to the DomainAnnotation entity.
func (tu *TranscriptUpdate) ClearDomains() *TranscriptUpdate {
	tu.mutation.ClearDomains()
	return tu
}

// RemoveDomainIDs removes the "domains" edge to DomainAnnotation entities by IDs.
func (tu *TranscriptUpdate) RemoveDomainIDs(ids ...string) *TranscriptUpdate {
	tu.mutation.RemoveDomainIDs(ids...)
	return tu
}

// RemoveDomains removes "domains" edges to DomainAnnotation entities.
func (tu *TranscriptUpdate) RemoveDomains(d ...*DomainAnnotation) *TranscriptUpdate {
	ids := make([]string, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return tu.RemoveDomainIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TranscriptUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tu.hooks) == 0 {
		if err = tu.check(); err != nil {
			return 0, err
		}
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TranscriptMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tu.check(); err != nil {
				return 0, err
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TranscriptUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TranscriptUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TranscriptUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TranscriptUpdate) check() error {
	if v, ok := tu.mutation.Start(); ok {
		if err := transcript.StartValidator(v); err != nil {
			return &ValidationError{Name: "start", err: fmt.Errorf(`ent: validator failed for field "Transcript.start": %w`, err)}
		}
	}
	if v, ok := tu.mutation.End(); ok {
		if err := transcript.EndValidator(v); err != nil {
			return &ValidationError{Name: "end", err: fmt.Errorf(`ent: validator failed for field "Transcript.end": %w`, err)}
		}
	}
	return nil
}

func (tu *TranscriptUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   transcript.Table,
			Columns: transcript.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: transcript.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Seqname(); ok {
		_spec.SetField(transcript.FieldSeqname, field.TypeString, value)
	}
	if value, ok := tu.mutation.Strand(); ok {
		_spec.SetField(transcript.FieldStrand, field.TypeString, value)
	}
	if value, ok := tu.mutation.GetType(); ok {
		_spec.SetField(transcript.FieldType, field.TypeString, value)
	}
	if value, ok := tu.mutation.Source(); ok {
		_spec.SetField(transcript.FieldSource, field.TypeString, value)
	}
	if value, ok := tu.mutation.Start(); ok {
		_spec.SetField(transcript.FieldStart, field.TypeInt32, value)
	}
	if value, ok := tu.mutation.AddedStart(); ok {
		_spec.AddField(transcript.FieldStart, field.TypeInt32, value)
	}
	if value, ok := tu.mutation.End(); ok {
		_spec.SetField(transcript.FieldEnd, field.TypeInt32, value)
	}
	if value, ok := tu.mutation.AddedEnd(); ok {
		_spec.AddField(transcript.FieldEnd, field.TypeInt32, value)
	}
	if value, ok := tu.mutation.Exon(); ok {
		_spec.SetField(transcript.FieldExon, field.TypeJSON, value)
	}
	if value, ok := tu.mutation.AppendedExon(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, transcript.FieldExon, value)
		})
	}
	if value, ok := tu.mutation.FivePrimeUtr(); ok {
		_spec.SetField(transcript.FieldFivePrimeUtr, field.TypeJSON, value)
	}
	if value, ok := tu.mutation.AppendedFivePrimeUtr(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, transcript.FieldFivePrimeUtr, value)
		})
	}
	if value, ok := tu.mutation.ThreePrimeUtr(); ok {
		_spec.SetField(transcript.FieldThreePrimeUtr, field.TypeJSON, value)
	}
	if value, ok := tu.mutation.AppendedThreePrimeUtr(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, transcript.FieldThreePrimeUtr, value)
		})
	}
	if value, ok := tu.mutation.Cds(); ok {
		_spec.SetField(transcript.FieldCds, field.TypeJSON, value)
	}
	if value, ok := tu.mutation.AppendedCds(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, transcript.FieldCds, value)
		})
	}
	if value, ok := tu.mutation.GenomicSequence(); ok {
		_spec.SetField(transcript.FieldGenomicSequence, field.TypeString, value)
	}
	if value, ok := tu.mutation.ExonSequence(); ok {
		_spec.SetField(transcript.FieldExonSequence, field.TypeString, value)
	}
	if value, ok := tu.mutation.CdsSequence(); ok {
		_spec.SetField(transcript.FieldCdsSequence, field.TypeString, value)
	}
	if value, ok := tu.mutation.ProteinSequence(); ok {
		_spec.SetField(transcript.FieldProteinSequence, field.TypeString, value)
	}
	if tu.mutation.LocusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transcript.LocusTable,
			Columns: []string{transcript.LocusColumn},
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
	if nodes := tu.mutation.LocusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transcript.LocusTable,
			Columns: []string{transcript.LocusColumn},
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
	if tu.mutation.GotermsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   transcript.GotermsTable,
			Columns: transcript.GotermsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: goterm.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedGotermsIDs(); len(nodes) > 0 && !tu.mutation.GotermsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   transcript.GotermsTable,
			Columns: transcript.GotermsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: goterm.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.GotermsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   transcript.GotermsTable,
			Columns: transcript.GotermsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: goterm.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.DomainsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   transcript.DomainsTable,
			Columns: transcript.DomainsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: domainannotation.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedDomainsIDs(); len(nodes) > 0 && !tu.mutation.DomainsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   transcript.DomainsTable,
			Columns: transcript.DomainsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: domainannotation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.DomainsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   transcript.DomainsTable,
			Columns: transcript.DomainsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: domainannotation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{transcript.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// TranscriptUpdateOne is the builder for updating a single Transcript entity.
type TranscriptUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TranscriptMutation
}

// SetSeqname sets the "seqname" field.
func (tuo *TranscriptUpdateOne) SetSeqname(s string) *TranscriptUpdateOne {
	tuo.mutation.SetSeqname(s)
	return tuo
}

// SetStrand sets the "strand" field.
func (tuo *TranscriptUpdateOne) SetStrand(s string) *TranscriptUpdateOne {
	tuo.mutation.SetStrand(s)
	return tuo
}

// SetType sets the "type" field.
func (tuo *TranscriptUpdateOne) SetType(s string) *TranscriptUpdateOne {
	tuo.mutation.SetType(s)
	return tuo
}

// SetSource sets the "source" field.
func (tuo *TranscriptUpdateOne) SetSource(s string) *TranscriptUpdateOne {
	tuo.mutation.SetSource(s)
	return tuo
}

// SetStart sets the "start" field.
func (tuo *TranscriptUpdateOne) SetStart(i int32) *TranscriptUpdateOne {
	tuo.mutation.ResetStart()
	tuo.mutation.SetStart(i)
	return tuo
}

// AddStart adds i to the "start" field.
func (tuo *TranscriptUpdateOne) AddStart(i int32) *TranscriptUpdateOne {
	tuo.mutation.AddStart(i)
	return tuo
}

// SetEnd sets the "end" field.
func (tuo *TranscriptUpdateOne) SetEnd(i int32) *TranscriptUpdateOne {
	tuo.mutation.ResetEnd()
	tuo.mutation.SetEnd(i)
	return tuo
}

// AddEnd adds i to the "end" field.
func (tuo *TranscriptUpdateOne) AddEnd(i int32) *TranscriptUpdateOne {
	tuo.mutation.AddEnd(i)
	return tuo
}

// SetExon sets the "exon" field.
func (tuo *TranscriptUpdateOne) SetExon(gr []gffio.GffRecord) *TranscriptUpdateOne {
	tuo.mutation.SetExon(gr)
	return tuo
}

// AppendExon appends gr to the "exon" field.
func (tuo *TranscriptUpdateOne) AppendExon(gr []gffio.GffRecord) *TranscriptUpdateOne {
	tuo.mutation.AppendExon(gr)
	return tuo
}

// SetFivePrimeUtr sets the "five_prime_utr" field.
func (tuo *TranscriptUpdateOne) SetFivePrimeUtr(gr []gffio.GffRecord) *TranscriptUpdateOne {
	tuo.mutation.SetFivePrimeUtr(gr)
	return tuo
}

// AppendFivePrimeUtr appends gr to the "five_prime_utr" field.
func (tuo *TranscriptUpdateOne) AppendFivePrimeUtr(gr []gffio.GffRecord) *TranscriptUpdateOne {
	tuo.mutation.AppendFivePrimeUtr(gr)
	return tuo
}

// SetThreePrimeUtr sets the "three_prime_utr" field.
func (tuo *TranscriptUpdateOne) SetThreePrimeUtr(gr []gffio.GffRecord) *TranscriptUpdateOne {
	tuo.mutation.SetThreePrimeUtr(gr)
	return tuo
}

// AppendThreePrimeUtr appends gr to the "three_prime_utr" field.
func (tuo *TranscriptUpdateOne) AppendThreePrimeUtr(gr []gffio.GffRecord) *TranscriptUpdateOne {
	tuo.mutation.AppendThreePrimeUtr(gr)
	return tuo
}

// SetCds sets the "cds" field.
func (tuo *TranscriptUpdateOne) SetCds(gr []gffio.GffRecord) *TranscriptUpdateOne {
	tuo.mutation.SetCds(gr)
	return tuo
}

// AppendCds appends gr to the "cds" field.
func (tuo *TranscriptUpdateOne) AppendCds(gr []gffio.GffRecord) *TranscriptUpdateOne {
	tuo.mutation.AppendCds(gr)
	return tuo
}

// SetGenomicSequence sets the "genomic_sequence" field.
func (tuo *TranscriptUpdateOne) SetGenomicSequence(s string) *TranscriptUpdateOne {
	tuo.mutation.SetGenomicSequence(s)
	return tuo
}

// SetExonSequence sets the "exon_sequence" field.
func (tuo *TranscriptUpdateOne) SetExonSequence(s string) *TranscriptUpdateOne {
	tuo.mutation.SetExonSequence(s)
	return tuo
}

// SetCdsSequence sets the "cds_sequence" field.
func (tuo *TranscriptUpdateOne) SetCdsSequence(s string) *TranscriptUpdateOne {
	tuo.mutation.SetCdsSequence(s)
	return tuo
}

// SetProteinSequence sets the "protein_sequence" field.
func (tuo *TranscriptUpdateOne) SetProteinSequence(s string) *TranscriptUpdateOne {
	tuo.mutation.SetProteinSequence(s)
	return tuo
}

// SetLocusID sets the "locus" edge to the Locus entity by ID.
func (tuo *TranscriptUpdateOne) SetLocusID(id string) *TranscriptUpdateOne {
	tuo.mutation.SetLocusID(id)
	return tuo
}

// SetNillableLocusID sets the "locus" edge to the Locus entity by ID if the given value is not nil.
func (tuo *TranscriptUpdateOne) SetNillableLocusID(id *string) *TranscriptUpdateOne {
	if id != nil {
		tuo = tuo.SetLocusID(*id)
	}
	return tuo
}

// SetLocus sets the "locus" edge to the Locus entity.
func (tuo *TranscriptUpdateOne) SetLocus(l *Locus) *TranscriptUpdateOne {
	return tuo.SetLocusID(l.ID)
}

// AddGotermIDs adds the "goterms" edge to the GoTerm entity by IDs.
func (tuo *TranscriptUpdateOne) AddGotermIDs(ids ...string) *TranscriptUpdateOne {
	tuo.mutation.AddGotermIDs(ids...)
	return tuo
}

// AddGoterms adds the "goterms" edges to the GoTerm entity.
func (tuo *TranscriptUpdateOne) AddGoterms(g ...*GoTerm) *TranscriptUpdateOne {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return tuo.AddGotermIDs(ids...)
}

// AddDomainIDs adds the "domains" edge to the DomainAnnotation entity by IDs.
func (tuo *TranscriptUpdateOne) AddDomainIDs(ids ...string) *TranscriptUpdateOne {
	tuo.mutation.AddDomainIDs(ids...)
	return tuo
}

// AddDomains adds the "domains" edges to the DomainAnnotation entity.
func (tuo *TranscriptUpdateOne) AddDomains(d ...*DomainAnnotation) *TranscriptUpdateOne {
	ids := make([]string, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return tuo.AddDomainIDs(ids...)
}

// Mutation returns the TranscriptMutation object of the builder.
func (tuo *TranscriptUpdateOne) Mutation() *TranscriptMutation {
	return tuo.mutation
}

// ClearLocus clears the "locus" edge to the Locus entity.
func (tuo *TranscriptUpdateOne) ClearLocus() *TranscriptUpdateOne {
	tuo.mutation.ClearLocus()
	return tuo
}

// ClearGoterms clears all "goterms" edges to the GoTerm entity.
func (tuo *TranscriptUpdateOne) ClearGoterms() *TranscriptUpdateOne {
	tuo.mutation.ClearGoterms()
	return tuo
}

// RemoveGotermIDs removes the "goterms" edge to GoTerm entities by IDs.
func (tuo *TranscriptUpdateOne) RemoveGotermIDs(ids ...string) *TranscriptUpdateOne {
	tuo.mutation.RemoveGotermIDs(ids...)
	return tuo
}

// RemoveGoterms removes "goterms" edges to GoTerm entities.
func (tuo *TranscriptUpdateOne) RemoveGoterms(g ...*GoTerm) *TranscriptUpdateOne {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return tuo.RemoveGotermIDs(ids...)
}

// ClearDomains clears all "domains" edges to the DomainAnnotation entity.
func (tuo *TranscriptUpdateOne) ClearDomains() *TranscriptUpdateOne {
	tuo.mutation.ClearDomains()
	return tuo
}

// RemoveDomainIDs removes the "domains" edge to DomainAnnotation entities by IDs.
func (tuo *TranscriptUpdateOne) RemoveDomainIDs(ids ...string) *TranscriptUpdateOne {
	tuo.mutation.RemoveDomainIDs(ids...)
	return tuo
}

// RemoveDomains removes "domains" edges to DomainAnnotation entities.
func (tuo *TranscriptUpdateOne) RemoveDomains(d ...*DomainAnnotation) *TranscriptUpdateOne {
	ids := make([]string, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return tuo.RemoveDomainIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TranscriptUpdateOne) Select(field string, fields ...string) *TranscriptUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Transcript entity.
func (tuo *TranscriptUpdateOne) Save(ctx context.Context) (*Transcript, error) {
	var (
		err  error
		node *Transcript
	)
	if len(tuo.hooks) == 0 {
		if err = tuo.check(); err != nil {
			return nil, err
		}
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TranscriptMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tuo.check(); err != nil {
				return nil, err
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tuo.mutation)
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

// SaveX is like Save, but panics if an error occurs.
func (tuo *TranscriptUpdateOne) SaveX(ctx context.Context) *Transcript {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TranscriptUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TranscriptUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TranscriptUpdateOne) check() error {
	if v, ok := tuo.mutation.Start(); ok {
		if err := transcript.StartValidator(v); err != nil {
			return &ValidationError{Name: "start", err: fmt.Errorf(`ent: validator failed for field "Transcript.start": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.End(); ok {
		if err := transcript.EndValidator(v); err != nil {
			return &ValidationError{Name: "end", err: fmt.Errorf(`ent: validator failed for field "Transcript.end": %w`, err)}
		}
	}
	return nil
}

func (tuo *TranscriptUpdateOne) sqlSave(ctx context.Context) (_node *Transcript, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   transcript.Table,
			Columns: transcript.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: transcript.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Transcript.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, transcript.FieldID)
		for _, f := range fields {
			if !transcript.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != transcript.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Seqname(); ok {
		_spec.SetField(transcript.FieldSeqname, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Strand(); ok {
		_spec.SetField(transcript.FieldStrand, field.TypeString, value)
	}
	if value, ok := tuo.mutation.GetType(); ok {
		_spec.SetField(transcript.FieldType, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Source(); ok {
		_spec.SetField(transcript.FieldSource, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Start(); ok {
		_spec.SetField(transcript.FieldStart, field.TypeInt32, value)
	}
	if value, ok := tuo.mutation.AddedStart(); ok {
		_spec.AddField(transcript.FieldStart, field.TypeInt32, value)
	}
	if value, ok := tuo.mutation.End(); ok {
		_spec.SetField(transcript.FieldEnd, field.TypeInt32, value)
	}
	if value, ok := tuo.mutation.AddedEnd(); ok {
		_spec.AddField(transcript.FieldEnd, field.TypeInt32, value)
	}
	if value, ok := tuo.mutation.Exon(); ok {
		_spec.SetField(transcript.FieldExon, field.TypeJSON, value)
	}
	if value, ok := tuo.mutation.AppendedExon(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, transcript.FieldExon, value)
		})
	}
	if value, ok := tuo.mutation.FivePrimeUtr(); ok {
		_spec.SetField(transcript.FieldFivePrimeUtr, field.TypeJSON, value)
	}
	if value, ok := tuo.mutation.AppendedFivePrimeUtr(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, transcript.FieldFivePrimeUtr, value)
		})
	}
	if value, ok := tuo.mutation.ThreePrimeUtr(); ok {
		_spec.SetField(transcript.FieldThreePrimeUtr, field.TypeJSON, value)
	}
	if value, ok := tuo.mutation.AppendedThreePrimeUtr(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, transcript.FieldThreePrimeUtr, value)
		})
	}
	if value, ok := tuo.mutation.Cds(); ok {
		_spec.SetField(transcript.FieldCds, field.TypeJSON, value)
	}
	if value, ok := tuo.mutation.AppendedCds(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, transcript.FieldCds, value)
		})
	}
	if value, ok := tuo.mutation.GenomicSequence(); ok {
		_spec.SetField(transcript.FieldGenomicSequence, field.TypeString, value)
	}
	if value, ok := tuo.mutation.ExonSequence(); ok {
		_spec.SetField(transcript.FieldExonSequence, field.TypeString, value)
	}
	if value, ok := tuo.mutation.CdsSequence(); ok {
		_spec.SetField(transcript.FieldCdsSequence, field.TypeString, value)
	}
	if value, ok := tuo.mutation.ProteinSequence(); ok {
		_spec.SetField(transcript.FieldProteinSequence, field.TypeString, value)
	}
	if tuo.mutation.LocusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transcript.LocusTable,
			Columns: []string{transcript.LocusColumn},
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
	if nodes := tuo.mutation.LocusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transcript.LocusTable,
			Columns: []string{transcript.LocusColumn},
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
	if tuo.mutation.GotermsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   transcript.GotermsTable,
			Columns: transcript.GotermsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: goterm.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedGotermsIDs(); len(nodes) > 0 && !tuo.mutation.GotermsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   transcript.GotermsTable,
			Columns: transcript.GotermsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: goterm.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.GotermsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   transcript.GotermsTable,
			Columns: transcript.GotermsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: goterm.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.DomainsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   transcript.DomainsTable,
			Columns: transcript.DomainsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: domainannotation.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedDomainsIDs(); len(nodes) > 0 && !tuo.mutation.DomainsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   transcript.DomainsTable,
			Columns: transcript.DomainsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: domainannotation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.DomainsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   transcript.DomainsTable,
			Columns: transcript.DomainsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: domainannotation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Transcript{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{transcript.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
