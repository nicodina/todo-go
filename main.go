package main

import (
	"fmt"
	"strings"
	"log"
	"net/http"
)

const port = "8080"

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	// Parse arguments
	r.ParseForm()
	
	// Let's print something on the server side
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
    	fmt.Println("key:", k)
    	fmt.Println("val:", strings.Join(v, ""))
	}

	// Send data back to the client
	fmt.Fprintf(w, "Hello Nico!")
}

func main() {

	log.Print("Running on port " + port)

	// Handlers
	http.Handle("/static/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/", sayhelloName)

	// Starting the server
	log.Fatal(http.ListenAndServe(":"+port, nil))

}


