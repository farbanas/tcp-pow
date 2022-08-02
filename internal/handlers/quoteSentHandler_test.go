package handlers_test

import (
	"tcp-pow/internal/handlers"
	"tcp-pow/internal/models"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Given a solution sent handler", func() {
	Describe("When handle is called", func() {
		var quoteSentHandler *handlers.QuoteSentHandler
		var data models.TCPData
		BeforeEach(func() {
			quoteSentHandler = handlers.NewQuoteSentHandler()
			data = models.TCPData{
				Prefix: "7",
				Nonce:  "abc",
			}
		})

		Context("And everything is fine", func() {
			It("Then handler should an error", func() {
				response, err := quoteSentHandler.Handle(data)
				Expect(response).To(BeEquivalentTo(models.TCPData{}))
				Expect(err).To(BeNil())
			})
		})
	})
})
