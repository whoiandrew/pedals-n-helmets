package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("/home/whoiandrew/golang-course/golang-homeworks1/golang-gin/front/templates/*")
	router.Static("/assets", "/home/whoiandrew/golang-course/golang-homeworks1/golang-gin/front/assets/")
	initRoutes(router)
	router.Run(":8081")
}
