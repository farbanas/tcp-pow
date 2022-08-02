package handlers

import (
	"errors"
	"tcp-pow/internal/models"
)

// ChallengeSentHandler handles processing the SendChallenge package type.
type ChallengeSentHandler struct {
	solutionCalculator SolutionCalculator
}

func NewChallengeSentHandler(solutionCalculator SolutionCalculator) *ChallengeSentHandler {
	return &ChallengeSentHandler{
		solutionCalculator: solutionCalculator,
	}
}

// Handle checks if nonce and difficulty are set in the data and then proceeds to calculate the solution (prefix).
// This handler is used on the client.
func (c *ChallengeSentHandler) Handle(response models.TCPData) (models.TCPData, error) {
	nonce := response.Nonce
	if nonce == "" {
		return models.TCPData{}, errors.New("nonce is empty")
	}

	difficulty := response.Difficulty
	if difficulty == 0 {
		return models.TCPData{}, errors.New("difficulty is zero")
	}

	solution, err := c.solutionCalculator.Calculate(nonce, response.Difficulty)
	if err != nil {
		return models.TCPData{}, err
	}

	request := models.TCPData{
		PackageType: models.SendSolution,
		Nonce:       nonce,
		Difficulty:  difficulty,
		Prefix:      solution,
	}

	return request, nil
}
