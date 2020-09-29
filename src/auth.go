package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

func authUser(c *gin.Context) (valid bool) {
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	usr := User{}
	c.Bind(&usr)

	fmt.Printf("logged in: %+v", usr)
	usrFromDB := getUserFromDB(session, usr.Nickname)

	if (usrFromDB != User{}) && CheckPasswordHash(usr.Password, usrFromDB.Password) {
		loggedInUsr := usrFromDB
		cookie, err := c.Cookie("auth_cookie")
		if _, ok := tokensCache.m[cookie]; (!ok ||err != nil || cookie == "NotSet") {
			cookie = ToHash(string(loggedInUsr.Nickname))
			tokensCache.Lock()
			tokensCache.m[cookie] = loggedInUsr
			tokensCache.Unlock()
			c.SetCookie("auth_cookie", cookie, 3600, "/", "localhost", false, false)
		}

		fmt.Printf("Cookie value: %s \n", cookie)
		fmt.Println(tokensCache.m)
		valid = true
	}
	return

}
