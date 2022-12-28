// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"genomedb/ent/keggorthlogy"
	"genomedb/ent/keggpathway"
	"genomedb/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// KeggOrthlogyUpdate is the builder for updating KeggOrthlogy entities.
type KeggOrthlogyUpdate struct {
	config
	hooks    []Hook
	mutation *KeggOrthlogyMutation
}

// Where appends a list predicates to the KeggOrthlogyUpdate builder.
func (kou *KeggOrthlogyUpdate) Where(ps ...predicate.KeggOrthlogy) *KeggOrthlogyUpdate {
	kou.mutation.Where(ps...)
	return kou
}

// SetName sets the "name" field.
func (kou *KeggOrthlogyUpdate) SetName(s string) *KeggOrthlogyUpdate {
	kou.mutation.SetName(s)
	return kou
}

// AddPathwayIDs adds the "pathways" edge to the KeggPathway entity by IDs.
func (kou *KeggOrthlogyUpdate) AddPathwayIDs(ids ...string) *KeggOrthlogyUpdate {
	kou.mutation.AddPathwayIDs(ids...)
	return kou
}

// AddPathways adds the "pathways" edges to the KeggPathway entity.
func (kou *KeggOrthlogyUpdate) AddPathways(k ...*KeggPathway) *KeggOrthlogyUpdate {
	ids := make([]string, len(k))
	for i := range k {
		ids[i] = k[i].ID
	}
	return kou.AddPathwayIDs(ids...)
}

// Mutation returns the KeggOrthlogyMutation object of the builder.
func (kou *KeggOrthlogyUpdate) Mutation() *KeggOrthlogyMutation {
	return kou.mutation
}

// ClearPathways clears all "pathways" edges to the KeggPathway entity.
func (kou *KeggOrthlogyUpdate) ClearPathways() *KeggOrthlogyUpdate {
	kou.mutation.ClearPathways()
	return kou
}

// RemovePathwayIDs removes the "pathways" edge to KeggPathway entities by IDs.
func (kou *KeggOrthlogyUpdate) RemovePathwayIDs(ids ...string) *KeggOrthlogyUpdate {
	kou.mutation.RemovePathwayIDs(ids...)
	return kou
}

// RemovePathways removes "pathways" edges to KeggPathway entities.
func (kou *KeggOrthlogyUpdate) RemovePathways(k ...*KeggPathway) *KeggOrthlogyUpdate {
	ids := make([]string, len(k))
	for i := range k {
		ids[i] = k[i].ID
	}
	return kou.RemovePathwayIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (kou *KeggOrthlogyUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(kou.hooks) == 0 {
		affected, err = kou.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*KeggOrthlogyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			kou.mutation = mutation
			affected, err = kou.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(kou.hooks) - 1; i >= 0; i-- {
			if kou.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = kou.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, kou.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (kou *KeggOrthlogyUpdate) SaveX(ctx context.Context) int {
	affected, err := kou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (kou *KeggOrthlogyUpdate) Exec(ctx context.Context) error {
	_, err := kou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kou *KeggOrthlogyUpdate) ExecX(ctx context.Context) {
	if err := kou.Exec(ctx); err != nil {
		panic(err)
	}
}

func (kou *KeggOrthlogyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   keggorthlogy.Table,
			Columns: keggorthlogy.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: keggorthlogy.FieldID,
			},
		},
	}
	if ps := kou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := kou.mutation.Name(); ok {
		_spec.SetField(keggorthlogy.FieldName, field.TypeString, value)
	}
	if kou.mutation.PathwaysCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   keggorthlogy.PathwaysTable,
			Columns: keggorthlogy.PathwaysPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: keggpathway.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := kou.mutation.RemovedPathwaysIDs(); len(nodes) > 0 && !kou.mutation.PathwaysCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   keggorthlogy.PathwaysTable,
			Columns: keggorthlogy.PathwaysPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: keggpathway.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := kou.mutation.PathwaysIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   keggorthlogy.PathwaysTable,
			Columns: keggorthlogy.PathwaysPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: keggpathway.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, kou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{keggorthlogy.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// KeggOrthlogyUpdateOne is the builder for updating a single KeggOrthlogy entity.
type KeggOrthlogyUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *KeggOrthlogyMutation
}

// SetName sets the "name" field.
func (kouo *KeggOrthlogyUpdateOne) SetName(s string) *KeggOrthlogyUpdateOne {
	kouo.mutation.SetName(s)
	return kouo
}

// AddPathwayIDs adds the "pathways" edge to the KeggPathway entity by IDs.
func (kouo *KeggOrthlogyUpdateOne) AddPathwayIDs(ids ...string) *KeggOrthlogyUpdateOne {
	kouo.mutation.AddPathwayIDs(ids...)
	return kouo
}

// AddPathways adds the "pathways" edges to the KeggPathway entity.
func (kouo *KeggOrthlogyUpdateOne) AddPathways(k ...*KeggPathway) *KeggOrthlogyUpdateOne {
	ids := make([]string, len(k))
	for i := range k {
		ids[i] = k[i].ID
	}
	return kouo.AddPathwayIDs(ids...)
}

// Mutation returns the KeggOrthlogyMutation object of the builder.
func (kouo *KeggOrthlogyUpdateOne) Mutation() *KeggOrthlogyMutation {
	return kouo.mutation
}

// ClearPathways clears all "pathways" edges to the KeggPathway entity.
func (kouo *KeggOrthlogyUpdateOne) ClearPathways() *KeggOrthlogyUpdateOne {
	kouo.mutation.ClearPathways()
	return kouo
}

// RemovePathwayIDs removes the "pathways" edge to KeggPathway entities by IDs.
func (kouo *KeggOrthlogyUpdateOne) RemovePathwayIDs(ids ...string) *KeggOrthlogyUpdateOne {
	kouo.mutation.RemovePathwayIDs(ids...)
	return kouo
}

// RemovePathways removes "pathways" edges to KeggPathway entities.
func (kouo *KeggOrthlogyUpdateOne) RemovePathways(k ...*KeggPathway) *KeggOrthlogyUpdateOne {
	ids := make([]string, len(k))
	for i := range k {
		ids[i] = k[i].ID
	}
	return kouo.RemovePathwayIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (kouo *KeggOrthlogyUpdateOne) Select(field string, fields ...string) *KeggOrthlogyUpdateOne {
	kouo.fields = append([]string{field}, fields...)
	return kouo
}

// Save executes the query and returns the updated KeggOrthlogy entity.
func (kouo *KeggOrthlogyUpdateOne) Save(ctx context.Context) (*KeggOrthlogy, error) {
	var (
		err  error
		node *KeggOrthlogy
	)
	if len(kouo.hooks) == 0 {
		node, err = kouo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*KeggOrthlogyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			kouo.mutation = mutation
			node, err = kouo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(kouo.hooks) - 1; i >= 0; i-- {
			if kouo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = kouo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, kouo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*KeggOrthlogy)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from KeggOrthlogyMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (kouo *KeggOrthlogyUpdateOne) SaveX(ctx context.Context) *KeggOrthlogy {
	node, err := kouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (kouo *KeggOrthlogyUpdateOne) Exec(ctx context.Context) error {
	_, err := kouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kouo *KeggOrthlogyUpdateOne) ExecX(ctx context.Context) {
	if err := kouo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (kouo *KeggOrthlogyUpdateOne) sqlSave(ctx context.Context) (_node *KeggOrthlogy, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   keggorthlogy.Table,
			Columns: keggorthlogy.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: keggorthlogy.FieldID,
			},
		},
	}
	id, ok := kouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "KeggOrthlogy.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := kouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, keggorthlogy.FieldID)
		for _, f := range fields {
			if !keggorthlogy.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != keggorthlogy.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := kouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := kouo.mutation.Name(); ok {
		_spec.SetField(keggorthlogy.FieldName, field.TypeString, value)
	}
	if kouo.mutation.PathwaysCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   keggorthlogy.PathwaysTable,
			Columns: keggorthlogy.PathwaysPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: keggpathway.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := kouo.mutation.RemovedPathwaysIDs(); len(nodes) > 0 && !kouo.mutation.PathwaysCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   keggorthlogy.PathwaysTable,
			Columns: keggorthlogy.PathwaysPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: keggpathway.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := kouo.mutation.PathwaysIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   keggorthlogy.PathwaysTable,
			Columns: keggorthlogy.PathwaysPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: keggpathway.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &KeggOrthlogy{config: kouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, kouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{keggorthlogy.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
