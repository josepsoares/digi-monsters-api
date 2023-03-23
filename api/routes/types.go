package digiroutes

import (
	digihandlers "josepsoares/digi-monsters-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func DigiTypesRouter(app fiber.Router) {
	app.Get("/type/:id", digihandlers.GetDigiType())
	app.Get("/type", digihandlers.IndexDigiTypes())
}
