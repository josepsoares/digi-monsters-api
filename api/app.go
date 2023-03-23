package main

import (
	digiroutes "josepsoares/digi-monsters-api/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/zeimedee/go-postgres/database"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	api := app.Group("/api")

	// setup routes
	digiroutes.DigimonsRouter(api)
	digiroutes.DigiLevelsRouter(api)
	digiroutes.DigiAttributesRouter(api)
	digiroutes.DigiTypesRouter(api)

	log.Fatal(app.Listen(":8080"))
}
