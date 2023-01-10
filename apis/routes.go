package apis

import (
	"keyi/apis/category"
	"keyi/apis/product"
	"keyi/auth"
	_ "keyi/docs"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func registerRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/api")
	})
	app.Get("/docs", func(c *fiber.Ctx) error {
		return c.Redirect("/docs/index.html")
	})
	app.Get("/docs/*", fiberSwagger.WrapHandler)
}

func RegisterRoutes(app *fiber.App) {
	registerRoutes(app)

	group := app.Group("/api")
	group.Get("/", Index)

	auth.RegisterRoutes(group)
	product.RegisterRoutes(group)
	category.RegisterRoutes(group)
}
