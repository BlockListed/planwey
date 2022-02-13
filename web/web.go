package web

import (
	"fmt"
	"log"
	"planwey/configuration"

	"github.com/gofiber/fiber/v2"
)

var App fiber.App = *fiber.New()

func Run() {
	defineRoutes()
	log.Fatal(App.Listen(fmt.Sprintf(":%s", configuration.Configuration.Web.Port)))
}

func defineRoutes() {
	api := App.Group("/api")
	apiRoutes(api)
}
