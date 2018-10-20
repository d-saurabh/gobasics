package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type SomeStruct struct {
	Age int
	Name string
}
func indexhandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w,"<h1>hello Go</h1>")
}
func somehandler(w http.ResponseWriter, r *http.Request)  {
	p := SomeStruct{Name:"saurabh",Age:30}
	t,_ := template.ParseFiles("htmltemplate.html")
	t.Execute(w,p)
}
func main()  {
	http.HandleFunc("/",indexhandler)
	http.HandleFunc("/someend/",somehandler)
	http.ListenAndServe(":8000",nil)
}
