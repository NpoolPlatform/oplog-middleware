// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/oplog-middleware/pkg/db/ent/oplog"
	"github.com/NpoolPlatform/oplog-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// OpLogUpdate is the builder for updating OpLog entities.
type OpLogUpdate struct {
	config
	hooks     []Hook
	mutation  *OpLogMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the OpLogUpdate builder.
func (olu *OpLogUpdate) Where(ps ...predicate.OpLog) *OpLogUpdate {
	olu.mutation.Where(ps...)
	return olu
}

// SetCreatedAt sets the "created_at" field.
func (olu *OpLogUpdate) SetCreatedAt(u uint32) *OpLogUpdate {
	olu.mutation.ResetCreatedAt()
	olu.mutation.SetCreatedAt(u)
	return olu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (olu *OpLogUpdate) SetNillableCreatedAt(u *uint32) *OpLogUpdate {
	if u != nil {
		olu.SetCreatedAt(*u)
	}
	return olu
}

// AddCreatedAt adds u to the "created_at" field.
func (olu *OpLogUpdate) AddCreatedAt(u int32) *OpLogUpdate {
	olu.mutation.AddCreatedAt(u)
	return olu
}

// SetUpdatedAt sets the "updated_at" field.
func (olu *OpLogUpdate) SetUpdatedAt(u uint32) *OpLogUpdate {
	olu.mutation.ResetUpdatedAt()
	olu.mutation.SetUpdatedAt(u)
	return olu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (olu *OpLogUpdate) AddUpdatedAt(u int32) *OpLogUpdate {
	olu.mutation.AddUpdatedAt(u)
	return olu
}

// SetDeletedAt sets the "deleted_at" field.
func (olu *OpLogUpdate) SetDeletedAt(u uint32) *OpLogUpdate {
	olu.mutation.ResetDeletedAt()
	olu.mutation.SetDeletedAt(u)
	return olu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (olu *OpLogUpdate) SetNillableDeletedAt(u *uint32) *OpLogUpdate {
	if u != nil {
		olu.SetDeletedAt(*u)
	}
	return olu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (olu *OpLogUpdate) AddDeletedAt(u int32) *OpLogUpdate {
	olu.mutation.AddDeletedAt(u)
	return olu
}

// SetAutoID sets the "auto_id" field.
func (olu *OpLogUpdate) SetAutoID(u uint32) *OpLogUpdate {
	olu.mutation.ResetAutoID()
	olu.mutation.SetAutoID(u)
	return olu
}

// AddAutoID adds u to the "auto_id" field.
func (olu *OpLogUpdate) AddAutoID(u int32) *OpLogUpdate {
	olu.mutation.AddAutoID(u)
	return olu
}

// SetAppID sets the "app_id" field.
func (olu *OpLogUpdate) SetAppID(u uuid.UUID) *OpLogUpdate {
	olu.mutation.SetAppID(u)
	return olu
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (olu *OpLogUpdate) SetNillableAppID(u *uuid.UUID) *OpLogUpdate {
	if u != nil {
		olu.SetAppID(*u)
	}
	return olu
}

// ClearAppID clears the value of the "app_id" field.
func (olu *OpLogUpdate) ClearAppID() *OpLogUpdate {
	olu.mutation.ClearAppID()
	return olu
}

// SetUserID sets the "user_id" field.
func (olu *OpLogUpdate) SetUserID(u uuid.UUID) *OpLogUpdate {
	olu.mutation.SetUserID(u)
	return olu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (olu *OpLogUpdate) SetNillableUserID(u *uuid.UUID) *OpLogUpdate {
	if u != nil {
		olu.SetUserID(*u)
	}
	return olu
}

// ClearUserID clears the value of the "user_id" field.
func (olu *OpLogUpdate) ClearUserID() *OpLogUpdate {
	olu.mutation.ClearUserID()
	return olu
}

// SetPath sets the "path" field.
func (olu *OpLogUpdate) SetPath(s string) *OpLogUpdate {
	olu.mutation.SetPath(s)
	return olu
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (olu *OpLogUpdate) SetNillablePath(s *string) *OpLogUpdate {
	if s != nil {
		olu.SetPath(*s)
	}
	return olu
}

// ClearPath clears the value of the "path" field.
func (olu *OpLogUpdate) ClearPath() *OpLogUpdate {
	olu.mutation.ClearPath()
	return olu
}

// SetMethod sets the "method" field.
func (olu *OpLogUpdate) SetMethod(s string) *OpLogUpdate {
	olu.mutation.SetMethod(s)
	return olu
}

// SetNillableMethod sets the "method" field if the given value is not nil.
func (olu *OpLogUpdate) SetNillableMethod(s *string) *OpLogUpdate {
	if s != nil {
		olu.SetMethod(*s)
	}
	return olu
}

// ClearMethod clears the value of the "method" field.
func (olu *OpLogUpdate) ClearMethod() *OpLogUpdate {
	olu.mutation.ClearMethod()
	return olu
}

// SetArguments sets the "arguments" field.
func (olu *OpLogUpdate) SetArguments(s string) *OpLogUpdate {
	olu.mutation.SetArguments(s)
	return olu
}

// SetNillableArguments sets the "arguments" field if the given value is not nil.
func (olu *OpLogUpdate) SetNillableArguments(s *string) *OpLogUpdate {
	if s != nil {
		olu.SetArguments(*s)
	}
	return olu
}

// ClearArguments clears the value of the "arguments" field.
func (olu *OpLogUpdate) ClearArguments() *OpLogUpdate {
	olu.mutation.ClearArguments()
	return olu
}

// SetCurValue sets the "cur_value" field.
func (olu *OpLogUpdate) SetCurValue(s string) *OpLogUpdate {
	olu.mutation.SetCurValue(s)
	return olu
}

// SetNillableCurValue sets the "cur_value" field if the given value is not nil.
func (olu *OpLogUpdate) SetNillableCurValue(s *string) *OpLogUpdate {
	if s != nil {
		olu.SetCurValue(*s)
	}
	return olu
}

// ClearCurValue clears the value of the "cur_value" field.
func (olu *OpLogUpdate) ClearCurValue() *OpLogUpdate {
	olu.mutation.ClearCurValue()
	return olu
}

// SetHumanReadable sets the "human_readable" field.
func (olu *OpLogUpdate) SetHumanReadable(s string) *OpLogUpdate {
	olu.mutation.SetHumanReadable(s)
	return olu
}

// SetNillableHumanReadable sets the "human_readable" field if the given value is not nil.
func (olu *OpLogUpdate) SetNillableHumanReadable(s *string) *OpLogUpdate {
	if s != nil {
		olu.SetHumanReadable(*s)
	}
	return olu
}

// ClearHumanReadable clears the value of the "human_readable" field.
func (olu *OpLogUpdate) ClearHumanReadable() *OpLogUpdate {
	olu.mutation.ClearHumanReadable()
	return olu
}

// SetResult sets the "result" field.
func (olu *OpLogUpdate) SetResult(s string) *OpLogUpdate {
	olu.mutation.SetResult(s)
	return olu
}

// SetNillableResult sets the "result" field if the given value is not nil.
func (olu *OpLogUpdate) SetNillableResult(s *string) *OpLogUpdate {
	if s != nil {
		olu.SetResult(*s)
	}
	return olu
}

// ClearResult clears the value of the "result" field.
func (olu *OpLogUpdate) ClearResult() *OpLogUpdate {
	olu.mutation.ClearResult()
	return olu
}

// SetFailReason sets the "fail_reason" field.
func (olu *OpLogUpdate) SetFailReason(s string) *OpLogUpdate {
	olu.mutation.SetFailReason(s)
	return olu
}

// SetNillableFailReason sets the "fail_reason" field if the given value is not nil.
func (olu *OpLogUpdate) SetNillableFailReason(s *string) *OpLogUpdate {
	if s != nil {
		olu.SetFailReason(*s)
	}
	return olu
}

// ClearFailReason clears the value of the "fail_reason" field.
func (olu *OpLogUpdate) ClearFailReason() *OpLogUpdate {
	olu.mutation.ClearFailReason()
	return olu
}

// SetElapsedMillisecs sets the "elapsed_millisecs" field.
func (olu *OpLogUpdate) SetElapsedMillisecs(u uint32) *OpLogUpdate {
	olu.mutation.ResetElapsedMillisecs()
	olu.mutation.SetElapsedMillisecs(u)
	return olu
}

// SetNillableElapsedMillisecs sets the "elapsed_millisecs" field if the given value is not nil.
func (olu *OpLogUpdate) SetNillableElapsedMillisecs(u *uint32) *OpLogUpdate {
	if u != nil {
		olu.SetElapsedMillisecs(*u)
	}
	return olu
}

// AddElapsedMillisecs adds u to the "elapsed_millisecs" field.
func (olu *OpLogUpdate) AddElapsedMillisecs(u int32) *OpLogUpdate {
	olu.mutation.AddElapsedMillisecs(u)
	return olu
}

// ClearElapsedMillisecs clears the value of the "elapsed_millisecs" field.
func (olu *OpLogUpdate) ClearElapsedMillisecs() *OpLogUpdate {
	olu.mutation.ClearElapsedMillisecs()
	return olu
}

// Mutation returns the OpLogMutation object of the builder.
func (olu *OpLogUpdate) Mutation() *OpLogMutation {
	return olu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (olu *OpLogUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := olu.defaults(); err != nil {
		return 0, err
	}
	if len(olu.hooks) == 0 {
		affected, err = olu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OpLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			olu.mutation = mutation
			affected, err = olu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(olu.hooks) - 1; i >= 0; i-- {
			if olu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = olu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, olu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (olu *OpLogUpdate) SaveX(ctx context.Context) int {
	affected, err := olu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (olu *OpLogUpdate) Exec(ctx context.Context) error {
	_, err := olu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (olu *OpLogUpdate) ExecX(ctx context.Context) {
	if err := olu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (olu *OpLogUpdate) defaults() error {
	if _, ok := olu.mutation.UpdatedAt(); !ok {
		if oplog.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized oplog.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := oplog.UpdateDefaultUpdatedAt()
		olu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (olu *OpLogUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *OpLogUpdate {
	olu.modifiers = append(olu.modifiers, modifiers...)
	return olu
}

func (olu *OpLogUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   oplog.Table,
			Columns: oplog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: oplog.FieldID,
			},
		},
	}
	if ps := olu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := olu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: oplog.FieldCreatedAt,
		})
	}
	if value, ok := olu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: oplog.FieldCreatedAt,
		})
	}
	if value, ok := olu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: oplog.FieldUpdatedAt,
		})
	}
	if value, ok := olu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: oplog.FieldUpdatedAt,
		})
	}
	if value, ok := olu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: oplog.FieldDeletedAt,
		})
	}
	if value, ok := olu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: oplog.FieldDeletedAt,
		})
	}
	if value, ok := olu.mutation.AutoID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: oplog.FieldAutoID,
		})
	}
	if value, ok := olu.mutation.AddedAutoID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: oplog.FieldAutoID,
		})
	}
	if value, ok := olu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: oplog.FieldAppID,
		})
	}
	if olu.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: oplog.FieldAppID,
		})
	}
	if value, ok := olu.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: oplog.FieldUserID,
		})
	}
	if olu.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: oplog.FieldUserID,
		})
	}
	if value, ok := olu.mutation.Path(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: oplog.FieldPath,
		})
	}
	if olu.mutation.PathCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: oplog.FieldPath,
		})
	}
	if value, ok := olu.mutation.Method(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: oplog.FieldMethod,
		})
	}
	if olu.mutation.MethodCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: oplog.FieldMethod,
		})
	}
	if value, ok := olu.mutation.Arguments(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: oplog.FieldArguments,
		})
	}
	if olu.mutation.ArgumentsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: oplog.FieldArguments,
		})
	}
	if value, ok := olu.mutation.CurValue(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: oplog.FieldCurValue,
		})
	}
	if olu.mutation.CurValueCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: oplog.FieldCurValue,
		})
	}
	if value, ok := olu.mutation.HumanReadable(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: oplog.FieldHumanReadable,
		})
	}
	if olu.mutation.HumanReadableCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: oplog.FieldHumanReadable,
		})
	}
	if value, ok := olu.mutation.Result(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: oplog.FieldResult,
		})
	}
	if olu.mutation.ResultCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: oplog.FieldResult,
		})
	}
	if value, ok := olu.mutation.FailReason(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: oplog.FieldFailReason,
		})
	}
	if olu.mutation.FailReasonCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: oplog.FieldFailReason,
		})
	}
	if value, ok := olu.mutation.ElapsedMillisecs(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: oplog.FieldElapsedMillisecs,
		})
	}
	if value, ok := olu.mutation.AddedElapsedMillisecs(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: oplog.FieldElapsedMillisecs,
		})
	}
	if olu.mutation.ElapsedMillisecsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: oplog.FieldElapsedMillisecs,
		})
	}
	_spec.Modifiers = olu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, olu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{oplog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// OpLogUpdateOne is the builder for updating a single OpLog entity.
type OpLogUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *OpLogMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (oluo *OpLogUpdateOne) SetCreatedAt(u uint32) *OpLogUpdateOne {
	oluo.mutation.ResetCreatedAt()
	oluo.mutation.SetCreatedAt(u)
	return oluo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (oluo *OpLogUpdateOne) SetNillableCreatedAt(u *uint32) *OpLogUpdateOne {
	if u != nil {
		oluo.SetCreatedAt(*u)
	}
	return oluo
}

// AddCreatedAt adds u to the "created_at" field.
func (oluo *OpLogUpdateOne) AddCreatedAt(u int32) *OpLogUpdateOne {
	oluo.mutation.AddCreatedAt(u)
	return oluo
}

// SetUpdatedAt sets the "updated_at" field.
func (oluo *OpLogUpdateOne) SetUpdatedAt(u uint32) *OpLogUpdateOne {
	oluo.mutation.ResetUpdatedAt()
	oluo.mutation.SetUpdatedAt(u)
	return oluo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (oluo *OpLogUpdateOne) AddUpdatedAt(u int32) *OpLogUpdateOne {
	oluo.mutation.AddUpdatedAt(u)
	return oluo
}

// SetDeletedAt sets the "deleted_at" field.
func (oluo *OpLogUpdateOne) SetDeletedAt(u uint32) *OpLogUpdateOne {
	oluo.mutation.ResetDeletedAt()
	oluo.mutation.SetDeletedAt(u)
	return oluo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (oluo *OpLogUpdateOne) SetNillableDeletedAt(u *uint32) *OpLogUpdateOne {
	if u != nil {
		oluo.SetDeletedAt(*u)
	}
	return oluo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (oluo *OpLogUpdateOne) AddDeletedAt(u int32) *OpLogUpdateOne {
	oluo.mutation.AddDeletedAt(u)
	return oluo
}

// SetAutoID sets the "auto_id" field.
func (oluo *OpLogUpdateOne) SetAutoID(u uint32) *OpLogUpdateOne {
	oluo.mutation.ResetAutoID()
	oluo.mutation.SetAutoID(u)
	return oluo
}

// AddAutoID adds u to the "auto_id" field.
func (oluo *OpLogUpdateOne) AddAutoID(u int32) *OpLogUpdateOne {
	oluo.mutation.AddAutoID(u)
	return oluo
}

// SetAppID sets the "app_id" field.
func (oluo *OpLogUpdateOne) SetAppID(u uuid.UUID) *OpLogUpdateOne {
	oluo.mutation.SetAppID(u)
	return oluo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (oluo *OpLogUpdateOne) SetNillableAppID(u *uuid.UUID) *OpLogUpdateOne {
	if u != nil {
		oluo.SetAppID(*u)
	}
	return oluo
}

// ClearAppID clears the value of the "app_id" field.
func (oluo *OpLogUpdateOne) ClearAppID() *OpLogUpdateOne {
	oluo.mutation.ClearAppID()
	return oluo
}

// SetUserID sets the "user_id" field.
func (oluo *OpLogUpdateOne) SetUserID(u uuid.UUID) *OpLogUpdateOne {
	oluo.mutation.SetUserID(u)
	return oluo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (oluo *OpLogUpdateOne) SetNillableUserID(u *uuid.UUID) *OpLogUpdateOne {
	if u != nil {
		oluo.SetUserID(*u)
	}
	return oluo
}

// ClearUserID clears the value of the "user_id" field.
func (oluo *OpLogUpdateOne) ClearUserID() *OpLogUpdateOne {
	oluo.mutation.ClearUserID()
	return oluo
}

// SetPath sets the "path" field.
func (oluo *OpLogUpdateOne) SetPath(s string) *OpLogUpdateOne {
	oluo.mutation.SetPath(s)
	return oluo
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (oluo *OpLogUpdateOne) SetNillablePath(s *string) *OpLogUpdateOne {
	if s != nil {
		oluo.SetPath(*s)
	}
	return oluo
}

// ClearPath clears the value of the "path" field.
func (oluo *OpLogUpdateOne) ClearPath() *OpLogUpdateOne {
	oluo.mutation.ClearPath()
	return oluo
}

// SetMethod sets the "method" field.
func (oluo *OpLogUpdateOne) SetMethod(s string) *OpLogUpdateOne {
	oluo.mutation.SetMethod(s)
	return oluo
}

// SetNillableMethod sets the "method" field if the given value is not nil.
func (oluo *OpLogUpdateOne) SetNillableMethod(s *string) *OpLogUpdateOne {
	if s != nil {
		oluo.SetMethod(*s)
	}
	return oluo
}

// ClearMethod clears the value of the "method" field.
func (oluo *OpLogUpdateOne) ClearMethod() *OpLogUpdateOne {
	oluo.mutation.ClearMethod()
	return oluo
}

// SetArguments sets the "arguments" field.
func (oluo *OpLogUpdateOne) SetArguments(s string) *OpLogUpdateOne {
	oluo.mutation.SetArguments(s)
	return oluo
}

// SetNillableArguments sets the "arguments" field if the given value is not nil.
func (oluo *OpLogUpdateOne) SetNillableArguments(s *string) *OpLogUpdateOne {
	if s != nil {
		oluo.SetArguments(*s)
	}
	return oluo
}

// ClearArguments clears the value of the "arguments" field.
func (oluo *OpLogUpdateOne) ClearArguments() *OpLogUpdateOne {
	oluo.mutation.ClearArguments()
	return oluo
}

// SetCurValue sets the "cur_value" field.
func (oluo *OpLogUpdateOne) SetCurValue(s string) *OpLogUpdateOne {
	oluo.mutation.SetCurValue(s)
	return oluo
}

// SetNillableCurValue sets the "cur_value" field if the given value is not nil.
func (oluo *OpLogUpdateOne) SetNillableCurValue(s *string) *OpLogUpdateOne {
	if s != nil {
		oluo.SetCurValue(*s)
	}
	return oluo
}

// ClearCurValue clears the value of the "cur_value" field.
func (oluo *OpLogUpdateOne) ClearCurValue() *OpLogUpdateOne {
	oluo.mutation.ClearCurValue()
	return oluo
}

// SetHumanReadable sets the "human_readable" field.
func (oluo *OpLogUpdateOne) SetHumanReadable(s string) *OpLogUpdateOne {
	oluo.mutation.SetHumanReadable(s)
	return oluo
}

// SetNillableHumanReadable sets the "human_readable" field if the given value is not nil.
func (oluo *OpLogUpdateOne) SetNillableHumanReadable(s *string) *OpLogUpdateOne {
	if s != nil {
		oluo.SetHumanReadable(*s)
	}
	return oluo
}

// ClearHumanReadable clears the value of the "human_readable" field.
func (oluo *OpLogUpdateOne) ClearHumanReadable() *OpLogUpdateOne {
	oluo.mutation.ClearHumanReadable()
	return oluo
}

// SetResult sets the "result" field.
func (oluo *OpLogUpdateOne) SetResult(s string) *OpLogUpdateOne {
	oluo.mutation.SetResult(s)
	return oluo
}

// SetNillableResult sets the "result" field if the given value is not nil.
func (oluo *OpLogUpdateOne) SetNillableResult(s *string) *OpLogUpdateOne {
	if s != nil {
		oluo.SetResult(*s)
	}
	return oluo
}

// ClearResult clears the value of the "result" field.
func (oluo *OpLogUpdateOne) ClearResult() *OpLogUpdateOne {
	oluo.mutation.ClearResult()
	return oluo
}

// SetFailReason sets the "fail_reason" field.
func (oluo *OpLogUpdateOne) SetFailReason(s string) *OpLogUpdateOne {
	oluo.mutation.SetFailReason(s)
	return oluo
}

// SetNillableFailReason sets the "fail_reason" field if the given value is not nil.
func (oluo *OpLogUpdateOne) SetNillableFailReason(s *string) *OpLogUpdateOne {
	if s != nil {
		oluo.SetFailReason(*s)
	}
	return oluo
}

// ClearFailReason clears the value of the "fail_reason" field.
func (oluo *OpLogUpdateOne) ClearFailReason() *OpLogUpdateOne {
	oluo.mutation.ClearFailReason()
	return oluo
}

// SetElapsedMillisecs sets the "elapsed_millisecs" field.
func (oluo *OpLogUpdateOne) SetElapsedMillisecs(u uint32) *OpLogUpdateOne {
	oluo.mutation.ResetElapsedMillisecs()
	oluo.mutation.SetElapsedMillisecs(u)
	return oluo
}

// SetNillableElapsedMillisecs sets the "elapsed_millisecs" field if the given value is not nil.
func (oluo *OpLogUpdateOne) SetNillableElapsedMillisecs(u *uint32) *OpLogUpdateOne {
	if u != nil {
		oluo.SetElapsedMillisecs(*u)
	}
	return oluo
}

// AddElapsedMillisecs adds u to the "elapsed_millisecs" field.
func (oluo *OpLogUpdateOne) AddElapsedMillisecs(u int32) *OpLogUpdateOne {
	oluo.mutation.AddElapsedMillisecs(u)
	return oluo
}

// ClearElapsedMillisecs clears the value of the "elapsed_millisecs" field.
func (oluo *OpLogUpdateOne) ClearElapsedMillisecs() *OpLogUpdateOne {
	oluo.mutation.ClearElapsedMillisecs()
	return oluo
}

// Mutation returns the OpLogMutation object of the builder.
func (oluo *OpLogUpdateOne) Mutation() *OpLogMutation {
	return oluo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (oluo *OpLogUpdateOne) Select(field string, fields ...string) *OpLogUpdateOne {
	oluo.fields = append([]string{field}, fields...)
	return oluo
}

// Save executes the query and returns the updated OpLog entity.
func (oluo *OpLogUpdateOne) Save(ctx context.Context) (*OpLog, error) {
	var (
		err  error
		node *OpLog
	)
	if err := oluo.defaults(); err != nil {
		return nil, err
	}
	if len(oluo.hooks) == 0 {
		node, err = oluo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OpLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			oluo.mutation = mutation
			node, err = oluo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(oluo.hooks) - 1; i >= 0; i-- {
			if oluo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = oluo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, oluo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*OpLog)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from OpLogMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (oluo *OpLogUpdateOne) SaveX(ctx context.Context) *OpLog {
	node, err := oluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (oluo *OpLogUpdateOne) Exec(ctx context.Context) error {
	_, err := oluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oluo *OpLogUpdateOne) ExecX(ctx context.Context) {
	if err := oluo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oluo *OpLogUpdateOne) defaults() error {
	if _, ok := oluo.mutation.UpdatedAt(); !ok {
		if oplog.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized oplog.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := oplog.UpdateDefaultUpdatedAt()
		oluo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (oluo *OpLogUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *OpLogUpdateOne {
	oluo.modifiers = append(oluo.modifiers, modifiers...)
	return oluo
}

func (oluo *OpLogUpdateOne) sqlSave(ctx context.Context) (_node *OpLog, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   oplog.Table,
			Columns: oplog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: oplog.FieldID,
			},
		},
	}
	id, ok := oluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OpLog.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := oluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, oplog.FieldID)
		for _, f := range fields {
			if !oplog.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != oplog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := oluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oluo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: oplog.FieldCreatedAt,
		})
	}
	if value, ok := oluo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: oplog.FieldCreatedAt,
		})
	}
	if value, ok := oluo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: oplog.FieldUpdatedAt,
		})
	}
	if value, ok := oluo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: oplog.FieldUpdatedAt,
		})
	}
	if value, ok := oluo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: oplog.FieldDeletedAt,
		})
	}
	if value, ok := oluo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: oplog.FieldDeletedAt,
		})
	}
	if value, ok := oluo.mutation.AutoID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: oplog.FieldAutoID,
		})
	}
	if value, ok := oluo.mutation.AddedAutoID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: oplog.FieldAutoID,
		})
	}
	if value, ok := oluo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: oplog.FieldAppID,
		})
	}
	if oluo.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: oplog.FieldAppID,
		})
	}
	if value, ok := oluo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: oplog.FieldUserID,
		})
	}
	if oluo.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: oplog.FieldUserID,
		})
	}
	if value, ok := oluo.mutation.Path(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: oplog.FieldPath,
		})
	}
	if oluo.mutation.PathCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: oplog.FieldPath,
		})
	}
	if value, ok := oluo.mutation.Method(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: oplog.FieldMethod,
		})
	}
	if oluo.mutation.MethodCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: oplog.FieldMethod,
		})
	}
	if value, ok := oluo.mutation.Arguments(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: oplog.FieldArguments,
		})
	}
	if oluo.mutation.ArgumentsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: oplog.FieldArguments,
		})
	}
	if value, ok := oluo.mutation.CurValue(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: oplog.FieldCurValue,
		})
	}
	if oluo.mutation.CurValueCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: oplog.FieldCurValue,
		})
	}
	if value, ok := oluo.mutation.HumanReadable(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: oplog.FieldHumanReadable,
		})
	}
	if oluo.mutation.HumanReadableCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: oplog.FieldHumanReadable,
		})
	}
	if value, ok := oluo.mutation.Result(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: oplog.FieldResult,
		})
	}
	if oluo.mutation.ResultCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: oplog.FieldResult,
		})
	}
	if value, ok := oluo.mutation.FailReason(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: oplog.FieldFailReason,
		})
	}
	if oluo.mutation.FailReasonCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: oplog.FieldFailReason,
		})
	}
	if value, ok := oluo.mutation.ElapsedMillisecs(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: oplog.FieldElapsedMillisecs,
		})
	}
	if value, ok := oluo.mutation.AddedElapsedMillisecs(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: oplog.FieldElapsedMillisecs,
		})
	}
	if oluo.mutation.ElapsedMillisecsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: oplog.FieldElapsedMillisecs,
		})
	}
	_spec.Modifiers = oluo.modifiers
	_node = &OpLog{config: oluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, oluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{oplog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
