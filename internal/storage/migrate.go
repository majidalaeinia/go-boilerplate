package storage

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
	"github.com/pkg/errors"
)

//go:embed migrations
var migrationFiles embed.FS

func EnsureMigrationsDone(driver database.Driver, dbName string) error {
	httpFS, err := fs.Sub(migrationFiles, "migrations")
	if err != nil {
		return err
	}

	srcDriver, err := httpfs.New(http.FS(httpFS), ".")
	if err != nil {
		return err
	}
	m, err := migrate.NewWithInstance(
		"httpfs",
		srcDriver,
		dbName, driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
