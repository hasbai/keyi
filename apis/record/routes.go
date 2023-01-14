package record

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app fiber.Router) {
	app.Get("/records/products", ListRecords)
	app.Post("/records/products/:id", AddRecord)
	app.Delete("/records/products/:id", DeleteRecord)
	app.Delete("/records/products", DeleteAllRecords)
}
