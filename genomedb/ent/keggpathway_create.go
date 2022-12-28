// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"genomedb/ent/keggorthlogy"
	"genomedb/ent/keggpathway"
	"genomedb/ent/keggreaction"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// KeggPathwayCreate is the builder for creating a KeggPathway entity.
type KeggPathwayCreate struct {
	config
	mutation *KeggPathwayMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (kpc *KeggPathwayCreate) SetName(s string) *KeggPathwayCreate {
	kpc.mutation.SetName(s)
	return kpc
}

// SetID sets the "id" field.
func (kpc *KeggPathwayCreate) SetID(s string) *KeggPathwayCreate {
	kpc.mutation.SetID(s)
	return kpc
}

// AddRelatingMapIDs adds the "relating_map" edge to the KeggPathway entity by IDs.
func (kpc *KeggPathwayCreate) AddRelatingMapIDs(ids ...string) *KeggPathwayCreate {
	kpc.mutation.AddRelatingMapIDs(ids...)
	return kpc
}

// AddRelatingMap adds the "relating_map" edges to the KeggPathway entity.
func (kpc *KeggPathwayCreate) AddRelatingMap(k ...*KeggPathway) *KeggPathwayCreate {
	ids := make([]string, len(k))
	for i := range k {
		ids[i] = k[i].ID
	}
	return kpc.AddRelatingMapIDs(ids...)
}

// AddRelatedMapIDs adds the "related_map" edge to the KeggPathway entity by IDs.
func (kpc *KeggPathwayCreate) AddRelatedMapIDs(ids ...string) *KeggPathwayCreate {
	kpc.mutation.AddRelatedMapIDs(ids...)
	return kpc
}

// AddRelatedMap adds the "related_map" edges to the KeggPathway entity.
func (kpc *KeggPathwayCreate) AddRelatedMap(k ...*KeggPathway) *KeggPathwayCreate {
	ids := make([]string, len(k))
	for i := range k {
		ids[i] = k[i].ID
	}
	return kpc.AddRelatedMapIDs(ids...)
}

// AddReactionIDs adds the "reactions" edge to the KeggReaction entity by IDs.
func (kpc *KeggPathwayCreate) AddReactionIDs(ids ...string) *KeggPathwayCreate {
	kpc.mutation.AddReactionIDs(ids...)
	return kpc
}

// AddReactions adds the "reactions" edges to the KeggReaction entity.
func (kpc *KeggPathwayCreate) AddReactions(k ...*KeggReaction) *KeggPathwayCreate {
	ids := make([]string, len(k))
	for i := range k {
		ids[i] = k[i].ID
	}
	return kpc.AddReactionIDs(ids...)
}

// AddOrthologyIDs adds the "orthologies" edge to the KeggOrthlogy entity by IDs.
func (kpc *KeggPathwayCreate) AddOrthologyIDs(ids ...string) *KeggPathwayCreate {
	kpc.mutation.AddOrthologyIDs(ids...)
	return kpc
}

// AddOrthologies adds the "orthologies" edges to the KeggOrthlogy entity.
func (kpc *KeggPathwayCreate) AddOrthologies(k ...*KeggOrthlogy) *KeggPathwayCreate {
	ids := make([]string, len(k))
	for i := range k {
		ids[i] = k[i].ID
	}
	return kpc.AddOrthologyIDs(ids...)
}

// Mutation returns the KeggPathwayMutation object of the builder.
func (kpc *KeggPathwayCreate) Mutation() *KeggPathwayMutation {
	return kpc.mutation
}

// Save creates the KeggPathway in the database.
func (kpc *KeggPathwayCreate) Save(ctx context.Context) (*KeggPathway, error) {
	var (
		err  error
		node *KeggPathway
	)
	if len(kpc.hooks) == 0 {
		if err = kpc.check(); err != nil {
			return nil, err
		}
		node, err = kpc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*KeggPathwayMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = kpc.check(); err != nil {
				return nil, err
			}
			kpc.mutation = mutation
			if node, err = kpc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(kpc.hooks) - 1; i >= 0; i-- {
			if kpc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = kpc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, kpc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*KeggPathway)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from KeggPathwayMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (kpc *KeggPathwayCreate) SaveX(ctx context.Context) *KeggPathway {
	v, err := kpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (kpc *KeggPathwayCreate) Exec(ctx context.Context) error {
	_, err := kpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kpc *KeggPathwayCreate) ExecX(ctx context.Context) {
	if err := kpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (kpc *KeggPathwayCreate) check() error {
	if _, ok := kpc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "KeggPathway.name"`)}
	}
	return nil
}

func (kpc *KeggPathwayCreate) sqlSave(ctx context.Context) (*KeggPathway, error) {
	_node, _spec := kpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, kpc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected KeggPathway.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (kpc *KeggPathwayCreate) createSpec() (*KeggPathway, *sqlgraph.CreateSpec) {
	var (
		_node = &KeggPathway{config: kpc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: keggpathway.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: keggpathway.FieldID,
			},
		}
	)
	_spec.OnConflict = kpc.conflict
	if id, ok := kpc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := kpc.mutation.Name(); ok {
		_spec.SetField(keggpathway.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if nodes := kpc.mutation.RelatingMapIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   keggpathway.RelatingMapTable,
			Columns: keggpathway.RelatingMapPrimaryKey,
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := kpc.mutation.RelatedMapIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   keggpathway.RelatedMapTable,
			Columns: keggpathway.RelatedMapPrimaryKey,
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := kpc.mutation.ReactionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   keggpathway.ReactionsTable,
			Columns: keggpathway.ReactionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: keggreaction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := kpc.mutation.OrthologiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   keggpathway.OrthologiesTable,
			Columns: keggpathway.OrthologiesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: keggorthlogy.FieldID,
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
//	client.KeggPathway.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.KeggPathwayUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (kpc *KeggPathwayCreate) OnConflict(opts ...sql.ConflictOption) *KeggPathwayUpsertOne {
	kpc.conflict = opts
	return &KeggPathwayUpsertOne{
		create: kpc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.KeggPathway.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (kpc *KeggPathwayCreate) OnConflictColumns(columns ...string) *KeggPathwayUpsertOne {
	kpc.conflict = append(kpc.conflict, sql.ConflictColumns(columns...))
	return &KeggPathwayUpsertOne{
		create: kpc,
	}
}

type (
	// KeggPathwayUpsertOne is the builder for "upsert"-ing
	//  one KeggPathway node.
	KeggPathwayUpsertOne struct {
		create *KeggPathwayCreate
	}

	// KeggPathwayUpsert is the "OnConflict" setter.
	KeggPathwayUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *KeggPathwayUpsert) SetName(v string) *KeggPathwayUpsert {
	u.Set(keggpathway.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *KeggPathwayUpsert) UpdateName() *KeggPathwayUpsert {
	u.SetExcluded(keggpathway.FieldName)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.KeggPathway.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(keggpathway.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *KeggPathwayUpsertOne) UpdateNewValues() *KeggPathwayUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(keggpathway.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.KeggPathway.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *KeggPathwayUpsertOne) Ignore() *KeggPathwayUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *KeggPathwayUpsertOne) DoNothing() *KeggPathwayUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the KeggPathwayCreate.OnConflict
// documentation for more info.
func (u *KeggPathwayUpsertOne) Update(set func(*KeggPathwayUpsert)) *KeggPathwayUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&KeggPathwayUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *KeggPathwayUpsertOne) SetName(v string) *KeggPathwayUpsertOne {
	return u.Update(func(s *KeggPathwayUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *KeggPathwayUpsertOne) UpdateName() *KeggPathwayUpsertOne {
	return u.Update(func(s *KeggPathwayUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *KeggPathwayUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for KeggPathwayCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *KeggPathwayUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *KeggPathwayUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: KeggPathwayUpsertOne.ID is not supported by MySQL driver. Use KeggPathwayUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *KeggPathwayUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// KeggPathwayCreateBulk is the builder for creating many KeggPathway entities in bulk.
type KeggPathwayCreateBulk struct {
	config
	builders []*KeggPathwayCreate
	conflict []sql.ConflictOption
}

// Save creates the KeggPathway entities in the database.
func (kpcb *KeggPathwayCreateBulk) Save(ctx context.Context) ([]*KeggPathway, error) {
	specs := make([]*sqlgraph.CreateSpec, len(kpcb.builders))
	nodes := make([]*KeggPathway, len(kpcb.builders))
	mutators := make([]Mutator, len(kpcb.builders))
	for i := range kpcb.builders {
		func(i int, root context.Context) {
			builder := kpcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*KeggPathwayMutation)
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
					_, err = mutators[i+1].Mutate(root, kpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = kpcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, kpcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, kpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (kpcb *KeggPathwayCreateBulk) SaveX(ctx context.Context) []*KeggPathway {
	v, err := kpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (kpcb *KeggPathwayCreateBulk) Exec(ctx context.Context) error {
	_, err := kpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kpcb *KeggPathwayCreateBulk) ExecX(ctx context.Context) {
	if err := kpcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.KeggPathway.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.KeggPathwayUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (kpcb *KeggPathwayCreateBulk) OnConflict(opts ...sql.ConflictOption) *KeggPathwayUpsertBulk {
	kpcb.conflict = opts
	return &KeggPathwayUpsertBulk{
		create: kpcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.KeggPathway.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (kpcb *KeggPathwayCreateBulk) OnConflictColumns(columns ...string) *KeggPathwayUpsertBulk {
	kpcb.conflict = append(kpcb.conflict, sql.ConflictColumns(columns...))
	return &KeggPathwayUpsertBulk{
		create: kpcb,
	}
}

// KeggPathwayUpsertBulk is the builder for "upsert"-ing
// a bulk of KeggPathway nodes.
type KeggPathwayUpsertBulk struct {
	create *KeggPathwayCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.KeggPathway.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(keggpathway.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *KeggPathwayUpsertBulk) UpdateNewValues() *KeggPathwayUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(keggpathway.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.KeggPathway.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *KeggPathwayUpsertBulk) Ignore() *KeggPathwayUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *KeggPathwayUpsertBulk) DoNothing() *KeggPathwayUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the KeggPathwayCreateBulk.OnConflict
// documentation for more info.
func (u *KeggPathwayUpsertBulk) Update(set func(*KeggPathwayUpsert)) *KeggPathwayUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&KeggPathwayUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *KeggPathwayUpsertBulk) SetName(v string) *KeggPathwayUpsertBulk {
	return u.Update(func(s *KeggPathwayUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *KeggPathwayUpsertBulk) UpdateName() *KeggPathwayUpsertBulk {
	return u.Update(func(s *KeggPathwayUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *KeggPathwayUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the KeggPathwayCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for KeggPathwayCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *KeggPathwayUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
