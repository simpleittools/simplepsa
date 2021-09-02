package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/simpleittools/simplepsa/database"
	"github.com/simpleittools/simplepsa/models"
)

// CreateWorkOrder is to create the work orders
func CreateWorkOrder(c *fiber.Ctx) error {
	workOrder := models.WorkOrder{}
	err := c.BodyParser(&workOrder)

	if err != nil {
		return err
	}
	//date := time.Now()
	//dateString := date.Format("2006.01.02")

	//workOrderName := models.WorkOrder{WorkOrderName: dateString}

	database.DB.Create(&workOrder)
	return c.JSON(workOrder)
}
