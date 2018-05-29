package main

import (
	"fmt"
	"net/http"
)

func index(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Hello World! %s", r.URL.Path[1:])
}

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.Handle("/", index)
	server := &http.Server {
		Addr:		"0.0.0.0:8080",
		Handler:	mux,
	}

	server.ListenAndServe()
}