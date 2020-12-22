package main

import (
	"net/http"
	"log"
)
func main() {

	const PORT = "8080"
	log.Fatal(http.ListenAndServe(":"+PORT, nil))

}