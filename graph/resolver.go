//go:generate go run github.com/99designs/gqlgen generate

package graph

import (
	"gqlgen-app/ent"
	"gqlgen-app/graph/model"
)

type Resolver struct {
	DB *ent.Client

	todos []*model.Todo
	users []*model.User
}
