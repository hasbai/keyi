package auth

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app fiber.Router) {
	app.Get("/tenants", ListTenants)

	app.Post("/login", Login)
	app.Post("/refresh", Refresh)
	app.Post("/register", Register)
	app.Post("/validate", Validate)
	app.Get("/users/:id/activate", Activate)

	app.Get("/users", ListUsers)
	app.Get("/users/:id", GetUser)

	app.Get("/users/:id/follow", ListFollow)
	app.Get("/users/:id/followed-by", ListFollowedBy)
	app.Post("/users/:id/follow/:f_id", AddFollow)
	app.Delete("/users/:id/follow/:f_id", DeleteFollow)
}
