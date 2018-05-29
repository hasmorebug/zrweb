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
	http.HandleFunc("/response/example1", writeExample)
	http.HandleFunc("/response/example2", writeHeaderExample)
	http.HandleFunc("/response/example3", headerExample)
	http.HandleFunc("/response/example4", jsonExample)
	/////////

	server.ListenAndServe()
}
