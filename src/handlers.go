package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	ai "github.com/night-codes/mgo-ai"
	"gopkg.in/mgo.v2"
)

var (
	tokensCache = Cache{m: make(map[string]User)}
)

func login(c *gin.Context) {
	if valid := authUser(c); valid {
		c.Redirect(http.StatusPermanentRedirect, "/")
	} else {
		c.String(http.StatusOK, "Wrong login or password")
	}

}

func logout(c *gin.Context) {
	cookie, err := c.Cookie("auth_cookie")
	if err != nil || cookie != "NotSet" {
		cookie = "NotSet"
		c.SetCookie("auth_cookie", cookie, 3600, "/", "localhost", false, false)
	}
	c.Redirect(http.StatusPermanentRedirect, "/")
}

func fullArticle(c *gin.Context) {
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer session.Close()

	id := c.Param("id")
	article := getArticleFromDB(session, id)
	loggedInUsr, _ := tokensCache.getUser(c, "auth_cookie")

	c.HTML(
		http.StatusOK,
		"article.html",
		obj{"content": article.Content,
			"title":     article.Title,
			"nickname":  loggedInUsr.Nickname,
			"articleId": id,
			"comments":  getArticleComments(session, id),
		},
	)
}

func deleteUser(c *gin.Context) {
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer session.Close()

	id := c.Param("id")
	deleteUserFromDB(session, id)
}

func adminPanel(c *gin.Context) {
	loggedInUsr, _ := tokensCache.getUser(c, "auth_cookie")
	if loggedInUsr.IsAdmin {
		session, err := mgo.Dial("mongodb://localhost:27017")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer session.Close()

		users := getAllUsers(session)

		c.HTML(
			http.StatusOK,
			"adminPanel.html",
			obj{
				"users": users,
			},
		)
	} else {
		c.Status(http.StatusForbidden)
	}
}

func homePage(c *gin.Context) {
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer session.Close()

	usr, isCookieSetted := tokensCache.getUser(c, "auth_cookie")

	if isCookieSetted {
		usr = getUserFromDB(session, usr.Nickname)
	}
	c.HTML(
		http.StatusOK,
		"index.html",
		obj{"payload": getArticlesFromDB(session),
			"nickname": usr.Nickname,
			"isAdmin":  usr.IsAdmin,
		},
	)
}

func register(c *gin.Context) {
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	usr := User{}
	c.Bind(&usr)
	usr.Password = ToHash(usr.Password)
	usrFromDB := getUserFromDB(session, usr.Nickname)
	if (usrFromDB == User{}) {
		err = addUserToDB(session, usr)
		if err != nil {
			fmt.Println(err.Error())
		}
		c.String(http.StatusOK, fmt.Sprintf("New user %v found, welcome!", string(usr.Nickname)))
	} else {
		c.String(http.StatusOK, fmt.Sprintf("user %v is already exists", string(usr.Nickname)))
	}
}

func addComment(c *gin.Context) {
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer session.Close()
	comment := Comment{} //req
	c.Bind(&comment)
	respComment := addCommentToDB(session, comment)
	fmt.Printf("comment before%v", comment)
	fmt.Printf("comment after%v", respComment)
	c.JSON(http.StatusOK, obj{
		"author":     respComment.Author,
		"content":    respComment.Content,
		"prettyTime": respComment.PrettyTime,
	})

}

func addArticle(c *gin.Context) {
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer session.Close()

	ai.Connect(session.DB("main").C("counters"))
	article := Article{} //req
	c.Bind(&article)
	respArticle := addArticleToDB(session, article)
	fmt.Printf("Article %+v", respArticle)
	c.JSON(http.StatusOK, obj{"id": respArticle.ID, "prettyTime": respArticle.PrettyTime})

}

func editUserInfo(c *gin.Context) {
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	user := User{}

	usr, _ := tokensCache.getUser(c, "auth_cookie")

	user.Nickname = usr.Nickname
	c.Bind(&user)
	err = updateUserInDB(session, user)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.Status(http.StatusOK)
}

func userPage(c *gin.Context) {
	usr, _ := tokensCache.getUser(c, "auth_cookie")
	c.HTML(
		http.StatusOK,
		"userPage.html",
		obj{
			"nickname":  usr.Nickname,
			"firstname": usr.Firstname,
			"lastname":  usr.Lastname,
			"age":       usr.Age,
			"bicycle":   usr.Bicycle,
		},
	)
}
