package httpx

import "github.com/gofiber/fiber/v2"

func JSON(c *fiber.Ctx, status int, v interface{}) error {
	return c.Status(status).JSON(v)
}

func Error(c *fiber.Ctx, status int, msg string) error {
	return c.Status(status).JSON(fiber.Map{"error": msg})
}
