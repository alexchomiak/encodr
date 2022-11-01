package resource

import (
	"strconv"
	"time"

	"github.com/alexchomiak/encodr/cmd/encodr/model"
	"github.com/skip2/go-qrcode"

	urllib "net/url"

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

	// * URL Decode value
	decodedUrl, err := urllib.QueryUnescape(url)

	size := c.Query("size")
	var sizeInt int64

	if size == "" {
		sizeInt = 256
	} else {
		sizeInt, _ = strconv.ParseInt(size, 10, 64)
	}

	// * Recovery level
	recovery := c.Query("recovery")
	var recoveryLevel qrcode.RecoveryLevel
	switch recovery {
	case "L":
		recoveryLevel = qrcode.Low
	case "M":
		recoveryLevel = qrcode.Medium
	case "H":
		recoveryLevel = qrcode.High
	case "Hplus":
		recoveryLevel = qrcode.Highest
	default:
		recoveryLevel = qrcode.High
	}

	file, err := q.Controller.Encode(decodedUrl, sizeInt, recoveryLevel)

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
