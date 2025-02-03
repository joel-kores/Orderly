package products

import (
	handlers "Orderly/internal/handlers/products"
	"github.com/gin-gonic/gin"
)

func SetupProductRoutes(router *gin.Engine, productHandler *handlers.ProductHandler) {
	productRoutes := router.Group("/products")
	{
		productRoutes.POST("/", productHandler.CreateProduct)
		productRoutes.GET("/", productHandler.GetAllProducts)
		productRoutes.GET("/:id", productHandler.GetProductByID)
		productRoutes.PUT("/:id", productHandler.UpdateProduct)
		productRoutes.DELETE("/:id", productHandler.DeleteProduct)
	}
}
