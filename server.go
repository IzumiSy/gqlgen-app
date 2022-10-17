package main

import (
	"gqlgen-app/ent"
	"gqlgen-app/graph"
	"gqlgen-app/graph/generated"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := ent.Open("sqlite", "./main.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	config := generated.Config{Resolvers: &graph.Resolver{DB: db}}
	gqlHandler := handler.New(generated.NewExecutableSchema(config))
	gqlHandler.AddTransport(transport.POST{})
	gqlHandler.Use(extension.AutomaticPersistedQuery{Cache: &graphql.MapCache{}})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", gqlHandler)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
