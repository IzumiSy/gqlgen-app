package ent

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
)

// ID is the resolver for the id field.
func (r *categoryResolver) ID(ctx context.Context, obj *Category) (int, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id int) (Noder, error) {
	panic(fmt.Errorf("not implemented: Node - node"))
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]Noder, error) {
	panic(fmt.Errorf("not implemented: Nodes - nodes"))
}

// Categories is the resolver for the categories field.
func (r *queryResolver) Categories(ctx context.Context) ([]*Category, error) {
	panic(fmt.Errorf("not implemented: Categories - categories"))
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*Todo, error) {
	panic(fmt.Errorf("not implemented: Todos - todos"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*User, error) {
    return r.DB.User.Query().All(ctx)
}

// ID is the resolver for the id field.
func (r *todoResolver) ID(ctx context.Context, obj *Todo) (int, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// ID is the resolver for the id field.
func (r *userResolver) ID(ctx context.Context, obj *User) (int, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// TodoIDs is the resolver for the todoIDs field.
func (r *createCategoryInputResolver) TodoIDs(ctx context.Context, obj *CreateCategoryInput, data []int) error {
	panic(fmt.Errorf("not implemented: TodoIDs - todoIDs"))
}

// AssigneeID is the resolver for the assigneeID field.
func (r *createTodoInputResolver) AssigneeID(ctx context.Context, obj *CreateTodoInput, data int) error {
	panic(fmt.Errorf("not implemented: AssigneeID - assigneeID"))
}

// CategoryIDs is the resolver for the categoryIDs field.
func (r *createTodoInputResolver) CategoryIDs(ctx context.Context, obj *CreateTodoInput, data []int) error {
	panic(fmt.Errorf("not implemented: CategoryIDs - categoryIDs"))
}

// TodoIDs is the resolver for the todoIDs field.
func (r *createUserInputResolver) TodoIDs(ctx context.Context, obj *CreateUserInput, data []int) error {
	panic(fmt.Errorf("not implemented: TodoIDs - todoIDs"))
}

// Category returns CategoryResolver implementation.
func (r *Resolver) Category() CategoryResolver { return &categoryResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Todo returns TodoResolver implementation.
func (r *Resolver) Todo() TodoResolver { return &todoResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

// CreateCategoryInput returns CreateCategoryInputResolver implementation.
func (r *Resolver) CreateCategoryInput() CreateCategoryInputResolver {
	return &createCategoryInputResolver{r}
}

// CreateTodoInput returns CreateTodoInputResolver implementation.
func (r *Resolver) CreateTodoInput() CreateTodoInputResolver { return &createTodoInputResolver{r} }

// CreateUserInput returns CreateUserInputResolver implementation.
func (r *Resolver) CreateUserInput() CreateUserInputResolver { return &createUserInputResolver{r} }

type categoryResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
type createCategoryInputResolver struct{ *Resolver }
type createTodoInputResolver struct{ *Resolver }
type createUserInputResolver struct{ *Resolver }
