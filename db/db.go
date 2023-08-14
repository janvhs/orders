// TODO: Rstructure when moving to a config
// TODO: I feel like moving to a ORM is not a bad idea
package db

import (
	"database/sql"
	"embed"

	"github.com/pressly/goose/v3"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/extra/bundebug"
	_ "modernc.org/sqlite"
)

const (
	driverName string = "sqlite"
	dialect    string = "sqlite"
)

//go:embed migrations/*.sql
var migrationsFs embed.FS
var migrationDir string = "migrations"

func Connect(dsn string, isDebug bool) (*bun.DB, error) {
	db, err := sql.Open(driverName, dsn)

	if err != nil {
		return nil, err
	}

	bunDb := bun.NewDB(db, sqlitedialect.New())

	if isDebug {
		bunDb.AddQueryHook(bundebug.NewQueryHook(
			bundebug.WithVerbose(true),
			bundebug.WithEnabled(true),
		))
	}

	if err := bunDb.Ping(); err != nil {
		return nil, err
	}

	return bunDb, nil
}

func Migrate(db *sql.DB) error {
	goose.SetBaseFS(migrationsFs)

	if err := goose.SetDialect(dialect); err != nil {
		return err
	}

	if err := goose.Up(db, migrationDir); err != nil {
		return err
	}

	return nil
}
