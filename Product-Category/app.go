package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang/Product-Category/routes"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
	})

	app.Static("/public", "./public", fiber.Static{ //http://localhost:3000/public OR http://localhost:3000/public/dog.jpeg
		Compress:  true,
		ByteRange: true,
		Browse:    true,
		MaxAge:    3600,
	})

	// reviewRouter := app.Group("api/review")
	productRouter := app.Group("/api/product")
	userRouter := app.Group("/api/user")
	cartRouter := app.Group("/api/cart")
	categoryRouter := app.Group("/api/category")
	reviewRouter := app.Group("/api/review")
	routes.ConfigReviewRouter(&reviewRouter)
	routes.ConfigProductRouter(&productRouter) //http://localhost:3000/api/book
	routes.ConfigUserRouter(&userRouter)
	routes.ConfigCartRouter(&cartRouter)
	routes.ConfigCategoryRouter(&categoryRouter)
	app.Listen(":3000")
}
