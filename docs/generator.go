package docs

import "github.com/gofiber/fiber/v2"

func ServeDocs(app *fiber.App) {
	app.Get("/docs", func(c *fiber.Ctx) error { return c.SendFile("./docs/scalar/index.html") })
	app.Get("/openapi.yaml", func(c *fiber.Ctx) error { return c.SendFile("./docs/openapi.yaml") })
}
