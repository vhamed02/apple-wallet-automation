package main

import (
	"log"
	"os"

	"github.com/apple-wallet-automation/internal/categorizer"
	"github.com/apple-wallet-automation/internal/config"
	"github.com/apple-wallet-automation/internal/handler"
	"github.com/apple-wallet-automation/internal/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "config.yml"
	}

	cfg, err := config.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	store, err := storage.New(cfg.Storage.DataDir)
	if err != nil {
		log.Fatalf("failed to initialize storage: %v", err)
	}

	cat := categorizer.New(cfg.Categories)

	recordHandler := handler.NewRecordHandler(cfg, store, cat)

	app := fiber.New(fiber.Config{
		AppName: "Apple Wallet Automation",
	})

	app.Use(recover.New())
	app.Use(logger.New())

	app.Post("/record/", recordHandler.Handle)
	app.Post("/record", recordHandler.Handle)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("server starting on port %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
