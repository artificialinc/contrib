// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/artificialinc/contrib/entgql/internal/todogotype/ent/pet"
	"github.com/artificialinc/contrib/entgql/internal/todogotype/ent/schema/uintgql"
	"entgo.io/ent/dialect/sql"
)

// Pet is the model entity for the Pet schema.
type Pet struct {
	config `json:"-"`
	// ID of the ent.
	ID uintgql.Uint64 `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Pet) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case pet.FieldID:
			values[i] = new(sql.NullInt64)
		case pet.FieldName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Pet", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Pet fields.
func (pe *Pet) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case pet.FieldID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				pe.ID = uintgql.Uint64(value.Int64)
			}
		case pet.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pe.Name = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Pet.
// Note that you need to call Pet.Unwrap() before calling this method if this Pet
// was returned from a transaction, and the transaction was committed or rolled back.
func (pe *Pet) Update() *PetUpdateOne {
	return (&PetClient{config: pe.config}).UpdateOne(pe)
}

// Unwrap unwraps the Pet entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pe *Pet) Unwrap() *Pet {
	_tx, ok := pe.config.driver.(*txDriver)
	if !ok {
		panic("ent: Pet is not a transactional entity")
	}
	pe.config.driver = _tx.drv
	return pe
}

// String implements the fmt.Stringer.
func (pe *Pet) String() string {
	var builder strings.Builder
	builder.WriteString("Pet(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pe.ID))
	builder.WriteString("name=")
	builder.WriteString(pe.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Pets is a parsable slice of Pet.
type Pets []*Pet

func (pe Pets) config(cfg config) {
	for _i := range pe {
		pe[_i].config = cfg
	}
}
