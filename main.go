package main

import (
	"go-crud-postgres/config"
	"go-crud-postgres/models"
	"go-crud-postgres/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	config.ConnectDB()

	// Auto-migrate the schema
	config.DB.AutoMigrate(&models.Employee{})

	routes.EmployeeRoutes(app)
	app.Listen(":3000")
}
