package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"updater-docker/api/presenter"
	"updater-docker/pkg/updater"

	"github.com/gofiber/fiber/v2"
)

func CreateDockerHandler(service updater.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req presenter.CreateRequest
		err := c.BodyParser(&req)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.UpdaterErrorResponse(errors.New("gagal membaca request body")))
		}

		if req.Path == "" {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UpdaterErrorResponse(errors.New("path tidak boleh kosong.")))
		}

		if req.ClientName == "" {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UpdaterErrorResponse(errors.New("client name tidak boleh kosong.")))
		}

		fmt.Printf("🔄 Mulai membuat ke image: %s\n", req.ClientName)
		crt, err := service.CreateDocker(req)

		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UpdaterErrorResponse(err))
		}

		return c.JSON(presenter.UpdaterSuccessResponse(string(crt)))
	}
}

func UpdateDocker(service updater.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req presenter.UpdaterRequest
		err := c.BodyParser(&req)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.UpdaterErrorResponse(errors.New("gagal membaca request body")))
		}

		if req.Path == "" {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UpdaterErrorResponse(errors.New("path tidak boleh kosong.")))
		}

		if req.NameDocker == "" {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UpdaterErrorResponse(errors.New("nama image tidak boleh kosong.")))
		}

		fmt.Printf("🔄 Mulai update ke image: %s\n", req.NameDocker)
		upd, err := service.UpdateDocker(req)

		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UpdaterErrorResponse(err))
		}

		return c.JSON(presenter.UpdaterSuccessResponse(string(upd)))
	}
}
