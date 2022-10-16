package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"gqlgen-app/graph/generated"
	"gqlgen-app/graph/model"
	"math/rand"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	var currentUser *model.User = nil
	for _, u := range r.users {
		if u.ID == input.UserID {
			currentUser = u
		}
	}
	if currentUser == nil {
		user := &model.User{
			ID:   input.UserID,
			Name: "User" + input.UserID,
		}
		r.users = append(r.users, user)
		currentUser = user
	}

	todo := &model.Todo{
		ID:     fmt.Sprintf("T%d", rand.Int()),
		Text:   input.Text,
		UserID: currentUser.ID,
		User:   currentUser,
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
    return r.users, nil
}

// User is the resolver for the user field.
func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return &model.User{
		ID:   obj.UserID,
		Name: obj.User.Name,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
