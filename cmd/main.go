package main

import (
	"encoding/json"
	"fmt"
	"log"
	"updater-docker/api/routes"
	"updater-docker/pkg/updater"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type UpdateRequest struct {
	NameDocker string `json:"name_docker"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccesResponse struct {
	Message string `json:"message"`
	Output  string `json:"output"`
}

func main() {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(cors.New())

	app.Use(logger.New())

	updateDockerService := updater.NewService()
	routes.BookRouter(app, updateDockerService)

	port := ":10000"
	fmt.Printf("🚀 Updater jalan di port %s\n", port)
	log.Fatal(app.Listen(port))
}
