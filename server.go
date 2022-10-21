package main

import (
	"gqlgen-app/ent"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

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

	/*
		gqlHandler := handler.New(ent.NewSchema(db))
		gqlHandler.AddTransport(transport.POST{})
		gqlHandler.Use(extension.AutomaticPersistedQuery{Cache: &graphql.MapCache{}})
	*/

	gqlHandler := handler.NewDefaultServer(ent.NewSchema(db))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", gqlHandler)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
