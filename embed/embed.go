package embed

import _ "embed"

//go:embed schema.graphql
var SchemaString string

//go:embed graphiql.html
var GraphiqlHTML []byte
