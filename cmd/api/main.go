// Package main is the entry point for the CartQL API server.
package main

import (
	"github.com/UjjwalBaranwal/CartQL/internal/config"
	"github.com/UjjwalBaranwal/CartQL/internal/database"
	"github.com/UjjwalBaranwal/CartQL/internal/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	// fmt.Println("Hello, World!")
	log := logger.New()
	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}
	db, err := database.New(&cfg.Database)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}
	mainDB, err := db.DB()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get database connection")
	}
	defer mainDB.Close()
	log.Info().Msg("Database connection established")
	gin.SetMode(cfg.Server.GinMode)
	log.Info().Msg("Starting Server")
}
