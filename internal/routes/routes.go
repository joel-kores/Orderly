package routes

import (
	handlers3 "Orderly/internal/handlers/auth"
	Handlers "Orderly/internal/handlers/categories"
	"Orderly/internal/handlers/order_items"
	handlers2 "Orderly/internal/handlers/orders"
	handlers "Orderly/internal/handlers/products"
	routes2 "Orderly/internal/routes/auth"
	"Orderly/internal/routes/categories"
	orderitemsroutes "Orderly/internal/routes/order_items"
	routes "Orderly/internal/routes/orders"

	"Orderly/internal/routes/products"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, productHandler *handlers.ProductHandler,
	categoryHandler *Handlers.CategoryHandler, orderItemHandler *order_items.OrderItemHandler,
	orderHandler *handlers2.OrderHandler, authHandler *handlers3.AuthHandler) {
	products.SetupProductRoutes(router, productHandler)
	categories.SetupCategoryRoutes(router, categoryHandler)
	orderitemsroutes.SetupOrderItemsRoutes(router, orderItemHandler)
	routes.SetupOrderRoutes(router, orderHandler)
	routes2.SetupAuthRoutes(router, authHandler)
}
