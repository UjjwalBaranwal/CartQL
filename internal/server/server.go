// Package server provides the implementation of the gRPC server for the application.
package server

import (
	"net/http"

	"github.com/UjjwalBaranwal/CartQL/internal/config"
	"github.com/UjjwalBaranwal/CartQL/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

// Server represents the server for the application, containing the configuration, database connection, and logger.
type Server struct {
	config         *config.Config
	db             *gorm.DB
	logger         *zerolog.Logger
	authService    *services.AuthService
	productService *services.ProductService
	userService    *services.UserService
	uploadService  *services.UploadService
}

// New creates and returns a new Server instance initialized with the provided configuration, database connection, and logger.
func New(
	cfg *config.Config,
	db *gorm.DB,
	logger *zerolog.Logger,
	authService *services.AuthService,
	productService *services.ProductService,
	userService *services.UserService,
	uploadService *services.UploadService) *Server {
	return &Server{
		config:         cfg,
		db:             db,
		logger:         logger,
		authService:    authService,
		productService: productService,
		userService:    userService,
		uploadService:  uploadService,
	}
}

// SetupRoutes sets up the HTTP routes for the server using the Gin framework and returns the configured router.
func (s *Server) SetupRoutes() *gin.Engine {
	router := gin.New()

	// Add Middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(s.corsMiddleware())

	// add routes
	router.GET("/health", s.healthCheck)
	router.Static("/uploads", "./uploads")
	api := router.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", s.register)
			auth.POST("/login", s.login)
			auth.POST("/refresh", s.refreshToken)
			auth.POST("/logout", s.logout)
		}
		protected := api.Group("/")
		protected.Use(s.authMiddleware())
		{
			// User routes
			users := protected.Group("/users")
			{
				userRoutes := users
				userRoutes.GET("/profile", s.getProfile)
				userRoutes.PUT("/profile", s.updateProfile)
			}

			// Category routes
			categories := protected.Group("/categories")
			{
				categoryRoutes := categories
				categoryRoutes.POST("/", s.adminMiddleware(), s.createCategory)
				categoryRoutes.PUT("/:id", s.adminMiddleware(), s.updateCategory)
				categoryRoutes.DELETE("/:id", s.adminMiddleware(), s.deleteCategory)
			}
			// product routes
			products := protected.Group("/products")
			{
				productRoutes := products
				productRoutes.POST("/", s.adminMiddleware(), s.createProduct)
				productRoutes.PUT("/:id", s.adminMiddleware(), s.updateProduct)
				productRoutes.DELETE("/:id", s.adminMiddleware(), s.deleteProduct)
				productRoutes.POST("/:id/images", s.adminMiddleware(), s.uploadProductImage)
			}
		}
		// public routes
		api.GET("/categories", s.getCategories)
		api.GET("/products", s.getProducts)
		api.GET("/products/:id", s.getProduct)
	}
	return router
}

func (s *Server) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (s *Server) corsMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
