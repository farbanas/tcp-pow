package calculators

import (
	"encoding/hex"
	"hash"
	"strconv"
	"strings"
)

// SolutionCalculator is used for solving the given cryptographic puzzle.
type SolutionCalculator struct {
	hasher hash.Hash
}

func NewSolutionCalculator(hasher hash.Hash) *SolutionCalculator {
	return &SolutionCalculator{hasher: hasher}
}

// Calculate uses the given nonce and difficulty to calculate a prefix with which the hash of prefix + nonce will have
// difficulty number of zeros.
func (s *SolutionCalculator) Calculate(nonce string, difficulty int) (string, error) {
	defer s.hasher.Reset()
	prefix := 0
	solutionPrefix := createPrefixZeroes(difficulty)

	for {
		prefixStr := strconv.Itoa(prefix)
		_, err := s.hasher.Write([]byte(prefixStr + nonce))
		if err != nil {
			return "", err
		}

		result := s.hasher.Sum(nil)
		hexResult := hex.EncodeToString(result)
		if strings.HasPrefix(hexResult, solutionPrefix) {
			return prefixStr, nil
		}

		s.hasher.Reset()
		prefix++
	}
}

// createPrefixZeroes is used for precomputing the string that only contains zeros, based on the difficulty.
// This is a small optimization to always validate the solution agains this precomputed string. If
// the solution has this string as a prefix, it means that the solution has enough zeros.
func createPrefixZeroes(difficulty int) string {
	precomputedZeroes := ""
	for i := 0; i < difficulty; i++ {
		precomputedZeroes += "0"
	}

	return precomputedZeroes
}
