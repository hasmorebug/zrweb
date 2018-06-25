package zlearn

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

/////////
func HttpRouterExample() {
	hrExample3()
	//hrExample2()
	//hrExmaple1()
}

///////////
// Basic Authentication
func BasicAuth(h httprouter.Handle, requiredUser, requiredPassword string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		user, password, hasAuth := r.BasicAuth()

		if hasAuth && user == requiredUser && password == requiredPassword {
			h(w, r, ps)
		} else {
			w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}

func Protect(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "Protected!\n")
}

func hrExample3() {
	user := "gordon"
	pass := "secret!"

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/protected/", BasicAuth(Protect, user, pass))

	log.Fatal(http.ListenAndServe(":8080", router))
}

//////////
// Multi-domain/Sub-domain
type HostSwitch map[string]http.Handler

func (hs HostSwitch) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler := hs[r.Host]; handler != nil {
		handler.ServeHTTP(w, r)
	} else {
		http.Error(w, "Forbidden", 403)
	}
}

func hrExample2() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	hs := make(HostSwitch)
	hs["localhost:8080"] = router

	log.Fatal(http.ListenAndServe(":8080", hs))
}

/////////
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Not protected!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "Protected!\n")
}

func hrExmaple1() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	log.Fatal(http.ListenAndServe(":8080", router))
}

// default router
func startServer() {
	//initHandler()

	server := http.Server{
		Addr: ":8080",
	}

	server.ListenAndServe()
}
