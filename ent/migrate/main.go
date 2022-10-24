//go:build ignore

package main

import (
	"context"
	"log"
	"os"

	_ "ariga.io/atlas/sql/sqlite"
	_ "github.com/mattn/go-sqlite3"

	"gqlgen-app/ent/migrate"

	atlas "ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("migration name is required. Use: 'go run -mod=mod ent/migrate/main.go <name>'")
	}

	dir, err := atlas.NewLocalDir("ent/migrate/migration")
	if err != nil {
		log.Fatal("failed creating migration directory: %s", err.Error())
	}

	opts := []schema.MigrateOption{
		schema.WithDir(dir),
		// schema.WithMigrationMode(schema.ModeReplay),
		schema.WithDialect(dialect.SQLite),
		schema.WithFormatter(atlas.DefaultFormatter),
	}

	ctx := context.Background()
	if err := migrate.NamedDiff(ctx, "sqlite3://main.db?_fk=1", os.Args[1], opts...); err != nil {
		log.Fatalf("failed generating migration file: %v", err)
	}
}
