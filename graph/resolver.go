//go:generate go run github.com/99designs/gqlgen generate

package graph

import (
	"gqlgen-app/ent"
)

type Resolver struct {
	DB *ent.Client
}
