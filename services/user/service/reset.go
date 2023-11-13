package service

import (
	"time"
	e "warehouseai/internal/errors"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type ResetUserPasswordRequest struct {
	Password string `json:"password"`
}

func ResetUserPassword(request ResetUserPasswordRequest, userId string, userUpdater UserUpdater, logger *logrus.Logger) *e.ErrorResponse {
	hash, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 12)
	request.Password = string(hash)

	if _, dbErr := userUpdater.RawUpdate(userId, request); dbErr != nil {
		logger.WithFields(logrus.Fields{"time": time.Now(), "error": dbErr.Payload}).Info("Update user password")
		return e.NewErrorResponseFromDBError(dbErr.ErrorType, dbErr.Message)
	}

	return nil
}
