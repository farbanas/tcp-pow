package handlers

import (
	"tcp-pow/internal/models"
)

// QuoteSentHandler is used for processing the SendQuote package type on the client side. It currently does nothing.
type QuoteSentHandler struct {
}

func NewQuoteSentHandler() *QuoteSentHandler {
	return &QuoteSentHandler{}
}

func (q *QuoteSentHandler) Handle(response models.TCPData) (models.TCPData, error) {
	return models.TCPData{}, nil
}
