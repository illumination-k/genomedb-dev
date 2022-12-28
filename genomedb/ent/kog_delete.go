// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"genomedb/ent/kog"
	"genomedb/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// KOGDelete is the builder for deleting a KOG entity.
type KOGDelete struct {
	config
	hooks    []Hook
	mutation *KOGMutation
}

// Where appends a list predicates to the KOGDelete builder.
func (kd *KOGDelete) Where(ps ...predicate.KOG) *KOGDelete {
	kd.mutation.Where(ps...)
	return kd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (kd *KOGDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(kd.hooks) == 0 {
		affected, err = kd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*KOGMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			kd.mutation = mutation
			affected, err = kd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(kd.hooks) - 1; i >= 0; i-- {
			if kd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = kd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, kd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (kd *KOGDelete) ExecX(ctx context.Context) int {
	n, err := kd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (kd *KOGDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: kog.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: kog.FieldID,
			},
		},
	}
	if ps := kd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, kd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// KOGDeleteOne is the builder for deleting a single KOG entity.
type KOGDeleteOne struct {
	kd *KOGDelete
}

// Exec executes the deletion query.
func (kdo *KOGDeleteOne) Exec(ctx context.Context) error {
	n, err := kdo.kd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{kog.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (kdo *KOGDeleteOne) ExecX(ctx context.Context) {
	kdo.kd.ExecX(ctx)
}
