package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

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
