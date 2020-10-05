package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

func authUser(c *gin.Context) (valid bool, resUser User) {
	usr := User{}
	c.Bind(&usr)
	respBody := dbHTTPReq("/getUserFromDB", usr.GetURLValues())

	usrFromDB := User{}
	json.Unmarshal(respBody, &usrFromDB)
	fmt.Printf("logged in: %+v", usrFromDB)

	if (usrFromDB != User{}) && CheckPasswordHash(usr.Password, usrFromDB.Password) {
		loggedInUsr := usrFromDB
		cookie, err := c.Cookie("auth_cookie")
		if _, ok := tokensCache.m[cookie]; !ok || err != nil || cookie == "NotSet" {
			cookie = ToHash(string(loggedInUsr.Nickname))
			tokensCache.Lock()
			tokensCache.m[cookie] = loggedInUsr
			tokensCache.Unlock()
			c.SetCookie("auth_cookie", cookie, 3600, "/", "localhost", false, false)
		}
		fmt.Printf("Cookie value: %s \n", cookie)
		fmt.Println(tokensCache.m)
		valid = true
		resUser = loggedInUsr
	}
	return

}
