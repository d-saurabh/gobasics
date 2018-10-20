package main

import (
	"io/ioutil"
	"net/http"
)

func main() {

	response, _ := http.Get("https://timesofindia.indiatimes.com/sitemap.cms")
	bytes, _ := ioutil.ReadAll(response.Body)
	string_body := string(bytes)
	println(string_body)
	response.Body.Close()

}
