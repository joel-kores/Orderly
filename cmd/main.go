package main

import (
	"Orderly/docs"
	handlers3 "Orderly/internal/handlers"
	Handlers "Orderly/internal/handlers/categories"
	"Orderly/internal/handlers/order_items"
	handlers2 "Orderly/internal/handlers/orders"
	handlers "Orderly/internal/handlers/products"
	"Orderly/internal/models"
	"Orderly/internal/repositories/categories"
	orderitem "Orderly/internal/repositories/order_items"
	repository2 "Orderly/internal/repositories/orders"
	repository "Orderly/internal/repositories/products"
	"Orderly/internal/routes"
	Services "Orderly/internal/services/categories"
	orderitemsservices "Orderly/internal/services/order_items"
	services2 "Orderly/internal/services/orders"
	services "Orderly/internal/services/products"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	docs.SwaggerInfo.Title = "Orderly API"
	docs.SwaggerInfo.Description = "This is a simple ecommerce API."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.BasePath = "/"

	dsn := "host=localhost user=admin password=password dbname=testOrderly port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err = db.AutoMigrate(&models.Product{}, &models.User{}, models.Category{}, models.Order{},
		models.OrderItem{}); err != nil {
		log.Fatal(err)
	}

	//repos
	productRepo := repository.NewProductRepository(db)
	categoryRepo := categories.NewCategoryRepository(db)
	orderItemRepo := orderitem.NewOrderItemRepository(db)
	orderRepo := repository2.NewOrderRepository(db)

	//services
	productService := services.NewProductService(productRepo)
	categoryService := Services.NewCategoryService(categoryRepo)
	orderItemService := orderitemsservices.NewOrderItemService(orderItemRepo)
	orderService := services2.NewOrderService(orderRepo)

	//handlers
	productHandler := handlers.NewProductHandler(productService)
	categoryHandler := Handlers.NewCategoryHandler(categoryService)
	orderItemHandler := order_items.NewOrderItemHandler(orderItemService)
	orderHandler := handlers2.NewOrderHandler(orderService)

	router := gin.Default()
	router.GET("/health", handlers3.Healthcheck)
	routes.SetupRoutes(router, productHandler, categoryHandler, orderItemHandler, orderHandler)
	
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Println("Server is running on :8080")
	if err = router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
