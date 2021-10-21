package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/gofiber/fiber/v2"
)

func service(c *fiber.Ctx, service, command string) error {
	cmd := exec.Command("systemctl", command, service)
	out, err := cmd.CombinedOutput()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			fmt.Printf("systemctl finished with non-zero: %v\n", exitErr)
		} else {
			fmt.Printf("failed to run systemctl: %v", err)
			os.Exit(1)
		}
	}
	o := "ok"
	if len(string(out)) > 0 {
		o = string(out)[:len(string(out))-1]
	}

	return c.JSON(o)
}
