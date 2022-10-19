package main

import (
	"database/sql"
	"gqlgen-app/drivers"
	"gqlgen-app/ent"
	"gqlgen-app/graph"
	"gqlgen-app/graph/generated"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"modernc.org/sqlite"
)

func init() {
	sql.Register("sqlite3", drivers.SqliteDriver{Driver: &sqlite.Driver{}})
}

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := ent.Open("sqlite3", "./main.db")
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
