package main

import (
	"net/http"
	"zrweb/golearn"
)

func init() {
	//initGorm()
	//initSqlx()
	//connectDB()
}

func main() {
	golearn.StringExample()
	//gopg.PgExample()
	//storeDataExample()
	//databaseAllExample()
	//formatExample()
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
