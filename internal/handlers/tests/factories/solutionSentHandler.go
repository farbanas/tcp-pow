package factories

import (
	"hash"
	"tcp-pow/internal/handlers"
	"tcp-pow/internal/handlers/tests/mocks"

	"github.com/golang/mock/gomock"
)

func CreateSolutionSentHandler(hasher hash.Hash, difficulty int, ctrl *gomock.Controller) (*handlers.SolutionSentHandler, *mocks.MockQuoteHandler) {
	quoteHandler := mocks.NewMockQuoteHandler(ctrl)
	solutionSentHandler := handlers.NewSolutionSentHandler(hasher, difficulty, quoteHandler)
	return solutionSentHandler, quoteHandler
}
