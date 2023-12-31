package service

import (
	"time"
	"warehouseai/user/dataservice"
	e "warehouseai/user/errors"
	m "warehouseai/user/model"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

func Create(userInfo m.User, user dataservice.UserInterface, logger *logrus.Logger) (*string, *e.ErrorResponse) {
	userInfo.ID = uuid.Must(uuid.NewV4())

	if err := user.Create(&userInfo); err != nil {
		logger.WithFields(logrus.Fields{"time": time.Now(), "error": err.Payload}).Info("User create")
		return nil, e.NewErrorResponseFromDBError(err.ErrorType, err.Message)
	}

	plainId := userInfo.ID.String()

	return &plainId, nil
}
