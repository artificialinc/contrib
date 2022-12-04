// Code generated by ent, DO NOT EDIT.

package oastypes

import (
	"fmt"

	"github.com/artificialinc/contrib/entoas/internal/oastypes/schema"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the oastypes type in the database.
	Label = "oas_types"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldInt holds the string denoting the int field in the database.
	FieldInt = "int"
	// FieldInt8 holds the string denoting the int8 field in the database.
	FieldInt8 = "int8"
	// FieldInt16 holds the string denoting the int16 field in the database.
	FieldInt16 = "int16"
	// FieldInt32 holds the string denoting the int32 field in the database.
	FieldInt32 = "int32"
	// FieldInt64 holds the string denoting the int64 field in the database.
	FieldInt64 = "int64"
	// FieldUint holds the string denoting the uint field in the database.
	FieldUint = "uint"
	// FieldUint8 holds the string denoting the uint8 field in the database.
	FieldUint8 = "uint8"
	// FieldUint16 holds the string denoting the uint16 field in the database.
	FieldUint16 = "uint16"
	// FieldUint32 holds the string denoting the uint32 field in the database.
	FieldUint32 = "uint32"
	// FieldUint64 holds the string denoting the uint64 field in the database.
	FieldUint64 = "uint64"
	// FieldFloat32 holds the string denoting the float32 field in the database.
	FieldFloat32 = "float32"
	// FieldFloat64 holds the string denoting the float64 field in the database.
	FieldFloat64 = "float64"
	// FieldStringField holds the string denoting the string_field field in the database.
	FieldStringField = "string_field"
	// FieldBool holds the string denoting the bool field in the database.
	FieldBool = "bool"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// FieldTime holds the string denoting the time field in the database.
	FieldTime = "time"
	// FieldText holds the string denoting the text field in the database.
	FieldText = "text"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldStrings holds the string denoting the strings field in the database.
	FieldStrings = "strings"
	// FieldInts holds the string denoting the ints field in the database.
	FieldInts = "ints"
	// FieldFloats holds the string denoting the floats field in the database.
	FieldFloats = "floats"
	// FieldBytes holds the string denoting the bytes field in the database.
	FieldBytes = "bytes"
	// FieldNicknames holds the string denoting the nicknames field in the database.
	FieldNicknames = "nicknames"
	// FieldJSONSlice holds the string denoting the json_slice field in the database.
	FieldJSONSlice = "json_slice"
	// FieldJSONObj holds the string denoting the json_obj field in the database.
	FieldJSONObj = "json_obj"
	// FieldOther holds the string denoting the other field in the database.
	FieldOther = "other"
	// Table holds the table name of the oastypes in the database.
	Table = "oas_types"
)

// Columns holds all SQL columns for oastypes fields.
var Columns = []string{
	FieldID,
	FieldInt,
	FieldInt8,
	FieldInt16,
	FieldInt32,
	FieldInt64,
	FieldUint,
	FieldUint8,
	FieldUint16,
	FieldUint32,
	FieldUint64,
	FieldFloat32,
	FieldFloat64,
	FieldStringField,
	FieldBool,
	FieldUUID,
	FieldTime,
	FieldText,
	FieldState,
	FieldStrings,
	FieldInts,
	FieldFloats,
	FieldBytes,
	FieldNicknames,
	FieldJSONSlice,
	FieldJSONObj,
	FieldOther,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultUUID holds the default value on creation for the "uuid" field.
	DefaultUUID func() uuid.UUID
	// DefaultOther holds the default value on creation for the "other" field.
	DefaultOther *schema.Link
)

// State defines the type for the "state" enum field.
type State string

// State values.
const (
	StateOn  State = "on"
	StateOff State = "off"
)

func (s State) String() string {
	return string(s)
}

// StateValidator is a validator for the "state" field enum values. It is called by the builders before save.
func StateValidator(s State) error {
	switch s {
	case StateOn, StateOff:
		return nil
	default:
		return fmt.Errorf("oastypes: invalid enum value for state field: %q", s)
	}
}
