package main

import (
	"fmt"
	"net/http"
)

func ServePages(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "GET":
		var path = r.URL.Path
		fmt.Println(path)

		if path == "/" {
			path = "../assets/pages/index.html"
		} else if path == "/learn" {
			path = "../assets/pages/learn.html"
		} else if path == "/about" {
			path = "../assets/pages/about.html"
		} else if path == "/quiz" {
			path = "../assets/pages/quiz.html"
		} else {
			path = ".." + path
		}

		http.ServeFile(w, r, path)
	}
}
