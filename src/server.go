package main

import (
	"fmt"
	"net/http"
)

func ServePages(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "text/html")

	switch r.Method {

	case "GET":
		var path = r.URL.Path
		fmt.Println(path)

		if path == "/" {
			path = "../website/index.html"
		} else if path == "/learn" {
			path = "../website/learn.html"
		} else if path == "/about" {
			path = "../website/about.html"
		} else if path == "/quiz" {
			path = "../website/quiz.html"
		} else {
			path = ".." + path
		}

		http.ServeFile(w, r, path)
	}
}
