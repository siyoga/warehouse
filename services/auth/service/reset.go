package service

import (
	"context"
	"fmt"
	"math/rand"
	"time"
	"warehouseai/auth/adapter"
	"warehouseai/auth/dataservice"
	"warehouseai/auth/model"
	e "warehouseai/internal/errors"
	"warehouseai/internal/gen"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type ResetAttemptRequest struct {
	Email string `json:"email"`
}

type ResetAttemptResponse struct {
	TokenId string `json:"token_id"`
}

type ResetConfirmRequest struct {
	UserId   string `json:"user_id"`
	Password string `json:"password"`
}

type ResetConfirmResponse struct {
	UserId *string `json:"user_id"`
}

type ResetVerifyResponse struct {
	UserId string `json:"user_id"`
}

func PasswordReset(request ResetConfirmRequest, resetTokenId string, user adapter.UserGrpcInterface, resetToken dataservice.ResetTokenInterface, logger *logrus.Logger, ctx context.Context) (*ResetConfirmResponse, *e.ErrorResponse) {
	existResetToken, dbErr := resetToken.Get(map[string]interface{}{"id": resetTokenId})

	if dbErr != nil {
		logger.WithFields(logrus.Fields{"time": time.Now(), "error": dbErr.Payload}).Info("Reset password")
		return nil, e.NewErrorResponseFromDBError(dbErr.ErrorType, dbErr.Message)
	}

	if err := resetToken.Delete(map[string]interface{}{"id": existResetToken.ID}); err != nil {
		logger.WithFields(logrus.Fields{"time": time.Now(), "error": err.Payload}).Info("Reset password")
		return nil, e.NewErrorResponseFromDBError(err.ErrorType, err.Message)
	}

	hash, hashErr := bcrypt.GenerateFromPassword([]byte(request.Password), 12)

	if hashErr != nil {
		logger.WithFields(logrus.Fields{"time": time.Now(), "error": hashErr.Error()}).Info("Reset password")
		return nil, e.NewErrorResponse(e.HttpInternalError, "Error while hashing new password.")
	}

	resp, gwErr := user.ResetPassword(ctx, &gen.ResetPasswordRequest{UserId: request.UserId, Password: string(hash)})

	if gwErr != nil {
		return nil, gwErr
	}

	return &ResetConfirmResponse{UserId: resp}, nil
}

func VerifyResetCode(verificationCode string, resetTokenId string, resetToken dataservice.ResetTokenInterface, logger *logrus.Logger) (*ResetVerifyResponse, *e.ErrorResponse) {
	existResetToken, err := resetToken.Get(map[string]interface{}{"id": resetTokenId})

	if err != nil {
		logger.WithFields(logrus.Fields{"time": time.Now(), "error": err.Payload}).Info("Verify reset token")
		return nil, e.NewErrorResponseFromDBError(err.ErrorType, err.Message)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existResetToken.Token), []byte(verificationCode)); err != nil {
		logger.WithFields(logrus.Fields{"time": time.Now(), "error": err.Error()}).Info("Verify reset token")
		return nil, e.NewErrorResponse(e.HttpBadRequest, "Invalid verification code.")
	}

	return &ResetVerifyResponse{UserId: existResetToken.UserId.String()}, nil
}

func SendResetEmail(req ResetAttemptRequest, resetToken dataservice.ResetTokenInterface, user adapter.UserGrpcInterface, broker adapter.MailProducerInterface, logger *logrus.Logger) (*ResetAttemptResponse, *e.ErrorResponse) {
	existUser, gwErr := user.GetByEmail(context.Background(), req.Email)

	if gwErr != nil {
		logger.WithFields(logrus.Fields{"time": time.Now(), "error": gwErr.ErrorMessage}).Info("Send reset token")
		return nil, gwErr
	}

	verificationCode := generateCode(8)
	hash, bcryptErr := bcrypt.GenerateFromPassword([]byte(verificationCode), 12)

	if bcryptErr != nil {
		return nil, e.NewErrorResponse(e.HttpInternalError, "Verification key encryption error")
	}

	newResetToken := &model.ResetToken{
		ID:        uuid.Must(uuid.NewV4()),
		UserId:    uuid.FromStringOrNil(existUser.Id),
		Token:     string(hash),
		ExpiresAt: time.Now().Add(time.Minute * 5),
		CreatedAt: time.Now(),
	}

	if err := resetToken.Create(newResetToken); err != nil {
		logger.WithFields(logrus.Fields{"time": time.Now(), "error": err.Payload}).Info("Send reset token")
		return nil, e.NewErrorResponseFromDBError(err.ErrorType, err.Message)
	}

	message := model.Email{
		To:      req.Email,
		Subject: "Восстановление пароля",
		Message: fmt.Sprintf(`
      Здраствуйте, %s!
      
      Мы получили запрос на восстановление пароля от Вашего аккаунта, связанного с почтой %s.
      Ваш код верификации: %s
      
      Если это не вы - проигнорируйте данное письмо.
      
      WarehouseAI Team
      `, existUser.Firstname, existUser.Email, verificationCode),
	}

	if err := broker.SendEmail(message); err != nil {
		logger.WithFields(logrus.Fields{"time": time.Now(), "error": err.Error()}).Info("Send email")
		return nil, e.NewErrorResponse(e.HttpInternalError, "Failed to send email.")
	}

	return &ResetAttemptResponse{TokenId: newResetToken.ID.String()}, nil
}

func generateCode(length int) string {
	charset := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	batch := make([]byte, length)

	for i := range batch {
		batch[i] = charset[rand.Intn(len(charset))]
	}

	return string(batch)
}
