package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/simpleittools/simplepsa/config"
	"github.com/simpleittools/simplepsa/database"
	"github.com/simpleittools/simplepsa/routes"
	"log"
)

func main() {
	database.Conn()
	app := fiber.New()

	port := config.GoDotEnvVariable("SERVER_PORT")

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Routing(app)

	log.Fatal(app.Listen(port))
}
