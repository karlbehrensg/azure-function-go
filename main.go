package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	listenAddr := ":3000"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}

	app := fiber.New()

	app.Get("/api/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(listenAddr))
}
