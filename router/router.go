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

	// System routes
	routes.NotFound(router)
	routes.NoMethods(router)
	routes.Deals(router)

	// Static Routes
	routes.PublicCss(router)
	routes.PublicImages(router)

	// Index view
	views.IndexView(router)

	return router
}
