// Code generated by ent, DO NOT EDIT.

package oastypes

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/artificialinc/contrib/entoas/internal/oastypes/oastypes"
	"github.com/artificialinc/contrib/entoas/internal/oastypes/schema"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// OASTypesCreate is the builder for creating a OASTypes entity.
type OASTypesCreate struct {
	config
	mutation *OASTypesMutation
	hooks    []Hook
}

// SetInt sets the "int" field.
func (otc *OASTypesCreate) SetInt(i int) *OASTypesCreate {
	otc.mutation.SetInt(i)
	return otc
}

// SetInt8 sets the "int8" field.
func (otc *OASTypesCreate) SetInt8(i int8) *OASTypesCreate {
	otc.mutation.SetInt8(i)
	return otc
}

// SetInt16 sets the "int16" field.
func (otc *OASTypesCreate) SetInt16(i int16) *OASTypesCreate {
	otc.mutation.SetInt16(i)
	return otc
}

// SetInt32 sets the "int32" field.
func (otc *OASTypesCreate) SetInt32(i int32) *OASTypesCreate {
	otc.mutation.SetInt32(i)
	return otc
}

// SetInt64 sets the "int64" field.
func (otc *OASTypesCreate) SetInt64(i int64) *OASTypesCreate {
	otc.mutation.SetInt64(i)
	return otc
}

// SetUint sets the "uint" field.
func (otc *OASTypesCreate) SetUint(u uint) *OASTypesCreate {
	otc.mutation.SetUint(u)
	return otc
}

// SetUint8 sets the "uint8" field.
func (otc *OASTypesCreate) SetUint8(u uint8) *OASTypesCreate {
	otc.mutation.SetUint8(u)
	return otc
}

// SetUint16 sets the "uint16" field.
func (otc *OASTypesCreate) SetUint16(u uint16) *OASTypesCreate {
	otc.mutation.SetUint16(u)
	return otc
}

// SetUint32 sets the "uint32" field.
func (otc *OASTypesCreate) SetUint32(u uint32) *OASTypesCreate {
	otc.mutation.SetUint32(u)
	return otc
}

// SetUint64 sets the "uint64" field.
func (otc *OASTypesCreate) SetUint64(u uint64) *OASTypesCreate {
	otc.mutation.SetUint64(u)
	return otc
}

// SetFloat32 sets the "float32" field.
func (otc *OASTypesCreate) SetFloat32(f float32) *OASTypesCreate {
	otc.mutation.SetFloat32(f)
	return otc
}

// SetFloat64 sets the "float64" field.
func (otc *OASTypesCreate) SetFloat64(f float64) *OASTypesCreate {
	otc.mutation.SetFloat64(f)
	return otc
}

// SetStringField sets the "string_field" field.
func (otc *OASTypesCreate) SetStringField(s string) *OASTypesCreate {
	otc.mutation.SetStringField(s)
	return otc
}

// SetBool sets the "bool" field.
func (otc *OASTypesCreate) SetBool(b bool) *OASTypesCreate {
	otc.mutation.SetBool(b)
	return otc
}

// SetUUID sets the "uuid" field.
func (otc *OASTypesCreate) SetUUID(u uuid.UUID) *OASTypesCreate {
	otc.mutation.SetUUID(u)
	return otc
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (otc *OASTypesCreate) SetNillableUUID(u *uuid.UUID) *OASTypesCreate {
	if u != nil {
		otc.SetUUID(*u)
	}
	return otc
}

// SetTime sets the "time" field.
func (otc *OASTypesCreate) SetTime(t time.Time) *OASTypesCreate {
	otc.mutation.SetTime(t)
	return otc
}

// SetText sets the "text" field.
func (otc *OASTypesCreate) SetText(s string) *OASTypesCreate {
	otc.mutation.SetText(s)
	return otc
}

// SetState sets the "state" field.
func (otc *OASTypesCreate) SetState(o oastypes.State) *OASTypesCreate {
	otc.mutation.SetState(o)
	return otc
}

// SetStrings sets the "strings" field.
func (otc *OASTypesCreate) SetStrings(s []string) *OASTypesCreate {
	otc.mutation.SetStrings(s)
	return otc
}

// SetInts sets the "ints" field.
func (otc *OASTypesCreate) SetInts(i []int) *OASTypesCreate {
	otc.mutation.SetInts(i)
	return otc
}

// SetFloats sets the "floats" field.
func (otc *OASTypesCreate) SetFloats(f []float64) *OASTypesCreate {
	otc.mutation.SetFloats(f)
	return otc
}

// SetBytes sets the "bytes" field.
func (otc *OASTypesCreate) SetBytes(b []byte) *OASTypesCreate {
	otc.mutation.SetBytes(b)
	return otc
}

// SetNicknames sets the "nicknames" field.
func (otc *OASTypesCreate) SetNicknames(s []string) *OASTypesCreate {
	otc.mutation.SetNicknames(s)
	return otc
}

// SetJSONSlice sets the "json_slice" field.
func (otc *OASTypesCreate) SetJSONSlice(h []http.Dir) *OASTypesCreate {
	otc.mutation.SetJSONSlice(h)
	return otc
}

// SetJSONObj sets the "json_obj" field.
func (otc *OASTypesCreate) SetJSONObj(u url.URL) *OASTypesCreate {
	otc.mutation.SetJSONObj(u)
	return otc
}

// SetOther sets the "other" field.
func (otc *OASTypesCreate) SetOther(s *schema.Link) *OASTypesCreate {
	otc.mutation.SetOther(s)
	return otc
}

// Mutation returns the OASTypesMutation object of the builder.
func (otc *OASTypesCreate) Mutation() *OASTypesMutation {
	return otc.mutation
}

// Save creates the OASTypes in the database.
func (otc *OASTypesCreate) Save(ctx context.Context) (*OASTypes, error) {
	var (
		err  error
		node *OASTypes
	)
	otc.defaults()
	if len(otc.hooks) == 0 {
		if err = otc.check(); err != nil {
			return nil, err
		}
		node, err = otc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OASTypesMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = otc.check(); err != nil {
				return nil, err
			}
			otc.mutation = mutation
			if node, err = otc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(otc.hooks) - 1; i >= 0; i-- {
			if otc.hooks[i] == nil {
				return nil, fmt.Errorf("oastypes: uninitialized hook (forgotten import oastypes/runtime?)")
			}
			mut = otc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, otc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*OASTypes)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from OASTypesMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (otc *OASTypesCreate) SaveX(ctx context.Context) *OASTypes {
	v, err := otc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (otc *OASTypesCreate) Exec(ctx context.Context) error {
	_, err := otc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (otc *OASTypesCreate) ExecX(ctx context.Context) {
	if err := otc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (otc *OASTypesCreate) defaults() {
	if _, ok := otc.mutation.UUID(); !ok {
		v := oastypes.DefaultUUID()
		otc.mutation.SetUUID(v)
	}
	if _, ok := otc.mutation.Other(); !ok {
		v := oastypes.DefaultOther
		otc.mutation.SetOther(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (otc *OASTypesCreate) check() error {
	if _, ok := otc.mutation.Int(); !ok {
		return &ValidationError{Name: "int", err: errors.New(`oastypes: missing required field "OASTypes.int"`)}
	}
	if _, ok := otc.mutation.Int8(); !ok {
		return &ValidationError{Name: "int8", err: errors.New(`oastypes: missing required field "OASTypes.int8"`)}
	}
	if _, ok := otc.mutation.Int16(); !ok {
		return &ValidationError{Name: "int16", err: errors.New(`oastypes: missing required field "OASTypes.int16"`)}
	}
	if _, ok := otc.mutation.Int32(); !ok {
		return &ValidationError{Name: "int32", err: errors.New(`oastypes: missing required field "OASTypes.int32"`)}
	}
	if _, ok := otc.mutation.Int64(); !ok {
		return &ValidationError{Name: "int64", err: errors.New(`oastypes: missing required field "OASTypes.int64"`)}
	}
	if _, ok := otc.mutation.Uint(); !ok {
		return &ValidationError{Name: "uint", err: errors.New(`oastypes: missing required field "OASTypes.uint"`)}
	}
	if _, ok := otc.mutation.Uint8(); !ok {
		return &ValidationError{Name: "uint8", err: errors.New(`oastypes: missing required field "OASTypes.uint8"`)}
	}
	if _, ok := otc.mutation.Uint16(); !ok {
		return &ValidationError{Name: "uint16", err: errors.New(`oastypes: missing required field "OASTypes.uint16"`)}
	}
	if _, ok := otc.mutation.Uint32(); !ok {
		return &ValidationError{Name: "uint32", err: errors.New(`oastypes: missing required field "OASTypes.uint32"`)}
	}
	if _, ok := otc.mutation.Uint64(); !ok {
		return &ValidationError{Name: "uint64", err: errors.New(`oastypes: missing required field "OASTypes.uint64"`)}
	}
	if _, ok := otc.mutation.Float32(); !ok {
		return &ValidationError{Name: "float32", err: errors.New(`oastypes: missing required field "OASTypes.float32"`)}
	}
	if _, ok := otc.mutation.Float64(); !ok {
		return &ValidationError{Name: "float64", err: errors.New(`oastypes: missing required field "OASTypes.float64"`)}
	}
	if _, ok := otc.mutation.StringField(); !ok {
		return &ValidationError{Name: "string_field", err: errors.New(`oastypes: missing required field "OASTypes.string_field"`)}
	}
	if _, ok := otc.mutation.Bool(); !ok {
		return &ValidationError{Name: "bool", err: errors.New(`oastypes: missing required field "OASTypes.bool"`)}
	}
	if _, ok := otc.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New(`oastypes: missing required field "OASTypes.uuid"`)}
	}
	if _, ok := otc.mutation.Time(); !ok {
		return &ValidationError{Name: "time", err: errors.New(`oastypes: missing required field "OASTypes.time"`)}
	}
	if _, ok := otc.mutation.Text(); !ok {
		return &ValidationError{Name: "text", err: errors.New(`oastypes: missing required field "OASTypes.text"`)}
	}
	if _, ok := otc.mutation.State(); !ok {
		return &ValidationError{Name: "state", err: errors.New(`oastypes: missing required field "OASTypes.state"`)}
	}
	if v, ok := otc.mutation.State(); ok {
		if err := oastypes.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`oastypes: validator failed for field "OASTypes.state": %w`, err)}
		}
	}
	if _, ok := otc.mutation.Strings(); !ok {
		return &ValidationError{Name: "strings", err: errors.New(`oastypes: missing required field "OASTypes.strings"`)}
	}
	if _, ok := otc.mutation.Ints(); !ok {
		return &ValidationError{Name: "ints", err: errors.New(`oastypes: missing required field "OASTypes.ints"`)}
	}
	if _, ok := otc.mutation.Floats(); !ok {
		return &ValidationError{Name: "floats", err: errors.New(`oastypes: missing required field "OASTypes.floats"`)}
	}
	if _, ok := otc.mutation.Bytes(); !ok {
		return &ValidationError{Name: "bytes", err: errors.New(`oastypes: missing required field "OASTypes.bytes"`)}
	}
	if _, ok := otc.mutation.Nicknames(); !ok {
		return &ValidationError{Name: "nicknames", err: errors.New(`oastypes: missing required field "OASTypes.nicknames"`)}
	}
	if _, ok := otc.mutation.JSONSlice(); !ok {
		return &ValidationError{Name: "json_slice", err: errors.New(`oastypes: missing required field "OASTypes.json_slice"`)}
	}
	if _, ok := otc.mutation.JSONObj(); !ok {
		return &ValidationError{Name: "json_obj", err: errors.New(`oastypes: missing required field "OASTypes.json_obj"`)}
	}
	if _, ok := otc.mutation.Other(); !ok {
		return &ValidationError{Name: "other", err: errors.New(`oastypes: missing required field "OASTypes.other"`)}
	}
	return nil
}

func (otc *OASTypesCreate) sqlSave(ctx context.Context) (*OASTypes, error) {
	_node, _spec := otc.createSpec()
	if err := sqlgraph.CreateNode(ctx, otc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (otc *OASTypesCreate) createSpec() (*OASTypes, *sqlgraph.CreateSpec) {
	var (
		_node = &OASTypes{config: otc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: oastypes.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: oastypes.FieldID,
			},
		}
	)
	if value, ok := otc.mutation.Int(); ok {
		_spec.SetField(oastypes.FieldInt, field.TypeInt, value)
		_node.Int = value
	}
	if value, ok := otc.mutation.Int8(); ok {
		_spec.SetField(oastypes.FieldInt8, field.TypeInt8, value)
		_node.Int8 = value
	}
	if value, ok := otc.mutation.Int16(); ok {
		_spec.SetField(oastypes.FieldInt16, field.TypeInt16, value)
		_node.Int16 = value
	}
	if value, ok := otc.mutation.Int32(); ok {
		_spec.SetField(oastypes.FieldInt32, field.TypeInt32, value)
		_node.Int32 = value
	}
	if value, ok := otc.mutation.Int64(); ok {
		_spec.SetField(oastypes.FieldInt64, field.TypeInt64, value)
		_node.Int64 = value
	}
	if value, ok := otc.mutation.Uint(); ok {
		_spec.SetField(oastypes.FieldUint, field.TypeUint, value)
		_node.Uint = value
	}
	if value, ok := otc.mutation.Uint8(); ok {
		_spec.SetField(oastypes.FieldUint8, field.TypeUint8, value)
		_node.Uint8 = value
	}
	if value, ok := otc.mutation.Uint16(); ok {
		_spec.SetField(oastypes.FieldUint16, field.TypeUint16, value)
		_node.Uint16 = value
	}
	if value, ok := otc.mutation.Uint32(); ok {
		_spec.SetField(oastypes.FieldUint32, field.TypeUint32, value)
		_node.Uint32 = value
	}
	if value, ok := otc.mutation.Uint64(); ok {
		_spec.SetField(oastypes.FieldUint64, field.TypeUint64, value)
		_node.Uint64 = value
	}
	if value, ok := otc.mutation.Float32(); ok {
		_spec.SetField(oastypes.FieldFloat32, field.TypeFloat32, value)
		_node.Float32 = value
	}
	if value, ok := otc.mutation.Float64(); ok {
		_spec.SetField(oastypes.FieldFloat64, field.TypeFloat64, value)
		_node.Float64 = value
	}
	if value, ok := otc.mutation.StringField(); ok {
		_spec.SetField(oastypes.FieldStringField, field.TypeString, value)
		_node.StringField = value
	}
	if value, ok := otc.mutation.Bool(); ok {
		_spec.SetField(oastypes.FieldBool, field.TypeBool, value)
		_node.Bool = value
	}
	if value, ok := otc.mutation.UUID(); ok {
		_spec.SetField(oastypes.FieldUUID, field.TypeUUID, value)
		_node.UUID = value
	}
	if value, ok := otc.mutation.Time(); ok {
		_spec.SetField(oastypes.FieldTime, field.TypeTime, value)
		_node.Time = value
	}
	if value, ok := otc.mutation.Text(); ok {
		_spec.SetField(oastypes.FieldText, field.TypeString, value)
		_node.Text = value
	}
	if value, ok := otc.mutation.State(); ok {
		_spec.SetField(oastypes.FieldState, field.TypeEnum, value)
		_node.State = value
	}
	if value, ok := otc.mutation.Strings(); ok {
		_spec.SetField(oastypes.FieldStrings, field.TypeJSON, value)
		_node.Strings = value
	}
	if value, ok := otc.mutation.Ints(); ok {
		_spec.SetField(oastypes.FieldInts, field.TypeJSON, value)
		_node.Ints = value
	}
	if value, ok := otc.mutation.Floats(); ok {
		_spec.SetField(oastypes.FieldFloats, field.TypeJSON, value)
		_node.Floats = value
	}
	if value, ok := otc.mutation.Bytes(); ok {
		_spec.SetField(oastypes.FieldBytes, field.TypeBytes, value)
		_node.Bytes = value
	}
	if value, ok := otc.mutation.Nicknames(); ok {
		_spec.SetField(oastypes.FieldNicknames, field.TypeJSON, value)
		_node.Nicknames = value
	}
	if value, ok := otc.mutation.JSONSlice(); ok {
		_spec.SetField(oastypes.FieldJSONSlice, field.TypeJSON, value)
		_node.JSONSlice = value
	}
	if value, ok := otc.mutation.JSONObj(); ok {
		_spec.SetField(oastypes.FieldJSONObj, field.TypeJSON, value)
		_node.JSONObj = value
	}
	if value, ok := otc.mutation.Other(); ok {
		_spec.SetField(oastypes.FieldOther, field.TypeOther, value)
		_node.Other = value
	}
	return _node, _spec
}

// OASTypesCreateBulk is the builder for creating many OASTypes entities in bulk.
type OASTypesCreateBulk struct {
	config
	builders []*OASTypesCreate
}

// Save creates the OASTypes entities in the database.
func (otcb *OASTypesCreateBulk) Save(ctx context.Context) ([]*OASTypes, error) {
	specs := make([]*sqlgraph.CreateSpec, len(otcb.builders))
	nodes := make([]*OASTypes, len(otcb.builders))
	mutators := make([]Mutator, len(otcb.builders))
	for i := range otcb.builders {
		func(i int, root context.Context) {
			builder := otcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OASTypesMutation)
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
					_, err = mutators[i+1].Mutate(root, otcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, otcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, otcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (otcb *OASTypesCreateBulk) SaveX(ctx context.Context) []*OASTypes {
	v, err := otcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (otcb *OASTypesCreateBulk) Exec(ctx context.Context) error {
	_, err := otcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (otcb *OASTypesCreateBulk) ExecX(ctx context.Context) {
	if err := otcb.Exec(ctx); err != nil {
		panic(err)
	}
}
