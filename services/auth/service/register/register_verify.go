package register

import (
	"context"
	"fmt"
	"time"
	"warehouseai/auth/adapter"
	"warehouseai/auth/dataservice"
	e "warehouseai/auth/errors"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type RegisterVerifyRequest struct {
	Token  string `json:"token"`
	UserId string `json:"user_id"`
}

type RegisterVerifyResponse struct {
	Verified bool `json:"verified"`
}

func validateVerifyRequest(request RegisterVerifyRequest) *e.ErrorResponse {
	if request.Token == "" || request.UserId == "" {
		return e.NewErrorResponse(e.HttpBadRequest, "One of the parameters is empty.")
	}

	return nil
}

func RegisterVerify(
	request RegisterVerifyRequest,
	user adapter.UserGrpcInterface,
	verificationToken dataservice.VerificationTokenInterface,
	logger *logrus.Logger,
) (*RegisterVerifyResponse, *e.ErrorResponse) {
	if err := validateVerifyRequest(request); err != nil {
		return nil, err
	}

	existVerificationToken, dbErr := verificationToken.Get(map[string]interface{}{"user_id": request.UserId})

	if dbErr != nil {
		logger.WithFields(logrus.Fields{"time": time.Now(), "error": dbErr.Payload}).Info("Register verify user")
		return nil, e.NewErrorResponseFromDBError(dbErr.ErrorType, dbErr.Message)
	}

	fmt.Println(time.Now())
	fmt.Println(existVerificationToken.ExpiresAt)

	// Удаляем токен, если он протух. Пользователю нужно отправлять запрос еще раз.
	if time.Now().After(existVerificationToken.ExpiresAt) {
		verificationToken.Delete(map[string]interface{}{"id": existVerificationToken.ID})
		return nil, e.NewErrorResponse(e.HttpBadRequest, "Invalid or expired verification token.")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existVerificationToken.Token), []byte(request.Token)); err != nil {
		logger.WithFields(logrus.Fields{"time": time.Now(), "error": err.Error()}).Info("Register verify user")
		return nil, e.NewErrorResponse(e.HttpBadRequest, "Invalid register verification key")
	}

	verified, gwErr := user.UpdateVerificationStatus(context.Background(), request.UserId)

	if gwErr != nil {
		logger.WithFields(logrus.Fields{"time": time.Now(), "error": gwErr.ErrorMessage}).Info("Register verify user")
		return nil, gwErr
	}

	if err := verificationToken.Delete(map[string]interface{}{"id": existVerificationToken.ID}); err != nil {
		logger.WithFields(logrus.Fields{"time": time.Now(), "error": err.Payload}).Info("Register verify user")
		return nil, e.NewErrorResponseFromDBError(err.ErrorType, err.Message)
	}

	return &RegisterVerifyResponse{Verified: verified}, nil
}
