// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"genomedb/ent/domainannotation"
	"genomedb/ent/transcript"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DomainAnnotationCreate is the builder for creating a DomainAnnotation entity.
type DomainAnnotationCreate struct {
	config
	mutation *DomainAnnotationMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetDescription sets the "description" field.
func (dac *DomainAnnotationCreate) SetDescription(s string) *DomainAnnotationCreate {
	dac.mutation.SetDescription(s)
	return dac
}

// SetAnalysis sets the "Analysis" field.
func (dac *DomainAnnotationCreate) SetAnalysis(d domainannotation.Analysis) *DomainAnnotationCreate {
	dac.mutation.SetAnalysis(d)
	return dac
}

// SetID sets the "id" field.
func (dac *DomainAnnotationCreate) SetID(s string) *DomainAnnotationCreate {
	dac.mutation.SetID(s)
	return dac
}

// AddTranscriptIDs adds the "transcripts" edge to the Transcript entity by IDs.
func (dac *DomainAnnotationCreate) AddTranscriptIDs(ids ...string) *DomainAnnotationCreate {
	dac.mutation.AddTranscriptIDs(ids...)
	return dac
}

// AddTranscripts adds the "transcripts" edges to the Transcript entity.
func (dac *DomainAnnotationCreate) AddTranscripts(t ...*Transcript) *DomainAnnotationCreate {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return dac.AddTranscriptIDs(ids...)
}

// Mutation returns the DomainAnnotationMutation object of the builder.
func (dac *DomainAnnotationCreate) Mutation() *DomainAnnotationMutation {
	return dac.mutation
}

// Save creates the DomainAnnotation in the database.
func (dac *DomainAnnotationCreate) Save(ctx context.Context) (*DomainAnnotation, error) {
	var (
		err  error
		node *DomainAnnotation
	)
	if len(dac.hooks) == 0 {
		if err = dac.check(); err != nil {
			return nil, err
		}
		node, err = dac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DomainAnnotationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dac.check(); err != nil {
				return nil, err
			}
			dac.mutation = mutation
			if node, err = dac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(dac.hooks) - 1; i >= 0; i-- {
			if dac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dac.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, dac.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*DomainAnnotation)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from DomainAnnotationMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dac *DomainAnnotationCreate) SaveX(ctx context.Context) *DomainAnnotation {
	v, err := dac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dac *DomainAnnotationCreate) Exec(ctx context.Context) error {
	_, err := dac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dac *DomainAnnotationCreate) ExecX(ctx context.Context) {
	if err := dac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dac *DomainAnnotationCreate) check() error {
	if _, ok := dac.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "DomainAnnotation.description"`)}
	}
	if _, ok := dac.mutation.Analysis(); !ok {
		return &ValidationError{Name: "Analysis", err: errors.New(`ent: missing required field "DomainAnnotation.Analysis"`)}
	}
	if v, ok := dac.mutation.Analysis(); ok {
		if err := domainannotation.AnalysisValidator(v); err != nil {
			return &ValidationError{Name: "Analysis", err: fmt.Errorf(`ent: validator failed for field "DomainAnnotation.Analysis": %w`, err)}
		}
	}
	return nil
}

func (dac *DomainAnnotationCreate) sqlSave(ctx context.Context) (*DomainAnnotation, error) {
	_node, _spec := dac.createSpec()
	if err := sqlgraph.CreateNode(ctx, dac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected DomainAnnotation.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (dac *DomainAnnotationCreate) createSpec() (*DomainAnnotation, *sqlgraph.CreateSpec) {
	var (
		_node = &DomainAnnotation{config: dac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: domainannotation.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: domainannotation.FieldID,
			},
		}
	)
	_spec.OnConflict = dac.conflict
	if id, ok := dac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := dac.mutation.Description(); ok {
		_spec.SetField(domainannotation.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := dac.mutation.Analysis(); ok {
		_spec.SetField(domainannotation.FieldAnalysis, field.TypeEnum, value)
		_node.Analysis = value
	}
	if nodes := dac.mutation.TranscriptsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   domainannotation.TranscriptsTable,
			Columns: domainannotation.TranscriptsPrimaryKey,
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
//	client.DomainAnnotation.Create().
//		SetDescription(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DomainAnnotationUpsert) {
//			SetDescription(v+v).
//		}).
//		Exec(ctx)
func (dac *DomainAnnotationCreate) OnConflict(opts ...sql.ConflictOption) *DomainAnnotationUpsertOne {
	dac.conflict = opts
	return &DomainAnnotationUpsertOne{
		create: dac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DomainAnnotation.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dac *DomainAnnotationCreate) OnConflictColumns(columns ...string) *DomainAnnotationUpsertOne {
	dac.conflict = append(dac.conflict, sql.ConflictColumns(columns...))
	return &DomainAnnotationUpsertOne{
		create: dac,
	}
}

type (
	// DomainAnnotationUpsertOne is the builder for "upsert"-ing
	//  one DomainAnnotation node.
	DomainAnnotationUpsertOne struct {
		create *DomainAnnotationCreate
	}

	// DomainAnnotationUpsert is the "OnConflict" setter.
	DomainAnnotationUpsert struct {
		*sql.UpdateSet
	}
)

// SetDescription sets the "description" field.
func (u *DomainAnnotationUpsert) SetDescription(v string) *DomainAnnotationUpsert {
	u.Set(domainannotation.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *DomainAnnotationUpsert) UpdateDescription() *DomainAnnotationUpsert {
	u.SetExcluded(domainannotation.FieldDescription)
	return u
}

// SetAnalysis sets the "Analysis" field.
func (u *DomainAnnotationUpsert) SetAnalysis(v domainannotation.Analysis) *DomainAnnotationUpsert {
	u.Set(domainannotation.FieldAnalysis, v)
	return u
}

// UpdateAnalysis sets the "Analysis" field to the value that was provided on create.
func (u *DomainAnnotationUpsert) UpdateAnalysis() *DomainAnnotationUpsert {
	u.SetExcluded(domainannotation.FieldAnalysis)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.DomainAnnotation.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(domainannotation.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DomainAnnotationUpsertOne) UpdateNewValues() *DomainAnnotationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(domainannotation.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DomainAnnotation.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *DomainAnnotationUpsertOne) Ignore() *DomainAnnotationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DomainAnnotationUpsertOne) DoNothing() *DomainAnnotationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DomainAnnotationCreate.OnConflict
// documentation for more info.
func (u *DomainAnnotationUpsertOne) Update(set func(*DomainAnnotationUpsert)) *DomainAnnotationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DomainAnnotationUpsert{UpdateSet: update})
	}))
	return u
}

// SetDescription sets the "description" field.
func (u *DomainAnnotationUpsertOne) SetDescription(v string) *DomainAnnotationUpsertOne {
	return u.Update(func(s *DomainAnnotationUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *DomainAnnotationUpsertOne) UpdateDescription() *DomainAnnotationUpsertOne {
	return u.Update(func(s *DomainAnnotationUpsert) {
		s.UpdateDescription()
	})
}

// SetAnalysis sets the "Analysis" field.
func (u *DomainAnnotationUpsertOne) SetAnalysis(v domainannotation.Analysis) *DomainAnnotationUpsertOne {
	return u.Update(func(s *DomainAnnotationUpsert) {
		s.SetAnalysis(v)
	})
}

// UpdateAnalysis sets the "Analysis" field to the value that was provided on create.
func (u *DomainAnnotationUpsertOne) UpdateAnalysis() *DomainAnnotationUpsertOne {
	return u.Update(func(s *DomainAnnotationUpsert) {
		s.UpdateAnalysis()
	})
}

// Exec executes the query.
func (u *DomainAnnotationUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DomainAnnotationCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DomainAnnotationUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *DomainAnnotationUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: DomainAnnotationUpsertOne.ID is not supported by MySQL driver. Use DomainAnnotationUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *DomainAnnotationUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// DomainAnnotationCreateBulk is the builder for creating many DomainAnnotation entities in bulk.
type DomainAnnotationCreateBulk struct {
	config
	builders []*DomainAnnotationCreate
	conflict []sql.ConflictOption
}

// Save creates the DomainAnnotation entities in the database.
func (dacb *DomainAnnotationCreateBulk) Save(ctx context.Context) ([]*DomainAnnotation, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dacb.builders))
	nodes := make([]*DomainAnnotation, len(dacb.builders))
	mutators := make([]Mutator, len(dacb.builders))
	for i := range dacb.builders {
		func(i int, root context.Context) {
			builder := dacb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DomainAnnotationMutation)
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
					_, err = mutators[i+1].Mutate(root, dacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = dacb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dacb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dacb *DomainAnnotationCreateBulk) SaveX(ctx context.Context) []*DomainAnnotation {
	v, err := dacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dacb *DomainAnnotationCreateBulk) Exec(ctx context.Context) error {
	_, err := dacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dacb *DomainAnnotationCreateBulk) ExecX(ctx context.Context) {
	if err := dacb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DomainAnnotation.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DomainAnnotationUpsert) {
//			SetDescription(v+v).
//		}).
//		Exec(ctx)
func (dacb *DomainAnnotationCreateBulk) OnConflict(opts ...sql.ConflictOption) *DomainAnnotationUpsertBulk {
	dacb.conflict = opts
	return &DomainAnnotationUpsertBulk{
		create: dacb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DomainAnnotation.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dacb *DomainAnnotationCreateBulk) OnConflictColumns(columns ...string) *DomainAnnotationUpsertBulk {
	dacb.conflict = append(dacb.conflict, sql.ConflictColumns(columns...))
	return &DomainAnnotationUpsertBulk{
		create: dacb,
	}
}

// DomainAnnotationUpsertBulk is the builder for "upsert"-ing
// a bulk of DomainAnnotation nodes.
type DomainAnnotationUpsertBulk struct {
	create *DomainAnnotationCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.DomainAnnotation.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(domainannotation.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DomainAnnotationUpsertBulk) UpdateNewValues() *DomainAnnotationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(domainannotation.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DomainAnnotation.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *DomainAnnotationUpsertBulk) Ignore() *DomainAnnotationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DomainAnnotationUpsertBulk) DoNothing() *DomainAnnotationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DomainAnnotationCreateBulk.OnConflict
// documentation for more info.
func (u *DomainAnnotationUpsertBulk) Update(set func(*DomainAnnotationUpsert)) *DomainAnnotationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DomainAnnotationUpsert{UpdateSet: update})
	}))
	return u
}

// SetDescription sets the "description" field.
func (u *DomainAnnotationUpsertBulk) SetDescription(v string) *DomainAnnotationUpsertBulk {
	return u.Update(func(s *DomainAnnotationUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *DomainAnnotationUpsertBulk) UpdateDescription() *DomainAnnotationUpsertBulk {
	return u.Update(func(s *DomainAnnotationUpsert) {
		s.UpdateDescription()
	})
}

// SetAnalysis sets the "Analysis" field.
func (u *DomainAnnotationUpsertBulk) SetAnalysis(v domainannotation.Analysis) *DomainAnnotationUpsertBulk {
	return u.Update(func(s *DomainAnnotationUpsert) {
		s.SetAnalysis(v)
	})
}

// UpdateAnalysis sets the "Analysis" field to the value that was provided on create.
func (u *DomainAnnotationUpsertBulk) UpdateAnalysis() *DomainAnnotationUpsertBulk {
	return u.Update(func(s *DomainAnnotationUpsert) {
		s.UpdateAnalysis()
	})
}

// Exec executes the query.
func (u *DomainAnnotationUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the DomainAnnotationCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DomainAnnotationCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DomainAnnotationUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
