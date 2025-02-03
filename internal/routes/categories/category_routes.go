package categories

import (
	Handlers "Orderly/internal/handlers/categories"
	"github.com/gin-gonic/gin"
)

func SetupCategoryRoutes(router *gin.Engine, categoryHandler *Handlers.CategoryHandler) {
	categoryGroup := router.Group("/categories")
	{
		categoryGroup.POST("/", categoryHandler.CreateCategory)
		categoryGroup.GET("/:id", categoryHandler.GetCategoryByID)
		categoryGroup.GET("/", categoryHandler.GetAllCategories)
		categoryGroup.PUT("/:id", categoryHandler.UpdateCategory)
		categoryGroup.DELETE("/:id", categoryHandler.DeleteCategory)
	}
}
