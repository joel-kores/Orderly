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
	"fmt"
	"github.com/coreos/go-oidc"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang.org/x/oauth2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"math"
	"os"
	"time"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func ConnectToDB(dsn string) (*gorm.DB, error) {
	const maxRetries = 5
	var db *gorm.DB
	var err error

	for attempt := 1; attempt <= maxRetries; attempt++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("Database connection successful.")
			return db, nil
		}
		log.Printf("Database connection failed (attempt %d/%d): %v", attempt, maxRetries, err)
		sleepTime := time.Second * time.Duration(math.Pow(2, float64(attempt)))
		log.Printf("Retrying in %.0f seconds...", sleepTime.Seconds())
		time.Sleep(sleepTime)
	}

	return nil, fmt.Errorf("failed to connect to database after %d attempts: %w", maxRetries, err)
}

func main() {
	LoadEnv()
	docs.SwaggerInfo.Title = "Orderly API"
	docs.SwaggerInfo.Description = "This is a simple ecommerce API."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.BasePath = "/"

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	db, err := ConnectToDB(dsn)
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

	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.GET("/health", handlers3.Healthcheck)
	routes.SetupRoutes(router, productHandler, categoryHandler, orderItemHandler, orderHandler, userHandler)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Println("Server is running on port", addr)
	if err = router.Run(addr); err != nil {
		log.Fatal(err)
	}
}
