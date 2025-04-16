package users

import (
	"github.com/ehsundar/go-boilerplate/internal/storage"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
)

type UsersServer interface {
	GetUsers(http.ResponseWriter, *http.Request)
}

type usersServer struct {
	pool *pgxpool.Pool

	querier storage.Querier
}

func NewUsersServer(pool *pgxpool.Pool, querier storage.Querier) UsersServer {
	return &usersServer{
		pool:    pool,
		querier: querier,
	}
}
