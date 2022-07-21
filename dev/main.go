package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/chneau/remoteav/common"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/samber/lo"
)

func main() {
	// microphones := lo.Must(av.GetMicrophones())
	// fmt.Print("\033[H\033[2J") // Clear screen
	// log.Printf("microphones: %#+v\n", microphones)
	// audio := make(chan []float32)
	// microphones[0].Stream(audio)
	// defer microphones[0].Close()
	// for frame := range audio {
	// 	log.Println("frame:", len(frame))
	// }

	log.SetFlags(log.LstdFlags | log.Llongfile)
	resolver := common.NewResolver()
	schema := graphql.MustParseSchema(common.SchemaString, resolver)

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.With(middleware.Logger).Handle("/graphql", &relay.Handler{Schema: schema})
	router.Get("/graphiql", graphiqlHandler)
	url := lo.Must(url.Parse("http://localhost:5173"))
	proxy := httputil.NewSingleHostReverseProxy(url)
	router.Get("/*", proxy.ServeHTTP)
	router.Get(common.VideoPath, common.StreamVideoHandler(resolver.VideoStream()))
	router.Get(common.AudioPath, common.StreamAudioHandler(resolver.AudioStream()))
	fmt.Println("Listening on port http://localhost:7777")
	lo.Must0(http.ListenAndServe(":7777", router))
}

func graphiqlHandler(w http.ResponseWriter, r *http.Request) {
	lo.Must(w.Write(common.GraphiqlHTML))
}
