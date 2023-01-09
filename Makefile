generate:
	go run github.com/99designs/gqlgen generate

start:
	go run server.go

migrate:
	migrate -database "mysql://root:root@tcp(localhost:3306)/hackernews" -path internal/pkg/db/migrations/mysql up

# GraphQL init
# go run github.com/99designs/gqlgen init