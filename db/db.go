// TODO: Rstructure when moving to a config
// TODO: I feel like moving to a ORM is not a bad idea
package db

import (
	"database/sql"
	"embed"

	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

const (
	driverName string = "sqlite"
	dialect    string = "sqlite"
)

//go:embed migrations/*.sql
var migrationsFs embed.FS
var migrationDir string = "migrations"

func Connect(dsn string) (*sqlx.DB, error) {
	db, err := sql.Open(driverName, dsn)

	if err != nil {
		return nil, err
	}

	return sqlx.NewDb(db, driverName), nil
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
