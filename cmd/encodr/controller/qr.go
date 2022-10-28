package controller

import qrcode "github.com/skip2/go-qrcode"

type QRCodeController struct{}

func (q *QRCodeController) Encode(url string, size int64) ([]byte, error) {
	bytes, err := qrcode.Encode(url, qrcode.Medium, int(size))
	return bytes, err
}
