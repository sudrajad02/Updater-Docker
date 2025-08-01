package presenter

import "github.com/gofiber/fiber"

type CreateRequest struct {
	Path       string `json:"path"`
	ClientName string `json:"client_name"`
}

type UpdaterRequest struct {
	NameDocker string `json:"name_docker"`
	Path       string `json:"path"`
}

func UpdaterSuccessResponse(output string) *fiber.Map {
	return &fiber.Map{
		"status":  true,
		"message": "Docker berhasil di start/update",
		"output":  output,
	}
}

func UpdaterErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"error":  err.Error(),
	}
}
