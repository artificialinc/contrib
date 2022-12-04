// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/artificialinc/contrib/entproto/internal/entprototest/ent/blogpost"
	"github.com/artificialinc/contrib/entproto/internal/entprototest/ent/image"
	"github.com/artificialinc/contrib/entproto/internal/entprototest/ent/predicate"
	"github.com/artificialinc/contrib/entproto/internal/entprototest/ent/skipedgeexample"
	"github.com/artificialinc/contrib/entproto/internal/entprototest/ent/user"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUserName sets the "user_name" field.
func (uu *UserUpdate) SetUserName(s string) *UserUpdate {
	uu.mutation.SetUserName(s)
	return uu
}

// SetStatus sets the "status" field.
func (uu *UserUpdate) SetStatus(u user.Status) *UserUpdate {
	uu.mutation.SetStatus(u)
	return uu
}

// SetUnnecessary sets the "unnecessary" field.
func (uu *UserUpdate) SetUnnecessary(s string) *UserUpdate {
	uu.mutation.SetUnnecessary(s)
	return uu
}

// SetNillableUnnecessary sets the "unnecessary" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUnnecessary(s *string) *UserUpdate {
	if s != nil {
		uu.SetUnnecessary(*s)
	}
	return uu
}

// ClearUnnecessary clears the value of the "unnecessary" field.
func (uu *UserUpdate) ClearUnnecessary() *UserUpdate {
	uu.mutation.ClearUnnecessary()
	return uu
}

// AddBlogPostIDs adds the "blog_posts" edge to the BlogPost entity by IDs.
func (uu *UserUpdate) AddBlogPostIDs(ids ...int) *UserUpdate {
	uu.mutation.AddBlogPostIDs(ids...)
	return uu
}

// AddBlogPosts adds the "blog_posts" edges to the BlogPost entity.
func (uu *UserUpdate) AddBlogPosts(b ...*BlogPost) *UserUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return uu.AddBlogPostIDs(ids...)
}

// SetProfilePicID sets the "profile_pic" edge to the Image entity by ID.
func (uu *UserUpdate) SetProfilePicID(id uuid.UUID) *UserUpdate {
	uu.mutation.SetProfilePicID(id)
	return uu
}

// SetNillableProfilePicID sets the "profile_pic" edge to the Image entity by ID if the given value is not nil.
func (uu *UserUpdate) SetNillableProfilePicID(id *uuid.UUID) *UserUpdate {
	if id != nil {
		uu = uu.SetProfilePicID(*id)
	}
	return uu
}

// SetProfilePic sets the "profile_pic" edge to the Image entity.
func (uu *UserUpdate) SetProfilePic(i *Image) *UserUpdate {
	return uu.SetProfilePicID(i.ID)
}

// SetSkipEdgeID sets the "skip_edge" edge to the SkipEdgeExample entity by ID.
func (uu *UserUpdate) SetSkipEdgeID(id int) *UserUpdate {
	uu.mutation.SetSkipEdgeID(id)
	return uu
}

// SetNillableSkipEdgeID sets the "skip_edge" edge to the SkipEdgeExample entity by ID if the given value is not nil.
func (uu *UserUpdate) SetNillableSkipEdgeID(id *int) *UserUpdate {
	if id != nil {
		uu = uu.SetSkipEdgeID(*id)
	}
	return uu
}

// SetSkipEdge sets the "skip_edge" edge to the SkipEdgeExample entity.
func (uu *UserUpdate) SetSkipEdge(s *SkipEdgeExample) *UserUpdate {
	return uu.SetSkipEdgeID(s.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearBlogPosts clears all "blog_posts" edges to the BlogPost entity.
func (uu *UserUpdate) ClearBlogPosts() *UserUpdate {
	uu.mutation.ClearBlogPosts()
	return uu
}

// RemoveBlogPostIDs removes the "blog_posts" edge to BlogPost entities by IDs.
func (uu *UserUpdate) RemoveBlogPostIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveBlogPostIDs(ids...)
	return uu
}

// RemoveBlogPosts removes "blog_posts" edges to BlogPost entities.
func (uu *UserUpdate) RemoveBlogPosts(b ...*BlogPost) *UserUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return uu.RemoveBlogPostIDs(ids...)
}

// ClearProfilePic clears the "profile_pic" edge to the Image entity.
func (uu *UserUpdate) ClearProfilePic() *UserUpdate {
	uu.mutation.ClearProfilePic()
	return uu
}

// ClearSkipEdge clears the "skip_edge" edge to the SkipEdgeExample entity.
func (uu *UserUpdate) ClearSkipEdge() *UserUpdate {
	uu.mutation.ClearSkipEdge()
	return uu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uu.hooks) == 0 {
		if err = uu.check(); err != nil {
			return 0, err
		}
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uu.check(); err != nil {
				return 0, err
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			if uu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Status(); ok {
		if err := user.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "User.status": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.UserName(); ok {
		_spec.SetField(user.FieldUserName, field.TypeString, value)
	}
	if value, ok := uu.mutation.Status(); ok {
		_spec.SetField(user.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := uu.mutation.Unnecessary(); ok {
		_spec.SetField(user.FieldUnnecessary, field.TypeString, value)
	}
	if uu.mutation.UnnecessaryCleared() {
		_spec.ClearField(user.FieldUnnecessary, field.TypeString)
	}
	if uu.mutation.BlogPostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.BlogPostsTable,
			Columns: []string{user.BlogPostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: blogpost.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedBlogPostsIDs(); len(nodes) > 0 && !uu.mutation.BlogPostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.BlogPostsTable,
			Columns: []string{user.BlogPostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: blogpost.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.BlogPostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.BlogPostsTable,
			Columns: []string{user.BlogPostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: blogpost.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.ProfilePicCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user.ProfilePicTable,
			Columns: []string{user.ProfilePicColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: image.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.ProfilePicIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user.ProfilePicTable,
			Columns: []string{user.ProfilePicColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: image.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.SkipEdgeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.SkipEdgeTable,
			Columns: []string{user.SkipEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: skipedgeexample.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.SkipEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.SkipEdgeTable,
			Columns: []string{user.SkipEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: skipedgeexample.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUserName sets the "user_name" field.
func (uuo *UserUpdateOne) SetUserName(s string) *UserUpdateOne {
	uuo.mutation.SetUserName(s)
	return uuo
}

// SetStatus sets the "status" field.
func (uuo *UserUpdateOne) SetStatus(u user.Status) *UserUpdateOne {
	uuo.mutation.SetStatus(u)
	return uuo
}

// SetUnnecessary sets the "unnecessary" field.
func (uuo *UserUpdateOne) SetUnnecessary(s string) *UserUpdateOne {
	uuo.mutation.SetUnnecessary(s)
	return uuo
}

// SetNillableUnnecessary sets the "unnecessary" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUnnecessary(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetUnnecessary(*s)
	}
	return uuo
}

// ClearUnnecessary clears the value of the "unnecessary" field.
func (uuo *UserUpdateOne) ClearUnnecessary() *UserUpdateOne {
	uuo.mutation.ClearUnnecessary()
	return uuo
}

// AddBlogPostIDs adds the "blog_posts" edge to the BlogPost entity by IDs.
func (uuo *UserUpdateOne) AddBlogPostIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddBlogPostIDs(ids...)
	return uuo
}

// AddBlogPosts adds the "blog_posts" edges to the BlogPost entity.
func (uuo *UserUpdateOne) AddBlogPosts(b ...*BlogPost) *UserUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return uuo.AddBlogPostIDs(ids...)
}

// SetProfilePicID sets the "profile_pic" edge to the Image entity by ID.
func (uuo *UserUpdateOne) SetProfilePicID(id uuid.UUID) *UserUpdateOne {
	uuo.mutation.SetProfilePicID(id)
	return uuo
}

// SetNillableProfilePicID sets the "profile_pic" edge to the Image entity by ID if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableProfilePicID(id *uuid.UUID) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetProfilePicID(*id)
	}
	return uuo
}

// SetProfilePic sets the "profile_pic" edge to the Image entity.
func (uuo *UserUpdateOne) SetProfilePic(i *Image) *UserUpdateOne {
	return uuo.SetProfilePicID(i.ID)
}

// SetSkipEdgeID sets the "skip_edge" edge to the SkipEdgeExample entity by ID.
func (uuo *UserUpdateOne) SetSkipEdgeID(id int) *UserUpdateOne {
	uuo.mutation.SetSkipEdgeID(id)
	return uuo
}

// SetNillableSkipEdgeID sets the "skip_edge" edge to the SkipEdgeExample entity by ID if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableSkipEdgeID(id *int) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetSkipEdgeID(*id)
	}
	return uuo
}

// SetSkipEdge sets the "skip_edge" edge to the SkipEdgeExample entity.
func (uuo *UserUpdateOne) SetSkipEdge(s *SkipEdgeExample) *UserUpdateOne {
	return uuo.SetSkipEdgeID(s.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearBlogPosts clears all "blog_posts" edges to the BlogPost entity.
func (uuo *UserUpdateOne) ClearBlogPosts() *UserUpdateOne {
	uuo.mutation.ClearBlogPosts()
	return uuo
}

// RemoveBlogPostIDs removes the "blog_posts" edge to BlogPost entities by IDs.
func (uuo *UserUpdateOne) RemoveBlogPostIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveBlogPostIDs(ids...)
	return uuo
}

// RemoveBlogPosts removes "blog_posts" edges to BlogPost entities.
func (uuo *UserUpdateOne) RemoveBlogPosts(b ...*BlogPost) *UserUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return uuo.RemoveBlogPostIDs(ids...)
}

// ClearProfilePic clears the "profile_pic" edge to the Image entity.
func (uuo *UserUpdateOne) ClearProfilePic() *UserUpdateOne {
	uuo.mutation.ClearProfilePic()
	return uuo
}

// ClearSkipEdge clears the "skip_edge" edge to the SkipEdgeExample entity.
func (uuo *UserUpdateOne) ClearSkipEdge() *UserUpdateOne {
	uuo.mutation.ClearSkipEdge()
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	if len(uuo.hooks) == 0 {
		if err = uuo.check(); err != nil {
			return nil, err
		}
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uuo.check(); err != nil {
				return nil, err
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			if uuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, uuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*User)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UserMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Status(); ok {
		if err := user.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "User.status": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.UserName(); ok {
		_spec.SetField(user.FieldUserName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Status(); ok {
		_spec.SetField(user.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := uuo.mutation.Unnecessary(); ok {
		_spec.SetField(user.FieldUnnecessary, field.TypeString, value)
	}
	if uuo.mutation.UnnecessaryCleared() {
		_spec.ClearField(user.FieldUnnecessary, field.TypeString)
	}
	if uuo.mutation.BlogPostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.BlogPostsTable,
			Columns: []string{user.BlogPostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: blogpost.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedBlogPostsIDs(); len(nodes) > 0 && !uuo.mutation.BlogPostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.BlogPostsTable,
			Columns: []string{user.BlogPostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: blogpost.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.BlogPostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.BlogPostsTable,
			Columns: []string{user.BlogPostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: blogpost.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.ProfilePicCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user.ProfilePicTable,
			Columns: []string{user.ProfilePicColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: image.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.ProfilePicIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user.ProfilePicTable,
			Columns: []string{user.ProfilePicColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: image.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.SkipEdgeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.SkipEdgeTable,
			Columns: []string{user.SkipEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: skipedgeexample.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.SkipEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.SkipEdgeTable,
			Columns: []string{user.SkipEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: skipedgeexample.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
