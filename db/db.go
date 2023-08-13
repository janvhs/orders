// TODO: Rstructure when moving to a config
// TODO: I feel like moving to a ORM is not a bad idea
package db

import (
	"database/sql"
	"embed"
	"sync"

	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

const (
	driverName string = "sqlite"
	dialect    string = "sqlite"
	file       string = "db.sqlite"
)

var db *sql.DB
var dbOnce sync.Once

//go:embed migrations/*.sql
var migrationsFs embed.FS
var migrationDir string = "migrations"

func open() (*sqlx.DB, error) {
	var err error
	dbOnce.Do(func() {
		db, err = sql.Open(driverName, file)
	})

	if err != nil {
		return nil, err
	}

	return sqlx.NewDb(db, driverName), nil
}

func Connect() (*sqlx.DB, error) {
	return open()
}

func Migrate() error {
	db, err := open()
	if err != nil {
		return err
	}

	goose.SetBaseFS(migrationsFs)

	if err := goose.SetDialect(dialect); err != nil {
		return err
	}

	if err := goose.Up(db.DB, migrationDir); err != nil {
		return err
	}

	return nil
}
