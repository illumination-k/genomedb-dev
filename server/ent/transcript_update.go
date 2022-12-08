// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"genomedb/ent/predicate"
	"genomedb/ent/transcript"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
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

// SetTranscriptId sets the "transcriptId" field.
func (tu *TranscriptUpdate) SetTranscriptId(s string) *TranscriptUpdate {
	tu.mutation.SetTranscriptId(s)
	return tu
}

// SetGene sets the "gene" field.
func (tu *TranscriptUpdate) SetGene(s string) *TranscriptUpdate {
	tu.mutation.SetGene(s)
	return tu
}

// SetMrna sets the "mrna" field.
func (tu *TranscriptUpdate) SetMrna(s string) *TranscriptUpdate {
	tu.mutation.SetMrna(s)
	return tu
}

// SetCds sets the "cds" field.
func (tu *TranscriptUpdate) SetCds(s string) *TranscriptUpdate {
	tu.mutation.SetCds(s)
	return tu
}

// SetProtein sets the "protein" field.
func (tu *TranscriptUpdate) SetProtein(s string) *TranscriptUpdate {
	tu.mutation.SetProtein(s)
	return tu
}

// Mutation returns the TranscriptMutation object of the builder.
func (tu *TranscriptUpdate) Mutation() *TranscriptMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TranscriptUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tu.hooks) == 0 {
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TranscriptMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
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

func (tu *TranscriptUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   transcript.Table,
			Columns: transcript.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
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
	if value, ok := tu.mutation.TranscriptId(); ok {
		_spec.SetField(transcript.FieldTranscriptId, field.TypeString, value)
	}
	if value, ok := tu.mutation.Gene(); ok {
		_spec.SetField(transcript.FieldGene, field.TypeString, value)
	}
	if value, ok := tu.mutation.Mrna(); ok {
		_spec.SetField(transcript.FieldMrna, field.TypeString, value)
	}
	if value, ok := tu.mutation.Cds(); ok {
		_spec.SetField(transcript.FieldCds, field.TypeString, value)
	}
	if value, ok := tu.mutation.Protein(); ok {
		_spec.SetField(transcript.FieldProtein, field.TypeString, value)
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

// SetTranscriptId sets the "transcriptId" field.
func (tuo *TranscriptUpdateOne) SetTranscriptId(s string) *TranscriptUpdateOne {
	tuo.mutation.SetTranscriptId(s)
	return tuo
}

// SetGene sets the "gene" field.
func (tuo *TranscriptUpdateOne) SetGene(s string) *TranscriptUpdateOne {
	tuo.mutation.SetGene(s)
	return tuo
}

// SetMrna sets the "mrna" field.
func (tuo *TranscriptUpdateOne) SetMrna(s string) *TranscriptUpdateOne {
	tuo.mutation.SetMrna(s)
	return tuo
}

// SetCds sets the "cds" field.
func (tuo *TranscriptUpdateOne) SetCds(s string) *TranscriptUpdateOne {
	tuo.mutation.SetCds(s)
	return tuo
}

// SetProtein sets the "protein" field.
func (tuo *TranscriptUpdateOne) SetProtein(s string) *TranscriptUpdateOne {
	tuo.mutation.SetProtein(s)
	return tuo
}

// Mutation returns the TranscriptMutation object of the builder.
func (tuo *TranscriptUpdateOne) Mutation() *TranscriptMutation {
	return tuo.mutation
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
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TranscriptMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
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

func (tuo *TranscriptUpdateOne) sqlSave(ctx context.Context) (_node *Transcript, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   transcript.Table,
			Columns: transcript.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
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
	if value, ok := tuo.mutation.TranscriptId(); ok {
		_spec.SetField(transcript.FieldTranscriptId, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Gene(); ok {
		_spec.SetField(transcript.FieldGene, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Mrna(); ok {
		_spec.SetField(transcript.FieldMrna, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Cds(); ok {
		_spec.SetField(transcript.FieldCds, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Protein(); ok {
		_spec.SetField(transcript.FieldProtein, field.TypeString, value)
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