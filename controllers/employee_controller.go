package controllers

import (
	"go-crud-postgres/config"
	"go-crud-postgres/models"

	"github.com/gofiber/fiber/v2"
)

func GetEmployees(c *fiber.Ctx) error {
	var employees []models.Employee
	if err := config.DB.Find(&employees).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.JSON(employees)
}

func GetEmployee(c *fiber.Ctx) error {
	id := c.Params("id")
	var employee models.Employee
	if err := config.DB.First(&employee, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err.Error())
	}
	return c.JSON(employee)
}

func CreateEmployee(c *fiber.Ctx) error {
	var employee models.Employee
	if err := c.BodyParser(&employee); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	if err := config.DB.Create(&employee).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.JSON(employee)
}

func UpdateEmployee(c *fiber.Ctx) error {
	id := c.Params("id")
	var employee models.Employee
	if err := config.DB.First(&employee, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err.Error())
	}
	if err := c.BodyParser(&employee); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	if err := config.DB.Save(&employee).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.JSON(employee)
}

func DeleteEmployee(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := config.DB.Delete(&models.Employee{}, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err.Error())
	}
	return c.SendStatus(fiber.StatusNoContent)
}
