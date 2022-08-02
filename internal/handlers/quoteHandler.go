package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"tcp-pow/internal/models"
)

// quoteHandler is used for getting a new quote.
type quoteHandler struct {
	url string
}

func NewQuoteHandler(url string) QuoteHandler {
	return &quoteHandler{
		url: url,
	}
}

// GetQuote gets the quote from the url and returns it.
func (q *quoteHandler) GetQuote() (string, error) {
	response, err := http.Get(q.url)
	if err != nil {
		return "", err
	}

	responseBody := response.Body
	defer responseBody.Close()

	bodyBytes, err := io.ReadAll(responseBody)
	if err != nil {
		return "", err
	}

	var quotes []models.Quote
	err = json.Unmarshal(bodyBytes, &quotes)
	if err != nil {
		return "", err
	}

	if len(quotes) < 1 {
		return "", errors.New("didn't receive any quote from the address")
	}

	return quotes[0].QuoteText, nil
}
