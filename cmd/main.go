package main

import (
	handlers "Orderly/internal/handlers/products"
	"Orderly/internal/models"
	repository "Orderly/internal/repositories/products"
	routes "Orderly/internal/routes/products"
	services "Orderly/internal/services/products"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "host=localhost user=postgres password=password dbname=test port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err = db.AutoMigrate(&models.Product{}); err != nil {
		log.Fatal(err)
	}

	productRepo := repository.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	router := gin.Default()
	routes.RegisterRoutes(router, productHandler)

	log.Println("Server is running on :8080")
	if err = router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
