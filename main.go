package main

import (
	"net/http"
	"log"
)
func main() {
	const PORT = "8080"
	log.Print("Running on port " + PORT)

	// Handlers
	http.HandleFunc("/", HomeTodo)

	// Starting the server 
	log.Fatal(http.ListenAndServe(":"+PORT, nil))

}

// HomeTodo writes back the path to the user
func HomeTodo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.Path))
}