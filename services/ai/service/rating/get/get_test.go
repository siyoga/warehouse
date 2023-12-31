package get

import (
	"math"
	"testing"
	dMock "warehouseai/ai/dataservice/mocks"
	e "warehouseai/ai/errors"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestRatingGet(t *testing.T) {
	ctl := gomock.NewController(t)

	aiMock := dMock.NewMockAiInterface(ctl)
	ratingMock := dMock.NewMockRatingInterface(ctl)
	logger := logrus.New()

	request := GetAIRatingRequest{AiId: uuid.Must(uuid.NewV4()).String()}
	expectRating := float64(4.55555)
	expectCount := int64(123)

	aiMock.EXPECT().Get(map[string]interface{}{"id": request.AiId}).Return(nil, nil).Times(1)
	ratingMock.EXPECT().GetAverageAiRating(request.AiId).Return(&expectRating, nil).Times(1)
	ratingMock.EXPECT().GetCountAiRating(request.AiId).Return(&expectCount, nil).Times(1)

	response, err := GetAIRating(request, aiMock, ratingMock, logger)

	require.Nil(t, err)
	require.Equal(t, &GetAIRatingResponse{AverageRating: math.Round(expectRating*100) / 100, RatingCount: expectCount}, response)
}

func TestRatingGetError(t *testing.T) {
	// TODO: Add ai not exist case
	cases := []struct {
		name          string
		request       GetAIRatingRequest
		expectedError *e.DBError
	}{
		{
			name:          "Internal error",
			request:       GetAIRatingRequest{AiId: uuid.Must(uuid.NewV4()).String()},
			expectedError: e.NewDBError(e.DbSystem, "Something went wrong.", "internal error"),
		},
	}

	for _, tCase := range cases {
		ctl := gomock.NewController(t)

		ratingMock := dMock.NewMockRatingInterface(ctl)
		aiMock := dMock.NewMockAiInterface(ctl)
		logger := logrus.New()

		t.Run(tCase.name, func(t *testing.T) {
			aiMock.EXPECT().Get(map[string]interface{}{"id": tCase.request.AiId}).Return(nil, nil).Times(1)
			ratingMock.EXPECT().GetAverageAiRating(tCase.request.AiId).Return(nil, tCase.expectedError).Times(1)
			rating, err := GetAIRating(tCase.request, aiMock, ratingMock, logger)

			require.Nil(t, rating)
			require.NotNil(t, err)
			require.Equal(t, e.NewErrorResponseFromDBError(tCase.expectedError.ErrorType, tCase.expectedError.Message), err)
		})
	}
}
