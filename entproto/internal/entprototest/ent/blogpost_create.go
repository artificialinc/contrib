// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/artificialinc/contrib/entproto/internal/entprototest/ent/blogpost"
	"github.com/artificialinc/contrib/entproto/internal/entprototest/ent/category"
	"github.com/artificialinc/contrib/entproto/internal/entprototest/ent/user"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BlogPostCreate is the builder for creating a BlogPost entity.
type BlogPostCreate struct {
	config
	mutation *BlogPostMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (bpc *BlogPostCreate) SetTitle(s string) *BlogPostCreate {
	bpc.mutation.SetTitle(s)
	return bpc
}

// SetBody sets the "body" field.
func (bpc *BlogPostCreate) SetBody(s string) *BlogPostCreate {
	bpc.mutation.SetBody(s)
	return bpc
}

// SetExternalID sets the "external_id" field.
func (bpc *BlogPostCreate) SetExternalID(i int) *BlogPostCreate {
	bpc.mutation.SetExternalID(i)
	return bpc
}

// SetAuthorID sets the "author" edge to the User entity by ID.
func (bpc *BlogPostCreate) SetAuthorID(id int) *BlogPostCreate {
	bpc.mutation.SetAuthorID(id)
	return bpc
}

// SetNillableAuthorID sets the "author" edge to the User entity by ID if the given value is not nil.
func (bpc *BlogPostCreate) SetNillableAuthorID(id *int) *BlogPostCreate {
	if id != nil {
		bpc = bpc.SetAuthorID(*id)
	}
	return bpc
}

// SetAuthor sets the "author" edge to the User entity.
func (bpc *BlogPostCreate) SetAuthor(u *User) *BlogPostCreate {
	return bpc.SetAuthorID(u.ID)
}

// AddCategoryIDs adds the "categories" edge to the Category entity by IDs.
func (bpc *BlogPostCreate) AddCategoryIDs(ids ...int) *BlogPostCreate {
	bpc.mutation.AddCategoryIDs(ids...)
	return bpc
}

// AddCategories adds the "categories" edges to the Category entity.
func (bpc *BlogPostCreate) AddCategories(c ...*Category) *BlogPostCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return bpc.AddCategoryIDs(ids...)
}

// Mutation returns the BlogPostMutation object of the builder.
func (bpc *BlogPostCreate) Mutation() *BlogPostMutation {
	return bpc.mutation
}

// Save creates the BlogPost in the database.
func (bpc *BlogPostCreate) Save(ctx context.Context) (*BlogPost, error) {
	var (
		err  error
		node *BlogPost
	)
	if len(bpc.hooks) == 0 {
		if err = bpc.check(); err != nil {
			return nil, err
		}
		node, err = bpc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BlogPostMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bpc.check(); err != nil {
				return nil, err
			}
			bpc.mutation = mutation
			if node, err = bpc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(bpc.hooks) - 1; i >= 0; i-- {
			if bpc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bpc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, bpc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*BlogPost)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from BlogPostMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bpc *BlogPostCreate) SaveX(ctx context.Context) *BlogPost {
	v, err := bpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bpc *BlogPostCreate) Exec(ctx context.Context) error {
	_, err := bpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bpc *BlogPostCreate) ExecX(ctx context.Context) {
	if err := bpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bpc *BlogPostCreate) check() error {
	if _, ok := bpc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "BlogPost.title"`)}
	}
	if _, ok := bpc.mutation.Body(); !ok {
		return &ValidationError{Name: "body", err: errors.New(`ent: missing required field "BlogPost.body"`)}
	}
	if _, ok := bpc.mutation.ExternalID(); !ok {
		return &ValidationError{Name: "external_id", err: errors.New(`ent: missing required field "BlogPost.external_id"`)}
	}
	return nil
}

func (bpc *BlogPostCreate) sqlSave(ctx context.Context) (*BlogPost, error) {
	_node, _spec := bpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bpc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (bpc *BlogPostCreate) createSpec() (*BlogPost, *sqlgraph.CreateSpec) {
	var (
		_node = &BlogPost{config: bpc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: blogpost.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: blogpost.FieldID,
			},
		}
	)
	if value, ok := bpc.mutation.Title(); ok {
		_spec.SetField(blogpost.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := bpc.mutation.Body(); ok {
		_spec.SetField(blogpost.FieldBody, field.TypeString, value)
		_node.Body = value
	}
	if value, ok := bpc.mutation.ExternalID(); ok {
		_spec.SetField(blogpost.FieldExternalID, field.TypeInt, value)
		_node.ExternalID = value
	}
	if nodes := bpc.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   blogpost.AuthorTable,
			Columns: []string{blogpost.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.blog_post_author = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bpc.mutation.CategoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   blogpost.CategoriesTable,
			Columns: blogpost.CategoriesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BlogPostCreateBulk is the builder for creating many BlogPost entities in bulk.
type BlogPostCreateBulk struct {
	config
	builders []*BlogPostCreate
}

// Save creates the BlogPost entities in the database.
func (bpcb *BlogPostCreateBulk) Save(ctx context.Context) ([]*BlogPost, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bpcb.builders))
	nodes := make([]*BlogPost, len(bpcb.builders))
	mutators := make([]Mutator, len(bpcb.builders))
	for i := range bpcb.builders {
		func(i int, root context.Context) {
			builder := bpcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BlogPostMutation)
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
					_, err = mutators[i+1].Mutate(root, bpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bpcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bpcb *BlogPostCreateBulk) SaveX(ctx context.Context) []*BlogPost {
	v, err := bpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bpcb *BlogPostCreateBulk) Exec(ctx context.Context) error {
	_, err := bpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bpcb *BlogPostCreateBulk) ExecX(ctx context.Context) {
	if err := bpcb.Exec(ctx); err != nil {
		panic(err)
	}
}
