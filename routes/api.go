package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rafaelbreno/work-at-olist/app/handler"
)

var r *gin.Engine

func init() {
	r = gin.Default()

	routes()

	r.Run(":8080")
}

func routes() {
	authorRoutes()
}

func authorRoutes() {
	authorH := handler.GetAuthorHandlers()
	a := r.Group("/authors")
	{
		a.POST("/", authorH.ImportCSV)
		a.GET("/", authorH.FindAll)
		a.GET("/:id", authorH.FindById)
	}
}
