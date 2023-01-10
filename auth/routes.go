package auth

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app fiber.Router) {
	app.Post("/login", Login)
	app.Post("/refresh", Refresh)
	app.Post("/register", Register)
	app.Post("/validate", Validate)
	app.Get("/users/:id/activate", Activate)

	app.Get("/tenants", ListTenants)
}

