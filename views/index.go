package views

import (
	"html/template"
	"net/http"
	"time"

	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
)

func IndexView(route *gin.Engine) {
	app := ginview.NewMiddleware(goview.Config{
		Root:      "templates/",
		Extension: ".html",
		Master:    "layouts/master",
		Partials: []string{"partials/link",
			"partials/meta",
			"partials/title",
			"partials/scripts_head",
			"partials/scripts_foot"},
		Funcs: template.FuncMap{
			"copy": func() string {
				return time.Now().Format("2006")
			},
		},
		DisableCache: true,
	})

	appGroup := route.Group("/", app)
	{
		{
			appGroup.GET("/", func(ctx *gin.Context) {
				ginview.HTML(ctx, http.StatusOK, "index", gin.H{
					"title": "Hypi App | Hello World!",
				})
			})

		}
	}
}
