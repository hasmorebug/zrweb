package main

import (
	"net/http"
	"zrweb/gopg"
)

func init() {
	//initGorm()
	//initSqlx()
	//connectDB()
}

func main() {
	gopg.QueryExample()
	//gopg.PgExample()
	//golearn.GoInterfaceExample()
	//golearn.StringExample()
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
