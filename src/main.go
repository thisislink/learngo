package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

//var Port = ":4000" //local test

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Println("$PORT must be set, defaulting to 4000")
		port = ":4000"
	}
	http.HandleFunc("/", ServePages)
	fmt.Println("Listening @:", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
