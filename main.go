package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	app.Post("/update", updateHandler)

	port := ":10000"
	fmt.Printf("🚀 Updater jalan di port %s\n", port)
	log.Fatal(app.Listen(port))
}

func updateHandler(c *fiber.Ctx) error {
	var req UpdateRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(ErrorResponse{
			Error: "Gagal membaca request body.",
		})
	}

	if req.NameDocker == "" {
		return c.Status(400).JSON(ErrorResponse{
			Error: "Nama image tidak boleh kosong.",
		})
	}

	fmt.Printf("🔄 Mulai update ke image: %s\n", req.NameDocker)

	commands := []string{
		"docker pull " + req.NameDocker,
		"docker compose up -d",
	}

	command := strings.Join(commands, " && ")

	cmd := exec.Command("/bin/sh", "-c", command)
	cmd.Dir = ".."

	output, err := cmd.CombinedOutput()
	if err != nil {
		return c.Status(500).JSON(ErrorResponse{
			Error: "Gagal update docker.",
		})
	}

	return c.Status(200).JSON(SuccesResponse{
		Message: "Berhasil update docker.",
		Output:  string(output),
	})
}
