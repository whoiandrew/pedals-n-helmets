package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	tokensCache = Cache{m: make(map[string]User)}
)

func login(c *gin.Context) {
	if valid, _ := authUser(c); valid {
		c.Redirect(http.StatusPermanentRedirect, "/")
	} else {
		c.HTML(http.StatusOK, "message.html", obj{
			"message": "Wrong login or passwrod",
		})
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
	id := c.Param("id")

	aRespBody := dbHTTPReq("/getArticleFromDB", url.Values{
		"id": {id},
	})

	cRespBody := dbHTTPReq("/getArticleCommentsFromDB", url.Values{
		"id": {id},
	})

	article, comments := Article{}, []Comment{}

	json.Unmarshal(aRespBody, &article)
	json.Unmarshal(cRespBody, &comments)

	if loggedInUsr, ok := tokensCache.getUser(c, "auth_cookie"); ok{
		c.HTML(
			http.StatusOK,
			"article.html",
			obj{"content": article.Content,
				"title":     article.Title,
				"nickname":  loggedInUsr.Nickname,
				"articleId": id,
				"comments":  comments,
			},
		)
	}
	c.Status(http.StatusOK)
}

func homePage(c *gin.Context) {
	usrFromDB := User{}
	articles := []Article{}

	articlesData := dbHTTPReq("/getArticlesFromDB", url.Values{})
	json.Unmarshal(articlesData, &articles)

	usr, isCookieSetted := tokensCache.getUser(c, "auth_cookie")

	if isCookieSetted {
		usrData := dbHTTPReq("/getUserFromDB", usr.GetURLValues())
		json.Unmarshal(usrData, &usrFromDB)
	}
	
	c.HTML(
		http.StatusOK,
		"index.html",
		obj{"payload": articles,
			"nickname": usrFromDB.Nickname,
			"isAdmin":  usrFromDB.IsAdmin,
			"isModer":  usrFromDB.IsModer,
		},
	)
}

func register(c *gin.Context) {
	usr := User{}
	c.Bind(&usr)

	usr.Password = ToHash(usr.Password)

	data := dbHTTPReq("/getUserFromDB", url.Values{"nickname": {usr.Nickname}})
	usrFromDB := User{}

	json.Unmarshal(data, &usrFromDB)

	if (usrFromDB == User{}) {
		dbHTTPReq("/addUserToDB", usr.GetURLValues())
		c.HTML(http.StatusOK, "message.html", obj{
			"message": fmt.Sprintf("New user %v found, welcome!", string(usr.Nickname)),
		})
	} else {
		c.HTML(http.StatusOK, "message.html", obj{
			"message": fmt.Sprintf("user %v is already exists", string(usr.Nickname)),
		})
	}
}

func addComment(c *gin.Context) {
	comment := Comment{} //req
	c.Bind(&comment)

	cTime := time.Now()
	comment.CreationTime = cTime.Format(time.RFC3339)
	comment.PrettyTime = cTime.Format("01-02-2006 15:04:05")

	dbHTTPReq("/addCommentToDB", comment.GetURLValues())

	c.JSON(http.StatusOK, obj{
		"author":     comment.Author,
		"content":    comment.Content,
		"prettyTime": comment.PrettyTime,
	})

}

func addArticle(c *gin.Context) {
	article := Article{}
	c.Bind(&article)

	respArticle := Article{}
	data := dbHTTPReq("/addArticleToDB", article.GetURLValues())

	json.Unmarshal(data, &respArticle)

	c.JSON(http.StatusOK, obj{"id": respArticle.ID, "prettyTime": respArticle.PrettyTime})
}

func editUserInfo(c *gin.Context) {
	user := User{}
	if usr, ok := tokensCache.getUser(c, "auth_cookie"); ok {
		c.Bind(&user)
		user.Nickname = usr.Nickname
		fmt.Println(user)
		dbHTTPReq("/updateUserFromDB", user.GetURLValues())
		c.Status(http.StatusOK)
	}
	c.Status(http.StatusNotFound)
}

func userPage(c *gin.Context) {
	if usr, ok := tokensCache.getUser(c, "auth_cookie"); ok {
		respBody := dbHTTPReq("/getUserFromDB", usr.GetURLValues())
		user := User{}
		json.Unmarshal(respBody, &user)
		c.HTML(
			http.StatusOK,
			"userPage.html",
			obj{
				"nickname":  user.Nickname,
				"firstname": user.Firstname,
				"lastname":  user.Lastname,
				"age":       user.Age,
				"bicycle":   user.Bicycle,
			},
		)
	}
	c.Status(http.StatusOK)
}

func deleteArticle(c *gin.Context) {
	id, _ := c.GetPostForm("id")
	dbHTTPReq("/deleteArticleFromDB", url.Values{"id": {id}})
	c.Status(http.StatusOK)
}

func editArticle(c *gin.Context) {
	id := c.Param("id")
	content, _ := c.GetPostForm("content")
	dbHTTPReq("/editArticleFromDB", url.Values{"id": {id}, "content": {content}})
	c.Status(http.StatusOK)
}

func rateArticle(c *gin.Context) {
	if usr, ok := tokensCache.getUser(c, "auth_cookie"); ok {
		id, _ := c.GetPostForm("id")
		rate, _ := c.GetPostForm("rate")
	
		dbHTTPReq("/updateRatesFromDB", url.Values{
			"id":       {id},
			"rate":     {rate},
			"nickname": {usr.Nickname},
		})
	}
	c.Status(http.StatusOK)
}
