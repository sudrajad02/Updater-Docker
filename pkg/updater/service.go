package updater

import (
	"fmt"
	"os/exec"
	"strings"
	"updater-docker/api/presenter"
)

type Service interface {
	CreateDocker(req presenter.CreateRequest) (string, error)
	UpdateDocker(req presenter.UpdaterRequest) ([]byte, error)
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) CreateDocker(req presenter.CreateRequest) (string, error) {
	// Prepare the command
	cmd := exec.Command("docker", "compose", "up", "-d")
	cmd.Dir = req.Path

	// Capture stdout and stderr
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", fmt.Errorf("error creating stdout pipe: %w", err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return "", fmt.Errorf("error creating stderr pipe: %w", err)
	}

	// Start the command
	if err := cmd.Start(); err != nil {
		return "", fmt.Errorf("error starting command: %w", err)
	}

	// Output stdout
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := stdout.Read(buf)
			if err != nil {
				break
			}
			fmt.Printf("🟢 Docker STDOUT: %s", buf[:n])
		}
	}()

	// Output stderr
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := stderr.Read(buf)
			if err != nil {
				break
			}
			fmt.Printf("🔴 Docker STDERR: %s", buf[:n])
		}
	}()

	// Wait for the command to finish
	if err := cmd.Wait(); err != nil {
		return "", fmt.Errorf("docker compose exited with error: %w", err)
	}

	// Return success message
	successMessage := fmt.Sprintf("✅ Company \"%s\" is now running\n", req.ClientName)
	return successMessage, nil
}

func (s *service) UpdateDocker(req presenter.UpdaterRequest) ([]byte, error) {
	commands := []string{
		"docker pull " + req.NameDocker,
		"docker compose up -d",
	}

	command := strings.Join(commands, " && ")

	cmd := exec.Command("/bin/sh", "-c", command)
	cmd.Dir = ".."

	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	return output, nil
}
