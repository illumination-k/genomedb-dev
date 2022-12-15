// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"genomedb/ent/cds"
	"genomedb/ent/predicate"
	"genomedb/ent/transcript"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CdsUpdate is the builder for updating Cds entities.
type CdsUpdate struct {
	config
	hooks    []Hook
	mutation *CdsMutation
}

// Where appends a list predicates to the CdsUpdate builder.
func (cu *CdsUpdate) Where(ps ...predicate.Cds) *CdsUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetSeqname sets the "seqname" field.
func (cu *CdsUpdate) SetSeqname(s string) *CdsUpdate {
	cu.mutation.SetSeqname(s)
	return cu
}

// SetStart sets the "start" field.
func (cu *CdsUpdate) SetStart(i int32) *CdsUpdate {
	cu.mutation.ResetStart()
	cu.mutation.SetStart(i)
	return cu
}

// AddStart adds i to the "start" field.
func (cu *CdsUpdate) AddStart(i int32) *CdsUpdate {
	cu.mutation.AddStart(i)
	return cu
}

// SetEnd sets the "end" field.
func (cu *CdsUpdate) SetEnd(i int32) *CdsUpdate {
	cu.mutation.ResetEnd()
	cu.mutation.SetEnd(i)
	return cu
}

// AddEnd adds i to the "end" field.
func (cu *CdsUpdate) AddEnd(i int32) *CdsUpdate {
	cu.mutation.AddEnd(i)
	return cu
}

// SetPhase sets the "phase" field.
func (cu *CdsUpdate) SetPhase(i int8) *CdsUpdate {
	cu.mutation.ResetPhase()
	cu.mutation.SetPhase(i)
	return cu
}

// AddPhase adds i to the "phase" field.
func (cu *CdsUpdate) AddPhase(i int8) *CdsUpdate {
	cu.mutation.AddPhase(i)
	return cu
}

// SetStrand sets the "strand" field.
func (cu *CdsUpdate) SetStrand(s string) *CdsUpdate {
	cu.mutation.SetStrand(s)
	return cu
}

// SetTranscriptID sets the "transcript" edge to the Transcript entity by ID.
func (cu *CdsUpdate) SetTranscriptID(id string) *CdsUpdate {
	cu.mutation.SetTranscriptID(id)
	return cu
}

// SetNillableTranscriptID sets the "transcript" edge to the Transcript entity by ID if the given value is not nil.
func (cu *CdsUpdate) SetNillableTranscriptID(id *string) *CdsUpdate {
	if id != nil {
		cu = cu.SetTranscriptID(*id)
	}
	return cu
}

// SetTranscript sets the "transcript" edge to the Transcript entity.
func (cu *CdsUpdate) SetTranscript(t *Transcript) *CdsUpdate {
	return cu.SetTranscriptID(t.ID)
}

// Mutation returns the CdsMutation object of the builder.
func (cu *CdsUpdate) Mutation() *CdsMutation {
	return cu.mutation
}

// ClearTranscript clears the "transcript" edge to the Transcript entity.
func (cu *CdsUpdate) ClearTranscript() *CdsUpdate {
	cu.mutation.ClearTranscript()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CdsUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		if err = cu.check(); err != nil {
			return 0, err
		}
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CdsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cu.check(); err != nil {
				return 0, err
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CdsUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CdsUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CdsUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CdsUpdate) check() error {
	if v, ok := cu.mutation.Start(); ok {
		if err := cds.StartValidator(v); err != nil {
			return &ValidationError{Name: "start", err: fmt.Errorf(`ent: validator failed for field "Cds.start": %w`, err)}
		}
	}
	if v, ok := cu.mutation.End(); ok {
		if err := cds.EndValidator(v); err != nil {
			return &ValidationError{Name: "end", err: fmt.Errorf(`ent: validator failed for field "Cds.end": %w`, err)}
		}
	}
	return nil
}

func (cu *CdsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cds.Table,
			Columns: cds.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: cds.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Seqname(); ok {
		_spec.SetField(cds.FieldSeqname, field.TypeString, value)
	}
	if value, ok := cu.mutation.Start(); ok {
		_spec.SetField(cds.FieldStart, field.TypeInt32, value)
	}
	if value, ok := cu.mutation.AddedStart(); ok {
		_spec.AddField(cds.FieldStart, field.TypeInt32, value)
	}
	if value, ok := cu.mutation.End(); ok {
		_spec.SetField(cds.FieldEnd, field.TypeInt32, value)
	}
	if value, ok := cu.mutation.AddedEnd(); ok {
		_spec.AddField(cds.FieldEnd, field.TypeInt32, value)
	}
	if value, ok := cu.mutation.Phase(); ok {
		_spec.SetField(cds.FieldPhase, field.TypeInt8, value)
	}
	if value, ok := cu.mutation.AddedPhase(); ok {
		_spec.AddField(cds.FieldPhase, field.TypeInt8, value)
	}
	if value, ok := cu.mutation.Strand(); ok {
		_spec.SetField(cds.FieldStrand, field.TypeString, value)
	}
	if cu.mutation.TranscriptCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cds.TranscriptTable,
			Columns: []string{cds.TranscriptColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: transcript.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.TranscriptIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cds.TranscriptTable,
			Columns: []string{cds.TranscriptColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: transcript.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cds.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CdsUpdateOne is the builder for updating a single Cds entity.
type CdsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CdsMutation
}

// SetSeqname sets the "seqname" field.
func (cuo *CdsUpdateOne) SetSeqname(s string) *CdsUpdateOne {
	cuo.mutation.SetSeqname(s)
	return cuo
}

// SetStart sets the "start" field.
func (cuo *CdsUpdateOne) SetStart(i int32) *CdsUpdateOne {
	cuo.mutation.ResetStart()
	cuo.mutation.SetStart(i)
	return cuo
}

// AddStart adds i to the "start" field.
func (cuo *CdsUpdateOne) AddStart(i int32) *CdsUpdateOne {
	cuo.mutation.AddStart(i)
	return cuo
}

// SetEnd sets the "end" field.
func (cuo *CdsUpdateOne) SetEnd(i int32) *CdsUpdateOne {
	cuo.mutation.ResetEnd()
	cuo.mutation.SetEnd(i)
	return cuo
}

// AddEnd adds i to the "end" field.
func (cuo *CdsUpdateOne) AddEnd(i int32) *CdsUpdateOne {
	cuo.mutation.AddEnd(i)
	return cuo
}

// SetPhase sets the "phase" field.
func (cuo *CdsUpdateOne) SetPhase(i int8) *CdsUpdateOne {
	cuo.mutation.ResetPhase()
	cuo.mutation.SetPhase(i)
	return cuo
}

// AddPhase adds i to the "phase" field.
func (cuo *CdsUpdateOne) AddPhase(i int8) *CdsUpdateOne {
	cuo.mutation.AddPhase(i)
	return cuo
}

// SetStrand sets the "strand" field.
func (cuo *CdsUpdateOne) SetStrand(s string) *CdsUpdateOne {
	cuo.mutation.SetStrand(s)
	return cuo
}

// SetTranscriptID sets the "transcript" edge to the Transcript entity by ID.
func (cuo *CdsUpdateOne) SetTranscriptID(id string) *CdsUpdateOne {
	cuo.mutation.SetTranscriptID(id)
	return cuo
}

// SetNillableTranscriptID sets the "transcript" edge to the Transcript entity by ID if the given value is not nil.
func (cuo *CdsUpdateOne) SetNillableTranscriptID(id *string) *CdsUpdateOne {
	if id != nil {
		cuo = cuo.SetTranscriptID(*id)
	}
	return cuo
}

// SetTranscript sets the "transcript" edge to the Transcript entity.
func (cuo *CdsUpdateOne) SetTranscript(t *Transcript) *CdsUpdateOne {
	return cuo.SetTranscriptID(t.ID)
}

// Mutation returns the CdsMutation object of the builder.
func (cuo *CdsUpdateOne) Mutation() *CdsMutation {
	return cuo.mutation
}

// ClearTranscript clears the "transcript" edge to the Transcript entity.
func (cuo *CdsUpdateOne) ClearTranscript() *CdsUpdateOne {
	cuo.mutation.ClearTranscript()
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CdsUpdateOne) Select(field string, fields ...string) *CdsUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Cds entity.
func (cuo *CdsUpdateOne) Save(ctx context.Context) (*Cds, error) {
	var (
		err  error
		node *Cds
	)
	if len(cuo.hooks) == 0 {
		if err = cuo.check(); err != nil {
			return nil, err
		}
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CdsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cuo.check(); err != nil {
				return nil, err
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Cds)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CdsMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CdsUpdateOne) SaveX(ctx context.Context) *Cds {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CdsUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CdsUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CdsUpdateOne) check() error {
	if v, ok := cuo.mutation.Start(); ok {
		if err := cds.StartValidator(v); err != nil {
			return &ValidationError{Name: "start", err: fmt.Errorf(`ent: validator failed for field "Cds.start": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.End(); ok {
		if err := cds.EndValidator(v); err != nil {
			return &ValidationError{Name: "end", err: fmt.Errorf(`ent: validator failed for field "Cds.end": %w`, err)}
		}
	}
	return nil
}

func (cuo *CdsUpdateOne) sqlSave(ctx context.Context) (_node *Cds, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cds.Table,
			Columns: cds.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: cds.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Cds.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cds.FieldID)
		for _, f := range fields {
			if !cds.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != cds.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Seqname(); ok {
		_spec.SetField(cds.FieldSeqname, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Start(); ok {
		_spec.SetField(cds.FieldStart, field.TypeInt32, value)
	}
	if value, ok := cuo.mutation.AddedStart(); ok {
		_spec.AddField(cds.FieldStart, field.TypeInt32, value)
	}
	if value, ok := cuo.mutation.End(); ok {
		_spec.SetField(cds.FieldEnd, field.TypeInt32, value)
	}
	if value, ok := cuo.mutation.AddedEnd(); ok {
		_spec.AddField(cds.FieldEnd, field.TypeInt32, value)
	}
	if value, ok := cuo.mutation.Phase(); ok {
		_spec.SetField(cds.FieldPhase, field.TypeInt8, value)
	}
	if value, ok := cuo.mutation.AddedPhase(); ok {
		_spec.AddField(cds.FieldPhase, field.TypeInt8, value)
	}
	if value, ok := cuo.mutation.Strand(); ok {
		_spec.SetField(cds.FieldStrand, field.TypeString, value)
	}
	if cuo.mutation.TranscriptCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cds.TranscriptTable,
			Columns: []string{cds.TranscriptColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: transcript.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.TranscriptIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cds.TranscriptTable,
			Columns: []string{cds.TranscriptColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: transcript.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Cds{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cds.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
