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

package todo

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"entgo.io/contrib/entgql/internal/todo/ent"
)

func (r *categoryResolver) Todos(ctx context.Context, obj *ent.Category, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy []*ent.TodoOrder, where *ent.TodoWhereInput) (*ent.TodoConnection, error) {
	return r.client.Category.QueryTodos(obj).
		Paginate(ctx, after, first, before, last,
			ent.WithTodoOrder(orderBy...),
			ent.WithTodoFilter(where.Filter),
		)
}

func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	return r.client.Noder(ctx, id)
}

func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids)
}

func (r *queryResolver) BillProducts(ctx context.Context) ([]*ent.BillProduct, error) {
	return r.client.BillProduct.Query().All(ctx)
}

func (r *queryResolver) Groups(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.GroupWhereInput) (*ent.GroupConnection, error) {
	return r.client.Group.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithGroupFilter(where.Filter),
		)
}

func (r *queryResolver) Todos(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy []*ent.TodoOrder, where *ent.TodoWhereInput) (*ent.TodoConnection, error) {
	return r.client.Todo.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithTodoOrder(orderBy...),
			ent.WithTodoFilter(where.Filter),
		)
}

func (r *queryResolver) Users(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.UserWhereInput) (*ent.UserConnection, error) {
	return r.client.User.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithUserFilter(where.Filter),
		)
}

func (r *todoResolver) Children(ctx context.Context, obj *ent.Todo, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy []*ent.TodoOrder, where *ent.TodoWhereInput) (*ent.TodoConnection, error) {
	return r.client.Todo.QueryChildren(obj).
		Paginate(ctx, after, first, before, last,
			ent.WithTodoOrder(orderBy...),
			ent.WithTodoFilter(where.Filter),
		)
}

// Category returns CategoryResolver implementation.
func (r *Resolver) Category() CategoryResolver { return &categoryResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Todo returns TodoResolver implementation.
func (r *Resolver) Todo() TodoResolver { return &todoResolver{r} }

// TodoWhereInput returns TodoWhereInputResolver implementation.
func (r *Resolver) TodoWhereInput() TodoWhereInputResolver { return &todoWhereInputResolver{r} }

type categoryResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
type todoWhereInputResolver struct{ *Resolver }
