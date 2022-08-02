package mappers

import "tcp-pow/internal/models"

// ErrorMapper is a simple implementation of wrapping errors into our internal models.TCPData.
type ErrorMapper struct {
}

func NewErrorMapper() *ErrorMapper {
	return &ErrorMapper{}
}

// MapToErrorResponse creates a models.TCPData that has the error message set in the Error field.
func (e *ErrorMapper) MapToErrorResponse(err error) models.TCPData {
	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}

	response := models.TCPData{
		Error: errMsg,
	}

	return response
}
