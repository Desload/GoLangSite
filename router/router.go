package router

import (
	"FirstSite/routes"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	routes.NotFound(router)
	routes.NoMethods(router)

	return router
}
