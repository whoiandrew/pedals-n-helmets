package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	ai "github.com/night-codes/mgo-ai"
)

var (
	ctl, _        = NewController()
	usersCache    = Cache{m: make(map[string]interface{})}
	articlesCache = Cache{m: make(map[string]interface{})}
)

func getArticleFromDB(c *gin.Context) {
	id, _ := c.GetPostForm("id")
	articlesCache.RLock()
	res := articlesCache.m[id]
	articlesCache.RUnlock()
	c.JSON(http.StatusOK, res)
}

func getArticleCommentsFromDB(c *gin.Context) {
	articleID, _ := c.GetPostForm("id")
	res := ctl.getArticleComments(articleID)
	c.JSON(http.StatusOK, res)
}

func addUserToDB(c *gin.Context) {
	u := User{}
	c.Bind(&u)
	ctl.addUser(u)
	usersCache.Lock()
	usersCache.m[u.Nickname] = u
	usersCache.Unlock()
	c.JSON(http.StatusOK, u)
}

func getUserFromDB(c *gin.Context) {
	u := User{}
	c.Bind(&u)
	usersCache.RLock()
	res := usersCache.m[u.Nickname]
	usersCache.RUnlock()
	c.JSON(http.StatusOK, res)
}

func getAllUsersFromDB(c *gin.Context) {
	usrs := make([]User, 0, len(usersCache.m))
	for _, v := range usersCache.m {
		if u, ok := v.(User); ok {
			usrs = append(usrs, u)
		}
	}
	c.JSON(http.StatusOK, usrs)
}

func deleteUserFromDB(c *gin.Context) {
	id, _ := c.GetPostForm("id")
	usersCache.DeleteItem(id)
	ctl.deleteUser(id)
	c.Status(http.StatusOK)
}

func getArticlesFromDB(c *gin.Context) {
	articles := make([]Article, 0, len(articlesCache.m))
	for _, v := range articlesCache.m {
		if a, ok := v.(Article); ok {
			articles = append(articles, a)
		}
	}
	c.JSON(http.StatusOK, articles)
}

func addCommentToDB(c *gin.Context) {
	comment := Comment{}
	c.Bind(&comment)
	ctl.addComment(comment)
	c.Status(http.StatusOK)
}

func addArticleToDB(c *gin.Context) {
	session := ctl.session.Clone()
	defer session.Close()
	article := Article{}
	c.Bind(&article)
	ai.Connect(session.DB("main").C("counters"))
	article.ID = strconv.FormatUint(ai.Next("articles"), 10)
	cTime := time.Now()
	article.CreationTime = cTime.Format(time.RFC3339)
	article.PrettyTime = cTime.Format("01-02-2006 15:04:05")
	articlesCache.Lock()
	articlesCache.m[article.ID] = article
	articlesCache.Unlock()
	article = ctl.addArticle(article)
	c.JSON(http.StatusOK, article)
}

func updateUserFromDB(c *gin.Context) {
	user := User{}
	c.Bind(&user)
	ctl.updateUser(user)
	usersCache.Lock()
	if u, ok := usersCache.m[user.Nickname].(User); ok {
		u.Firstname = user.Firstname
		u.Lastname = user.Lastname
		u.Age = user.Age
		u.Bicycle = user.Bicycle
		usersCache.m[user.Nickname] = u
	}
	usersCache.Unlock()
	c.Status(http.StatusOK)

}

func updateRatesFromDB(c *gin.Context) {
	id, _ := c.GetPostForm("id")
	nickname, _ := c.GetPostForm("nickname")
	rate, _ := c.GetPostForm("rate")
	ctl.updateRates(id, nickname, rate)
	c.Status(http.StatusOK)
}

func editArticleFromDB(c *gin.Context) {
	id, _ := c.GetPostForm("id")
	content, _ := c.GetPostForm("content")
	ctl.editArticle(id, content)
	articlesCache.Lock()
	if a, ok := articlesCache.m[id].(Article); ok {
		a.Content = content
		articlesCache.m[id] = a
	}
	articlesCache.Unlock()
	c.Status(http.StatusOK)

}

func deleteArticleFromDB(c *gin.Context) {
	id, _ := c.GetPostForm("id")
	ctl.deleteArticle(id)
	articlesCache.DeleteItem(id)
	c.Status(http.StatusOK)
}
