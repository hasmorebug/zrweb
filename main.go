package main

import (
	"net/http"
)

func main() {
	jsonFunc()
	//startServer()
}

func jsonFunc() {
	encodeJson()
	//createJson()
	//decodeJson()
	//parseJson()
}

//
func startServer() {
	initHandler()

	server := http.Server{
		Addr: ":8080",
	}

	server.ListenAndServe()
}

//
func initHandler() {
	files := http.FileServer(http.Dir("./template"))

	http.Handle("/static/", http.StripPrefix("/static/", files))

	/////////
	http.HandleFunc("/", index)

	//////////
	http.HandleFunc("/request/example1", headers)
	http.HandleFunc("/request/example2", body)
	http.HandleFunc("/request/example3", process)
	http.HandleFunc("/request/example4", process2)
	http.HandleFunc("/request/example5", process3)
	http.HandleFunc("/request/example6", process4)
	http.HandleFunc("/request/example7", process5)
	//////////

	/////////
	http.HandleFunc("/cookie/example1", setCookie)
	http.HandleFunc("/cookie/example2", setCookie2)
	http.HandleFunc("/cookie/example3", getCookie)
	http.HandleFunc("/cookie/example4", getCookie2)
	http.HandleFunc("/cookie/example5", setMessage)
	http.HandleFunc("/cookie/example6", showMessage)
	//////////

	/////////
	http.HandleFunc("/response/example1", writeExample)
	http.HandleFunc("/response/example2", writeHeaderExample)
	http.HandleFunc("/response/example3", headerExample)
	http.HandleFunc("/response/example4", jsonExample)
	/////////
}
