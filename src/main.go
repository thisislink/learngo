package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

//var Port = ":4000" //local test

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Println("$PORT must be set, defaulting to 9000")
		port = "9000"
	}

	/*
		http.HandleFunc("/", ServePages)
		fmt.Println("Listening @:", port)
		log.Println(http.ListenAndServe(":"+port, nil))
	*/
	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("../assets/pages/*.html")
	router.Static("../assets/pages", "pages")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.Run(":" + port)
}
