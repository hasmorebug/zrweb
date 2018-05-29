package main

import (
	"net/http"
)

func main() {
	server := http.Server{
		Addr: ":8080",
	}

	/////////
	http.HandleFunc("/", index)
	/////////
	http.HandleFunc("/cookie/example1", setCookie)
	http.HandleFunc("/cookie/example2", getCookie)
	http.HandleFunc("/cookie/example3", getCookie2)
	http.HandleFunc("/cookie/example4", setMessage)
	http.HandleFunc("/cookie/example5", showMessage)
	//////////

	/////////
	http.HandleFunc("/response/example1", writeExample)
	http.HandleFunc("/response/example2", writeHeaderExample)
	http.HandleFunc("/response/example3", headerExample)
	http.HandleFunc("/response/example4", jsonExample)
	/////////

	server.ListenAndServe()
}
