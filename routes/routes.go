package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func NotFound(route *gin.Engine) {
	route.NoRoute(func(c *gin.Context) {
		c.AbortWithStatusJSON(404, "Not Found")
	})
}

func NoMethods(route *gin.Engine) {
	route.NoMethod(func(c *gin.Context) {
		c.AbortWithStatusJSON(405, "not allowed")
	})
}

func Deals(route *gin.Engine) {
	route.GET("/home", func(c *gin.Context) {
		fmt.Println(("Here"))
	})
}

func PublicImages(route *gin.Engine) {
	route.Static("/public/images", "./public/images")
}

func PublicCss(route *gin.Engine) {
	route.Static("/public/css", "./public/css")
}
