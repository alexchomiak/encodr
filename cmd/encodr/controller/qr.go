package controller

import qrcode "github.com/skip2/go-qrcode"

type QRCodeController struct{}

func (q *QRCodeController) Encode(url string, size int64, recovery qrcode.RecoveryLevel) ([]byte, error) {
	bytes, err := qrcode.Encode(url, recovery, int(size))
	return bytes, err
}
