package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/chneau/remoteav/common"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/samber/lo"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	resolver := common.NewResolver()
	schema := graphql.MustParseSchema(common.SchemaString, resolver)

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(middleware.CleanPath)
	router.Use(middleware.GetHead)
	router.Use(middleware.StripSlashes)
	router.Use(middleware.Compress(5))
	router.Use(middleware.Timeout(3 * time.Second))

	router.With(middleware.Logger).Handle("/graphql", &relay.Handler{Schema: schema})
	router.Get("/graphiql", graphiqlHandler)
	router.Get("/*", httputil.NewSingleHostReverseProxy(lo.Must(url.Parse("http://localhost:5173"))).ServeHTTP)
	router.Get(common.VideoPath, common.StreamVideoHandler(resolver.VideoStream()))
	router.Get(common.AudioPath, common.StreamAudioHandler(resolver.AudioStream()))

	fmt.Println("Listening on port http://localhost:7777")
	lo.Must0(http.ListenAndServe(":7777", router))
}

func graphiqlHandler(w http.ResponseWriter, r *http.Request) {
	lo.Must(w.Write(common.GraphiqlHTML))
}
