package main

import (
	"fmt"
	"strconv"
	"time"

	ai "github.com/night-codes/mgo-ai"

	"gopkg.in/mgo.v2"
)

//Process errors!!!!

func addUserToDB(s *mgo.Session, u User) (err error) {
	c := s.DB("main").C("users")
	err = c.Insert(u)
	return
}

func getUserFromDB(s *mgo.Session, nickname string) (res User) {
	c := s.DB("main").C("users")
	c.FindId(nickname).One(&res)
	return
}

func getArticlesFromDB(s *mgo.Session) (res []Article) {
	c := s.DB("main").C("articles")
	c.Find(nil).All(&res)
	return
}

func getArticleFromDB(s *mgo.Session, id string) (res Article) {
	c := s.DB("main").C("articles")
	c.FindId(id).One(&res)
	return
}

func getAllUsers(s *mgo.Session) (res []User) {
	c := s.DB("main").C("users")
	c.Find(nil).All(&res)
	return
}

func updateUserInDB(s *mgo.Session, u User) (err error) {
	c := s.DB("main").C("users")
	fmt.Println(u)
	err = c.Update(obj{"_id": u.Nickname}, obj{
		"$set": obj{
			"firstname": u.Firstname,
			"lastname":  u.Lastname,
			"age":       u.Age,
			"bicycle":   u.Bicycle,
		},
	})
	return
}

func deleteArticleFromDB(s *mgo.Session, articleId string) {
	c := s.DB("main").C("articles")
	c.RemoveId(articleId)
	
}

func deleteUserFromDB(s *mgo.Session, id string) {
	c := s.DB("main").C("users")
	c.RemoveId(id)
	fmt.Printf("%v has deleted", id)
}

func addCommentToDB(s *mgo.Session, comment Comment) Comment {
	cTime := time.Now()
	comment.CreationTime = cTime.Format(time.RFC3339)
	comment.PrettyTime = cTime.Format("01-02-2006 15:04:05")
	c := s.DB("main").C("comments")
	c.Insert(comment)
	return comment
}

func addArticleToDB(s *mgo.Session, article Article) Article {
	c := s.DB("main").C("articles")

	ai.Connect(s.DB("main").C("counters"))
	cTime := time.Now()
	article.ID = strconv.FormatUint(ai.Next("articles"), 10)
	article.CreationTime = cTime.Format(time.RFC3339)
	article.PrettyTime = cTime.Format("01-02-2006 15:04:05")
	c.Insert(article)
	return article
}

func getArticleComments(s *mgo.Session, articleID string) (comments []Comment) {
	c := s.DB("main").C("comments")
	c.Find(obj{"articleId": articleID}).All(&comments)
	return
}
