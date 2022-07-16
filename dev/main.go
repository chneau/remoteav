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
	resolver := common.NewResolver(lo.Must(camera.GetCameras()))
	schema := graphql.MustParseSchema(common.SchemaString, resolver)

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.With(middleware.Logger).Handle("/graphql", &relay.Handler{Schema: schema})
	router.Get("/graphiql", graphiqlHandler)
	router.Get("/*", proxy)
	router.Get("/stream", common.StreamHandlerfunc(resolver.ImageStream()))
	fmt.Println("Listening on port http://localhost:7777")
	lo.Must0(http.ListenAndServe(":7777", router))
}

func graphiqlHandler(w http.ResponseWriter, r *http.Request) {
	lo.Must(w.Write(common.GraphiqlHTML))
}
