package routes

import (
	"github.com/Leonardo-lucas-pereira/Api_go_biblioteca/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		books := main.Group("books")
		{
			books.GET("/:id", controllers.ShowBook)
			books.GET("/", controllers.ShowBooks)
			books.POST("/", controllers.CreateBook)
			books.PUT("/:id", controllers.UpdateBook)
			books.DELETE("/:id", controllers.DeleteBook)
		}
	}
	return router
}
