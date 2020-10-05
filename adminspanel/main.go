package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("/home/whoiandrew/golang-course/golang-homeworks1/golang-gin/adminspanel/templates/*")
	r.Static("/assets", "/home/whoiandrew/golang-course/golang-homeworks1/golang-gin/adminspanel/assets/")
	r.POST("/", adminPanel)
	r.POST("/deleteUser/:id", deleteUser)
	r.Run(":8083")

}

func dbHTTPReq(endp string, data url.Values) (body []byte) {
	domain := "http://localhost:8082"
	resp, err := http.PostForm(domain+endp, data)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return
}

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

func adminPanel(c *gin.Context) {
	users := []User{}
	data := dbHTTPReq("/getAllUsersFromDB", url.Values{})
	json.Unmarshal(data, &users)
	c.HTML(
		http.StatusOK,
		"adminPanel.html",
		obj{
			"users": users,
		},
	)
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")
	dbHTTPReq("/deleteUserFromDB", url.Values{
		"id": {id},
	})
}
