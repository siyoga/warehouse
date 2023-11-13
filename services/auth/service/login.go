package service

import (
	"context"
	"time"
	"warehouseai/auth/adapter"
	"warehouseai/auth/dataservice"
	"warehouseai/auth/model"
	e "warehouseai/internal/errors"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(req *LoginRequest, user adapter.UserGrpcInterface, session dataservice.SessionInterface, logger *logrus.Logger, ctx context.Context) (*model.Session, *e.ErrorResponse) {
	existUser, gwErr := user.GetByEmail(context.Background(), req.Email)

	if gwErr != nil {
		logger.WithFields(logrus.Fields{"time": time.Now(), "error": gwErr.ErrorMessage}).Info("Login user")
		return nil, gwErr
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existUser.Password), []byte(req.Password)); err != nil {
		logger.WithFields(logrus.Fields{"time": time.Now(), "error": err.Error()}).Info("Login user")
		return nil, e.NewErrorResponse(e.HttpBadRequest, "Invalid credentials")
	}

	newSession, sessErr := session.Create(ctx, existUser.Id)

	if sessErr != nil {
		logger.WithFields(logrus.Fields{"time": time.Now(), "error": sessErr}).Info("Login user")
		return nil, e.NewErrorResponseFromDBError(sessErr.ErrorType, sessErr.Message)
	}

	return newSession, nil
}
