package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/nicodina/memory"
	"github.com/nicodina/session"
)

const port = "8080"

var sessionManager *session.Manager

func init() {
	sessionManager, _ = session.NewManager("memory", "gosessionid", 3600)
	go sessionManager.GC()
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	// Parse arguments
	r.ParseForm()

	// Let's print something on the server side
	log.Println("path", r.URL.Path)
	log.Println("scheme", r.URL.Scheme)
	for k, v := range r.Form {
		log.Println("key:", k)
		log.Println("val:", strings.Join(v, ""))
	}

	// Send data back to the client
	fmt.Fprintf(w, "Hello Nico!")
}

func login(w http.ResponseWriter, r *http.Request) {

	log.Println("Login ", r.Method, " request ...")

	w.Header().Set("Content-Type", "text/html")
	sess := sessionManager.SessionStart(w, r)

	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.html")
		t.Execute(w, sess.Get("username"))
	} else {
		r.ParseForm()

		username := template.HTMLEscapeString(r.Form.Get("username"))
		password := template.HTMLEscapeString(r.Form.Get("password"))

		// Let's simply log data server side
		log.Println("Username: ", username)
		log.Println("Password: ", password)

		sess.Set("username", username)
		http.Redirect(w, r, "/", 302)
	}
}

func uploadFile(w http.ResponseWriter, r *http.Request) {

	log.Println("Upload ", r.Method, " request ...")

	if r.Method == "GET" {
		t, _ := template.ParseFiles("upload.html")
		t.Execute(w, createToken())
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

func createToken() string {

	crutime := time.Now().Unix()
	h := md5.New()

	io.WriteString(h, strconv.FormatInt(crutime, 10))

	return fmt.Sprintf("%x", h.Sum(nil))
}

func main() {

	log.Print("Running on port " + port)

	// Handlers
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", uploadFile)
	http.HandleFunc("/", sayhelloName)

	http.Handle("/static/", http.FileServer(http.Dir("public")))

	// Starting the server
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
