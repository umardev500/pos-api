package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/pos-api/internal/container"
	"github.com/umardev500/pos-api/pkg"
)

const defaultPort = "8080"

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	if err := godotenv.Load(); err != nil {
		log.Error().Err(err).Msg("Error loading .env file")
	}
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	app := setupServer()
	gormDB := pkg.NewGormDB()
	v := pkg.NewValidator()
	containers := container.NewRegistryContainer(gormDB, v)
	pkg.NewRouter(app, containers).Handle()
	ch := make(chan error, 1)
	go func() {
		port := getPort()
		log.Info().Msg("Listening on port " + port)
		ch <- app.Listen(":" + port)
	}()

	select {
	case <-ctx.Done():
		log.Info().Msg("Shutting down server...")
		app.Shutdown()
	case err := <-ch:
		log.Error().Err(err).Msg("Server encountered an error")
	}
}

func setupServer() *fiber.App {
	return fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	return port
}
