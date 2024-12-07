package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"errors"
	"github.com/ehsundar/go-boilerplate/internal/items"
	"github.com/ehsundar/go-boilerplate/internal/storage"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	pool, err := storage.NewConnectionPool(ctx, config.PostgresConn)
	if err != nil {
		return err
	}
	defer pool.Close()

	_, err = storage.NewRedisClient(ctx, config.RedisConn)
	if err != nil {
		return err
	}

	querier := storage.New()

	itemsServer := items.NewItemsServer(pool, querier)

	mux := http.NewServeMux()

	server := &http.Server{
		Addr:        config.ServerAddr,
		Handler:     mux,
		BaseContext: func(listener net.Listener) context.Context { return ctx },
	}

	registerRoutes(mux, itemsServer)

	go func() {
		slog.Info("Starting server on %s", config.ServerAddr)
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			slog.Error("HTTP server error: %v", err)
		}
		slog.Info("Stopped serving new connections")
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	slog.Info("Shutting down server")
	return server.Shutdown(shutdownCtx)
}

func registerRoutes(mux *http.ServeMux, itemsServer items.ItemsServer) {
	mux.HandleFunc("/items", itemsServer.GetItems)
}
