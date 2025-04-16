migrate:
	migrate --database "postgres://boilerplate_user:boilerplate_password@localhost:5432/boilerplate?sslmode=disable" --path internal/storage/migrations up

sqlc:
	sqlc generate
