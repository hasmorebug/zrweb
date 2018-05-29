package main

import (
	"fmt"
	"github.com/gin-gonic/gin/json"
	"net/http"
)

//
func index(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Hello World! %s", r.URL.Path[1:])
}

///////// ResponseWriter /////////
//
type Post struct {
	User    string
	Threads []string
}

//
func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html><head><title>Go Web Programming</title></head><body><h1>Hello World</h1></body></html>`

	w.Write([]byte(str))
}

//
func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintln(w, "No such service, try next door")
}

//
func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "https://bing.com")
	w.WriteHeader(http.StatusFound)
}

//
func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := Post{
		User:    "Sau Sheong",
		Threads: []string{"first", "second", "third"},
	}
	json, _ := json.Marshal(post)
	w.Write(json)
}
