package main

import (
	"fmt"
	"log"
	"net/http"
)

var Port = ":4000"

func main() {
	http.HandleFunc("/", ServePages)
	fmt.Println("Listening @:", Port)
	log.Fatal(http.ListenAndServe(Port, nil))

}
