package routes

import (
	"updater-docker/api/handlers"
	"updater-docker/pkg/updater"

	"github.com/gofiber/fiber/v2"
)

// BookRouter is the Router for GoFiber App
func BookRouter(app fiber.Router, service updater.Service) {
	app.Post("/create", handlers.CreateDockerHandler(service))
	app.Post("/update", handlers.UpdateDocker(service))
}
