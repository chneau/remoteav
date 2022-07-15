//go:build embed

package common

import "embed"

//go:embed dist
var FrontendDist embed.FS
