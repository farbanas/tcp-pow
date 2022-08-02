package handlers

import (
	"encoding/hex"
	"hash"
	"strings"
	"tcp-pow/internal/models"
)

// SolutionSentHandler handles the SendSolution package type on the server side.
type SolutionSentHandler struct {
	hasher           hash.Hash
	difficulty       int
	precomputedZeros string
	quoteHandler     QuoteHandler
}

func NewSolutionSentHandler(hasher hash.Hash, difficulty int, quoteHandler QuoteHandler) *SolutionSentHandler {
	precomputedZeros := ""
	for i := 0; i < difficulty; i++ {
		precomputedZeros += "0"
	}

	return &SolutionSentHandler{
		hasher:           hasher,
		difficulty:       difficulty,
		precomputedZeros: precomputedZeros,
		quoteHandler:     quoteHandler,
	}
}

// Handle checks for the existence of nonce and prefix and then proceeds to calculate the hash and check if it really starts
// with enough zeros (defined by the difficulty parameter).
// Errors in this handler are wrapped as response data and returned to the client because they are connected to the client instead
// of being internal.
func (s *SolutionSentHandler) Handle(data models.TCPData) (models.TCPData, error) {
	if data.Nonce == "" {
		responseData := models.TCPData{
			PackageType: models.Error,
			Nonce:       data.Nonce,
			Difficulty:  data.Difficulty,
			Prefix:      data.Prefix,
			Error:       "received an empty nonce",
		}

		return responseData, nil
	}

	if data.Prefix == "" {
		responseData := models.TCPData{
			PackageType: models.Error,
			Nonce:       data.Nonce,
			Difficulty:  data.Difficulty,
			Prefix:      data.Prefix,
			Error:       "received an empty prefix",
		}

		return responseData, nil
	}

	checksum, err := s.calculateChecksum([]byte(data.Prefix + data.Nonce))
	if err != nil {
		return models.TCPData{}, err
	}

	if !strings.HasPrefix(checksum, s.precomputedZeros) {
		responseData := models.TCPData{
			PackageType: models.SolutionIncorrect,
			Nonce:       data.Nonce,
			Difficulty:  data.Difficulty,
			Prefix:      data.Prefix,
		}

		return responseData, nil
	}

	quote, err := s.quoteHandler.GetQuote()
	if err != nil {
		return models.TCPData{}, err
	}

	responseData := models.TCPData{
		PackageType: models.SendQuote,
		Quote:       quote,
	}

	return responseData, nil
}

func (s *SolutionSentHandler) calculateChecksum(data []byte) (string, error) {
	defer s.hasher.Reset()

	_, err := s.hasher.Write(data)
	if err != nil {
		return "", err
	}

	checksum := s.hasher.Sum(nil)
	hexChecksum := hex.EncodeToString(checksum)

	return hexChecksum, nil
}
