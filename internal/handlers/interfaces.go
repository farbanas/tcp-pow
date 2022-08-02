package handlers

import (
	"net"
	"tcp-pow/internal/models"
)

//go:generate mockgen -destination=./tests/mocks/connectionMapper.go -package mocks . ConnectionMapper
type ConnectionMapper interface {
	MapToInternalData(conn net.Conn) (models.TCPData, bool, error)
	MapFromInternalData(models.TCPData) ([]byte, error)
}

//go:generate mockgen -destination=./tests/mocks/errorMapper.go -package mocks . ErrorMapper
type ErrorMapper interface {
	MapToErrorResponse(err error) models.TCPData
}

//go:generate mockgen -destination=./tests/mocks/nonceValidator.go -package mocks . NonceValidator
type NonceValidator interface {
	IsNonceValid(tcpData models.TCPData) bool
}

//go:generate mockgen -destination=./tests/mocks/tcpHandler.go -package mocks . TCPHandler
type TCPHandler interface {
	Handle(models.TCPData) (models.TCPData, error)
}

//go:generate mockgen -destination=./tests/mocks/cache.go -package mocks . Cache
type Cache interface {
	Get(address string) (string, error)
	Set(address string, nonce string)
}

//go:generate mockgen -destination=./tests/mocks/nonceCalculator.go -package mocks . NonceCalculator
type NonceCalculator interface {
	GetNewValue() string
}

//go:generate mockgen -destination=./tests/mocks/solutionCalculator.go -package mocks . SolutionCalculator
type SolutionCalculator interface {
	Calculate(nonce string, difficulty int) (string, error)
}

//go:generate mockgen -destination=./tests/mocks/quoteHandler.go -package mocks . QuoteHandler
type QuoteHandler interface {
	GetQuote() (string, error)
}
