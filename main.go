package main

import (
	"fmt"
	"io/fs"
	"net/http"
	"time"

	"github.com/chneau/remoteav/av"
	"github.com/chneau/remoteav/common"
	"github.com/chneau/remoteav/dist"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/samber/lo"
)

func main() {
	resolver := common.NewResolver(lo.Must(av.GetCameras()))
	schema := graphql.MustParseSchema(common.SchemaString, resolver)
	dist := http.FileServer(http.FS(lo.Must(fs.Sub(dist.FrontendDist, "dist"))))

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(middleware.CleanPath)
	router.Use(middleware.GetHead)
	router.Use(middleware.StripSlashes)
	router.Use(middleware.Compress(5))
	router.Use(middleware.Timeout(3 * time.Second))

	router.Handle("/graphql", &relay.Handler{Schema: schema})
	router.Handle("/*", dist)
	router.Get("/stream", common.StreamHandlerfunc(resolver.ImageStream()))

	fmt.Println("Listening on port http://localhost:7777")
	lo.Must0(http.ListenAndServe(":7777", router))
}
