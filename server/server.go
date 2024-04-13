package server

import (
	"fmt"
	"net"

	. "../utils"
)

type Server struct {
	JoinedConnections Set[string]
	RequestsToSend    Queue[Request]
	ReceivedRequests  Queue[Request]
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
func (serv *Server) processRequest(connAddr string, receivedBuffer []byte) {
	statusCode := receivedBuffer[0]

	switch statusCode {
	case PING:
		serv.JoinedConnections.Add(connAddr)
	}
}

func (serv *Server) Start() error {
	addr := fmt.Sprintf("%s:%d", serv.Address, serv.Port)
	listener, err := net.Listen("tcp", addr)

	if err != nil {
		return fmt.Errorf("Failed to start server")
	}

	defer listener.Close()

	var reqToSend Request
	toFetchRequest := true
	var failedRequestCount uint8 = 0

	for {
		if toFetchRequest && serv.RequestsToSend.Len != 0 {
			reqToSend, _ = serv.ReceivedRequests.Dequeue()
		}

		conn, err := listener.Accept()
		connAddr := conn.LocalAddr().String()

		if err != nil {
			fmt.Println("Failed to accept connection: " + connAddr)
			continue
		}

		receivedBuffer := make([]byte, 1024)
		n, err := conn.Read(receivedBuffer)

		if err != nil {
			fmt.Println("Failed to read from " + connAddr)
			continue
		}

		serv.processRequest(connAddr, receivedBuffer[:n])

		if reqToSend.Addr == connAddr {
			_, err = conn.Write(reqToSend.ContentBuffer)

			if err != nil {
				fmt.Println("Failed to send bytes to " + connAddr)
				continue
			}

			failedRequestCount = 0
			toFetchRequest = true
		} else if failedRequestCount > failedSendRequestLimit {
			// Drop the connection if Addr is not connected else enqueue it

			if serv.JoinedConnections.Contains(reqToSend.Addr) {
				serv.RequestsToSend.Enqueue(reqToSend)
			}

			toFetchRequest = true
			failedRequestCount = 0
		} else {
			failedRequestCount++
		}
	}
}
