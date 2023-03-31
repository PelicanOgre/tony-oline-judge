package test

import (
	"crypto/tls"
	"net/smtp"
	"testing"

	"github.com/jordan-wright/email"
)

func TestSendEmail(t *testing.T) {
	e := email.NewEmail()
	e.From = "Get <tojproject@163.com>"
	e.To = []string{"kizzy8452@gmail.com"}
	e.Subject = "Verification code sending test"
	e.HTML = []byte("Your TOJ verification code: <b>123456</b>")
	// 返回 EOF 时，关闭SSL重试
	err := e.SendWithTLS("smtp.163.com:465",
		smtp.PlainAuth("", "tojproject@163.com", "HFTHAYCSPVUMZHEI", "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		t.Fatal(err)
	}
}
