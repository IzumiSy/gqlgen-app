//go:build ignore

package main

import (
	"context"
	"database/sql"
	"gqlgen-app/drivers"
	"log"
	"os"

	_ "ariga.io/atlas/sql/sqlite"
	"modernc.org/sqlite"

	atlas "ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/examples/migration/ent/migrate"
)

func init() {
	sql.Register("sqlite3", drivers.SqliteDriver{Driver: &sqlite.Driver{}})
}

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
		schema.WithMigrationMode(schema.ModeReplay),
		schema.WithDialect(dialect.SQLite),
		schema.WithFormatter(atlas.DefaultFormatter),
	}

	ctx := context.Background()
	if err := migrate.NamedDiff(ctx, "sqlite3://main.db", os.Args[1], opts...); err != nil {
		log.Fatalf("failed generating migration file: %v", err)
	}
}
