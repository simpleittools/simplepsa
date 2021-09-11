package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/simpleittools/simplepsa/database"
	"github.com/simpleittools/simplepsa/models"
	"gorm.io/gorm/clause"
)

func GetWorkOrders(c *fiber.Ctx) error {
	var workOrders []models.WorkOrder

	database.DB.Preload("Clients").Preload(clause.Associations).Find(&workOrders)

	return c.JSON(workOrders)
}

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

//func GetAllWorkOrder(c *fiber.Ctx) error {
//	var workOrder []models.WorkOrder
//
//	database.DB.Preload("Clients").Preload(clause.Associations).Find(&workOrder)
//
//	return c.JSON(workOrder)
//}
