package create

import (
	"time"
	db "warehouse/src/internal/database"
	pg "warehouse/src/internal/database/postgresdb"
	"warehouse/src/internal/utils/httputils"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

type Request struct {
	Name        string                 `json:"name"`
	AiID        uuid.UUID              `json:"ai_id"`
	Payload     map[string]interface{} `json:"payload"`
	PayloadType pg.PayloadType         `json:"payload_type"`
	InputType   pg.IOType              `json:"input_type"`
	OutputType  pg.IOType              `json:"output_type"`
	RequestType pg.RequestScheme       `json:"request_type"`
	URL         string                 `json:"url"`
}

type CommandCreator interface {
	Add(item *pg.Command) *db.DBError
}

func CreateCommand(commandCreds *Request, commandCreator CommandCreator, logger *logrus.Logger) *httputils.ErrorResponse {
	newCommand := &pg.Command{
		ID:            uuid.Must(uuid.NewV4()),
		Name:          commandCreds.Name,
		AIID:          commandCreds.AiID,
		RequestScheme: commandCreds.RequestType,
		InputType:     commandCreds.InputType,
		OutputType:    commandCreds.OutputType,
		Payload:       commandCreds.Payload,
		PayloadType:   commandCreds.PayloadType,
		URL:           commandCreds.URL,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if dbErr := commandCreator.Add(newCommand); dbErr != nil {
		logger.WithFields(logrus.Fields{"time": time.Now(), "error": dbErr.Payload}).Info("Add new command to AI")
		return httputils.NewErrorResponseFromDBError(dbErr.ErrorType, dbErr.Message)
	}

	return nil
}
