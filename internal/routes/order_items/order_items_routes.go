package order_items_routes

import (
	"Orderly/internal/handlers/order_items"
	"github.com/gin-gonic/gin"
)

func SetupOrderItemsRoutes(router *gin.Engine, orderItemHandler *order_items.OrderItemHandler) {
	orderItemRoutes := router.Group("/order-items")
	{
		orderItemRoutes.POST("/", orderItemHandler.CreateOrderItem)
		orderItemRoutes.GET("/", orderItemHandler.GetAllOrderItems)
		orderItemRoutes.GET("/:id", orderItemHandler.GetOrderItemByID)
		orderItemRoutes.PUT("/:id", orderItemHandler.UpdateOrderItem)
		orderItemRoutes.DELETE("/:id", orderItemHandler.DeleteOrderItem)
	}
}
