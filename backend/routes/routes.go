package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/simpleittools/simplepsa/controllers"
)

func Routing(app *fiber.App) {

	workOrdersApi := app.Group("/api/workorder")
	workOrdersApi.Post("/create", controllers.CreateWorkOrder)
	workOrdersApi.Get("/", controllers.GetWorkOrders)

	clientsApi := app.Group("/api/client")
	clientsApi.Post("/create", controllers.CreateClient)
	clientsApi.Get("/", controllers.GetAllClients)
}
