package items

import (
	"github.com/ehsundar/go-boilerplate/internal/storage"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
)

type ItemsServer interface {
	GetItems(http.ResponseWriter, *http.Request)
}

type itemsServer struct {
	pool *pgxpool.Pool

	querier storage.Querier
}

func NewItemsServer(pool *pgxpool.Pool, querier storage.Querier) ItemsServer {
	return &itemsServer{
		pool:    pool,
		querier: querier,
	}
}
