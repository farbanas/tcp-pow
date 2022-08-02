package handlers_test

import (
	"crypto/sha256"
	"errors"
	"hash"
	"tcp-pow/internal/handlers"
	"tcp-pow/internal/handlers/tests/factories"
	"tcp-pow/internal/handlers/tests/mocks"
	"tcp-pow/internal/models"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Given a solution sent handler", func() {
	Describe("When handle is called", func() {
		var ctrl *gomock.Controller
		var solutionSentHandler *handlers.SolutionSentHandler
		var hasher hash.Hash
		var difficulty int
		var quoteHandler *mocks.MockQuoteHandler
		var data models.TCPData

		BeforeEach(func() {
			ctrl = gomock.NewController(GinkgoT())
			hasher = sha256.New()
			difficulty = 1
			solutionSentHandler, quoteHandler = factories.CreateSolutionSentHandler(hasher, difficulty, ctrl)
			data = models.TCPData{
				Prefix:     "7",
				Nonce:      "abc",
				Difficulty: difficulty,
			}
		})

		Context("And nonce is an empty string", func() {
			BeforeEach(func() {
				data.Nonce = ""
			})

			It("Then handler should return response data with error message and error package type", func() {
				expectedResponse := models.TCPData{
					PackageType: models.Error,
					Nonce:       data.Nonce,
					Difficulty:  data.Difficulty,
					Prefix:      data.Prefix,
					Error:       "received an empty nonce",
				}

				response, err := solutionSentHandler.Handle(data)
				Expect(response).To(BeEquivalentTo(expectedResponse))
				Expect(err).To(BeNil())
			})
		})

		Context("And prefix is an empty string", func() {
			BeforeEach(func() {
				data.Prefix = ""
			})

			It("Then handler should return response data with error message and error package type", func() {
				expectedResponse := models.TCPData{
					PackageType: models.Error,
					Nonce:       data.Nonce,
					Difficulty:  data.Difficulty,
					Prefix:      data.Prefix,
					Error:       "received an empty prefix",
				}

				response, err := solutionSentHandler.Handle(data)
				Expect(response).To(BeEquivalentTo(expectedResponse))
				Expect(err).To(BeNil())
			})
		})

		Context("And solution is not valid", func() {
			BeforeEach(func() {
				data.Prefix = "5"
			})

			It("Then handler should return response data with solution incorrect package type", func() {
				expectedResponse := models.TCPData{
					PackageType: models.SolutionIncorrect,
					Nonce:       data.Nonce,
					Difficulty:  data.Difficulty,
					Prefix:      data.Prefix,
				}

				response, err := solutionSentHandler.Handle(data)
				Expect(response).To(BeEquivalentTo(expectedResponse))
				Expect(err).To(BeNil())
			})
		})

		Context("And solution is valid", func() {
			When("Quote handler returns an error", func() {
				BeforeEach(func() {
					quoteHandler.EXPECT().GetQuote().Return("", errors.New(""))
				})

				It("Then handler should return an error", func() {
					response, err := solutionSentHandler.Handle(data)
					Expect(response).To(BeEquivalentTo(models.TCPData{}))
					Expect(err).ToNot(BeNil())
				})
			})

			When("Quote handler returns a quote", func() {
				BeforeEach(func() {
					quoteHandler.EXPECT().GetQuote().Return("test", nil)
				})

				It("Then handler should return data with send quote package type and a quote", func() {
					expectedResponse := models.TCPData{
						PackageType: models.SendQuote,
						Quote:       "test",
					}

					response, err := solutionSentHandler.Handle(data)
					Expect(response).To(BeEquivalentTo(expectedResponse))
					Expect(err).To(BeNil())
				})
			})
		})
	})
})
