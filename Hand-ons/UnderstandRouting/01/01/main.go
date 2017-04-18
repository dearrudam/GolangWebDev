package main

import (
	"net/http"
	"io"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Here's the index page")
}
func dog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Here's the dog page")
}
func me(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello Max")
}
