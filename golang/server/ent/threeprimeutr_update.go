// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"genomedb/ent/predicate"
	"genomedb/ent/threeprimeutr"
	"genomedb/ent/transcript"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ThreePrimeUtrUpdate is the builder for updating ThreePrimeUtr entities.
type ThreePrimeUtrUpdate struct {
	config
	hooks    []Hook
	mutation *ThreePrimeUtrMutation
}

// Where appends a list predicates to the ThreePrimeUtrUpdate builder.
func (tpuu *ThreePrimeUtrUpdate) Where(ps ...predicate.ThreePrimeUtr) *ThreePrimeUtrUpdate {
	tpuu.mutation.Where(ps...)
	return tpuu
}

// SetSeqname sets the "seqname" field.
func (tpuu *ThreePrimeUtrUpdate) SetSeqname(s string) *ThreePrimeUtrUpdate {
	tpuu.mutation.SetSeqname(s)
	return tpuu
}

// SetStart sets the "start" field.
func (tpuu *ThreePrimeUtrUpdate) SetStart(i int32) *ThreePrimeUtrUpdate {
	tpuu.mutation.ResetStart()
	tpuu.mutation.SetStart(i)
	return tpuu
}

// AddStart adds i to the "start" field.
func (tpuu *ThreePrimeUtrUpdate) AddStart(i int32) *ThreePrimeUtrUpdate {
	tpuu.mutation.AddStart(i)
	return tpuu
}

// SetEnd sets the "end" field.
func (tpuu *ThreePrimeUtrUpdate) SetEnd(i int32) *ThreePrimeUtrUpdate {
	tpuu.mutation.ResetEnd()
	tpuu.mutation.SetEnd(i)
	return tpuu
}

// AddEnd adds i to the "end" field.
func (tpuu *ThreePrimeUtrUpdate) AddEnd(i int32) *ThreePrimeUtrUpdate {
	tpuu.mutation.AddEnd(i)
	return tpuu
}

// SetStrand sets the "strand" field.
func (tpuu *ThreePrimeUtrUpdate) SetStrand(s string) *ThreePrimeUtrUpdate {
	tpuu.mutation.SetStrand(s)
	return tpuu
}

// SetTranscriptID sets the "transcript" edge to the Transcript entity by ID.
func (tpuu *ThreePrimeUtrUpdate) SetTranscriptID(id string) *ThreePrimeUtrUpdate {
	tpuu.mutation.SetTranscriptID(id)
	return tpuu
}

// SetNillableTranscriptID sets the "transcript" edge to the Transcript entity by ID if the given value is not nil.
func (tpuu *ThreePrimeUtrUpdate) SetNillableTranscriptID(id *string) *ThreePrimeUtrUpdate {
	if id != nil {
		tpuu = tpuu.SetTranscriptID(*id)
	}
	return tpuu
}

// SetTranscript sets the "transcript" edge to the Transcript entity.
func (tpuu *ThreePrimeUtrUpdate) SetTranscript(t *Transcript) *ThreePrimeUtrUpdate {
	return tpuu.SetTranscriptID(t.ID)
}

// Mutation returns the ThreePrimeUtrMutation object of the builder.
func (tpuu *ThreePrimeUtrUpdate) Mutation() *ThreePrimeUtrMutation {
	return tpuu.mutation
}

// ClearTranscript clears the "transcript" edge to the Transcript entity.
func (tpuu *ThreePrimeUtrUpdate) ClearTranscript() *ThreePrimeUtrUpdate {
	tpuu.mutation.ClearTranscript()
	return tpuu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tpuu *ThreePrimeUtrUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tpuu.hooks) == 0 {
		if err = tpuu.check(); err != nil {
			return 0, err
		}
		affected, err = tpuu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ThreePrimeUtrMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tpuu.check(); err != nil {
				return 0, err
			}
			tpuu.mutation = mutation
			affected, err = tpuu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tpuu.hooks) - 1; i >= 0; i-- {
			if tpuu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tpuu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tpuu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tpuu *ThreePrimeUtrUpdate) SaveX(ctx context.Context) int {
	affected, err := tpuu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tpuu *ThreePrimeUtrUpdate) Exec(ctx context.Context) error {
	_, err := tpuu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tpuu *ThreePrimeUtrUpdate) ExecX(ctx context.Context) {
	if err := tpuu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tpuu *ThreePrimeUtrUpdate) check() error {
	if v, ok := tpuu.mutation.Start(); ok {
		if err := threeprimeutr.StartValidator(v); err != nil {
			return &ValidationError{Name: "start", err: fmt.Errorf(`ent: validator failed for field "ThreePrimeUtr.start": %w`, err)}
		}
	}
	if v, ok := tpuu.mutation.End(); ok {
		if err := threeprimeutr.EndValidator(v); err != nil {
			return &ValidationError{Name: "end", err: fmt.Errorf(`ent: validator failed for field "ThreePrimeUtr.end": %w`, err)}
		}
	}
	return nil
}

func (tpuu *ThreePrimeUtrUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   threeprimeutr.Table,
			Columns: threeprimeutr.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: threeprimeutr.FieldID,
			},
		},
	}
	if ps := tpuu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tpuu.mutation.Seqname(); ok {
		_spec.SetField(threeprimeutr.FieldSeqname, field.TypeString, value)
	}
	if value, ok := tpuu.mutation.Start(); ok {
		_spec.SetField(threeprimeutr.FieldStart, field.TypeInt32, value)
	}
	if value, ok := tpuu.mutation.AddedStart(); ok {
		_spec.AddField(threeprimeutr.FieldStart, field.TypeInt32, value)
	}
	if value, ok := tpuu.mutation.End(); ok {
		_spec.SetField(threeprimeutr.FieldEnd, field.TypeInt32, value)
	}
	if value, ok := tpuu.mutation.AddedEnd(); ok {
		_spec.AddField(threeprimeutr.FieldEnd, field.TypeInt32, value)
	}
	if value, ok := tpuu.mutation.Strand(); ok {
		_spec.SetField(threeprimeutr.FieldStrand, field.TypeString, value)
	}
	if tpuu.mutation.TranscriptCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   threeprimeutr.TranscriptTable,
			Columns: []string{threeprimeutr.TranscriptColumn},
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
	if nodes := tpuu.mutation.TranscriptIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   threeprimeutr.TranscriptTable,
			Columns: []string{threeprimeutr.TranscriptColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, tpuu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{threeprimeutr.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ThreePrimeUtrUpdateOne is the builder for updating a single ThreePrimeUtr entity.
type ThreePrimeUtrUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ThreePrimeUtrMutation
}

// SetSeqname sets the "seqname" field.
func (tpuuo *ThreePrimeUtrUpdateOne) SetSeqname(s string) *ThreePrimeUtrUpdateOne {
	tpuuo.mutation.SetSeqname(s)
	return tpuuo
}

// SetStart sets the "start" field.
func (tpuuo *ThreePrimeUtrUpdateOne) SetStart(i int32) *ThreePrimeUtrUpdateOne {
	tpuuo.mutation.ResetStart()
	tpuuo.mutation.SetStart(i)
	return tpuuo
}

// AddStart adds i to the "start" field.
func (tpuuo *ThreePrimeUtrUpdateOne) AddStart(i int32) *ThreePrimeUtrUpdateOne {
	tpuuo.mutation.AddStart(i)
	return tpuuo
}

// SetEnd sets the "end" field.
func (tpuuo *ThreePrimeUtrUpdateOne) SetEnd(i int32) *ThreePrimeUtrUpdateOne {
	tpuuo.mutation.ResetEnd()
	tpuuo.mutation.SetEnd(i)
	return tpuuo
}

// AddEnd adds i to the "end" field.
func (tpuuo *ThreePrimeUtrUpdateOne) AddEnd(i int32) *ThreePrimeUtrUpdateOne {
	tpuuo.mutation.AddEnd(i)
	return tpuuo
}

// SetStrand sets the "strand" field.
func (tpuuo *ThreePrimeUtrUpdateOne) SetStrand(s string) *ThreePrimeUtrUpdateOne {
	tpuuo.mutation.SetStrand(s)
	return tpuuo
}

// SetTranscriptID sets the "transcript" edge to the Transcript entity by ID.
func (tpuuo *ThreePrimeUtrUpdateOne) SetTranscriptID(id string) *ThreePrimeUtrUpdateOne {
	tpuuo.mutation.SetTranscriptID(id)
	return tpuuo
}

// SetNillableTranscriptID sets the "transcript" edge to the Transcript entity by ID if the given value is not nil.
func (tpuuo *ThreePrimeUtrUpdateOne) SetNillableTranscriptID(id *string) *ThreePrimeUtrUpdateOne {
	if id != nil {
		tpuuo = tpuuo.SetTranscriptID(*id)
	}
	return tpuuo
}

// SetTranscript sets the "transcript" edge to the Transcript entity.
func (tpuuo *ThreePrimeUtrUpdateOne) SetTranscript(t *Transcript) *ThreePrimeUtrUpdateOne {
	return tpuuo.SetTranscriptID(t.ID)
}

// Mutation returns the ThreePrimeUtrMutation object of the builder.
func (tpuuo *ThreePrimeUtrUpdateOne) Mutation() *ThreePrimeUtrMutation {
	return tpuuo.mutation
}

// ClearTranscript clears the "transcript" edge to the Transcript entity.
func (tpuuo *ThreePrimeUtrUpdateOne) ClearTranscript() *ThreePrimeUtrUpdateOne {
	tpuuo.mutation.ClearTranscript()
	return tpuuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tpuuo *ThreePrimeUtrUpdateOne) Select(field string, fields ...string) *ThreePrimeUtrUpdateOne {
	tpuuo.fields = append([]string{field}, fields...)
	return tpuuo
}

// Save executes the query and returns the updated ThreePrimeUtr entity.
func (tpuuo *ThreePrimeUtrUpdateOne) Save(ctx context.Context) (*ThreePrimeUtr, error) {
	var (
		err  error
		node *ThreePrimeUtr
	)
	if len(tpuuo.hooks) == 0 {
		if err = tpuuo.check(); err != nil {
			return nil, err
		}
		node, err = tpuuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ThreePrimeUtrMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tpuuo.check(); err != nil {
				return nil, err
			}
			tpuuo.mutation = mutation
			node, err = tpuuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tpuuo.hooks) - 1; i >= 0; i-- {
			if tpuuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tpuuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tpuuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*ThreePrimeUtr)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ThreePrimeUtrMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tpuuo *ThreePrimeUtrUpdateOne) SaveX(ctx context.Context) *ThreePrimeUtr {
	node, err := tpuuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tpuuo *ThreePrimeUtrUpdateOne) Exec(ctx context.Context) error {
	_, err := tpuuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tpuuo *ThreePrimeUtrUpdateOne) ExecX(ctx context.Context) {
	if err := tpuuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tpuuo *ThreePrimeUtrUpdateOne) check() error {
	if v, ok := tpuuo.mutation.Start(); ok {
		if err := threeprimeutr.StartValidator(v); err != nil {
			return &ValidationError{Name: "start", err: fmt.Errorf(`ent: validator failed for field "ThreePrimeUtr.start": %w`, err)}
		}
	}
	if v, ok := tpuuo.mutation.End(); ok {
		if err := threeprimeutr.EndValidator(v); err != nil {
			return &ValidationError{Name: "end", err: fmt.Errorf(`ent: validator failed for field "ThreePrimeUtr.end": %w`, err)}
		}
	}
	return nil
}

func (tpuuo *ThreePrimeUtrUpdateOne) sqlSave(ctx context.Context) (_node *ThreePrimeUtr, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   threeprimeutr.Table,
			Columns: threeprimeutr.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: threeprimeutr.FieldID,
			},
		},
	}
	id, ok := tpuuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ThreePrimeUtr.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tpuuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, threeprimeutr.FieldID)
		for _, f := range fields {
			if !threeprimeutr.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != threeprimeutr.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tpuuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tpuuo.mutation.Seqname(); ok {
		_spec.SetField(threeprimeutr.FieldSeqname, field.TypeString, value)
	}
	if value, ok := tpuuo.mutation.Start(); ok {
		_spec.SetField(threeprimeutr.FieldStart, field.TypeInt32, value)
	}
	if value, ok := tpuuo.mutation.AddedStart(); ok {
		_spec.AddField(threeprimeutr.FieldStart, field.TypeInt32, value)
	}
	if value, ok := tpuuo.mutation.End(); ok {
		_spec.SetField(threeprimeutr.FieldEnd, field.TypeInt32, value)
	}
	if value, ok := tpuuo.mutation.AddedEnd(); ok {
		_spec.AddField(threeprimeutr.FieldEnd, field.TypeInt32, value)
	}
	if value, ok := tpuuo.mutation.Strand(); ok {
		_spec.SetField(threeprimeutr.FieldStrand, field.TypeString, value)
	}
	if tpuuo.mutation.TranscriptCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   threeprimeutr.TranscriptTable,
			Columns: []string{threeprimeutr.TranscriptColumn},
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
	if nodes := tpuuo.mutation.TranscriptIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   threeprimeutr.TranscriptTable,
			Columns: []string{threeprimeutr.TranscriptColumn},
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
	_node = &ThreePrimeUtr{config: tpuuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tpuuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{threeprimeutr.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}