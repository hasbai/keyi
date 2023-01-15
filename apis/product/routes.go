package product

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app fiber.Router) {
	app.Get("/products/records", ListRecords)
	app.Get("/products/favorite", ListFavorites)
	app.Get("/products/:id", GetProduct)
	app.Get("/categories/:id/products", ListProducts)
	app.Get("/users/:id/products", ListUserProducts)
	app.Get("/users/:id/products/:type", ListUserProductsType)

	app.Post("/categories/:id/products", AddProduct)
	app.Post("/products/:id/favorite", AddFavorite)
	app.Post("/products/:id/records", AddRecord)

	app.Put("/products/:id", ModifyProduct)

	app.Delete("/products/records", DeleteAllRecords)
	app.Delete("/products/favorite", DeleteAllFavorites)
	app.Delete("/products/:id/records", DeleteRecord)
	app.Delete("/products/:id/favorite", DeleteFavorite)
	app.Delete("/products/:id", DeleteProduct)
}
