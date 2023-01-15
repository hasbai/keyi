package product

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app fiber.Router) {
	app.Get("/products/favorite", ListFavorites)
	app.Get("/products/:id", GetProduct)
	app.Get("/categories/:id/products", ListProducts)

	app.Post("/categories/:id/products", AddProduct)
	app.Post("/products/:id/favorite", AddFavorite)

	app.Put("/products/:id", ModifyProduct)

	app.Delete("/products/favorite", DeleteAllFavorites)
	app.Delete("/products/:id/favorite", DeleteFavorite)
	app.Delete("/products/:id", DeleteProduct)
}
