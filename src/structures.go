package main

import (
	"net/url"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

type urlValuesGetter interface {
	getUrlValues()
}

//User represents user's structure
type User struct {
	Nickname  string `form:"nickname" bson:"_id" json:"nickname"`
	Password  string `form:"password" bson:"password" json:"password"`
	Firstname string `form:"firstname" bson:"firstname" json:"firstname"`
	Lastname  string `form:"lastname" bson:"lastname" json:"lastname"`
	Age       int    `form:"age" bson:"age" json:"age"`
	Bicycle   string `form:"bicycle" bson:"bicycle" json:"bicycle"`
	IsAdmin   bool   `form:"isAdmin" bson:"isAdmin" json:"isAdmin"`
	IsModer   bool   `form:"isModer" bson:"isModer" json:"isModer"`
}

//GetURLValues returns url.Values struct's implementation
func (usr *User) GetURLValues() url.Values {
	return url.Values{
		"nickname":  {usr.Nickname},
		"password":  {usr.Password},
		"firstname": {usr.Firstname},
		"lastname":  {usr.Lastname},
		"age":       {strconv.Itoa(usr.Age)},
		"bicycle":   {usr.Bicycle},
		"isModer":   {strconv.FormatBool(usr.IsModer)},
		"isAdmin":   {strconv.FormatBool(usr.IsAdmin)},
	}
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
	Rates        obj    `form:"rates" json:"rates" bson:"rates"`
}

//CountRates returns based on likes/dislikes total rating
func CountRates(rates obj) (res int) {
	for _, v := range rates {
		if v == "like" {
			res++
		} else {
			res--
		}
	}
	return
}

//GetURLValues returns url.Values struct's implementation
func (a *Article) GetURLValues() url.Values {
	return url.Values{
		"id":           {a.ID},
		"author":       {a.Author},
		"title":        {a.Title},
		"content":      {a.Content},
		"creationTime": {a.CreationTime},
		"prettyTime":   {a.PrettyTime},
	}
}

//Comment represents comment's structure
type Comment struct {
	ArticleID    string `form:"articleId" json:"articleId" bson:"articleId"`
	Author       string `form:"author" json:"author" bson:"author"`
	Content      string `form:"content" json:"content" bson:"content"`
	CreationTime string `form:"creationTime" json:"creationTime" bson:"creationTime"`
	PrettyTime   string `form:"prettyTime" json:"prettyTime" bson:"prettyTime"`
}

//GetURLValues returns url.Values struct's implementation
func (c *Comment) GetURLValues() url.Values {
	return url.Values{
		"articleId":    {c.ArticleID},
		"author":       {c.Author},
		"content":      {c.Content},
		"creationTime": {c.CreationTime},
		"prettyTime":   {c.PrettyTime},
	}
}
