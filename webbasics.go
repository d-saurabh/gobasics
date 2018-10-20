package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>hello there</h1>")
	fmt.Fprintf(w, "<p>whoa</p>")

}
func main() {

	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":8000", nil)

}
