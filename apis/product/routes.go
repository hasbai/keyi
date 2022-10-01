package product

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app fiber.Router) {
	app.Get("/categories/:id/products", ListProducts)
	app.Post("/categories/:id/products", AddProduct)
	app.Get("/products/:id", GetProduct)
	app.Put("/products/:id", ModifyProduct)
	app.Delete("/products/:id", DeleteProduct)
}
