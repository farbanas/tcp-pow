package mappers_test

import (
	"errors"
	"tcp-pow/internal/mappers"
	"tcp-pow/internal/models"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Given a error mapper", func() {
	Describe("When map to error response is called", func() {
		var errorMapper *mappers.ErrorMapper
		var err error

		BeforeEach(func() {
			errorMapper = mappers.NewErrorMapper()
			err = errors.New("this is an error")
		})

		Context("And error is nil", func() {

			It("Then mapper should return tcp data with error field an empty string", func() {
				data := errorMapper.MapToErrorResponse(nil)
				Expect(data).To(BeEquivalentTo(models.TCPData{}))
			})
		})

		Context("And error is not nil", func() {

			It("Then mapper should return tcp data with an error field that equals to the error message", func() {
				expectedData := models.TCPData{
					Error: "this is an error",
				}

				data := errorMapper.MapToErrorResponse(err)
				Expect(data).To(BeEquivalentTo(expectedData))
			})
		})
	})
})
