package validators_test

import (
	"errors"
	"tcp-pow/internal/models"
	"tcp-pow/internal/validators"
	"tcp-pow/internal/validators/tests/factories"
	"tcp-pow/internal/validators/tests/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestValidators(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Validator Suite")
}

var _ = Describe("Given a nonce validator", func() {
	Describe("When is nonce valid is called", func() {
		var ctrl *gomock.Controller
		var nonceValidator *validators.NonceValidator
		var cache *mocks.MockCache
		var tcpData models.TCPData

		BeforeEach(func() {
			ctrl = gomock.NewController(GinkgoT())
			cache, nonceValidator = factories.CreateNonceValidator(ctrl)
			tcpData = models.TCPData{
				Address: "abc",
				Nonce:   "test",
			}
		})

		AfterEach(func() {
			ctrl.Finish()
		})

		Context("And address is not found in cache", func() {
			BeforeEach(func() {
				cache.EXPECT().Get(tcpData.Address).Return("", errors.New(""))
			})

			It("Then validator should return true (nonce valid)", func() {
				valid := nonceValidator.IsNonceValid(tcpData)
				Expect(valid).To(BeTrue())
			})
		})

		Context("And address is found in cache", func() {
			BeforeEach(func() {
				cache.EXPECT().Get(tcpData.Address).Return("test", nil)
			})

			When("Nonce from cache is not equal to the one in data", func() {
				BeforeEach(func() {
					tcpData.Nonce = "notatest"
				})

				It("Then validator should return false", func() {
					valid := nonceValidator.IsNonceValid(tcpData)
					Expect(valid).To(BeFalse())
				})
			})

			When("Nonce from cache is equal to the one in data", func() {

				It("Then validator should return true", func() {
					valid := nonceValidator.IsNonceValid(tcpData)
					Expect(valid).To(BeTrue())
				})
			})
		})
	})
})
