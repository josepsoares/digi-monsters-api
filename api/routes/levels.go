package digiroutes

import (
	digihandlers "josepsoares/digi-monsters-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func DigiLevelsRouter(app fiber.Router) {
	app.Get("/level/:id", digihandlers.GetDigiLevel())
	app.Get("/level", digihandlers.IndexDigiLevels())
}
