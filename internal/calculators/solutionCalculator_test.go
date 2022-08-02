package calculators_test

import (
	"crypto/sha256"
	"encoding/hex"
	"hash"
	"tcp-pow/internal/calculators"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Given a solution calculator", func() {
	Describe("When calculate is called", func() {
		var solutionCalculator *calculators.SolutionCalculator
		var hasher hash.Hash
		var nonce string
		var difficulty int
		BeforeEach(func() {
			hasher = sha256.New()
			solutionCalculator = calculators.NewSolutionCalculator(hasher)
			nonce = "test"
			difficulty = 3
		})

		Context("And difficulty is 0", func() {
			BeforeEach(func() {
				difficulty = 0
			})

			It("Then calculator should return a string with no leading zeroes", func() {
				prefix, err := solutionCalculator.Calculate(nonce, difficulty)
				Expect(prefix).To(BeEquivalentTo("0"))
				Expect(err).To(BeNil())
			})
		})

		Context("And difficulty is 3", func() {

			It("Then calculator should return a string with 3 leading zeroes", func() {
				prefix, err := solutionCalculator.Calculate(nonce, difficulty)
				Expect(prefix).ToNot(BeEquivalentTo("0"))
				Expect(err).To(BeNil())

				_, err = hasher.Write([]byte(prefix + nonce))
				Expect(err).To(BeNil())

				result := hasher.Sum(nil)
				hexResult := hex.EncodeToString(result)
				Expect(hexResult).To(HavePrefix("000"))
			})
		})
	})
})
