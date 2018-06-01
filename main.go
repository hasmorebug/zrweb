package main

import (
	"net/http"
)

func init() {
	connectDB()
}


func main() {
	//databaseExample()
	//xmlFunc()
	//jsonFunc()
	//startServer()
}

//
func startServer() {
	//initHandler()

	server := http.Server{
		Addr: ":8080",
	}

	server.ListenAndServe()
}
