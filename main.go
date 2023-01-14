package main

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/yasserhejlah/goVosProject/database"
)

func main() {
	database.Connect()
	app := fiber.New()
	/* app.Use(logger.New())
	app.Use(cors.New())
	router.SetupRoutes(app)
	// handle unavailable route
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	}) */
	app.Listen(":8080")
}
