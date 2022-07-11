package main

import (
	_ "embed"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
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
	resolver := &Resolver{}
	schema := graphql.MustParseSchema(schemaString, resolver)
	resolver.cameras = getCameras()
	runRouter(schema)
}

func getCameras() []*Camera {
	files := lo.Must(ioutil.ReadDir("/dev/"))
	result := []*Camera{}
	for _, file := range files {
		fileName := file.Name()
		if len(fileName) <= 5 || fileName[:5] != "video" {
			continue
		}
		cameraNumber := lo.Must(strconv.Atoi(fileName[5:]))
		result = append(result, &Camera{id: int32(cameraNumber)})
	}
	return result
}

func runRouter(schema *graphql.Schema) {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(3 * time.Second))

	r.With(middleware.Logger).Get("/graphiql", func(w http.ResponseWriter, r *http.Request) {
		lo.Must(w.Write(graphiqlHTML))
	})
	r.With(middleware.Logger).Handle("/graphql", &relay.Handler{Schema: schema})
	r.Get("/*", proxy)

	fmt.Println("Listening on port http://localhost:7777")
	lo.Must0(http.ListenAndServe(":7777", r))
}
