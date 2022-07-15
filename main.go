package main

import (
	"fmt"
	"io/fs"
	"net/http"
	"time"

	"github.com/chneau/remoteav/camera"
	"github.com/chneau/remoteav/common"
	"github.com/chneau/remoteav/dist"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/samber/lo"
)

func main() {
	resolver := &common.Resolver{Cameras_: lo.Must(camera.GetCameras())}
	schema := graphql.MustParseSchema(common.SchemaString, resolver)
	dist := http.FileServer(http.FS(lo.Must(fs.Sub(dist.FrontendDist, "dist"))))

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.CleanPath)
	r.Use(middleware.GetHead)
	r.Use(middleware.StripSlashes)
	r.Use(middleware.Compress(5))
	r.Use(middleware.Timeout(3 * time.Second))

	r.Handle("/graphql", &relay.Handler{Schema: schema})
	r.Handle("/*", dist)

	fmt.Println("Listening on port http://localhost:7777")
	lo.Must0(http.ListenAndServe(":7777", r))
}
