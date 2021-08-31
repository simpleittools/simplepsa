package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	app := fiber.New()

	port := goDotEnvVariable("PORT")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	dbType := goDotEnvVariable("DBTYPE")

	fmt.Println(dbType)

	log.Fatal(app.Listen(port))
}
