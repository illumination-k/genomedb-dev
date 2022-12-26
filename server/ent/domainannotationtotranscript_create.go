// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"genomedb/ent/domainannotation"
	"genomedb/ent/domainannotationtotranscript"
	"genomedb/ent/transcript"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DomainAnnotationToTranscriptCreate is the builder for creating a DomainAnnotationToTranscript entity.
type DomainAnnotationToTranscriptCreate struct {
	config
	mutation *DomainAnnotationToTranscriptMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetDomainAnnotationID sets the "domain_annotation_id" field.
func (dattc *DomainAnnotationToTranscriptCreate) SetDomainAnnotationID(s string) *DomainAnnotationToTranscriptCreate {
	dattc.mutation.SetDomainAnnotationID(s)
	return dattc
}

// SetTranscriptID sets the "transcript_id" field.
func (dattc *DomainAnnotationToTranscriptCreate) SetTranscriptID(s string) *DomainAnnotationToTranscriptCreate {
	dattc.mutation.SetTranscriptID(s)
	return dattc
}

// SetStart sets the "start" field.
func (dattc *DomainAnnotationToTranscriptCreate) SetStart(i int32) *DomainAnnotationToTranscriptCreate {
	dattc.mutation.SetStart(i)
	return dattc
}

// SetStop sets the "stop" field.
func (dattc *DomainAnnotationToTranscriptCreate) SetStop(i int32) *DomainAnnotationToTranscriptCreate {
	dattc.mutation.SetStop(i)
	return dattc
}

// SetScore sets the "score" field.
func (dattc *DomainAnnotationToTranscriptCreate) SetScore(f float64) *DomainAnnotationToTranscriptCreate {
	dattc.mutation.SetScore(f)
	return dattc
}

// SetDomainID sets the "domain" edge to the DomainAnnotation entity by ID.
func (dattc *DomainAnnotationToTranscriptCreate) SetDomainID(id string) *DomainAnnotationToTranscriptCreate {
	dattc.mutation.SetDomainID(id)
	return dattc
}

// SetDomain sets the "domain" edge to the DomainAnnotation entity.
func (dattc *DomainAnnotationToTranscriptCreate) SetDomain(d *DomainAnnotation) *DomainAnnotationToTranscriptCreate {
	return dattc.SetDomainID(d.ID)
}

// SetTranscript sets the "transcript" edge to the Transcript entity.
func (dattc *DomainAnnotationToTranscriptCreate) SetTranscript(t *Transcript) *DomainAnnotationToTranscriptCreate {
	return dattc.SetTranscriptID(t.ID)
}

// Mutation returns the DomainAnnotationToTranscriptMutation object of the builder.
func (dattc *DomainAnnotationToTranscriptCreate) Mutation() *DomainAnnotationToTranscriptMutation {
	return dattc.mutation
}

// Save creates the DomainAnnotationToTranscript in the database.
func (dattc *DomainAnnotationToTranscriptCreate) Save(ctx context.Context) (*DomainAnnotationToTranscript, error) {
	var (
		err  error
		node *DomainAnnotationToTranscript
	)
	if len(dattc.hooks) == 0 {
		if err = dattc.check(); err != nil {
			return nil, err
		}
		node, err = dattc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DomainAnnotationToTranscriptMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dattc.check(); err != nil {
				return nil, err
			}
			dattc.mutation = mutation
			if node, err = dattc.sqlSave(ctx); err != nil {
				return nil, err
			}
			return node, err
		})
		for i := len(dattc.hooks) - 1; i >= 0; i-- {
			if dattc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dattc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, dattc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*DomainAnnotationToTranscript)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from DomainAnnotationToTranscriptMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dattc *DomainAnnotationToTranscriptCreate) SaveX(ctx context.Context) *DomainAnnotationToTranscript {
	v, err := dattc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dattc *DomainAnnotationToTranscriptCreate) Exec(ctx context.Context) error {
	_, err := dattc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dattc *DomainAnnotationToTranscriptCreate) ExecX(ctx context.Context) {
	if err := dattc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dattc *DomainAnnotationToTranscriptCreate) check() error {
	if _, ok := dattc.mutation.DomainAnnotationID(); !ok {
		return &ValidationError{Name: "domain_annotation_id", err: errors.New(`ent: missing required field "DomainAnnotationToTranscript.domain_annotation_id"`)}
	}
	if _, ok := dattc.mutation.TranscriptID(); !ok {
		return &ValidationError{Name: "transcript_id", err: errors.New(`ent: missing required field "DomainAnnotationToTranscript.transcript_id"`)}
	}
	if _, ok := dattc.mutation.Start(); !ok {
		return &ValidationError{Name: "start", err: errors.New(`ent: missing required field "DomainAnnotationToTranscript.start"`)}
	}
	if v, ok := dattc.mutation.Start(); ok {
		if err := domainannotationtotranscript.StartValidator(v); err != nil {
			return &ValidationError{Name: "start", err: fmt.Errorf(`ent: validator failed for field "DomainAnnotationToTranscript.start": %w`, err)}
		}
	}
	if _, ok := dattc.mutation.Stop(); !ok {
		return &ValidationError{Name: "stop", err: errors.New(`ent: missing required field "DomainAnnotationToTranscript.stop"`)}
	}
	if v, ok := dattc.mutation.Stop(); ok {
		if err := domainannotationtotranscript.StopValidator(v); err != nil {
			return &ValidationError{Name: "stop", err: fmt.Errorf(`ent: validator failed for field "DomainAnnotationToTranscript.stop": %w`, err)}
		}
	}
	if _, ok := dattc.mutation.Score(); !ok {
		return &ValidationError{Name: "score", err: errors.New(`ent: missing required field "DomainAnnotationToTranscript.score"`)}
	}
	if _, ok := dattc.mutation.DomainID(); !ok {
		return &ValidationError{Name: "domain", err: errors.New(`ent: missing required edge "DomainAnnotationToTranscript.domain"`)}
	}
	if _, ok := dattc.mutation.TranscriptID(); !ok {
		return &ValidationError{Name: "transcript", err: errors.New(`ent: missing required edge "DomainAnnotationToTranscript.transcript"`)}
	}
	return nil
}

func (dattc *DomainAnnotationToTranscriptCreate) sqlSave(ctx context.Context) (*DomainAnnotationToTranscript, error) {
	_node, _spec := dattc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dattc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

func (dattc *DomainAnnotationToTranscriptCreate) createSpec() (*DomainAnnotationToTranscript, *sqlgraph.CreateSpec) {
	var (
		_node = &DomainAnnotationToTranscript{config: dattc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: domainannotationtotranscript.Table,
		}
	)
	_spec.OnConflict = dattc.conflict
	if value, ok := dattc.mutation.Start(); ok {
		_spec.SetField(domainannotationtotranscript.FieldStart, field.TypeInt32, value)
		_node.Start = value
	}
	if value, ok := dattc.mutation.Stop(); ok {
		_spec.SetField(domainannotationtotranscript.FieldStop, field.TypeInt32, value)
		_node.Stop = value
	}
	if value, ok := dattc.mutation.Score(); ok {
		_spec.SetField(domainannotationtotranscript.FieldScore, field.TypeFloat64, value)
		_node.Score = value
	}
	if nodes := dattc.mutation.DomainIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   domainannotationtotranscript.DomainTable,
			Columns: []string{domainannotationtotranscript.DomainColumn},
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
		_node.DomainAnnotationID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dattc.mutation.TranscriptIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   domainannotationtotranscript.TranscriptTable,
			Columns: []string{domainannotationtotranscript.TranscriptColumn},
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
		_node.TranscriptID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DomainAnnotationToTranscript.Create().
//		SetDomainAnnotationID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DomainAnnotationToTranscriptUpsert) {
//			SetDomainAnnotationID(v+v).
//		}).
//		Exec(ctx)
func (dattc *DomainAnnotationToTranscriptCreate) OnConflict(opts ...sql.ConflictOption) *DomainAnnotationToTranscriptUpsertOne {
	dattc.conflict = opts
	return &DomainAnnotationToTranscriptUpsertOne{
		create: dattc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DomainAnnotationToTranscript.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dattc *DomainAnnotationToTranscriptCreate) OnConflictColumns(columns ...string) *DomainAnnotationToTranscriptUpsertOne {
	dattc.conflict = append(dattc.conflict, sql.ConflictColumns(columns...))
	return &DomainAnnotationToTranscriptUpsertOne{
		create: dattc,
	}
}

type (
	// DomainAnnotationToTranscriptUpsertOne is the builder for "upsert"-ing
	//  one DomainAnnotationToTranscript node.
	DomainAnnotationToTranscriptUpsertOne struct {
		create *DomainAnnotationToTranscriptCreate
	}

	// DomainAnnotationToTranscriptUpsert is the "OnConflict" setter.
	DomainAnnotationToTranscriptUpsert struct {
		*sql.UpdateSet
	}
)

// SetDomainAnnotationID sets the "domain_annotation_id" field.
func (u *DomainAnnotationToTranscriptUpsert) SetDomainAnnotationID(v string) *DomainAnnotationToTranscriptUpsert {
	u.Set(domainannotationtotranscript.FieldDomainAnnotationID, v)
	return u
}

// UpdateDomainAnnotationID sets the "domain_annotation_id" field to the value that was provided on create.
func (u *DomainAnnotationToTranscriptUpsert) UpdateDomainAnnotationID() *DomainAnnotationToTranscriptUpsert {
	u.SetExcluded(domainannotationtotranscript.FieldDomainAnnotationID)
	return u
}

// SetTranscriptID sets the "transcript_id" field.
func (u *DomainAnnotationToTranscriptUpsert) SetTranscriptID(v string) *DomainAnnotationToTranscriptUpsert {
	u.Set(domainannotationtotranscript.FieldTranscriptID, v)
	return u
}

// UpdateTranscriptID sets the "transcript_id" field to the value that was provided on create.
func (u *DomainAnnotationToTranscriptUpsert) UpdateTranscriptID() *DomainAnnotationToTranscriptUpsert {
	u.SetExcluded(domainannotationtotranscript.FieldTranscriptID)
	return u
}

// SetStart sets the "start" field.
func (u *DomainAnnotationToTranscriptUpsert) SetStart(v int32) *DomainAnnotationToTranscriptUpsert {
	u.Set(domainannotationtotranscript.FieldStart, v)
	return u
}

// UpdateStart sets the "start" field to the value that was provided on create.
func (u *DomainAnnotationToTranscriptUpsert) UpdateStart() *DomainAnnotationToTranscriptUpsert {
	u.SetExcluded(domainannotationtotranscript.FieldStart)
	return u
}

// AddStart adds v to the "start" field.
func (u *DomainAnnotationToTranscriptUpsert) AddStart(v int32) *DomainAnnotationToTranscriptUpsert {
	u.Add(domainannotationtotranscript.FieldStart, v)
	return u
}

// SetStop sets the "stop" field.
func (u *DomainAnnotationToTranscriptUpsert) SetStop(v int32) *DomainAnnotationToTranscriptUpsert {
	u.Set(domainannotationtotranscript.FieldStop, v)
	return u
}

// UpdateStop sets the "stop" field to the value that was provided on create.
func (u *DomainAnnotationToTranscriptUpsert) UpdateStop() *DomainAnnotationToTranscriptUpsert {
	u.SetExcluded(domainannotationtotranscript.FieldStop)
	return u
}

// AddStop adds v to the "stop" field.
func (u *DomainAnnotationToTranscriptUpsert) AddStop(v int32) *DomainAnnotationToTranscriptUpsert {
	u.Add(domainannotationtotranscript.FieldStop, v)
	return u
}

// SetScore sets the "score" field.
func (u *DomainAnnotationToTranscriptUpsert) SetScore(v float64) *DomainAnnotationToTranscriptUpsert {
	u.Set(domainannotationtotranscript.FieldScore, v)
	return u
}

// UpdateScore sets the "score" field to the value that was provided on create.
func (u *DomainAnnotationToTranscriptUpsert) UpdateScore() *DomainAnnotationToTranscriptUpsert {
	u.SetExcluded(domainannotationtotranscript.FieldScore)
	return u
}

// AddScore adds v to the "score" field.
func (u *DomainAnnotationToTranscriptUpsert) AddScore(v float64) *DomainAnnotationToTranscriptUpsert {
	u.Add(domainannotationtotranscript.FieldScore, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.DomainAnnotationToTranscript.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *DomainAnnotationToTranscriptUpsertOne) UpdateNewValues() *DomainAnnotationToTranscriptUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DomainAnnotationToTranscript.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *DomainAnnotationToTranscriptUpsertOne) Ignore() *DomainAnnotationToTranscriptUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DomainAnnotationToTranscriptUpsertOne) DoNothing() *DomainAnnotationToTranscriptUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DomainAnnotationToTranscriptCreate.OnConflict
// documentation for more info.
func (u *DomainAnnotationToTranscriptUpsertOne) Update(set func(*DomainAnnotationToTranscriptUpsert)) *DomainAnnotationToTranscriptUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DomainAnnotationToTranscriptUpsert{UpdateSet: update})
	}))
	return u
}

// SetDomainAnnotationID sets the "domain_annotation_id" field.
func (u *DomainAnnotationToTranscriptUpsertOne) SetDomainAnnotationID(v string) *DomainAnnotationToTranscriptUpsertOne {
	return u.Update(func(s *DomainAnnotationToTranscriptUpsert) {
		s.SetDomainAnnotationID(v)
	})
}

// UpdateDomainAnnotationID sets the "domain_annotation_id" field to the value that was provided on create.
func (u *DomainAnnotationToTranscriptUpsertOne) UpdateDomainAnnotationID() *DomainAnnotationToTranscriptUpsertOne {
	return u.Update(func(s *DomainAnnotationToTranscriptUpsert) {
		s.UpdateDomainAnnotationID()
	})
}

// SetTranscriptID sets the "transcript_id" field.
func (u *DomainAnnotationToTranscriptUpsertOne) SetTranscriptID(v string) *DomainAnnotationToTranscriptUpsertOne {
	return u.Update(func(s *DomainAnnotationToTranscriptUpsert) {
		s.SetTranscriptID(v)
	})
}

// UpdateTranscriptID sets the "transcript_id" field to the value that was provided on create.
func (u *DomainAnnotationToTranscriptUpsertOne) UpdateTranscriptID() *DomainAnnotationToTranscriptUpsertOne {
	return u.Update(func(s *DomainAnnotationToTranscriptUpsert) {
		s.UpdateTranscriptID()
	})
}

// SetStart sets the "start" field.
func (u *DomainAnnotationToTranscriptUpsertOne) SetStart(v int32) *DomainAnnotationToTranscriptUpsertOne {
	return u.Update(func(s *DomainAnnotationToTranscriptUpsert) {
		s.SetStart(v)
	})
}

// AddStart adds v to the "start" field.
func (u *DomainAnnotationToTranscriptUpsertOne) AddStart(v int32) *DomainAnnotationToTranscriptUpsertOne {
	return u.Update(func(s *DomainAnnotationToTranscriptUpsert) {
		s.AddStart(v)
	})
}

// UpdateStart sets the "start" field to the value that was provided on create.
func (u *DomainAnnotationToTranscriptUpsertOne) UpdateStart() *DomainAnnotationToTranscriptUpsertOne {
	return u.Update(func(s *DomainAnnotationToTranscriptUpsert) {
		s.UpdateStart()
	})
}

// SetStop sets the "stop" field.
func (u *DomainAnnotationToTranscriptUpsertOne) SetStop(v int32) *DomainAnnotationToTranscriptUpsertOne {
	return u.Update(func(s *DomainAnnotationToTranscriptUpsert) {
		s.SetStop(v)
	})
}

// AddStop adds v to the "stop" field.
func (u *DomainAnnotationToTranscriptUpsertOne) AddStop(v int32) *DomainAnnotationToTranscriptUpsertOne {
	return u.Update(func(s *DomainAnnotationToTranscriptUpsert) {
		s.AddStop(v)
	})
}

// UpdateStop sets the "stop" field to the value that was provided on create.
func (u *DomainAnnotationToTranscriptUpsertOne) UpdateStop() *DomainAnnotationToTranscriptUpsertOne {
	return u.Update(func(s *DomainAnnotationToTranscriptUpsert) {
		s.UpdateStop()
	})
}

// SetScore sets the "score" field.
func (u *DomainAnnotationToTranscriptUpsertOne) SetScore(v float64) *DomainAnnotationToTranscriptUpsertOne {
	return u.Update(func(s *DomainAnnotationToTranscriptUpsert) {
		s.SetScore(v)
	})
}

// AddScore adds v to the "score" field.
func (u *DomainAnnotationToTranscriptUpsertOne) AddScore(v float64) *DomainAnnotationToTranscriptUpsertOne {
	return u.Update(func(s *DomainAnnotationToTranscriptUpsert) {
		s.AddScore(v)
	})
}

// UpdateScore sets the "score" field to the value that was provided on create.
func (u *DomainAnnotationToTranscriptUpsertOne) UpdateScore() *DomainAnnotationToTranscriptUpsertOne {
	return u.Update(func(s *DomainAnnotationToTranscriptUpsert) {
		s.UpdateScore()
	})
}

// Exec executes the query.
func (u *DomainAnnotationToTranscriptUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DomainAnnotationToTranscriptCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DomainAnnotationToTranscriptUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// DomainAnnotationToTranscriptCreateBulk is the builder for creating many DomainAnnotationToTranscript entities in bulk.
type DomainAnnotationToTranscriptCreateBulk struct {
	config
	builders []*DomainAnnotationToTranscriptCreate
	conflict []sql.ConflictOption
}

// Save creates the DomainAnnotationToTranscript entities in the database.
func (dattcb *DomainAnnotationToTranscriptCreateBulk) Save(ctx context.Context) ([]*DomainAnnotationToTranscript, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dattcb.builders))
	nodes := make([]*DomainAnnotationToTranscript, len(dattcb.builders))
	mutators := make([]Mutator, len(dattcb.builders))
	for i := range dattcb.builders {
		func(i int, root context.Context) {
			builder := dattcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DomainAnnotationToTranscriptMutation)
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
					_, err = mutators[i+1].Mutate(root, dattcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = dattcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dattcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
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
		if _, err := mutators[0].Mutate(ctx, dattcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dattcb *DomainAnnotationToTranscriptCreateBulk) SaveX(ctx context.Context) []*DomainAnnotationToTranscript {
	v, err := dattcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dattcb *DomainAnnotationToTranscriptCreateBulk) Exec(ctx context.Context) error {
	_, err := dattcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dattcb *DomainAnnotationToTranscriptCreateBulk) ExecX(ctx context.Context) {
	if err := dattcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DomainAnnotationToTranscript.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DomainAnnotationToTranscriptUpsert) {
//			SetDomainAnnotationID(v+v).
//		}).
//		Exec(ctx)
func (dattcb *DomainAnnotationToTranscriptCreateBulk) OnConflict(opts ...sql.ConflictOption) *DomainAnnotationToTranscriptUpsertBulk {
	dattcb.conflict = opts
	return &DomainAnnotationToTranscriptUpsertBulk{
		create: dattcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DomainAnnotationToTranscript.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dattcb *DomainAnnotationToTranscriptCreateBulk) OnConflictColumns(columns ...string) *DomainAnnotationToTranscriptUpsertBulk {
	dattcb.conflict = append(dattcb.conflict, sql.ConflictColumns(columns...))
	return &DomainAnnotationToTranscriptUpsertBulk{
		create: dattcb,
	}
}

// DomainAnnotationToTranscriptUpsertBulk is the builder for "upsert"-ing
// a bulk of DomainAnnotationToTranscript nodes.
type DomainAnnotationToTranscriptUpsertBulk struct {
	create *DomainAnnotationToTranscriptCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.DomainAnnotationToTranscript.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *DomainAnnotationToTranscriptUpsertBulk) UpdateNewValues() *DomainAnnotationToTranscriptUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DomainAnnotationToTranscript.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *DomainAnnotationToTranscriptUpsertBulk) Ignore() *DomainAnnotationToTranscriptUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DomainAnnotationToTranscriptUpsertBulk) DoNothing() *DomainAnnotationToTranscriptUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DomainAnnotationToTranscriptCreateBulk.OnConflict
// documentation for more info.
func (u *DomainAnnotationToTranscriptUpsertBulk) Update(set func(*DomainAnnotationToTranscriptUpsert)) *DomainAnnotationToTranscriptUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DomainAnnotationToTranscriptUpsert{UpdateSet: update})
	}))
	return u
}

// SetDomainAnnotationID sets the "domain_annotation_id" field.
func (u *DomainAnnotationToTranscriptUpsertBulk) SetDomainAnnotationID(v string) *DomainAnnotationToTranscriptUpsertBulk {
	return u.Update(func(s *DomainAnnotationToTranscriptUpsert) {
		s.SetDomainAnnotationID(v)
	})
}

// UpdateDomainAnnotationID sets the "domain_annotation_id" field to the value that was provided on create.
func (u *DomainAnnotationToTranscriptUpsertBulk) UpdateDomainAnnotationID() *DomainAnnotationToTranscriptUpsertBulk {
	return u.Update(func(s *DomainAnnotationToTranscriptUpsert) {
		s.UpdateDomainAnnotationID()
	})
}

// SetTranscriptID sets the "transcript_id" field.
func (u *DomainAnnotationToTranscriptUpsertBulk) SetTranscriptID(v string) *DomainAnnotationToTranscriptUpsertBulk {
	return u.Update(func(s *DomainAnnotationToTranscriptUpsert) {
		s.SetTranscriptID(v)
	})
}

// UpdateTranscriptID sets the "transcript_id" field to the value that was provided on create.
func (u *DomainAnnotationToTranscriptUpsertBulk) UpdateTranscriptID() *DomainAnnotationToTranscriptUpsertBulk {
	return u.Update(func(s *DomainAnnotationToTranscriptUpsert) {
		s.UpdateTranscriptID()
	})
}

// SetStart sets the "start" field.
func (u *DomainAnnotationToTranscriptUpsertBulk) SetStart(v int32) *DomainAnnotationToTranscriptUpsertBulk {
	return u.Update(func(s *DomainAnnotationToTranscriptUpsert) {
		s.SetStart(v)
	})
}

// AddStart adds v to the "start" field.
func (u *DomainAnnotationToTranscriptUpsertBulk) AddStart(v int32) *DomainAnnotationToTranscriptUpsertBulk {
	return u.Update(func(s *DomainAnnotationToTranscriptUpsert) {
		s.AddStart(v)
	})
}

// UpdateStart sets the "start" field to the value that was provided on create.
func (u *DomainAnnotationToTranscriptUpsertBulk) UpdateStart() *DomainAnnotationToTranscriptUpsertBulk {
	return u.Update(func(s *DomainAnnotationToTranscriptUpsert) {
		s.UpdateStart()
	})
}

// SetStop sets the "stop" field.
func (u *DomainAnnotationToTranscriptUpsertBulk) SetStop(v int32) *DomainAnnotationToTranscriptUpsertBulk {
	return u.Update(func(s *DomainAnnotationToTranscriptUpsert) {
		s.SetStop(v)
	})
}

// AddStop adds v to the "stop" field.
func (u *DomainAnnotationToTranscriptUpsertBulk) AddStop(v int32) *DomainAnnotationToTranscriptUpsertBulk {
	return u.Update(func(s *DomainAnnotationToTranscriptUpsert) {
		s.AddStop(v)
	})
}

// UpdateStop sets the "stop" field to the value that was provided on create.
func (u *DomainAnnotationToTranscriptUpsertBulk) UpdateStop() *DomainAnnotationToTranscriptUpsertBulk {
	return u.Update(func(s *DomainAnnotationToTranscriptUpsert) {
		s.UpdateStop()
	})
}

// SetScore sets the "score" field.
func (u *DomainAnnotationToTranscriptUpsertBulk) SetScore(v float64) *DomainAnnotationToTranscriptUpsertBulk {
	return u.Update(func(s *DomainAnnotationToTranscriptUpsert) {
		s.SetScore(v)
	})
}

// AddScore adds v to the "score" field.
func (u *DomainAnnotationToTranscriptUpsertBulk) AddScore(v float64) *DomainAnnotationToTranscriptUpsertBulk {
	return u.Update(func(s *DomainAnnotationToTranscriptUpsert) {
		s.AddScore(v)
	})
}

// UpdateScore sets the "score" field to the value that was provided on create.
func (u *DomainAnnotationToTranscriptUpsertBulk) UpdateScore() *DomainAnnotationToTranscriptUpsertBulk {
	return u.Update(func(s *DomainAnnotationToTranscriptUpsert) {
		s.UpdateScore()
	})
}

// Exec executes the query.
func (u *DomainAnnotationToTranscriptUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the DomainAnnotationToTranscriptCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DomainAnnotationToTranscriptCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DomainAnnotationToTranscriptUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
