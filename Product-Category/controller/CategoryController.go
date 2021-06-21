package controller

import (
	"github.com/gofiber/fiber/v2"
	repo "github.com/golang/Product-Category/repository"
)

func GetAllCategories(c *fiber.Ctx) error {
	return c.JSON(repo.Categories.GetAllCategories())
}

func GetCategoryById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	category, err := repo.Categories.FindCategoryById(int64(id))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	return c.JSON(category)
}
