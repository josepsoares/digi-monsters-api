package digihandlers

import (
	"josepsoares/digi-monsters-api/pkg/models"

	"github.com/gofiber/fiber/v2"
	"github.com/zeimedee/go-postgres/database"
)

func GetDigiAttribute() fiber.Handler {
	return func(c *fiber.Ctx) error {
		digiAttribute := []models.Attribute{}
		id := new(models.Attribute)

		if err := c.BodyParser(id); err != nil {
			return c.Status(400).JSON(err.Error())
		}

		database.DB.Db.Where("id = ?", id.ID).Find(&digiAttribute)

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
			"data":   digiAttribute,
			"error":  nil,
		})
	}
}

// GetBooks is handler/controller which lists all Books from the BookShop
func IndexDigiAttributes() fiber.Handler {
	return func(c *fiber.Ctx) error {
		digiAttributes := []models.Attribute{}
		database.DB.Db.Find(&digiAttributes)

		return c.JSON(&fiber.Map{
			"status": true,
			"data":   digiAttributes,
			"error":  nil,
		})
	}
}
