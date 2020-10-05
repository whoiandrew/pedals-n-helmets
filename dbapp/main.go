package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	fillCaches()
	InitRouters(r)
	r.Run(":8082")
}
