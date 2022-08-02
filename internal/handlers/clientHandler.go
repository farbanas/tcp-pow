package handlers

import (
	"errors"
	"fmt"
	"net"
	"tcp-pow/internal/models"
)

// ClientHandler is the entrypoint handler for the client. It specifies two important actions: sending the initial request and responding
// to other data received from the server.
type ClientHandler struct {
	connectionMapper         ConnectionMapper
	challengeSentHandler     TCPHandler
	quoteSentHandler         TCPHandler
	solutionIncorrectHandler TCPHandler
	errorMapper              ErrorMapper
}

func NewClientHandler(
	connectionMapper ConnectionMapper,
	challengeSentHandler TCPHandler,
	quoteSentHandler TCPHandler,
	solutionIncorrectHandler TCPHandler,
	errorMapper ErrorMapper,
) *ClientHandler {
	return &ClientHandler{
		connectionMapper:         connectionMapper,
		challengeSentHandler:     challengeSentHandler,
		quoteSentHandler:         quoteSentHandler,
		solutionIncorrectHandler: solutionIncorrectHandler,
		errorMapper:              errorMapper,
	}
}

// SendInitialRequest sends an empty request that has RequestChallenge package type to the server.
func (c *ClientHandler) SendInitialRequest(conn net.Conn) error {
	request := models.TCPData{
		PackageType: models.RequestChallenge,
	}

	requestBytes, err := c.connectionMapper.MapFromInternalData(request)
	if err != nil {
		return err
	}

	requestBytes = append(requestBytes, '\n')

	_, err = conn.Write(requestBytes)
	if err != nil {
		return err
	}

	return nil
}

// Handle routes the data to the correct handler based on the package type.
func (c *ClientHandler) Handle(conn net.Conn) ([]byte, bool, error) {
	response, done, err := c.connectionMapper.MapToInternalData(conn)
	if err != nil {
		return nil, false, err
	}

	if done {
		return nil, true, nil
	}

	request, done, err := c.CalculateClientRequest(response)
	if err != nil {
		return nil, false, err
	}

	fmt.Printf("request:%+v\n", request)

	marshaledRequest, err := c.connectionMapper.MapFromInternalData(request)
	if err != nil {
		return nil, false, err
	}

	return marshaledRequest, done, nil
}

func (c *ClientHandler) CalculateClientRequest(response models.TCPData) (models.TCPData, bool, error) {
	var request models.TCPData
	var err error

	switch response.PackageType {
	case models.SendChallenge:
		request, err = c.challengeSentHandler.Handle(response)
		if err != nil {
			return models.TCPData{}, false, err
		}

	case models.SendQuote:
		c.quoteSentHandler.Handle(response)
		return models.TCPData{}, true, nil

	case models.SolutionIncorrect:
		request, err = c.solutionIncorrectHandler.Handle(response)
		if err != nil {
			return models.TCPData{}, false, err
		}

	case models.Error:
		err = errors.New(response.Error)
		request = c.errorMapper.MapToErrorResponse(err)
		return request, false, nil

	case models.Undefined:
		err = errors.New("received undefined package type")
		request = c.errorMapper.MapToErrorResponse(err)
		return request, false, nil

	default:
		err = errors.New("received unknown package type")
		request = c.errorMapper.MapToErrorResponse(err)
		return request, false, nil
	}

	return request, false, nil
}
