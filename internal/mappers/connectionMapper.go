package mappers

import (
	"bufio"
	"encoding/json"
	"net"
	"tcp-pow/internal/models"
)

// ConnectionMapper is used for mapping the data that we receive from the outside to internal application data and vice-versa.
type ConnectionMapper struct {
}

func NewConnectionMapper() *ConnectionMapper {
	return &ConnectionMapper{}
}

// MapToInternalData maps the net.Conn model to models.TCPData (internal model). Since this is not perfectly decomposed,
// it returns a bool as a second return parameter. This bool specifies whether the connection is closed or not.
func (c *ConnectionMapper) MapToInternalData(conn net.Conn) (models.TCPData, bool, error) {
	data, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return models.TCPData{}, true, nil
	}

	var tcpData models.TCPData
	err = json.Unmarshal([]byte(data), &tcpData)
	if err != nil {
		return models.TCPData{}, false, err
	}

	return tcpData, false, nil
}

// MapFromInternalData maps the models.TCPData to bytes, for easy writing to connection.
func (c *ConnectionMapper) MapFromInternalData(data models.TCPData) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}
