package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/simpleittools/simplepsa/database"
	"github.com/simpleittools/simplepsa/models"
)

func CreateClient(c *fiber.Ctx) error {
	client := models.Client{}
	err := c.BodyParser(&client)

	if err != nil {
		return err
	}
	//date := time.Now()
	//dateString := date.Format("2006.01.02")

	//workOrderName := models.WorkOrder{WorkOrderName: dateString}

	database.DB.Create(&client)
	return c.JSON(client)
}

func GetAllClients(c *fiber.Ctx) error {
	var clients []models.Client
	database.DB.Find(&clients)
	return c.JSON(clients)
}
