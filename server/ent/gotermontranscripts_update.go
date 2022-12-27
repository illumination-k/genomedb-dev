// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"genomedb/ent/goterm"
	"genomedb/ent/gotermontranscripts"
	"genomedb/ent/predicate"
	"genomedb/ent/transcript"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GoTermOnTranscriptsUpdate is the builder for updating GoTermOnTranscripts entities.
type GoTermOnTranscriptsUpdate struct {
	config
	hooks    []Hook
	mutation *GoTermOnTranscriptsMutation
}

// Where appends a list predicates to the GoTermOnTranscriptsUpdate builder.
func (gtotu *GoTermOnTranscriptsUpdate) Where(ps ...predicate.GoTermOnTranscripts) *GoTermOnTranscriptsUpdate {
	gtotu.mutation.Where(ps...)
	return gtotu
}

// SetEvidenceCode sets the "evidence_code" field.
func (gtotu *GoTermOnTranscriptsUpdate) SetEvidenceCode(s string) *GoTermOnTranscriptsUpdate {
	gtotu.mutation.SetEvidenceCode(s)
	return gtotu
}

// SetGoTermID sets the "go_term_id" field.
func (gtotu *GoTermOnTranscriptsUpdate) SetGoTermID(s string) *GoTermOnTranscriptsUpdate {
	gtotu.mutation.SetGoTermID(s)
	return gtotu
}

// SetTranscriptID sets the "transcript_id" field.
func (gtotu *GoTermOnTranscriptsUpdate) SetTranscriptID(s string) *GoTermOnTranscriptsUpdate {
	gtotu.mutation.SetTranscriptID(s)
	return gtotu
}

// SetGoTerm sets the "go_term" edge to the GoTerm entity.
func (gtotu *GoTermOnTranscriptsUpdate) SetGoTerm(g *GoTerm) *GoTermOnTranscriptsUpdate {
	return gtotu.SetGoTermID(g.ID)
}

// SetTranscript sets the "transcript" edge to the Transcript entity.
func (gtotu *GoTermOnTranscriptsUpdate) SetTranscript(t *Transcript) *GoTermOnTranscriptsUpdate {
	return gtotu.SetTranscriptID(t.ID)
}

// Mutation returns the GoTermOnTranscriptsMutation object of the builder.
func (gtotu *GoTermOnTranscriptsUpdate) Mutation() *GoTermOnTranscriptsMutation {
	return gtotu.mutation
}

// ClearGoTerm clears the "go_term" edge to the GoTerm entity.
func (gtotu *GoTermOnTranscriptsUpdate) ClearGoTerm() *GoTermOnTranscriptsUpdate {
	gtotu.mutation.ClearGoTerm()
	return gtotu
}

// ClearTranscript clears the "transcript" edge to the Transcript entity.
func (gtotu *GoTermOnTranscriptsUpdate) ClearTranscript() *GoTermOnTranscriptsUpdate {
	gtotu.mutation.ClearTranscript()
	return gtotu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gtotu *GoTermOnTranscriptsUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(gtotu.hooks) == 0 {
		if err = gtotu.check(); err != nil {
			return 0, err
		}
		affected, err = gtotu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GoTermOnTranscriptsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gtotu.check(); err != nil {
				return 0, err
			}
			gtotu.mutation = mutation
			affected, err = gtotu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gtotu.hooks) - 1; i >= 0; i-- {
			if gtotu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gtotu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gtotu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (gtotu *GoTermOnTranscriptsUpdate) SaveX(ctx context.Context) int {
	affected, err := gtotu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gtotu *GoTermOnTranscriptsUpdate) Exec(ctx context.Context) error {
	_, err := gtotu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gtotu *GoTermOnTranscriptsUpdate) ExecX(ctx context.Context) {
	if err := gtotu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gtotu *GoTermOnTranscriptsUpdate) check() error {
	if _, ok := gtotu.mutation.GoTermID(); gtotu.mutation.GoTermCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GoTermOnTranscripts.go_term"`)
	}
	if _, ok := gtotu.mutation.TranscriptID(); gtotu.mutation.TranscriptCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GoTermOnTranscripts.transcript"`)
	}
	return nil
}

func (gtotu *GoTermOnTranscriptsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   gotermontranscripts.Table,
			Columns: gotermontranscripts.Columns,
			CompositeID: []*sqlgraph.FieldSpec{
				{
					Type:   field.TypeString,
					Column: gotermontranscripts.FieldGoTermID,
				},
				{
					Type:   field.TypeString,
					Column: gotermontranscripts.FieldTranscriptID,
				},
			},
		},
	}
	if ps := gtotu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gtotu.mutation.EvidenceCode(); ok {
		_spec.SetField(gotermontranscripts.FieldEvidenceCode, field.TypeString, value)
	}
	if gtotu.mutation.GoTermCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   gotermontranscripts.GoTermTable,
			Columns: []string{gotermontranscripts.GoTermColumn},
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
	if nodes := gtotu.mutation.GoTermIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   gotermontranscripts.GoTermTable,
			Columns: []string{gotermontranscripts.GoTermColumn},
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
	if gtotu.mutation.TranscriptCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   gotermontranscripts.TranscriptTable,
			Columns: []string{gotermontranscripts.TranscriptColumn},
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
	if nodes := gtotu.mutation.TranscriptIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   gotermontranscripts.TranscriptTable,
			Columns: []string{gotermontranscripts.TranscriptColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, gtotu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{gotermontranscripts.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// GoTermOnTranscriptsUpdateOne is the builder for updating a single GoTermOnTranscripts entity.
type GoTermOnTranscriptsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GoTermOnTranscriptsMutation
}

// SetEvidenceCode sets the "evidence_code" field.
func (gtotuo *GoTermOnTranscriptsUpdateOne) SetEvidenceCode(s string) *GoTermOnTranscriptsUpdateOne {
	gtotuo.mutation.SetEvidenceCode(s)
	return gtotuo
}

// SetGoTermID sets the "go_term_id" field.
func (gtotuo *GoTermOnTranscriptsUpdateOne) SetGoTermID(s string) *GoTermOnTranscriptsUpdateOne {
	gtotuo.mutation.SetGoTermID(s)
	return gtotuo
}

// SetTranscriptID sets the "transcript_id" field.
func (gtotuo *GoTermOnTranscriptsUpdateOne) SetTranscriptID(s string) *GoTermOnTranscriptsUpdateOne {
	gtotuo.mutation.SetTranscriptID(s)
	return gtotuo
}

// SetGoTerm sets the "go_term" edge to the GoTerm entity.
func (gtotuo *GoTermOnTranscriptsUpdateOne) SetGoTerm(g *GoTerm) *GoTermOnTranscriptsUpdateOne {
	return gtotuo.SetGoTermID(g.ID)
}

// SetTranscript sets the "transcript" edge to the Transcript entity.
func (gtotuo *GoTermOnTranscriptsUpdateOne) SetTranscript(t *Transcript) *GoTermOnTranscriptsUpdateOne {
	return gtotuo.SetTranscriptID(t.ID)
}

// Mutation returns the GoTermOnTranscriptsMutation object of the builder.
func (gtotuo *GoTermOnTranscriptsUpdateOne) Mutation() *GoTermOnTranscriptsMutation {
	return gtotuo.mutation
}

// ClearGoTerm clears the "go_term" edge to the GoTerm entity.
func (gtotuo *GoTermOnTranscriptsUpdateOne) ClearGoTerm() *GoTermOnTranscriptsUpdateOne {
	gtotuo.mutation.ClearGoTerm()
	return gtotuo
}

// ClearTranscript clears the "transcript" edge to the Transcript entity.
func (gtotuo *GoTermOnTranscriptsUpdateOne) ClearTranscript() *GoTermOnTranscriptsUpdateOne {
	gtotuo.mutation.ClearTranscript()
	return gtotuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (gtotuo *GoTermOnTranscriptsUpdateOne) Select(field string, fields ...string) *GoTermOnTranscriptsUpdateOne {
	gtotuo.fields = append([]string{field}, fields...)
	return gtotuo
}

// Save executes the query and returns the updated GoTermOnTranscripts entity.
func (gtotuo *GoTermOnTranscriptsUpdateOne) Save(ctx context.Context) (*GoTermOnTranscripts, error) {
	var (
		err  error
		node *GoTermOnTranscripts
	)
	if len(gtotuo.hooks) == 0 {
		if err = gtotuo.check(); err != nil {
			return nil, err
		}
		node, err = gtotuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GoTermOnTranscriptsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gtotuo.check(); err != nil {
				return nil, err
			}
			gtotuo.mutation = mutation
			node, err = gtotuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(gtotuo.hooks) - 1; i >= 0; i-- {
			if gtotuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gtotuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, gtotuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*GoTermOnTranscripts)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from GoTermOnTranscriptsMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (gtotuo *GoTermOnTranscriptsUpdateOne) SaveX(ctx context.Context) *GoTermOnTranscripts {
	node, err := gtotuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (gtotuo *GoTermOnTranscriptsUpdateOne) Exec(ctx context.Context) error {
	_, err := gtotuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gtotuo *GoTermOnTranscriptsUpdateOne) ExecX(ctx context.Context) {
	if err := gtotuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gtotuo *GoTermOnTranscriptsUpdateOne) check() error {
	if _, ok := gtotuo.mutation.GoTermID(); gtotuo.mutation.GoTermCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GoTermOnTranscripts.go_term"`)
	}
	if _, ok := gtotuo.mutation.TranscriptID(); gtotuo.mutation.TranscriptCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GoTermOnTranscripts.transcript"`)
	}
	return nil
}

func (gtotuo *GoTermOnTranscriptsUpdateOne) sqlSave(ctx context.Context) (_node *GoTermOnTranscripts, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   gotermontranscripts.Table,
			Columns: gotermontranscripts.Columns,
			CompositeID: []*sqlgraph.FieldSpec{
				{
					Type:   field.TypeString,
					Column: gotermontranscripts.FieldGoTermID,
				},
				{
					Type:   field.TypeString,
					Column: gotermontranscripts.FieldTranscriptID,
				},
			},
		},
	}
	if id, ok := gtotuo.mutation.GoTermID(); !ok {
		return nil, &ValidationError{Name: "go_term_id", err: errors.New(`ent: missing "GoTermOnTranscripts.go_term_id" for update`)}
	} else {
		_spec.Node.CompositeID[0].Value = id
	}
	if id, ok := gtotuo.mutation.TranscriptID(); !ok {
		return nil, &ValidationError{Name: "transcript_id", err: errors.New(`ent: missing "GoTermOnTranscripts.transcript_id" for update`)}
	} else {
		_spec.Node.CompositeID[1].Value = id
	}
	if fields := gtotuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, len(fields))
		for i, f := range fields {
			if !gotermontranscripts.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			_spec.Node.Columns[i] = f
		}
	}
	if ps := gtotuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gtotuo.mutation.EvidenceCode(); ok {
		_spec.SetField(gotermontranscripts.FieldEvidenceCode, field.TypeString, value)
	}
	if gtotuo.mutation.GoTermCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   gotermontranscripts.GoTermTable,
			Columns: []string{gotermontranscripts.GoTermColumn},
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
	if nodes := gtotuo.mutation.GoTermIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   gotermontranscripts.GoTermTable,
			Columns: []string{gotermontranscripts.GoTermColumn},
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
	if gtotuo.mutation.TranscriptCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   gotermontranscripts.TranscriptTable,
			Columns: []string{gotermontranscripts.TranscriptColumn},
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
	if nodes := gtotuo.mutation.TranscriptIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   gotermontranscripts.TranscriptTable,
			Columns: []string{gotermontranscripts.TranscriptColumn},
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
	_node = &GoTermOnTranscripts{config: gtotuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, gtotuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{gotermontranscripts.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}