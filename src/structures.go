package main

import (
	"sync"

	"github.com/gin-gonic/gin"
)

//User represents user's structure
type User struct {
	Nickname  string `form:"nickname" bson:"_id" json:"nickname"`
	Password  string `form:"password" bson:"password" json:"password"`
	Firstname string `form:"firstname" bson:"firstname" json:"firstname"`
	Lastname  string `form:"lastname" bson:"lastname" json:"lastname"`
	Age       int    `form:"age" bson:"age" json:"age"`
	Bicycle   string `form:"bicycle" bson:"bicycle" json:"bicycle"`
	IsAdmin   bool   `form:"isAdmin" bson:"isAdmin" json:"isAdmin"`
}

type obj map[string]interface{}

//Cache represents cache's structure based on map and mutex to able safe concurent interaction
type Cache struct {
	m map[string]User
	sync.RWMutex
}

func (c *Cache) getUser(ctx *gin.Context, key string) (usr User, found bool) {
	cookie, _ := ctx.Cookie(key)
	tokensCache.RLock()
	usr, found = tokensCache.m[cookie]
	tokensCache.RUnlock()
	return
}

//Article represents article's structure
type Article struct {
	ID           string `form:"ID" json:"id" bson:"_id"`
	Author       string `form:"author" json:"author" bson:"author"`
	Title        string `form:"title" json:"title" bson:"title"`
	Content      string `form:"content" json:"content" bson:"content"`
	CreationTime string `form:"creationTime" json:"creationTime" bson:"creationTime"`
	PrettyTime   string `form:"prettyTime" json:"prettyTime" bson:"prettyTime"`
}

//Comment represents comment's structure
type Comment struct {
	ArticleID    string `form:"articleId" json:"articleId" bson:"articleId"`
	Author       string `form:"author" json:"author" bson:"author"`
	Content      string `form:"content" json:"content" bson:"content"`
	CreationTime string `form:"creationTime" json:"creationTime" bson:"creationTime"`
	PrettyTime   string `form:"prettyTime" json:"prettyTime" bson:"prettyTime"`
}
