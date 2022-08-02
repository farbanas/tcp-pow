package handlers_test

import (
	"tcp-pow/internal/handlers"
	"tcp-pow/internal/handlers/tests/factories"
	"tcp-pow/internal/handlers/tests/mocks"
	"tcp-pow/internal/models"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestHandlers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handlers Suite")
}

var _ = Describe("Given a challenge requested handler", func() {
	Describe("When handle is called", func() {
		var ctrl *gomock.Controller
		var challengeRequestedHandler *handlers.ChallengeRequestedHandler
		var nonceCalculator *mocks.MockNonceCalculator
		var difficulty int
		var data models.TCPData

		BeforeEach(func() {
			ctrl = gomock.NewController(GinkgoT())
			challengeRequestedHandler, nonceCalculator, difficulty = factories.CreateChallengeRequestedHandler(ctrl)
			data = models.TCPData{}
		})

		Context("And everything is fine", func() {

			It("Then handler should return data with send challenge package type, nonce and difficulty", func() {
				nonceCalculator.EXPECT().GetNewValue().Return("abc")
				expectedResponse := models.TCPData{
					PackageType: models.SendChallenge,
					Nonce:       "abc",
					Difficulty:  difficulty,
				}

				response, err := challengeRequestedHandler.Handle(data)
				Expect(response).To(BeEquivalentTo(expectedResponse))
				Expect(err).To(BeNil())
			})
		})
	})
})
