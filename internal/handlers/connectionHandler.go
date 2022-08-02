package handlers

import (
	"errors"
	"fmt"
	"net"
	"tcp-pow/internal/models"
)

// ConnectionHandler is the main handler for the server.
type ConnectionHandler struct {
	connectionMapper        ConnectionMapper
	errorMapper             ErrorMapper
	nonceValidator          NonceValidator
	requestChallengeHandler TCPHandler
	solutionSentHandler     TCPHandler
}

func NewConnectionHandler(
	connectionMapper ConnectionMapper,
	errorMapper ErrorMapper,
	nonceValidator NonceValidator,
	requestChallengeHandler TCPHandler,
	solutionSentHandler TCPHandler,
) *ConnectionHandler {
	return &ConnectionHandler{
		connectionMapper:        connectionMapper,
		errorMapper:             errorMapper,
		nonceValidator:          nonceValidator,
		requestChallengeHandler: requestChallengeHandler,
		solutionSentHandler:     solutionSentHandler,
	}
}

// Handle reads the data from the connection, transforms it to the internal model and routes the data to the appropriate handler
// depending on which package type is found in the data.
func (c *ConnectionHandler) Handle(conn net.Conn) ([]byte, bool, error) {
	tcpData, done, err := c.connectionMapper.MapToInternalData(conn)
	if err != nil {
		return nil, false, err
	}

	if done {
		return nil, true, nil
	}

	fmt.Printf("request:%+v\n", tcpData)

	response, err := c.CalculateResponse(tcpData)
	if err != nil {
		return nil, false, err
	}

	fmt.Printf("response:%+v\n", response)

	responseBytes, err := c.connectionMapper.MapFromInternalData(response)
	if err != nil {
		return nil, false, err
	}

	return responseBytes, false, nil
}

func (c *ConnectionHandler) CalculateResponse(tcpData models.TCPData) (models.TCPData, error) {
	var response models.TCPData
	var err error

	if !c.nonceValidator.IsNonceValid(tcpData) {
		err = fmt.Errorf("nonce %s is not valid for client %s", tcpData.Nonce, tcpData.Address)
		response = c.errorMapper.MapToErrorResponse(err)
		return response, nil
	}

	switch tcpData.PackageType {
	case models.RequestChallenge:
		response, err = c.requestChallengeHandler.Handle(tcpData)
		if err != nil {
			return models.TCPData{}, err
		}

	case models.SendSolution:
		response, err = c.solutionSentHandler.Handle(tcpData)
		if err != nil {
			return models.TCPData{}, err
		}

	case models.Undefined:
		err = errors.New("received undefined package type")
		response = c.errorMapper.MapToErrorResponse(err)
		return response, nil

	default:
		err = errors.New("received unknown package type")
		response = c.errorMapper.MapToErrorResponse(err)
		return response, nil
	}

	return response, nil
}
