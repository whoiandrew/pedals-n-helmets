package main

import "sync"

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

type obj map[string]interface{}

//Cache represents cache's structure based on map and mutex to able safe concurent interaction
type Cache struct {
	m map[string]interface{}
	sync.RWMutex
}

//DeleteItem allows delete items from cache using RWmutex
func (c *Cache) DeleteItem(key string) {
	c.Lock()
	delete(c.m, key)
	c.Unlock()
}

//Fill adds data to chache
func (c *Cache) Fill(i interface{}) {
	switch v := i.(type) {
	case []User:
		for _, value := range v {
			c.Lock()
			c.m[value.Nickname] = value
			c.Unlock()
		}
	case []Article:
		for _, value := range v {
			c.Lock()
			c.m[value.ID] = value
			c.Unlock()
		}
	}
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
	NumRate      int    `form:"numRate" json:"numRate" bson:"numRate"`
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

//Comment represents comment's structure
type Comment struct {
	ArticleID    string `form:"articleId" json:"articleId" bson:"articleId"`
	Author       string `form:"author" json:"author" bson:"author"`
	Content      string `form:"content" json:"content" bson:"content"`
	CreationTime string `form:"creationTime" json:"creationTime" bson:"creationTime"`
	PrettyTime   string `form:"prettyTime" json:"prettyTime" bson:"prettyTime"`
}
