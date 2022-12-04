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

package verysecret

import (
	"github.com/artificialinc/contrib/entgql/internal/todopulid/ent/schema/pulid"
)

const (
	// Label holds the string label denoting the verysecret type in the database.
	Label = "very_secret"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// Table holds the table name of the verysecret in the database.
	Table = "very_secrets"
)

// Columns holds all SQL columns for verysecret fields.
var Columns = []string{
	FieldID,
	FieldPassword,
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
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() pulid.ID
)
