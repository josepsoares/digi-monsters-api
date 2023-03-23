package digihandlers

import (
	"josepsoares/digi-monsters-api/pkg/models"

	"github.com/gofiber/fiber/v2"
	"github.com/zeimedee/go-postgres/database"
)

func GetDigiLevel() fiber.Handler {
	return func(c *fiber.Ctx) error {
		digiLevel := []models.Level{}
		id := new(models.Level)

		if err := c.BodyParser(id); err != nil {
			return c.Status(400).JSON(err.Error())
		}

		database.DB.Db.Where("id = ?", id.ID).Find(&digiLevel)

		/* if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(&fiber.Map{
				"status": false,
				"data":   "",
				"error":  err.Error(),
			})
		} */

		return c.Status(200).JSON(&fiber.Map{
			"status": true,
			"data":   digiLevel,
			"error":  nil,
		})
	}
}

// GetBooks is handler/controller which lists all Books from the BookShop
func IndexDigiLevels() fiber.Handler {
	return func(c *fiber.Ctx) error {
		digiLevels := []models.Level{}
		database.DB.Db.Find(&digiLevels)

		return c.JSON(&fiber.Map{
			"status": true,
			"data":   digiLevels,
			"error":  nil,
		})
	}
}
