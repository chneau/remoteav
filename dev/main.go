package main

import (
	"fmt"
	"net/http"

	"github.com/chneau/remoteav/camera"
	"github.com/chneau/remoteav/common"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/samber/lo"
)

func main() {
	resolver := &common.Resolver{}
	schema := graphql.MustParseSchema(common.SchemaString, resolver)
	resolver.Cameras_ = lo.Must(camera.GetCameras())
	runRouter(schema)
}

func runRouter(schema *graphql.Schema) {
	relayHandler := &relay.Handler{Schema: schema}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	r.With(middleware.Logger).Handle("/graphql", relayHandler)
	r.Get("/graphiql", graphiqlHandler)
	r.Get("/*", proxy)

	fmt.Println("Listening on port http://localhost:7777")
	lo.Must0(http.ListenAndServe(":7777", r))
}

func graphiqlHandler(w http.ResponseWriter, r *http.Request) {
	lo.Must(w.Write(common.GraphiqlHTML))
}
