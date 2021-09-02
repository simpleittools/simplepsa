package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/simpleittools/simplepsa/controllers"
)

func Routing(app *fiber.App) {
	app.Post("/api/workorder/create", controllers.CreateWorkOrder)
}
