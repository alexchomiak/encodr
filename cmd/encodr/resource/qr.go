package resource

import (
	"strconv"
	"time"

	"github.com/alexchomiak/encodr/cmd/encodr/model"

	"github.com/alexchomiak/encodr/cmd/encodr/controller"
	"github.com/gofiber/fiber/v2"
)

type QRCodeResource struct {
	Controller controller.QRCodeController
}

func (q *QRCodeResource) getQRCode(c *fiber.Ctx) error {
	url := c.Params("url")
	if url == "" {
		c.JSON(&model.ErrorResponse{
			TimeStamp: time.Now(),
			Message:   "URL is required for QR code generation",
		})
		return c.SendStatus(400)

	}

	size := c.Query("size")
	var sizeInt int64

	if size == "" {
		sizeInt = 256
	} else {
		sizeInt, _ = strconv.ParseInt(size, 10, 64)
	}

	file, err := q.Controller.Encode(url, sizeInt)

	if err != nil {
		c.JSON(&model.ErrorResponse{
			TimeStamp: time.Now(),
			Message:   err.Error(),
		})
		return c.SendStatus(500)
	} else {
		c.Status(200)
		c.Context().SetContentType("image/png")
		return c.Send(file)
	}

}

func NewQRCodeResource() *QRCodeResource {
	return &QRCodeResource{
		Controller: controller.QRCodeController{},
	}
}

func (q *QRCodeResource) BindRoutes(app *fiber.App) {
	app.Get("/qrcode/:url", q.getQRCode)
}
