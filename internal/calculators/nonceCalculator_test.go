package calculators_test

import (
	"tcp-pow/internal/calculators"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCalculators(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Calculator Suite")
}

var _ = Describe("Given a nonce calculator", func() {
	Describe("When get new value is called", func() {
		var nonceCalculator *calculators.NonceCalculator
		var length int
		var previousNonce string

		Context("And length is 0", func() {
			BeforeEach(func() {
				length = 0
				nonceCalculator = calculators.NewNonceCalculator(length)
			})

			It("Then calculator should return an empty string", func() {
				nonce := nonceCalculator.GetNewValue()
				Expect(nonce).To(BeEquivalentTo(""))
			})
		})

		Context("And length is 5", func() {
			BeforeEach(func() {
				length = 5
				nonceCalculator = calculators.NewNonceCalculator(length)
			})

			It("Then calculator should return a string of length 5", func() {
				nonce := nonceCalculator.GetNewValue()
				Expect(len(nonce)).To(BeEquivalentTo(length))
				previousNonce = nonce
			})

			Context("And get new value is called again with length 5", func() {

				It("Then calculator should return a different string of length 5", func() {
					nonce := nonceCalculator.GetNewValue()
					Expect(len(nonce)).To(BeEquivalentTo(length))
					Expect(nonce).ToNot(BeEquivalentTo(previousNonce))
				})
			})
		})

	})
})
