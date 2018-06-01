package main

import (
	"net/http"
)

func init() {
	initGorm()
	//initSqlx()
	//connectDB()
}

func main() {
	ormExample2()
	//ormExample()
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
