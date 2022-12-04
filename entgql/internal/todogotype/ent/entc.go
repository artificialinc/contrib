// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

//go:build ignore
// +build ignore

package main

import (
	"log"

	"github.com/artificialinc/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema/field"
)

func main() {
	ex, err := entgql.NewExtension(
		entgql.WithWhereInputs(true),
		// This option is disabled in this example,
		// because the schema file is edited by the
		// internal/todo/ent example.
		//
		// entgql.WithSchemaPath("../todo.graphql"),
		//
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}

	err = entc.Generate("./ent/schema", &gen.Config{
		IDType: &field.TypeInfo{Type: field.TypeString},
		Header: `
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
		`,
	}, entc.Extensions(ex))
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
