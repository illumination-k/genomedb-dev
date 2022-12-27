// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"genomedb/ent/domainannotation"
	"genomedb/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DomainAnnotationDelete is the builder for deleting a DomainAnnotation entity.
type DomainAnnotationDelete struct {
	config
	hooks    []Hook
	mutation *DomainAnnotationMutation
}

// Where appends a list predicates to the DomainAnnotationDelete builder.
func (dad *DomainAnnotationDelete) Where(ps ...predicate.DomainAnnotation) *DomainAnnotationDelete {
	dad.mutation.Where(ps...)
	return dad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dad *DomainAnnotationDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(dad.hooks) == 0 {
		affected, err = dad.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DomainAnnotationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dad.mutation = mutation
			affected, err = dad.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(dad.hooks) - 1; i >= 0; i-- {
			if dad.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dad.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dad.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (dad *DomainAnnotationDelete) ExecX(ctx context.Context) int {
	n, err := dad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dad *DomainAnnotationDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: domainannotation.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: domainannotation.FieldID,
			},
		},
	}
	if ps := dad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, dad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// DomainAnnotationDeleteOne is the builder for deleting a single DomainAnnotation entity.
type DomainAnnotationDeleteOne struct {
	dad *DomainAnnotationDelete
}

// Exec executes the deletion query.
func (dado *DomainAnnotationDeleteOne) Exec(ctx context.Context) error {
	n, err := dado.dad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{domainannotation.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (dado *DomainAnnotationDeleteOne) ExecX(ctx context.Context) {
	dado.dad.ExecX(ctx)
}