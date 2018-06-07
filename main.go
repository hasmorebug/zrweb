package main

import (
	"net/http"
)

func init() {
	//initGorm()
	//initSqlx()
	//connectDB()
}

func main() {
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
