package servers

import "net"

const protocol = "tcp"

// Server is a simple wrapper around net tcp server functionalities.
type Server struct {
	port     string
	listener net.Listener
}

func NewServer(port string) (*Server, error) {
	listener, err := net.Listen(protocol, port)
	if err != nil {
		return nil, err
	}

	return &Server{
		port:     port,
		listener: listener,
	}, nil
}

func (s *Server) GetConnection() (net.Conn, error) {
	connection, err := s.listener.Accept()
	if err != nil {
		return nil, err
	}

	return connection, nil
}
