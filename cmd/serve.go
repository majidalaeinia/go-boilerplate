package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/ehsundar/go-boilerplate/internal/storage"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use: "serve",
	RunE: func(cmd *cobra.Command, args []string) error {
		return serve()
	},
}

func serve() error {
	ctx := context.Background()

	config, err := LoadConfig()
	if err != nil {
		return err
	}

	_, err = storage.NewConnectionPool(ctx, config.PostgresConn)
	if err != nil {
		return err
	}

	_, err = storage.NewRedisClient(ctx, config.RedisConn)
	if err != nil {
		return err
	}

	return nil
}
