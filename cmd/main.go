package main

import (
	Handlers "Orderly/internal/handlers/categories"
	"Orderly/internal/handlers/order_items"
	handlers "Orderly/internal/handlers/products"
	"Orderly/internal/models"
	"Orderly/internal/repositories/categories"
	orderitem "Orderly/internal/repositories/order_items"
	repository "Orderly/internal/repositories/products"
	"Orderly/internal/routes"
	orderitemsservices "Orderly/internal/services/order_items"

	Services "Orderly/internal/services/categories"
	services "Orderly/internal/services/products"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
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

	//services
	productService := services.NewProductService(productRepo)
	categoryService := Services.NewCategoryService(categoryRepo)
	orderItemService := orderitemsservices.NewOrderItemService(orderItemRepo)

	//handlers
	productHandler := handlers.NewProductHandler(productService)
	categoryHandler := Handlers.NewCategoryHandler(categoryService)
	orderItemHandler := order_items.NewOrderItemHandler(orderItemService)

	router := gin.Default()
	routes.SetupRoutes(router, productHandler, categoryHandler, orderItemHandler)
	log.Println("Server is running on :8080")
	if err = router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
