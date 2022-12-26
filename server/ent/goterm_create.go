// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"genomedb/ent/goterm"
	"genomedb/ent/transcript"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GoTermCreate is the builder for creating a GoTerm entity.
type GoTermCreate struct {
	config
	mutation *GoTermMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetNamespace sets the "namespace" field.
func (gtc *GoTermCreate) SetNamespace(_go goterm.Namespace) *GoTermCreate {
	gtc.mutation.SetNamespace(_go)
	return gtc
}

// SetName sets the "name" field.
func (gtc *GoTermCreate) SetName(s string) *GoTermCreate {
	gtc.mutation.SetName(s)
	return gtc
}

// SetLevel sets the "level" field.
func (gtc *GoTermCreate) SetLevel(i int32) *GoTermCreate {
	gtc.mutation.SetLevel(i)
	return gtc
}

// SetDepth sets the "depth" field.
func (gtc *GoTermCreate) SetDepth(i int32) *GoTermCreate {
	gtc.mutation.SetDepth(i)
	return gtc
}

// SetID sets the "id" field.
func (gtc *GoTermCreate) SetID(s string) *GoTermCreate {
	gtc.mutation.SetID(s)
	return gtc
}

// SetParentID sets the "parent" edge to the GoTerm entity by ID.
func (gtc *GoTermCreate) SetParentID(id string) *GoTermCreate {
	gtc.mutation.SetParentID(id)
	return gtc
}

// SetNillableParentID sets the "parent" edge to the GoTerm entity by ID if the given value is not nil.
func (gtc *GoTermCreate) SetNillableParentID(id *string) *GoTermCreate {
	if id != nil {
		gtc = gtc.SetParentID(*id)
	}
	return gtc
}

// SetParent sets the "parent" edge to the GoTerm entity.
func (gtc *GoTermCreate) SetParent(g *GoTerm) *GoTermCreate {
	return gtc.SetParentID(g.ID)
}

// AddChildIDs adds the "children" edge to the GoTerm entity by IDs.
func (gtc *GoTermCreate) AddChildIDs(ids ...string) *GoTermCreate {
	gtc.mutation.AddChildIDs(ids...)
	return gtc
}

// AddChildren adds the "children" edges to the GoTerm entity.
func (gtc *GoTermCreate) AddChildren(g ...*GoTerm) *GoTermCreate {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gtc.AddChildIDs(ids...)
}

// AddTranscriptIDs adds the "transcripts" edge to the Transcript entity by IDs.
func (gtc *GoTermCreate) AddTranscriptIDs(ids ...string) *GoTermCreate {
	gtc.mutation.AddTranscriptIDs(ids...)
	return gtc
}

// AddTranscripts adds the "transcripts" edges to the Transcript entity.
func (gtc *GoTermCreate) AddTranscripts(t ...*Transcript) *GoTermCreate {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return gtc.AddTranscriptIDs(ids...)
}

// Mutation returns the GoTermMutation object of the builder.
func (gtc *GoTermCreate) Mutation() *GoTermMutation {
	return gtc.mutation
}

// Save creates the GoTerm in the database.
func (gtc *GoTermCreate) Save(ctx context.Context) (*GoTerm, error) {
	var (
		err  error
		node *GoTerm
	)
	if len(gtc.hooks) == 0 {
		if err = gtc.check(); err != nil {
			return nil, err
		}
		node, err = gtc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GoTermMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gtc.check(); err != nil {
				return nil, err
			}
			gtc.mutation = mutation
			if node, err = gtc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(gtc.hooks) - 1; i >= 0; i-- {
			if gtc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gtc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, gtc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*GoTerm)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from GoTermMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (gtc *GoTermCreate) SaveX(ctx context.Context) *GoTerm {
	v, err := gtc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gtc *GoTermCreate) Exec(ctx context.Context) error {
	_, err := gtc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gtc *GoTermCreate) ExecX(ctx context.Context) {
	if err := gtc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gtc *GoTermCreate) check() error {
	if _, ok := gtc.mutation.Namespace(); !ok {
		return &ValidationError{Name: "namespace", err: errors.New(`ent: missing required field "GoTerm.namespace"`)}
	}
	if v, ok := gtc.mutation.Namespace(); ok {
		if err := goterm.NamespaceValidator(v); err != nil {
			return &ValidationError{Name: "namespace", err: fmt.Errorf(`ent: validator failed for field "GoTerm.namespace": %w`, err)}
		}
	}
	if _, ok := gtc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "GoTerm.name"`)}
	}
	if _, ok := gtc.mutation.Level(); !ok {
		return &ValidationError{Name: "level", err: errors.New(`ent: missing required field "GoTerm.level"`)}
	}
	if _, ok := gtc.mutation.Depth(); !ok {
		return &ValidationError{Name: "depth", err: errors.New(`ent: missing required field "GoTerm.depth"`)}
	}
	return nil
}

func (gtc *GoTermCreate) sqlSave(ctx context.Context) (*GoTerm, error) {
	_node, _spec := gtc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gtc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected GoTerm.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (gtc *GoTermCreate) createSpec() (*GoTerm, *sqlgraph.CreateSpec) {
	var (
		_node = &GoTerm{config: gtc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: goterm.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: goterm.FieldID,
			},
		}
	)
	_spec.OnConflict = gtc.conflict
	if id, ok := gtc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := gtc.mutation.Namespace(); ok {
		_spec.SetField(goterm.FieldNamespace, field.TypeEnum, value)
		_node.Namespace = value
	}
	if value, ok := gtc.mutation.Name(); ok {
		_spec.SetField(goterm.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := gtc.mutation.Level(); ok {
		_spec.SetField(goterm.FieldLevel, field.TypeInt32, value)
		_node.Level = value
	}
	if value, ok := gtc.mutation.Depth(); ok {
		_spec.SetField(goterm.FieldDepth, field.TypeInt32, value)
		_node.Depth = value
	}
	if nodes := gtc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   goterm.ParentTable,
			Columns: []string{goterm.ParentColumn},
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
		_node.go_term_children = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gtc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   goterm.ChildrenTable,
			Columns: []string{goterm.ChildrenColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gtc.mutation.TranscriptsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   goterm.TranscriptsTable,
			Columns: goterm.TranscriptsPrimaryKey,
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.GoTerm.Create().
//		SetNamespace(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GoTermUpsert) {
//			SetNamespace(v+v).
//		}).
//		Exec(ctx)
func (gtc *GoTermCreate) OnConflict(opts ...sql.ConflictOption) *GoTermUpsertOne {
	gtc.conflict = opts
	return &GoTermUpsertOne{
		create: gtc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.GoTerm.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (gtc *GoTermCreate) OnConflictColumns(columns ...string) *GoTermUpsertOne {
	gtc.conflict = append(gtc.conflict, sql.ConflictColumns(columns...))
	return &GoTermUpsertOne{
		create: gtc,
	}
}

type (
	// GoTermUpsertOne is the builder for "upsert"-ing
	//  one GoTerm node.
	GoTermUpsertOne struct {
		create *GoTermCreate
	}

	// GoTermUpsert is the "OnConflict" setter.
	GoTermUpsert struct {
		*sql.UpdateSet
	}
)

// SetNamespace sets the "namespace" field.
func (u *GoTermUpsert) SetNamespace(v goterm.Namespace) *GoTermUpsert {
	u.Set(goterm.FieldNamespace, v)
	return u
}

// UpdateNamespace sets the "namespace" field to the value that was provided on create.
func (u *GoTermUpsert) UpdateNamespace() *GoTermUpsert {
	u.SetExcluded(goterm.FieldNamespace)
	return u
}

// SetName sets the "name" field.
func (u *GoTermUpsert) SetName(v string) *GoTermUpsert {
	u.Set(goterm.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *GoTermUpsert) UpdateName() *GoTermUpsert {
	u.SetExcluded(goterm.FieldName)
	return u
}

// SetLevel sets the "level" field.
func (u *GoTermUpsert) SetLevel(v int32) *GoTermUpsert {
	u.Set(goterm.FieldLevel, v)
	return u
}

// UpdateLevel sets the "level" field to the value that was provided on create.
func (u *GoTermUpsert) UpdateLevel() *GoTermUpsert {
	u.SetExcluded(goterm.FieldLevel)
	return u
}

// AddLevel adds v to the "level" field.
func (u *GoTermUpsert) AddLevel(v int32) *GoTermUpsert {
	u.Add(goterm.FieldLevel, v)
	return u
}

// SetDepth sets the "depth" field.
func (u *GoTermUpsert) SetDepth(v int32) *GoTermUpsert {
	u.Set(goterm.FieldDepth, v)
	return u
}

// UpdateDepth sets the "depth" field to the value that was provided on create.
func (u *GoTermUpsert) UpdateDepth() *GoTermUpsert {
	u.SetExcluded(goterm.FieldDepth)
	return u
}

// AddDepth adds v to the "depth" field.
func (u *GoTermUpsert) AddDepth(v int32) *GoTermUpsert {
	u.Add(goterm.FieldDepth, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.GoTerm.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(goterm.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *GoTermUpsertOne) UpdateNewValues() *GoTermUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(goterm.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.GoTerm.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *GoTermUpsertOne) Ignore() *GoTermUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GoTermUpsertOne) DoNothing() *GoTermUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GoTermCreate.OnConflict
// documentation for more info.
func (u *GoTermUpsertOne) Update(set func(*GoTermUpsert)) *GoTermUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GoTermUpsert{UpdateSet: update})
	}))
	return u
}

// SetNamespace sets the "namespace" field.
func (u *GoTermUpsertOne) SetNamespace(v goterm.Namespace) *GoTermUpsertOne {
	return u.Update(func(s *GoTermUpsert) {
		s.SetNamespace(v)
	})
}

// UpdateNamespace sets the "namespace" field to the value that was provided on create.
func (u *GoTermUpsertOne) UpdateNamespace() *GoTermUpsertOne {
	return u.Update(func(s *GoTermUpsert) {
		s.UpdateNamespace()
	})
}

// SetName sets the "name" field.
func (u *GoTermUpsertOne) SetName(v string) *GoTermUpsertOne {
	return u.Update(func(s *GoTermUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *GoTermUpsertOne) UpdateName() *GoTermUpsertOne {
	return u.Update(func(s *GoTermUpsert) {
		s.UpdateName()
	})
}

// SetLevel sets the "level" field.
func (u *GoTermUpsertOne) SetLevel(v int32) *GoTermUpsertOne {
	return u.Update(func(s *GoTermUpsert) {
		s.SetLevel(v)
	})
}

// AddLevel adds v to the "level" field.
func (u *GoTermUpsertOne) AddLevel(v int32) *GoTermUpsertOne {
	return u.Update(func(s *GoTermUpsert) {
		s.AddLevel(v)
	})
}

// UpdateLevel sets the "level" field to the value that was provided on create.
func (u *GoTermUpsertOne) UpdateLevel() *GoTermUpsertOne {
	return u.Update(func(s *GoTermUpsert) {
		s.UpdateLevel()
	})
}

// SetDepth sets the "depth" field.
func (u *GoTermUpsertOne) SetDepth(v int32) *GoTermUpsertOne {
	return u.Update(func(s *GoTermUpsert) {
		s.SetDepth(v)
	})
}

// AddDepth adds v to the "depth" field.
func (u *GoTermUpsertOne) AddDepth(v int32) *GoTermUpsertOne {
	return u.Update(func(s *GoTermUpsert) {
		s.AddDepth(v)
	})
}

// UpdateDepth sets the "depth" field to the value that was provided on create.
func (u *GoTermUpsertOne) UpdateDepth() *GoTermUpsertOne {
	return u.Update(func(s *GoTermUpsert) {
		s.UpdateDepth()
	})
}

// Exec executes the query.
func (u *GoTermUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GoTermCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GoTermUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *GoTermUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: GoTermUpsertOne.ID is not supported by MySQL driver. Use GoTermUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *GoTermUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// GoTermCreateBulk is the builder for creating many GoTerm entities in bulk.
type GoTermCreateBulk struct {
	config
	builders []*GoTermCreate
	conflict []sql.ConflictOption
}

// Save creates the GoTerm entities in the database.
func (gtcb *GoTermCreateBulk) Save(ctx context.Context) ([]*GoTerm, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gtcb.builders))
	nodes := make([]*GoTerm, len(gtcb.builders))
	mutators := make([]Mutator, len(gtcb.builders))
	for i := range gtcb.builders {
		func(i int, root context.Context) {
			builder := gtcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GoTermMutation)
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
					_, err = mutators[i+1].Mutate(root, gtcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = gtcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gtcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, gtcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gtcb *GoTermCreateBulk) SaveX(ctx context.Context) []*GoTerm {
	v, err := gtcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gtcb *GoTermCreateBulk) Exec(ctx context.Context) error {
	_, err := gtcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gtcb *GoTermCreateBulk) ExecX(ctx context.Context) {
	if err := gtcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.GoTerm.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GoTermUpsert) {
//			SetNamespace(v+v).
//		}).
//		Exec(ctx)
func (gtcb *GoTermCreateBulk) OnConflict(opts ...sql.ConflictOption) *GoTermUpsertBulk {
	gtcb.conflict = opts
	return &GoTermUpsertBulk{
		create: gtcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.GoTerm.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (gtcb *GoTermCreateBulk) OnConflictColumns(columns ...string) *GoTermUpsertBulk {
	gtcb.conflict = append(gtcb.conflict, sql.ConflictColumns(columns...))
	return &GoTermUpsertBulk{
		create: gtcb,
	}
}

// GoTermUpsertBulk is the builder for "upsert"-ing
// a bulk of GoTerm nodes.
type GoTermUpsertBulk struct {
	create *GoTermCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.GoTerm.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(goterm.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *GoTermUpsertBulk) UpdateNewValues() *GoTermUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(goterm.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.GoTerm.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *GoTermUpsertBulk) Ignore() *GoTermUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GoTermUpsertBulk) DoNothing() *GoTermUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GoTermCreateBulk.OnConflict
// documentation for more info.
func (u *GoTermUpsertBulk) Update(set func(*GoTermUpsert)) *GoTermUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GoTermUpsert{UpdateSet: update})
	}))
	return u
}

// SetNamespace sets the "namespace" field.
func (u *GoTermUpsertBulk) SetNamespace(v goterm.Namespace) *GoTermUpsertBulk {
	return u.Update(func(s *GoTermUpsert) {
		s.SetNamespace(v)
	})
}

// UpdateNamespace sets the "namespace" field to the value that was provided on create.
func (u *GoTermUpsertBulk) UpdateNamespace() *GoTermUpsertBulk {
	return u.Update(func(s *GoTermUpsert) {
		s.UpdateNamespace()
	})
}

// SetName sets the "name" field.
func (u *GoTermUpsertBulk) SetName(v string) *GoTermUpsertBulk {
	return u.Update(func(s *GoTermUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *GoTermUpsertBulk) UpdateName() *GoTermUpsertBulk {
	return u.Update(func(s *GoTermUpsert) {
		s.UpdateName()
	})
}

// SetLevel sets the "level" field.
func (u *GoTermUpsertBulk) SetLevel(v int32) *GoTermUpsertBulk {
	return u.Update(func(s *GoTermUpsert) {
		s.SetLevel(v)
	})
}

// AddLevel adds v to the "level" field.
func (u *GoTermUpsertBulk) AddLevel(v int32) *GoTermUpsertBulk {
	return u.Update(func(s *GoTermUpsert) {
		s.AddLevel(v)
	})
}

// UpdateLevel sets the "level" field to the value that was provided on create.
func (u *GoTermUpsertBulk) UpdateLevel() *GoTermUpsertBulk {
	return u.Update(func(s *GoTermUpsert) {
		s.UpdateLevel()
	})
}

// SetDepth sets the "depth" field.
func (u *GoTermUpsertBulk) SetDepth(v int32) *GoTermUpsertBulk {
	return u.Update(func(s *GoTermUpsert) {
		s.SetDepth(v)
	})
}

// AddDepth adds v to the "depth" field.
func (u *GoTermUpsertBulk) AddDepth(v int32) *GoTermUpsertBulk {
	return u.Update(func(s *GoTermUpsert) {
		s.AddDepth(v)
	})
}

// UpdateDepth sets the "depth" field to the value that was provided on create.
func (u *GoTermUpsertBulk) UpdateDepth() *GoTermUpsertBulk {
	return u.Update(func(s *GoTermUpsert) {
		s.UpdateDepth()
	})
}

// Exec executes the query.
func (u *GoTermUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the GoTermCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GoTermCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GoTermUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
