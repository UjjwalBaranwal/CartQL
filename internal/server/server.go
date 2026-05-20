// Package server provides the implementation of the gRPC server for the application.
package server

import (
	"net/http"

	"github.com/UjjwalBaranwal/CartQL/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

// Server represents the server for the application, containing the configuration, database connection, and logger.
type Server struct {
	config *config.Config
	db     *gorm.DB
	logger zerolog.Logger
}

// New creates and returns a new Server instance initialized with the provided configuration, database connection, and logger.
func New(cfg *config.Config, db *gorm.DB, logger zerolog.Logger) *Server {
	return &Server{
		config: cfg,
		db:     db,
		logger: logger,
	}
}

func (s *Server) SetupRoutes() *gin.Engine {
	router := gin.New()

	//Add Middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(s.corsMiddleware())

	//add routes
	router.GET("/health", s.healthCheck)

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
