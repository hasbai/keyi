package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"keyi/auth"
	"keyi/config"
)

func RegisterMiddlewares(app *fiber.App) {
	app.Use(recover.New())
	if config.Config.Mode != "perf" {
		app.Use(logger.New())
	}
	if config.Config.Debug {
		app.Use(pprof.New())
	}
	app.Use(getUserID)
}

func getUserID(c *fiber.Ctx) error {
	authorization := c.Get("Authorization", "")
	if authorization == "" { // token can be in either header or cookie
		authorization = c.Cookies("access")
	}

	var claims *auth.MyClaims
	var err error
	if len(authorization) < 7 {
		claims = &auth.MyClaims{}
	} else {
		claims, err = auth.ParseToken(authorization[7:]) // extract "Bearer "
		if err != nil {
			return err
		}
	}

	c.Locals("claims", claims)

	return c.Next()
}
