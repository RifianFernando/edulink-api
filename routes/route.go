package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/skripsi-be/controllers"
	// "github.com/skripsi-be/middleware"
)

func Route(router *gin.Engine) {
	// Initialize Version
	apiV1 := router.Group("/api/v1")
	{

		// Student CRUD
		product := apiV1.Group("/student")
		{
			// product.GET("/", controllers.GetAllStudent) // every product in every store
			// product.GET("/:id", controllers.GetStudent)
			product.POST("/create", controllers.CreateStudent)
			// product.PUT("/update/:id", middleware.HaveStore(), routes.UpdateProduct)
			// product.DELETE("/delete/:id", middleware.HaveStore(), routes.DeleteProduct)
		}
	}
}
