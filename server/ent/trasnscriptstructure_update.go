// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"genomedb/ent/predicate"
	"genomedb/ent/trasnscriptstructure"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TrasnscriptStructureUpdate is the builder for updating TrasnscriptStructure entities.
type TrasnscriptStructureUpdate struct {
	config
	hooks    []Hook
	mutation *TrasnscriptStructureMutation
}

// Where appends a list predicates to the TrasnscriptStructureUpdate builder.
func (tsu *TrasnscriptStructureUpdate) Where(ps ...predicate.TrasnscriptStructure) *TrasnscriptStructureUpdate {
	tsu.mutation.Where(ps...)
	return tsu
}

// SetTranscriptId sets the "transcriptId" field.
func (tsu *TrasnscriptStructureUpdate) SetTranscriptId(s string) *TrasnscriptStructureUpdate {
	tsu.mutation.SetTranscriptId(s)
	return tsu
}

// SetFeature sets the "feature" field.
func (tsu *TrasnscriptStructureUpdate) SetFeature(s string) *TrasnscriptStructureUpdate {
	tsu.mutation.SetFeature(s)
	return tsu
}

// SetSeqname sets the "seqname" field.
func (tsu *TrasnscriptStructureUpdate) SetSeqname(s string) *TrasnscriptStructureUpdate {
	tsu.mutation.SetSeqname(s)
	return tsu
}

// SetStart sets the "start" field.
func (tsu *TrasnscriptStructureUpdate) SetStart(i int32) *TrasnscriptStructureUpdate {
	tsu.mutation.ResetStart()
	tsu.mutation.SetStart(i)
	return tsu
}

// AddStart adds i to the "start" field.
func (tsu *TrasnscriptStructureUpdate) AddStart(i int32) *TrasnscriptStructureUpdate {
	tsu.mutation.AddStart(i)
	return tsu
}

// SetEnd sets the "end" field.
func (tsu *TrasnscriptStructureUpdate) SetEnd(i int32) *TrasnscriptStructureUpdate {
	tsu.mutation.ResetEnd()
	tsu.mutation.SetEnd(i)
	return tsu
}

// AddEnd adds i to the "end" field.
func (tsu *TrasnscriptStructureUpdate) AddEnd(i int32) *TrasnscriptStructureUpdate {
	tsu.mutation.AddEnd(i)
	return tsu
}

// SetStrand sets the "strand" field.
func (tsu *TrasnscriptStructureUpdate) SetStrand(s string) *TrasnscriptStructureUpdate {
	tsu.mutation.SetStrand(s)
	return tsu
}

// Mutation returns the TrasnscriptStructureMutation object of the builder.
func (tsu *TrasnscriptStructureUpdate) Mutation() *TrasnscriptStructureMutation {
	return tsu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tsu *TrasnscriptStructureUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tsu.hooks) == 0 {
		if err = tsu.check(); err != nil {
			return 0, err
		}
		affected, err = tsu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TrasnscriptStructureMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tsu.check(); err != nil {
				return 0, err
			}
			tsu.mutation = mutation
			affected, err = tsu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tsu.hooks) - 1; i >= 0; i-- {
			if tsu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tsu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tsu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tsu *TrasnscriptStructureUpdate) SaveX(ctx context.Context) int {
	affected, err := tsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tsu *TrasnscriptStructureUpdate) Exec(ctx context.Context) error {
	_, err := tsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tsu *TrasnscriptStructureUpdate) ExecX(ctx context.Context) {
	if err := tsu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tsu *TrasnscriptStructureUpdate) check() error {
	if v, ok := tsu.mutation.Start(); ok {
		if err := trasnscriptstructure.StartValidator(v); err != nil {
			return &ValidationError{Name: "start", err: fmt.Errorf(`ent: validator failed for field "TrasnscriptStructure.start": %w`, err)}
		}
	}
	if v, ok := tsu.mutation.End(); ok {
		if err := trasnscriptstructure.EndValidator(v); err != nil {
			return &ValidationError{Name: "end", err: fmt.Errorf(`ent: validator failed for field "TrasnscriptStructure.end": %w`, err)}
		}
	}
	return nil
}

func (tsu *TrasnscriptStructureUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   trasnscriptstructure.Table,
			Columns: trasnscriptstructure.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: trasnscriptstructure.FieldID,
			},
		},
	}
	if ps := tsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tsu.mutation.TranscriptId(); ok {
		_spec.SetField(trasnscriptstructure.FieldTranscriptId, field.TypeString, value)
	}
	if value, ok := tsu.mutation.Feature(); ok {
		_spec.SetField(trasnscriptstructure.FieldFeature, field.TypeString, value)
	}
	if value, ok := tsu.mutation.Seqname(); ok {
		_spec.SetField(trasnscriptstructure.FieldSeqname, field.TypeString, value)
	}
	if value, ok := tsu.mutation.Start(); ok {
		_spec.SetField(trasnscriptstructure.FieldStart, field.TypeInt32, value)
	}
	if value, ok := tsu.mutation.AddedStart(); ok {
		_spec.AddField(trasnscriptstructure.FieldStart, field.TypeInt32, value)
	}
	if value, ok := tsu.mutation.End(); ok {
		_spec.SetField(trasnscriptstructure.FieldEnd, field.TypeInt32, value)
	}
	if value, ok := tsu.mutation.AddedEnd(); ok {
		_spec.AddField(trasnscriptstructure.FieldEnd, field.TypeInt32, value)
	}
	if value, ok := tsu.mutation.Strand(); ok {
		_spec.SetField(trasnscriptstructure.FieldStrand, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{trasnscriptstructure.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// TrasnscriptStructureUpdateOne is the builder for updating a single TrasnscriptStructure entity.
type TrasnscriptStructureUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TrasnscriptStructureMutation
}

// SetTranscriptId sets the "transcriptId" field.
func (tsuo *TrasnscriptStructureUpdateOne) SetTranscriptId(s string) *TrasnscriptStructureUpdateOne {
	tsuo.mutation.SetTranscriptId(s)
	return tsuo
}

// SetFeature sets the "feature" field.
func (tsuo *TrasnscriptStructureUpdateOne) SetFeature(s string) *TrasnscriptStructureUpdateOne {
	tsuo.mutation.SetFeature(s)
	return tsuo
}

// SetSeqname sets the "seqname" field.
func (tsuo *TrasnscriptStructureUpdateOne) SetSeqname(s string) *TrasnscriptStructureUpdateOne {
	tsuo.mutation.SetSeqname(s)
	return tsuo
}

// SetStart sets the "start" field.
func (tsuo *TrasnscriptStructureUpdateOne) SetStart(i int32) *TrasnscriptStructureUpdateOne {
	tsuo.mutation.ResetStart()
	tsuo.mutation.SetStart(i)
	return tsuo
}

// AddStart adds i to the "start" field.
func (tsuo *TrasnscriptStructureUpdateOne) AddStart(i int32) *TrasnscriptStructureUpdateOne {
	tsuo.mutation.AddStart(i)
	return tsuo
}

// SetEnd sets the "end" field.
func (tsuo *TrasnscriptStructureUpdateOne) SetEnd(i int32) *TrasnscriptStructureUpdateOne {
	tsuo.mutation.ResetEnd()
	tsuo.mutation.SetEnd(i)
	return tsuo
}

// AddEnd adds i to the "end" field.
func (tsuo *TrasnscriptStructureUpdateOne) AddEnd(i int32) *TrasnscriptStructureUpdateOne {
	tsuo.mutation.AddEnd(i)
	return tsuo
}

// SetStrand sets the "strand" field.
func (tsuo *TrasnscriptStructureUpdateOne) SetStrand(s string) *TrasnscriptStructureUpdateOne {
	tsuo.mutation.SetStrand(s)
	return tsuo
}

// Mutation returns the TrasnscriptStructureMutation object of the builder.
func (tsuo *TrasnscriptStructureUpdateOne) Mutation() *TrasnscriptStructureMutation {
	return tsuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tsuo *TrasnscriptStructureUpdateOne) Select(field string, fields ...string) *TrasnscriptStructureUpdateOne {
	tsuo.fields = append([]string{field}, fields...)
	return tsuo
}

// Save executes the query and returns the updated TrasnscriptStructure entity.
func (tsuo *TrasnscriptStructureUpdateOne) Save(ctx context.Context) (*TrasnscriptStructure, error) {
	var (
		err  error
		node *TrasnscriptStructure
	)
	if len(tsuo.hooks) == 0 {
		if err = tsuo.check(); err != nil {
			return nil, err
		}
		node, err = tsuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TrasnscriptStructureMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tsuo.check(); err != nil {
				return nil, err
			}
			tsuo.mutation = mutation
			node, err = tsuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tsuo.hooks) - 1; i >= 0; i-- {
			if tsuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tsuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tsuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*TrasnscriptStructure)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TrasnscriptStructureMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tsuo *TrasnscriptStructureUpdateOne) SaveX(ctx context.Context) *TrasnscriptStructure {
	node, err := tsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tsuo *TrasnscriptStructureUpdateOne) Exec(ctx context.Context) error {
	_, err := tsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tsuo *TrasnscriptStructureUpdateOne) ExecX(ctx context.Context) {
	if err := tsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tsuo *TrasnscriptStructureUpdateOne) check() error {
	if v, ok := tsuo.mutation.Start(); ok {
		if err := trasnscriptstructure.StartValidator(v); err != nil {
			return &ValidationError{Name: "start", err: fmt.Errorf(`ent: validator failed for field "TrasnscriptStructure.start": %w`, err)}
		}
	}
	if v, ok := tsuo.mutation.End(); ok {
		if err := trasnscriptstructure.EndValidator(v); err != nil {
			return &ValidationError{Name: "end", err: fmt.Errorf(`ent: validator failed for field "TrasnscriptStructure.end": %w`, err)}
		}
	}
	return nil
}

func (tsuo *TrasnscriptStructureUpdateOne) sqlSave(ctx context.Context) (_node *TrasnscriptStructure, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   trasnscriptstructure.Table,
			Columns: trasnscriptstructure.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: trasnscriptstructure.FieldID,
			},
		},
	}
	id, ok := tsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TrasnscriptStructure.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, trasnscriptstructure.FieldID)
		for _, f := range fields {
			if !trasnscriptstructure.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != trasnscriptstructure.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tsuo.mutation.TranscriptId(); ok {
		_spec.SetField(trasnscriptstructure.FieldTranscriptId, field.TypeString, value)
	}
	if value, ok := tsuo.mutation.Feature(); ok {
		_spec.SetField(trasnscriptstructure.FieldFeature, field.TypeString, value)
	}
	if value, ok := tsuo.mutation.Seqname(); ok {
		_spec.SetField(trasnscriptstructure.FieldSeqname, field.TypeString, value)
	}
	if value, ok := tsuo.mutation.Start(); ok {
		_spec.SetField(trasnscriptstructure.FieldStart, field.TypeInt32, value)
	}
	if value, ok := tsuo.mutation.AddedStart(); ok {
		_spec.AddField(trasnscriptstructure.FieldStart, field.TypeInt32, value)
	}
	if value, ok := tsuo.mutation.End(); ok {
		_spec.SetField(trasnscriptstructure.FieldEnd, field.TypeInt32, value)
	}
	if value, ok := tsuo.mutation.AddedEnd(); ok {
		_spec.AddField(trasnscriptstructure.FieldEnd, field.TypeInt32, value)
	}
	if value, ok := tsuo.mutation.Strand(); ok {
		_spec.SetField(trasnscriptstructure.FieldStrand, field.TypeString, value)
	}
	_node = &TrasnscriptStructure{config: tsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{trasnscriptstructure.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
