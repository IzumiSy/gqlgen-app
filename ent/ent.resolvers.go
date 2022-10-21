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

// ID is the resolver for the id field.
func (r *todoResolver) ID(ctx context.Context, obj *Todo) (int, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// ID is the resolver for the id field.
func (r *userResolver) ID(ctx context.Context, obj *User) (int, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// Category returns CategoryResolver implementation.
func (r *Resolver) Category() CategoryResolver { return &categoryResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Todo returns TodoResolver implementation.
func (r *Resolver) Todo() TodoResolver { return &todoResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type categoryResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
