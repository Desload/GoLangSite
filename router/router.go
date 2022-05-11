package router

import (
	"FirstSite/routes"
	"FirstSite/views"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	routes.NotFound(router)
	routes.NoMethods(router)

	views.IndexView(router)

	return router
}
