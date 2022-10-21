//go:build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	ex, err := entgql.NewExtension(
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("ent.graphql"),
	)
	if err != nil {
		log.Fatalf("failed creating entgql extension: %s", err.Error())
	}

	opts := []entc.Option{
		entc.Extensions(ex),
	}
	if err := entc.Generate("./ent/schema", &gen.Config{}, opts...); err != nil {
		log.Fatalf("Failed running ent codegen: %s", err.Error())
	}
}
