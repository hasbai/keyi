package category

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app fiber.Router) {
	app.Get("/categories/:id", GetCategory)
	app.Get("/categories", ListCategories)
	app.Post("/categories", AddCategory)
	app.Put("/categories/:id", ModifyCategory)
	app.Delete("/categories/:id", DeleteCategory)
}
