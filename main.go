package main

import (
	"net/http"
	"log"
)
func main() {
	const PORT = "8080"
	log.Print("Running on port " + PORT)

	// Handlers
	http.HandleFunc("/tasks/", RetrieveTasks)
	http.HandleFunc("/", HomeTodo)

	// Starting the server 
	log.Fatal(http.ListenAndServe(":"+PORT, nil))

}

// HomeTodo writes back the path to the user
func HomeTodo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.Path))
}

// RetrieveTasks gets a specific task
func RetrieveTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id := r.URL.Path[len("/tasks/"):]
		w.Write([]byte("Retrieving task number " + id))
	}
}