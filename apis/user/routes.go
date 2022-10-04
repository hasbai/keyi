package user

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app fiber.Router) {
	app.Post("/verify/email", VerifyEmail)
	app.Post("/users", RegisterUser)
	app.Put("/users/:id", ModifyUser)
}
