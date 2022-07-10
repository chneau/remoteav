package main

import (
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
)

func main() {
	r := gin.Default()

	lo.Must0(r.Run(":7777"))
}
