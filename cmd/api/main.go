package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/d1joseph/palaestra/internal/adapters/api"
	"github.com/d1joseph/palaestra/internal/adapters/repository"
	"github.com/d1joseph/palaestra/internal/core/services"
)

func main() {
	// Create a new Fiber app
	app := fiber.New(fiber.Config{
		AppName: "Palaestra API",
	})

	// Add middleware
	app.Use(logger.New())
	app.Use(recover.New())

	// Initialize repository
	userRepo := repository.NewMemoryUserRepository()

	// Initialize service
	userService := services.NewUserService(userRepo)

	// Initialize API handler
	apiHandler := api.NewAPIHandler(userService)

	// Register API handler with the server
	api.RegisterHandlers(app, apiHandler)

	// Start the server in a goroutine
	go func() {
		log.Println("Starting server on :8080")
		if err := app.Listen(":8080"); err != nil {
			log.Printf("Server error: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	if err := app.Shutdown(); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}
	log.Println("Server gracefully stopped")
}
