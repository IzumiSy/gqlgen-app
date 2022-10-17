package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"gqlgen-app/graph/generated"
	"gqlgen-app/graph/model"
	"math/rand"

	"github.com/google/uuid"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, err
	}

	// TODO: do rollback on error and do commit at last
	tx, err := r.DB.Tx(ctx)
	if err != nil {
		return nil, err
	}

	currentUser, err := tx.User.Get(ctx, userID)
	if err != nil {
		return nil, err
	}

	//
	// TODO: insert a new user when the user with the ID in the request does not exists on DB
	//

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
