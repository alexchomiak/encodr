package main

import (
	"time"

	"github.com/alexchomiak/encodr/cmd/encodr/model"
	"github.com/alexchomiak/encodr/cmd/encodr/resource"

	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		c.JSON(&model.HealthCheckResponse{
			TimeStamp: time.Now(),
			Status:    "OK",
		})
		return c.SendStatus(200)
	})

	var port string
	val, ok := os.LookupEnv("ENCODR_PORT")
	if ok {
		port = val
	} else {
		port = "80"
	}

	resource.NewQRCodeResource().BindRoutes(app)

	app.Listen(":" + port)
}
