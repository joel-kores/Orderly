package main

import (
	"Orderly/docs"
	handlers3 "Orderly/internal/handlers"
	handlers4 "Orderly/internal/handlers/auth"
	Handlers "Orderly/internal/handlers/categories"
	"Orderly/internal/handlers/order_items"
	handlers2 "Orderly/internal/handlers/orders"
	handlers "Orderly/internal/handlers/products"
	"Orderly/internal/models"
	repositories "Orderly/internal/repositories/auth"
	"Orderly/internal/repositories/categories"
	orderitem "Orderly/internal/repositories/order_items"
	repository2 "Orderly/internal/repositories/orders"
	repository "Orderly/internal/repositories/products"
	"Orderly/internal/routes"
	services3 "Orderly/internal/services/auth"
	Services "Orderly/internal/services/categories"
	orderitemsservices "Orderly/internal/services/order_items"
	services2 "Orderly/internal/services/orders"
	services "Orderly/internal/services/products"
	"context"
	"github.com/coreos/go-oidc"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang.org/x/oauth2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	LoadEnv()
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

	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	redirectURL := os.Getenv("GOOGLE_REDIRECT_URL")
	providerURL := "https://accounts.google.com"

	// Initialize the OIDC provider
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, providerURL)
	if err != nil {
		log.Fatalf("Failed to initialize provider: %v", err)
	}

	// Configure the OAuth2 config
	oauth2Config := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
	}

	// Initialize the verifier
	verifier := provider.Verifier(&oidc.Config{ClientID: clientID})

	//repos
	userRepo := repositories.NewUserRepository(db)
	productRepo := repository.NewProductRepository(db)
	categoryRepo := categories.NewCategoryRepository(db)
	orderItemRepo := orderitem.NewOrderItemRepository(db)
	orderRepo := repository2.NewOrderRepository(db)

	//services
	userService := services3.NewAuthService(userRepo)
	productService := services.NewProductService(productRepo)
	categoryService := Services.NewCategoryService(categoryRepo)
	orderItemService := orderitemsservices.NewOrderItemService(orderItemRepo)
	orderService := services2.NewOrderService(orderRepo)

	//handlers
	userHandler := handlers4.NewAuthHandler(oauth2Config, verifier, userService)
	productHandler := handlers.NewProductHandler(productService)
	categoryHandler := Handlers.NewCategoryHandler(categoryService)
	orderItemHandler := order_items.NewOrderItemHandler(orderItemService)
	orderHandler := handlers2.NewOrderHandler(orderService)

	router := gin.Default()
	router.GET("/health", handlers3.Healthcheck)
	routes.SetupRoutes(router, productHandler, categoryHandler, orderItemHandler, orderHandler, userHandler)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Println("Server is running on :8080")
	if err = router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
