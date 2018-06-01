package main

import (
	"net/http"
)

func init() {
	//connectDB()
	initSqlx()
}

func main() {
	ormExample()
	//databaseExample2()
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
