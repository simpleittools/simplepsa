package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/simpleittools/simplepsa/config"
	"github.com/simpleittools/simplepsa/database"
	"log"
)

func main() {
	database.Conn()
	app := fiber.New()

	port := config.GoDotEnvVariable("SERVER_PORT")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(port))
}
