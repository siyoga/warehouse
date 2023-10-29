package get

import (
	"time"
	db "warehouse/src/internal/database"
	pg "warehouse/src/internal/database/postgresdb"
	"warehouse/src/internal/utils/httputils"

	"github.com/sirupsen/logrus"
)

type Request struct {
	ID string `json:"id"`
}

type AIProvider interface {
	GetOneBy(key string, value interface{}) (*pg.AI, *db.DBError)
}

func GetByID(getRequest Request, aiProvider AIProvider, logger *logrus.Logger) (*pg.AI, *httputils.ErrorResponse) {
	existAI, dbErr := aiProvider.GetOneBy("id", getRequest.ID)

	if dbErr != nil {
		logger.WithFields(logrus.Fields{"time": time.Now(), "error": dbErr.Payload}).Info("Create AI")
		return nil, httputils.NewErrorResponseFromDBError(dbErr.ErrorType, dbErr.Message)
	}

	return existAI, nil
}
