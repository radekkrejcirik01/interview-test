package rest

import (
	"interview_test/pkg/rest/controller"

	"github.com/gofiber/fiber/v2"
)

func Create() *fiber.App {
	app := fiber.New()

	app.Post("records", controller.CreateRecord)
	app.Get("records/:id", controller.GetRecord)
	app.Put("records/:id", controller.UpdateRecord)
	app.Delete("records/:id", controller.DeleteRecord)
	app.Get("readyz", controller.HealthCheck)

	return app
}
