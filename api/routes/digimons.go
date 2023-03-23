package digiroutes

import (
	digihandlers "josepsoares/digi-monsters-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func DigimonsRouter(app fiber.Router) {
	app.Get("/digimon/:id", digihandlers.GetDigimon())
	app.Get("/digimon", digihandlers.IndexDigimons())
	app.Get("/digimon/partner", digihandlers.IndexPartnerDigimons())
}
