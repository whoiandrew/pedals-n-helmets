package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
)
//Controller represents controller's struct
type Controller struct {
	session *mgo.Session
}

//NewController initializes new mgov2 connection
func NewController() (*Controller, error) {
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		return nil, err
	}
	session.SetMode(mgo.Monotonic, true)
	return &Controller{
		session: session,
	}, nil
}

func (ctl *Controller) getArticle(id string) (res Article) {
	session := ctl.session.Clone()
	defer session.Close()
	col := session.DB("main").C("articles")
	err := col.FindId(id).One(&res)
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}

func (ctl *Controller) getArticleComments(articleID string) (res []Comment) {
	session := ctl.session.Clone()
	defer session.Close()
	col := session.DB("main").C("comments")
	col.Find(obj{"articleId": articleID}).All(&res)
	return
}

func (ctl *Controller) addUser(u User) {
	session := ctl.session.Clone()
	defer session.Close()
	col := session.DB("main").C("users")
	col.Insert(u)
}

func (ctl *Controller) getUser(id string) (u User) {
	session := ctl.session.Clone()
	defer session.Close()
	col := session.DB("main").C("users")
	fmt.Println("nickname  ", id)
	col.Find(obj{"_id": id}).One(&u)
	return
}

func (ctl *Controller) getAllUsers() (usrs []User) {
	session := ctl.session.Clone()
	defer session.Close()
	col := session.DB("main").C("users")
	col.Find(obj{}).All(&usrs)
	return
}

func (ctl *Controller) deleteUser(id string) {
	session := ctl.session.Clone()
	defer session.Close()
	col := session.DB("main").C("users")
	col.RemoveId(id)
}

func (ctl *Controller) getArticles() (articles []Article) {
	session := ctl.session.Clone()
	defer session.Close()
	col := session.DB("main").C("articles")
	err := col.Find(obj{}).All(&articles)
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}

func (ctl *Controller) addComment(comment Comment) {
	session := ctl.session.Clone()
	defer session.Close()
	col := session.DB("main").C("comments")
	col.Insert(comment)
}

func (ctl *Controller) addArticle(article Article) Article {
	session := ctl.session.Clone()
	defer session.Close()
	col := session.DB("main").C("articles")
	col.Insert(article)
	return article
}

func (ctl *Controller) updateUser(user User) {
	session := ctl.session.Clone()
	defer session.Close()
	col := session.DB("main").C("users")
	col.Update(
		obj{"_id": user.Nickname},
		obj{
			"$set": obj{
				"firstname": user.Firstname,
				"lastname":  user.Lastname,
				"age":       user.Age,
				"bicycle":   user.Bicycle,
			},
		})
}

func (ctl *Controller) updateRates(id, nickname, rate string) {
	session := ctl.session.Clone()
	defer session.Close()
	col := session.DB("main").C("articles")
	fmt.Println(id, nickname, rate)
	col.Update(obj{"_id": id}, obj{"$set": obj{fmt.Sprintf("rates.%v", nickname): rate}})
}

func (ctl *Controller) editArticle(id, content string) {
	session := ctl.session.Clone()
	defer session.Close()
	col := session.DB("main").C("articles")
	col.Update(obj{"_id": id}, obj{"$set": obj{"content": content}})
}

func (ctl *Controller) deleteArticle(id string) {
	session := ctl.session.Clone()
	defer session.Close()
	col := session.DB("main").C("articles")
	col.RemoveId(id)
}
