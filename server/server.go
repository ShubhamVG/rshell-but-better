package server

import (
	"fmt"
	"net"

	"../utils"
)

type Server struct {
	JoinedConnections utils.Set[string]
	SendRequests      utils.Queue[string]
	ReceivedRequests  utils.Queue[string]
	Address           string
	Port              uint
	NextPort          uint
}

func NewServer(
	addr string,
	port, nextPort uint,
) Server {
	return Server{
		Address:  addr,
		Port:     port,
		NextPort: nextPort,
	}
}

// TODO
func (serv *Server) processRequest(receivedBuffer []byte) []byte {
	return receivedBuffer
}

func (serv *Server) Start() error {
	addr := fmt.Sprintf("%s:%d", serv.Address, serv.Port)
	listener, err := net.Listen("tcp", addr)

	if err != nil {
		// TODO
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()

		if err != nil {
			// TODO
		}

		receivedBuffer := make([]byte, 1024)
		n, err := conn.Read(receivedBuffer)

		if err != nil {
			// TODO
		}

		replyBytes := serv.processRequest(receivedBuffer[:n])
		_, err = conn.Write(replyBytes)

		if err != nil {
			// TODO
		}
	}
}
