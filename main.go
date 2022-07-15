package main

import (
	"fmt"
	"io/fs"
	"net/http"
	"time"

	"github.com/chneau/remoteav/camera"
	"github.com/chneau/remoteav/common"
	"github.com/chneau/remoteav/dist"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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
	r.Use(middleware.Timeout(3 * time.Second))

	r.With(middleware.Logger).Get("/graphiql", graphiqlHandler)
	r.With(middleware.Logger).Handle("/graphql", relayHandler)
	dist := lo.Must(fs.Sub(dist.FrontendDist, "/dist"))
	r.Handle("/*", http.FileServer(http.FS(dist)))

	fmt.Println("Listening on port http://localhost:7777")
	lo.Must0(http.ListenAndServe(":7777", r))
}

func graphiqlHandler(w http.ResponseWriter, r *http.Request) {
	lo.Must(w.Write(common.GraphiqlHTML))
}
