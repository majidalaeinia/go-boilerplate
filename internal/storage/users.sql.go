// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: users.sql

package storage

import (
	"context"
)

const createUser = `-- name: CreateUser :one
insert into users (name)
values ($1)
    returning id
`

func (q *Queries) CreateUser(ctx context.Context, db DBTX, name string) (int32, error) {
	row := db.QueryRow(ctx, createUser, name)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const getUser = `-- name: GetUser :one
select id, name
from users
where id = $1
`

func (q *Queries) GetUser(ctx context.Context, db DBTX, id int32) (User, error) {
	row := db.QueryRow(ctx, getUser, id)
	var i User
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getUsers = `-- name: GetUsers :many
select id, name
from users
`

func (q *Queries) GetUsers(ctx context.Context, db DBTX) ([]User, error) {
	rows, err := db.Query(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
