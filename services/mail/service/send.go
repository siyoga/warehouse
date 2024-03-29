package service

import (
	"time"
	"warehouseai/mail/model"

	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

func SendEmail(
	from string,
	email model.Email,
	dialer *gomail.Dialer,
	logger *logrus.Logger,
) error {
	m := gomail.NewMessage()

	m.SetHeader("From", m.FormatAddress(from, "WarehouseAI Team"))
	m.SetHeader("To", email.To)
	m.SetHeader("Subject", email.Subject)
	m.SetBody("text/plain", email.Message)

	if err := dialer.DialAndSend(m); err != nil {
		logger.WithFields(logrus.Fields{"time": time.Now(), "error": err.Error()}).Info("Send email")
		return err
	}

	return nil
}
