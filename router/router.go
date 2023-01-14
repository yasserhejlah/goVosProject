package router

import (
	//	"github.com/yasserhejlah/goVosProject/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/yasserhejlah/goVosProject/handler"

)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {
	// grouping
	api := app.Group("/api")
	v1 := api.Group("/user")
	// routes
	v1.Get("/", handler.GetAllUsers)
	v1.Get("/:id", handler.GetSingleUser /* middleware.Authorize */)
	v1.Post("/", handler.CreateUser)
	v1.Put("/:id", handler.UpdateUser)
	v1.Delete("/:id", handler.DeleteUserByID)
}
