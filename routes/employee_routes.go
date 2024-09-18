package routes

import (
	"go-crud-postgres/controllers"

	"github.com/gofiber/fiber/v2"
)

func EmployeeRoutes(app *fiber.App) {
	api := app.Group("/api/employees")
	api.Get("/", controllers.GetEmployees)
	api.Get("/:id", controllers.GetEmployee)
	api.Post("/", controllers.CreateEmployee)
	api.Put("/:id", controllers.UpdateEmployee)
	api.Delete("/:id", controllers.DeleteEmployee)
}
