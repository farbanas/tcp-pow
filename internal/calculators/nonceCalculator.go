package calculators

import (
	"math/rand"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// NonceCalculator generates a random string of given length.
type NonceCalculator struct {
	length int
}

func NewNonceCalculator(length int) *NonceCalculator {
	rand.Seed(time.Now().UnixNano())
	return &NonceCalculator{
		length: length,
	}
}

// GetNewValue returns a random string of given length.
func (n *NonceCalculator) GetNewValue() string {
	b := make([]byte, n.length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
