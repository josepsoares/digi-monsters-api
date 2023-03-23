package digihandlers

import (
	"josepsoares/digi-monsters-api/pkg/models"

	"github.com/gofiber/fiber/v2"
	"github.com/zeimedee/go-postgres/database"
)

func GetDigiType() fiber.Handler {
	return func(c *fiber.Ctx) error {
		digiType := []models.Type{}
		id := new(models.Type)

		if err := c.BodyParser(id); err != nil {
			return c.Status(400).JSON(err.Error())
		}

		database.DB.Db.Where("id = ?", id.ID).Find(&digiType)

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
			"data":   digiType,
			"error":  nil,
		})
	}
}

// GetBooks is handler/controller which lists all Books from the BookShop
func IndexDigiTypes() fiber.Handler {
	return func(c *fiber.Ctx) error {
		digiTypes := []models.Type{}
		database.DB.Db.Find(&digiTypes)

		return c.JSON(&fiber.Map{
			"status": true,
			"data":   digiTypes,
			"error":  nil,
		})
	}
}
