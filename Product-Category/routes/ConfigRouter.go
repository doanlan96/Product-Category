package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang/Product-Category/controller"
)

func ConfigProductRouter(router *fiber.Router) {
	//Return all Products
	(*router).Get("/", controller.GetAllProduct)

	(*router).Get("/history", controller.GetAllHistoryPrices)

	(*router).Get("/:id", controller.GetProductById)

	(*router).Delete("/:id", controller.DeleteProductById)

	(*router).Patch("", controller.UpdateProductById)

	(*router).Post("", controller.CreateProduct)

	(*router).Put("", controller.UpsertProduct)

}

func ConfigUserRouter(router *fiber.Router) {
	//Return all Products
	(*router).Get("/", controller.GetAllUser)

	(*router).Get("/:id", controller.GetUserById)

	(*router).Delete("/:id", controller.DeleteUserById)

	(*router).Patch("", controller.UpdateUserById)

	(*router).Post("", controller.CreateUser)

	(*router).Put("", controller.UpsertUser)
}

func ConfigCartRouter(router *fiber.Router) {
	//Return all Products
	(*router).Get("/", controller.GetAllCart)

	(*router).Get("/:id", controller.GetCartById)

	(*router).Delete("/:id", controller.DeleteCartById)

	(*router).Patch("", controller.UpdateCartById)

	(*router).Post("", controller.CreateCart)

	(*router).Put("", controller.UpsertCart)
}

func ConfigCategoryRouter(router *fiber.Router) {
	//Return all Categories
	(*router).Get("/", controller.GetAllCategories)

	(*router).Get("/:id", controller.GetCategoryById)
}

func ConfigReviewRouter(router *fiber.Router) {
	(*router).Get("/", controller.GetAllReview)
	(*router).Post("", controller.CreateReview)
	(*router).Delete("/:id", controller.DeleteReviewById)
	(*router).Get("/average/:id", controller.AverageRating)
}
