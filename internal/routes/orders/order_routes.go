package routes

import (
	handlers "Orderly/internal/handlers/orders"
	"github.com/gin-gonic/gin"
)

func SetupOrderRoutes(router *gin.Engine, orderHandler *handlers.OrderHandler) {
	orderRoutes := router.Group("/orders")
	{
		orderRoutes.POST("/", orderHandler.CreateOrder)
		orderRoutes.GET("/", orderHandler.GetAllOrders)
		orderRoutes.GET("/:id", orderHandler.GetOrderByID)
		orderRoutes.PUT("/:id", orderHandler.UpdateOrder)
		orderRoutes.DELETE("/:id", orderHandler.DeleteOrder)
	}
}
