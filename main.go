package main

import (
	_ "embed"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/samber/lo"
)

//go:embed schema.graphql
var schemaString string

//go:embed graphiql.html
var graphiqlHTML []byte

func main() {
	schema := graphql.MustParseSchema(schemaString, &Resolver{})

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(10 * time.Second))

	r.Get("/graphiql", func(w http.ResponseWriter, r *http.Request) {
		lo.Must(w.Write(graphiqlHTML))
	})
	r.Handle("/graphql", &relay.Handler{Schema: schema})
	r.Get("/*", proxy)

	fmt.Println("Listening on port http://localhost:7777")
	lo.Must0(http.ListenAndServe(":7777", r))
}
