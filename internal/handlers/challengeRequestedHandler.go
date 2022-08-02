package handlers

import (
	"tcp-pow/internal/models"
)

// ChallengeRequestedHandler handles processing the RequestChallenge package type.
type ChallengeRequestedHandler struct {
	nonceCalculator NonceCalculator
	difficulty      int
}

func NewChallengeRequestedHandler(nonceCalculator NonceCalculator, difficulty int) *ChallengeRequestedHandler {
	return &ChallengeRequestedHandler{
		nonceCalculator: nonceCalculator,
		difficulty:      difficulty,
	}
}

// Handle gets the new nonce value and creates new data that has the SendChalleng package type, nonce and difficulty assigned.
func (c *ChallengeRequestedHandler) Handle(data models.TCPData) (models.TCPData, error) {
	nonce := c.nonceCalculator.GetNewValue()

	responseData := models.TCPData{
		PackageType: models.SendChallenge,
		Nonce:       nonce,
		Difficulty:  c.difficulty,
	}

	return responseData, nil
}
