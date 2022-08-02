package handlers_test

import (
	"errors"
	"tcp-pow/internal/handlers"
	"tcp-pow/internal/handlers/tests/factories"
	"tcp-pow/internal/handlers/tests/mocks"
	"tcp-pow/internal/models"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Given a challenge sent handler", func() {
	Describe("When handle is called", func() {
		var ctrl *gomock.Controller
		var challengeSentHandler *handlers.ChallengeSentHandler
		var solutionCalculator *mocks.MockSolutionCalculator
		var data models.TCPData

		BeforeEach(func() {
			ctrl = gomock.NewController(GinkgoT())
			challengeSentHandler, solutionCalculator = factories.CreateChallengeSentHandler(ctrl)
			data = models.TCPData{
				Nonce:      "abc",
				Difficulty: 3,
			}
		})

		Context("And nonce is an empty string", func() {
			BeforeEach(func() {
				data = models.TCPData{
					Nonce: "",
				}
			})

			It("Then handler should return an error", func() {
				response, err := challengeSentHandler.Handle(data)
				Expect(response).To(BeEquivalentTo(models.TCPData{}))
				Expect(err).ToNot(BeNil())
			})
		})

		Context("And difficulty is zero", func() {
			BeforeEach(func() {
				data = models.TCPData{
					Difficulty: 0,
				}
			})

			It("Then handler should return an error", func() {
				response, err := challengeSentHandler.Handle(data)
				Expect(response).To(BeEquivalentTo(models.TCPData{}))
				Expect(err).ToNot(BeNil())
			})
		})

		Context("And solution calculator returns an error", func() {
			BeforeEach(func() {
				solutionCalculator.EXPECT().Calculate("abc", 3).Return("", errors.New(""))
			})

			It("Then handler should return an error", func() {
				response, err := challengeSentHandler.Handle(data)
				Expect(response).To(BeEquivalentTo(models.TCPData{}))
				Expect(err).ToNot(BeNil())
			})
		})

		Context("And everything is fine", func() {
			BeforeEach(func() {
				solutionCalculator.EXPECT().Calculate("abc", 3).Return("5", nil)
			})

			It("Then handler should return data with send solution package type, nonce, difficulty and prefix", func() {
				expectedResponse := models.TCPData{
					PackageType: models.SendSolution,
					Nonce:       "abc",
					Difficulty:  3,
					Prefix:      "5",
				}

				response, err := challengeSentHandler.Handle(data)
				Expect(response).To(BeEquivalentTo(expectedResponse))
				Expect(err).To(BeNil())
			})
		})
	})
})
