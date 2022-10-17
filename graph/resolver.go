//go:generate go run github.com/99designs/gqlgen generate

package graph

import (
	"database/sql"
	"gqlgen-app/graph/model"
)

type Resolver struct {
	DB *sql.DB

	todos []*model.Todo
	users []*model.User
}
