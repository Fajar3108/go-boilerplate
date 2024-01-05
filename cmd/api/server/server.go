package server

import (
	"github.com/Fajar3108/go-boilerplate/internal/config"
	"github.com/Fajar3108/go-boilerplate/internal/database"
	"github.com/Fajar3108/go-boilerplate/internal/http/routes"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

func NewApp() error {
	_, err := database.SetupDatabaseConnection()
	if err != nil {
		return err
	}

	app := fiber.New(fiber.Config{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  10 * time.Second,
	})

	routes.SetupRoutes(app)

	err = StartServer(app)
	if err != nil {
		return err
	}

	return nil
}

func StartServer(app *fiber.App) error {
	port := strconv.Itoa(config.AppConfig.Port)

	err := app.Listen(":" + port)
	if err != nil {
		return err
	}

	return nil
}
