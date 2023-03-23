package digiroutes

import (
	digihandlers "josepsoares/digi-monsters-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func DigiAttributesRouter(app fiber.Router) {
	app.Get("/attribute/:id", digihandlers.GetDigiAttribute())
	app.Get("/attribute", digihandlers.IndexDigiAttributes())
}
