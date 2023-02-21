package blogcontroller

import (
	"net/http"

	"github.com/Shunux-Stuxnet/Blog_CRUD/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Show(c *fiber.Ctx) error {
	id := c.Params("id")
	var blog models.Blog
	if err := models.DB.First(&blog, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "No record found",
			})
		}
	}

	return c.JSON(blog)
}

func Index(c *fiber.Ctx) error {
	var blogs []models.Blog
	models.DB.Find(&blogs)
	return c.Status(fiber.StatusOK).JSON(blogs)
}

func Create(c *fiber.Ctx) error {
	var blog models.Blog
	if err := c.BodyParser(&blog); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if err := models.DB.Create(&blog).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(blog)
}

func Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var blog models.Blog
	if err := c.BodyParser(&blog); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if models.DB.Where("id = ?", id).Updates(&blog).RowsAffected == 0 {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"message": "ID doesn't exit, data updation failed",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data updated",
	})
}

func Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	var blog models.Blog

	if models.DB.Delete(&blog, id).RowsAffected == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Data not found",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Data Successfully deleted",
	})
}
