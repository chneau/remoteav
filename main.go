//go:build embed

package main

import (
	"fmt"

	"github.com/chneau/remoteav/camera"
	"github.com/chneau/remoteav/common"
	"github.com/chneau/remoteav/embed"
	"github.com/graph-gophers/graphql-go"
	"github.com/samber/lo"
)

func main() {
	resolver := &common.Resolver{}
	schema := graphql.MustParseSchema(embed.SchemaString, resolver)
	resolver.Cameras_ = lo.Must(camera.GetCameras())
	fmt.Println("schema:", schema)
}
