package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/SantiagoBedoya/hackernews/graph"
	"github.com/SantiagoBedoya/hackernews/internal/auth"
	database "github.com/SantiagoBedoya/hackernews/internal/pkg/db/mysql"
	"github.com/go-chi/chi"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()
	router.Use(auth.Middleware())

	database.InitDB()
	defer database.CloseDB()
	database.Migrate()

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
