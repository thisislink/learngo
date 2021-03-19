package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

//var Port = ":4000" //local test

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		log.Println("$PORT must be set, defaulting to 9000")
		port = "9000"
	}

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", ServePages)
	router.HandleFunc("/learn", ServePages)
	router.HandleFunc("/about", ServePages)
	router.HandleFunc("/quiz", ServePages)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
