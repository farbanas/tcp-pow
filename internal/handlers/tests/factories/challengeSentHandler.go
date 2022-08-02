package factories

import (
	"tcp-pow/internal/handlers"
	"tcp-pow/internal/handlers/tests/mocks"

	"github.com/golang/mock/gomock"
)

func CreateChallengeSentHandler(ctrl *gomock.Controller) (*handlers.ChallengeSentHandler, *mocks.MockSolutionCalculator) {
	solutionCalculator := mocks.NewMockSolutionCalculator(ctrl)
	challengeSentHandler := handlers.NewChallengeSentHandler(solutionCalculator)
	return challengeSentHandler, solutionCalculator
}
