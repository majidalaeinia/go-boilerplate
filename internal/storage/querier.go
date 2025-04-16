// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package storage

import (
	"context"
)

type Querier interface {
	CreateItem(ctx context.Context, db DBTX, name string) (int32, error)
	CreateUser(ctx context.Context, db DBTX, name string) (int32, error)
	GetItem(ctx context.Context, db DBTX, id int32) (Item, error)
	GetItems(ctx context.Context, db DBTX) ([]Item, error)
	GetUser(ctx context.Context, db DBTX, id int32) (User, error)
	GetUsers(ctx context.Context, db DBTX) ([]User, error)
}

var _ Querier = (*Queries)(nil)
