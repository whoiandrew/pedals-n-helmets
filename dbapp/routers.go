package main

import (
	"github.com/gin-gonic/gin"
)

func InitRouters(r *gin.Engine) {
	r.POST("/addUserToDB", addUserToDB)
	r.POST("/getUserFromDB", getUserFromDB)
	r.POST("/getAllUsersFromDB", getAllUsersFromDB)
	r.POST("/getArticleFromDB", getArticleFromDB)
	r.POST("/getArticleCommentsFromDB", getArticleCommentsFromDB)
	r.POST("/deleteUserFromDB", deleteUserFromDB)
	r.POST("/getArticlesFromDB", getArticlesFromDB)
	r.POST("/addCommentToDB", addCommentToDB)
	r.POST("/addArticleToDB", addArticleToDB)
	r.POST("/updateUserFromDB", updateUserFromDB)
	r.POST("/updateRatesFromDB", updateRatesFromDB)
	r.POST("/deleteArticleFromDB", deleteArticleFromDB)
	r.POST("/editArticleFromDB", editArticleFromDB)

}
