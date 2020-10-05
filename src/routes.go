package main

import (
	"github.com/gin-gonic/gin"
)

func initRoutes(r *gin.Engine) {
	r.Any("/", homePage)
	r.POST("/login", login)
	r.POST("/register", register)
	r.GET("/article/:id", fullArticle)
	r.POST("/article/:id", fullArticle)
	r.POST("/logout", logout)
	r.POST("/addComment", addComment)
	r.POST("/addArticle", addArticle)
	r.POST("/userPage", userPage)
	r.POST("/editUserInfo", editUserInfo)
	r.POST("/rateArticle", rateArticle)
	r.POST("/deleteArticle", deleteArticle)
	r.POST("/editArticle/:id", editArticle)
}
