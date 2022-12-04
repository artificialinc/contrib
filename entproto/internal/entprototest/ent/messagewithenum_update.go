// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/artificialinc/contrib/entproto/internal/entprototest/ent/messagewithenum"
	"github.com/artificialinc/contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MessageWithEnumUpdate is the builder for updating MessageWithEnum entities.
type MessageWithEnumUpdate struct {
	config
	hooks    []Hook
	mutation *MessageWithEnumMutation
}

// Where appends a list predicates to the MessageWithEnumUpdate builder.
func (mweu *MessageWithEnumUpdate) Where(ps ...predicate.MessageWithEnum) *MessageWithEnumUpdate {
	mweu.mutation.Where(ps...)
	return mweu
}

// SetEnumType sets the "enum_type" field.
func (mweu *MessageWithEnumUpdate) SetEnumType(mt messagewithenum.EnumType) *MessageWithEnumUpdate {
	mweu.mutation.SetEnumType(mt)
	return mweu
}

// SetNillableEnumType sets the "enum_type" field if the given value is not nil.
func (mweu *MessageWithEnumUpdate) SetNillableEnumType(mt *messagewithenum.EnumType) *MessageWithEnumUpdate {
	if mt != nil {
		mweu.SetEnumType(*mt)
	}
	return mweu
}

// SetEnumWithoutDefault sets the "enum_without_default" field.
func (mweu *MessageWithEnumUpdate) SetEnumWithoutDefault(mwd messagewithenum.EnumWithoutDefault) *MessageWithEnumUpdate {
	mweu.mutation.SetEnumWithoutDefault(mwd)
	return mweu
}

// Mutation returns the MessageWithEnumMutation object of the builder.
func (mweu *MessageWithEnumUpdate) Mutation() *MessageWithEnumMutation {
	return mweu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mweu *MessageWithEnumUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mweu.hooks) == 0 {
		if err = mweu.check(); err != nil {
			return 0, err
		}
		affected, err = mweu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageWithEnumMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mweu.check(); err != nil {
				return 0, err
			}
			mweu.mutation = mutation
			affected, err = mweu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mweu.hooks) - 1; i >= 0; i-- {
			if mweu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mweu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mweu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mweu *MessageWithEnumUpdate) SaveX(ctx context.Context) int {
	affected, err := mweu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mweu *MessageWithEnumUpdate) Exec(ctx context.Context) error {
	_, err := mweu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mweu *MessageWithEnumUpdate) ExecX(ctx context.Context) {
	if err := mweu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mweu *MessageWithEnumUpdate) check() error {
	if v, ok := mweu.mutation.EnumType(); ok {
		if err := messagewithenum.EnumTypeValidator(v); err != nil {
			return &ValidationError{Name: "enum_type", err: fmt.Errorf(`ent: validator failed for field "MessageWithEnum.enum_type": %w`, err)}
		}
	}
	if v, ok := mweu.mutation.EnumWithoutDefault(); ok {
		if err := messagewithenum.EnumWithoutDefaultValidator(v); err != nil {
			return &ValidationError{Name: "enum_without_default", err: fmt.Errorf(`ent: validator failed for field "MessageWithEnum.enum_without_default": %w`, err)}
		}
	}
	return nil
}

func (mweu *MessageWithEnumUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   messagewithenum.Table,
			Columns: messagewithenum.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: messagewithenum.FieldID,
			},
		},
	}
	if ps := mweu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mweu.mutation.EnumType(); ok {
		_spec.SetField(messagewithenum.FieldEnumType, field.TypeEnum, value)
	}
	if value, ok := mweu.mutation.EnumWithoutDefault(); ok {
		_spec.SetField(messagewithenum.FieldEnumWithoutDefault, field.TypeEnum, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mweu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{messagewithenum.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// MessageWithEnumUpdateOne is the builder for updating a single MessageWithEnum entity.
type MessageWithEnumUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MessageWithEnumMutation
}

// SetEnumType sets the "enum_type" field.
func (mweuo *MessageWithEnumUpdateOne) SetEnumType(mt messagewithenum.EnumType) *MessageWithEnumUpdateOne {
	mweuo.mutation.SetEnumType(mt)
	return mweuo
}

// SetNillableEnumType sets the "enum_type" field if the given value is not nil.
func (mweuo *MessageWithEnumUpdateOne) SetNillableEnumType(mt *messagewithenum.EnumType) *MessageWithEnumUpdateOne {
	if mt != nil {
		mweuo.SetEnumType(*mt)
	}
	return mweuo
}

// SetEnumWithoutDefault sets the "enum_without_default" field.
func (mweuo *MessageWithEnumUpdateOne) SetEnumWithoutDefault(mwd messagewithenum.EnumWithoutDefault) *MessageWithEnumUpdateOne {
	mweuo.mutation.SetEnumWithoutDefault(mwd)
	return mweuo
}

// Mutation returns the MessageWithEnumMutation object of the builder.
func (mweuo *MessageWithEnumUpdateOne) Mutation() *MessageWithEnumMutation {
	return mweuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mweuo *MessageWithEnumUpdateOne) Select(field string, fields ...string) *MessageWithEnumUpdateOne {
	mweuo.fields = append([]string{field}, fields...)
	return mweuo
}

// Save executes the query and returns the updated MessageWithEnum entity.
func (mweuo *MessageWithEnumUpdateOne) Save(ctx context.Context) (*MessageWithEnum, error) {
	var (
		err  error
		node *MessageWithEnum
	)
	if len(mweuo.hooks) == 0 {
		if err = mweuo.check(); err != nil {
			return nil, err
		}
		node, err = mweuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageWithEnumMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mweuo.check(); err != nil {
				return nil, err
			}
			mweuo.mutation = mutation
			node, err = mweuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(mweuo.hooks) - 1; i >= 0; i-- {
			if mweuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mweuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, mweuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*MessageWithEnum)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from MessageWithEnumMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (mweuo *MessageWithEnumUpdateOne) SaveX(ctx context.Context) *MessageWithEnum {
	node, err := mweuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mweuo *MessageWithEnumUpdateOne) Exec(ctx context.Context) error {
	_, err := mweuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mweuo *MessageWithEnumUpdateOne) ExecX(ctx context.Context) {
	if err := mweuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mweuo *MessageWithEnumUpdateOne) check() error {
	if v, ok := mweuo.mutation.EnumType(); ok {
		if err := messagewithenum.EnumTypeValidator(v); err != nil {
			return &ValidationError{Name: "enum_type", err: fmt.Errorf(`ent: validator failed for field "MessageWithEnum.enum_type": %w`, err)}
		}
	}
	if v, ok := mweuo.mutation.EnumWithoutDefault(); ok {
		if err := messagewithenum.EnumWithoutDefaultValidator(v); err != nil {
			return &ValidationError{Name: "enum_without_default", err: fmt.Errorf(`ent: validator failed for field "MessageWithEnum.enum_without_default": %w`, err)}
		}
	}
	return nil
}

func (mweuo *MessageWithEnumUpdateOne) sqlSave(ctx context.Context) (_node *MessageWithEnum, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   messagewithenum.Table,
			Columns: messagewithenum.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: messagewithenum.FieldID,
			},
		},
	}
	id, ok := mweuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MessageWithEnum.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mweuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, messagewithenum.FieldID)
		for _, f := range fields {
			if !messagewithenum.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != messagewithenum.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mweuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mweuo.mutation.EnumType(); ok {
		_spec.SetField(messagewithenum.FieldEnumType, field.TypeEnum, value)
	}
	if value, ok := mweuo.mutation.EnumWithoutDefault(); ok {
		_spec.SetField(messagewithenum.FieldEnumWithoutDefault, field.TypeEnum, value)
	}
	_node = &MessageWithEnum{config: mweuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mweuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{messagewithenum.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
