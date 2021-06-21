package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/Product-Category/model"
	repo "github.com/golang/Product-Category/repository"
)

func GetAllProduct(c *fiber.Ctx) error {
	AverageRating(c)
	return c.JSON(repo.Products.GetAllProducts())
}

func GetProductById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	AverageRating(c)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	product, err := repo.Products.FindProductById(int64(id))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	return c.JSON(product)
}

func DeleteProductById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	AverageRating(c)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	err = repo.Products.DeleteProductById(int64(id))
	// repo.Reviews.DeleteReviewByProductId(int64(id))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	} else {
		return c.SendString("delete successfully")
	}
}

func CreateProduct(c *fiber.Ctx) error {
	product := new(model.Product)

	err := c.BodyParser(&product)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	AverageRating(c)
	productId := repo.Products.CreateNewProduct(product)
	return c.SendString(fmt.Sprintf("New Product is created successfully with id = %d", productId))

}

func UpdateProductById(c *fiber.Ctx) error {
	updatedProduct := new(model.Product)
	oldProduct, _ := repo.Products.FindProductById(updatedProduct.Id)
	err := c.BodyParser(&updatedProduct)
	// id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	if updatedProduct.Price != oldProduct.Price {
		HistoryPrice := new(model.HistoryPrice)
		HistoryPrice.ProductId = updatedProduct.Id
		HistoryPrice.OldPrice = updatedProduct.Price
		err := c.BodyParser(&HistoryPrice)
		// if error
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "Cannot parse JSON",
				"error":   err,
			})
		}
		errr := repo.HistoryPrice.CreateNewHistoryPrice(HistoryPrice)
		if errr != nil {
			return c.Status(404).SendString(err.Error())
		}
	}
	AverageRating(c)
	err = repo.Products.UpdateProductById(updatedProduct)

	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	return c.SendString(fmt.Sprintf("Product with id = %d is successfully updated", updatedProduct.Id))
}

func UpsertProduct(c *fiber.Ctx) error {
	product := new(model.Product)

	err := c.BodyParser(&product)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	oldProduct, _ := repo.Products.FindProductById(product.Id)
	fmt.Println(oldProduct)
	if product.Price != oldProduct.Price {
		HistoryPrice := new(model.HistoryPrice)
		HistoryPrice.ProductId = product.Id
		HistoryPrice.OldPrice = oldProduct.Price
		err := c.BodyParser(&HistoryPrice)
		// if error
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "Cannot parse JSON",
				"error":   err,
			})
		}
		errr := repo.HistoryPrice.CreateNewHistoryPrice(HistoryPrice)
		if errr != nil {
			return c.Status(404).SendString(err.Error())
		}
	}
	AverageRating(c)
	id := repo.Products.UpsertProduct(product)
	return c.SendString(fmt.Sprintf("Product with id = %d is successfully upserted", id))
}

func GetAllHistoryPrices(c *fiber.Ctx) error {
	return c.JSON(repo.HistoryPrice.GetAllHistoryPrices())
}
