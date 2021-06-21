package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/Product-Category/model"
	repo "github.com/golang/Product-Category/repository"
)

var products []model.Product

func GetAllReview(c *fiber.Ctx) error {
	AverageRating(c)
	return c.JSON(repo.Reviews.GetAllReviews())
}

func DeleteReviewById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	AverageRating(c)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	err = repo.Reviews.DeleteReviewById(int64(id))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	} else {
		return c.SendString("delete successfully")
	}
}

func CreateReview(c *fiber.Ctx) error {
	review := new(model.Review)

	err := c.BodyParser(review)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	reviewId := repo.Reviews.CreateNewReview(review)
	AverageRating(c)
	return c.SendString(fmt.Sprintf("New review is created successfully with id = %d", reviewId))

}

func UpsertReview(c *fiber.Ctx) error {
	review := new(model.Review)

	err := c.BodyParser(&review)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	id := repo.Reviews.UpsertReview(review)
	AverageRating(c)
	return c.SendString(fmt.Sprintf("Review with id = %d is successfully upserted", id))
}

func AverageRating(c *fiber.Ctx) error {
	listReview := repo.Reviews.GetAllReviews()
	for _, review := range listReview {
		product, err := repo.Products.FindProductById(review.ProductId)
		if err != nil {
			return c.JSON(fiber.Map{
				"message": "Not found this Product id",
			})
		}
		result := repo.Reviews.AverageRating()
		product.AverageRating = result[int64(review.ProductId)]
		repo.Products.UpsertProduct(product)
	}
	return c.SendString(fmt.Sprintf("successfully"))
}
