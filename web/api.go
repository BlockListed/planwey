package web

import "github.com/gofiber/fiber/v2"

func apiRoutes(router fiber.Router) {
	router.Get("/", func(c *fiber.Ctx) error {
		c.SendString("Hello world")
		return nil
	})
}
