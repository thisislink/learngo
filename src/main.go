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

	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	router.Use(gin.Logger())
	router.Static("/website", "../website")

	router.LoadHTMLGlob("../website/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.html", nil)
	})
	router.GET("/learn", func(c *gin.Context) {
		c.HTML(http.StatusOK, "learn.html", nil)
	})
	router.GET("/quiz", func(c *gin.Context) {
		c.HTML(http.StatusOK, "quiz.html", nil)
	})

	router.Run(":" + port)
}
