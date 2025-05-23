// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lightsparkdev/spark/so/ent/predicate"
	"github.com/lightsparkdev/spark/so/ent/tokentransaction"
)

// TokenTransactionDelete is the builder for deleting a TokenTransaction entity.
type TokenTransactionDelete struct {
	config
	hooks    []Hook
	mutation *TokenTransactionMutation
}

// Where appends a list predicates to the TokenTransactionDelete builder.
func (ttd *TokenTransactionDelete) Where(ps ...predicate.TokenTransaction) *TokenTransactionDelete {
	ttd.mutation.Where(ps...)
	return ttd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ttd *TokenTransactionDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ttd.sqlExec, ttd.mutation, ttd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ttd *TokenTransactionDelete) ExecX(ctx context.Context) int {
	n, err := ttd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ttd *TokenTransactionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(tokentransaction.Table, sqlgraph.NewFieldSpec(tokentransaction.FieldID, field.TypeUUID))
	if ps := ttd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ttd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ttd.mutation.done = true
	return affected, err
}

// TokenTransactionDeleteOne is the builder for deleting a single TokenTransaction entity.
type TokenTransactionDeleteOne struct {
	ttd *TokenTransactionDelete
}

// Where appends a list predicates to the TokenTransactionDelete builder.
func (ttdo *TokenTransactionDeleteOne) Where(ps ...predicate.TokenTransaction) *TokenTransactionDeleteOne {
	ttdo.ttd.mutation.Where(ps...)
	return ttdo
}

// Exec executes the deletion query.
func (ttdo *TokenTransactionDeleteOne) Exec(ctx context.Context) error {
	n, err := ttdo.ttd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{tokentransaction.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ttdo *TokenTransactionDeleteOne) ExecX(ctx context.Context) {
	if err := ttdo.Exec(ctx); err != nil {
		panic(err)
	}
}
