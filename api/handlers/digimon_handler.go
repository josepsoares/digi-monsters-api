package digihandlers

import (
	"josepsoares/digi-monsters-api/pkg/models"

	"github.com/gofiber/fiber/v2"
	"github.com/zeimedee/go-postgres/database"
)

func GetDigimon() fiber.Handler {
	return func(c *fiber.Ctx) error {
		digimon := []models.Digimon{}
		id := new(models.Digimon)

		if err := c.BodyParser(id); err != nil {
			return c.Status(400).JSON(err.Error())
		}

		database.DB.Db.Where("id = ?", id.ID).Find(&digimon)

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
			"data":   digimon,
			"error":  nil,
		})
	}
}

// GetBooks is handler/controller which lists all Books from the BookShop
func IndexDigimons() fiber.Handler {
	return func(c *fiber.Ctx) error {
		digimons := []models.Digimon{}
		database.DB.Db.Find(&digimons)

		return c.JSON(&fiber.Map{
			"status": true,
			"data":   digimons,
			"error":  nil,
		})
	}
}

func IndexPartnerDigimons() fiber.Handler {
	return func(c *fiber.Ctx) error {
		digimon := []models.Digimon{}
		id := new(models.Digimon)

		if err := c.BodyParser(id); err != nil {
			return c.Status(400).JSON(err.Error())
		}

		database.DB.Db.Where("id = ?", id.ID).Find(&digimon)

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
			"data":   digimon,
			"error":  nil,
		})
	}
}
