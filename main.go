package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLFiles("HTML/mainPage.html")
	router.Static("/css", "CSS")

	router.GET("/", func(context *gin.Context) {
		context.HTML(200, "mainPage.html", gin.H{
			"title": "Main page",
		})
	})

	router.Run()
}
