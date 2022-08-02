package factories

import (
	"tcp-pow/internal/handlers"
	"tcp-pow/internal/handlers/tests/mocks"

	"github.com/golang/mock/gomock"
)

func CreateChallengeRequestedHandler(ctrl *gomock.Controller) (*handlers.ChallengeRequestedHandler, *mocks.MockNonceCalculator, int) {
	nonceCalculator := mocks.NewMockNonceCalculator(ctrl)
	difficulty := 5
	challengeRequestedHandler := handlers.NewChallengeRequestedHandler(nonceCalculator, difficulty)
	return challengeRequestedHandler, nonceCalculator, difficulty
}
