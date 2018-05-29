package main

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin/json"
	"io/ioutil"
	"net/http"
	"time"
)

//
func index(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Hello World! %s", r.URL.Path[1:])
}

///////// ResponseWriter /////////
//
type Post struct {
	User    string
	Threads []string
}

//
func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html><head><title>Go Web Programming</title></head><body><h1>Hello World</h1></body></html>`
	w.Write([]byte(str))
}

//
func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintln(w, "No such service, try next door")
}

//
func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "https://bing.com")
	w.WriteHeader(http.StatusFound)
}

//
func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := Post{
		User:    "Sau Sheong",
		Threads: []string{"first", "second", "third"},
	}
	js, _ := json.Marshal(post)
	w.Write(js)
}

/////////Cookie/////////
func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "Go Web Programming",
		HttpOnly: true,
	}

	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "Manning Publications Co",
		HttpOnly: true,
	}

	w.Header().Set("Set-Cookie", c1.String())
	w.Header().Add("Set-Cookie", c2.String())
}

func setCookie2(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "Go Web Programming",
		HttpOnly: true,
	}

	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "Manning Publications Co",
		HttpOnly: true,
	}

	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}

//
func getCookie(w http.ResponseWriter, r *http.Request) {
	h := r.Header["Cookie"]
	fmt.Fprintln(w, h)
}

//
func getCookie2(w http.ResponseWriter, r *http.Request) {
	c1, err := r.Cookie("first_cookie")
	if err != nil {
		fmt.Fprintln(w, "cannot get the first cookie")
	}

	c2 := r.Cookies()
	fmt.Fprintln(w, c1)
	fmt.Fprintln(w, c2)

}

//
func setMessage(w http.ResponseWriter, r *http.Request) {
	msg := []byte("Hello World")
	c := http.Cookie{
		Name:  "flash",
		Value: base64.URLEncoding.EncodeToString(msg),
	}

	http.SetCookie(w, &c)
}

func showMessage(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("flash")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintln(w, "No message found")
		}
	} else {
		rc := http.Cookie{
			Name:    "flash",
			MaxAge:  -1,
			Expires: time.Unix(1, 0),
		}
		http.SetCookie(w, &rc)
		val, _ := base64.URLEncoding.DecodeString(c.Value)
		fmt.Fprintln(w, string(val))
	}
}

/////////Request/////////
//
func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	fmt.Fprintln(w, h)
	fmt.Fprintln(w, h["Accept-Encoding"])
	fmt.Fprintln(w, h.Get("Accept-Encoding"))
}

func body(w http.ResponseWriter, r *http.Request) {
	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	fmt.Fprintln(w, body)
}

func process(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w, r.Form)
}

func process2(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w, r.PostForm)
}

func process3(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024)
	fmt.Fprintln(w, r.MultipartForm)
}

func process4(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024)
	fmt.Fprintln(w, r.MultipartForm)

	fileHeader := r.MultipartForm.File["uploaded"]
	if fileHeader != nil {
		file, err := fileHeader[0].Open()
		if err == nil {
			data, err := ioutil.ReadAll(file)
			if err == nil {
				fmt.Fprintln(w, data)
			}
		}
	} else {
		fmt.Fprintln(w, "upload error")
	}
}

func process5(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("uploaded")
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, data)
		}
	}
}
