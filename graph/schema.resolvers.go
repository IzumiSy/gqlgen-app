package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"gqlgen-app/graph/generated"
	"gqlgen-app/graph/model"

	"github.com/google/uuid"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, errors.New("User ID must be UUID format")
	}

	tx, err := r.DB.Tx(ctx)
	if err != nil {
		return nil, err
	}

	currentUser, err := tx.User.Get(ctx, userID)
	if err != nil {
		return nil, tryRollback(tx, err)
	} else if currentUser == nil {
		return nil, errors.New("Specified user ID does not exist")
	}

	createdTodo, err := tx.Todo.Create().
		SetText(input.Text).
		SetAssignee(currentUser).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	assignee := createdTodo.Edges.Assignee
	return &model.Todo{
		ID:     createdTodo.ID.String(),
		Text:   createdTodo.Text,
		UserID: assignee.ID.String(),
		User: &model.User{
			ID:   assignee.ID.String(),
			Name: assignee.Name,
		},
	}, nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	tx, err := r.DB.Tx(ctx)
	if err != nil {
		return nil, err
	}

	newUser, err := tx.User.Create().SetName(input.Name).Save(ctx)
	if err != nil {
		return nil, tryRollback(tx, err)
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &model.User{
		ID:   newUser.ID.String(),
		Name: newUser.Name,
	}, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	todos, err := r.DB.Todo.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	var todosResp []*model.Todo
	for _, todo := range todos {
		assignee := todo.Edges.Assignee
		todosResp = append(todosResp, &model.Todo{
			ID:     todo.ID.String(),
			Text:   todo.Text,
			UserID: assignee.ID.String(),
			User: &model.User{
				ID:   assignee.ID.String(),
				Name: assignee.Name,
			},
		})
	}

	return todosResp, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	users, err := r.DB.User.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	var usersResp []*model.User
	for _, user := range users {
		usersResp = append(usersResp, &model.User{
			ID:   user.ID.String(),
			Name: user.Name,
		})
	}

	return usersResp, nil
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
