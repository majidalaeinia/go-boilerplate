package cmd

import (
	"github.com/golang-migrate/migrate/v4/database/pgx/v5"
	"github.com/spf13/cobra"

	"github.com/ehsundar/go-boilerplate/internal/storage"
)

func init() {
	rootCmd.AddCommand(MigrateCmd)
}

var (
	MigrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: "Migrate the database",
		RunE: func(_ *cobra.Command, _ []string) error {
			config, err := LoadConfig()
			if err != nil {
				return err
			}

			postgres := pgx.Postgres{}
			driver, err := postgres.Open(config.PostgresConn)
			if err != nil {
				return err
			}
			defer driver.Close()

			return storage.EnsureMigrationsDone(driver, "boilerplate")
		},
	}
)
