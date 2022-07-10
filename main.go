package main

import (
	_ "embed"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/graphql-go"
	"github.com/samber/lo"
)

//go:embed schema.graphql
var schemaString string
var schema *graphql.Schema

//go:embed graphiql.html
var graphiqlHTML []byte

func init() {
	// schema = graphql.MustParseSchema(schemaString, &Resolver{})
}

func main() {
	r := gin.Default()

	r.GET("/graphiql", func(c *gin.Context) {
		lo.Must0(c.Writer.Write(graphiqlHTML))
	})
	r.Any("/*proxyPath", proxy)

	fmt.Println("Listening on port http://localhost:7777")
	lo.Must0(r.Run(":7777"))
}
