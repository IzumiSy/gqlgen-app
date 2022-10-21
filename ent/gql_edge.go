// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (c *Category) Todos(ctx context.Context) (result []*Todo, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = c.NamedTodos(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = c.Edges.TodosOrErr()
	}
	if IsNotLoaded(err) {
		result, err = c.QueryTodos().All(ctx)
	}
	return result, err
}

func (t *Todo) Assignee(ctx context.Context) (*User, error) {
	result, err := t.Edges.AssigneeOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryAssignee().Only(ctx)
	}
	return result, err
}

func (t *Todo) Categories(ctx context.Context) (result []*Category, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = t.NamedCategories(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = t.Edges.CategoriesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = t.QueryCategories().All(ctx)
	}
	return result, err
}

func (u *User) Todos(ctx context.Context) (result []*Todo, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedTodos(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.TodosOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryTodos().All(ctx)
	}
	return result, err
}