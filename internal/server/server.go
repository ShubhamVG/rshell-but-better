package server

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	. "../data_structures"
	. "../network"
)

// ============================Exportables==========================

type Server struct {
	JoinedConnections map[UniqueConnAddr]net.Conn
	ReceivedResponses Queue[Response] // this is stupid
	Address           string
	Port              uint
}

func NewServer(addr string, port uint) Server {
	joinedConnections := map[UniqueConnAddr]net.Conn{}
	receivedResponses := NewQueue[Response]()
	return Server{
		Address:           addr,
		Port:              port,
		JoinedConnections: joinedConnections,
		ReceivedResponses: receivedResponses,
	}
}

func (srvr *Server) Send(reqPtr *Request) error {
	req := *reqPtr
	uniqAddr := req.UniqueAddr

	// TODO: Write documentation
	if conn, ok := srvr.JoinedConnections[uniqAddr]; ok {
		payload := []byte{req.Status}
		reqBuffer := []byte(req.Payload)
		payload = append(payload, reqBuffer...)
		payloadBuffer := []byte(req.Payload)
		_, err := conn.Write(payloadBuffer)

		if err != nil {
			return err
		}

		return nil
	}

	return fmt.Errorf("Connection not in JoinedConnection")
}

func (srvr *Server) Start() error {
	addr := fmt.Sprintf("%s:%d", srvr.Address, srvr.Port)
	listener, err := net.Listen("tcp", addr)

	if err != nil {
		return err
	}

	defer srvr.destructivelyCloseAllConnections() // maybe useless
	defer listener.Close()

	go srvr.acceptConnections(&listener)
	go srvr.handleReceives()

	// Prevents the process from exiting right away
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sig

	return nil
}

// ==============================Internals============================

// Accept connections, add them to the JoinedConnections map
func (srvr *Server) acceptConnections(lstnrPtr *net.Listener) {
	listener := *lstnrPtr

	for {
		conn, err := listener.Accept()

		if err != nil {
			// TODO
		}

		uniqAddr := GetUniqueConnAddr(&conn)
		srvr.JoinedConnections[uniqAddr] = conn
	}
}

func (srvr *Server) destructivelyCloseAllConnections() {
	for _, conn := range srvr.JoinedConnections {
		conn.Close()
	}
}

// Receive and call processResponse
func (srvr *Server) handleReceives() {
	buffer := make([]byte, 1024)

	for uniqAddr, conn := range srvr.JoinedConnections {
		// buffer := make([]byte, 1024)
		conn.SetReadDeadline(time.Now().Add(readTimeLimit))
		n, err := conn.Read(buffer)

		switch err {
		case os.ErrDeadlineExceeded:
			continue
		case net.ErrClosed:
			// TODO (maybe remove the connection?)
			continue
		}

		receivedResponse := ParseIntoResponse(uniqAddr, buffer[:n])
		srvr.ReceivedResponses.Enqueue(receivedResponse)
		srvr.processResponse(receivedResponse)
	}
}

func (srvr *Server) processResponse(response Response) {
	switch response.Status {
	case PING:
		// TODO
	case REQUESTING_CLOSE:
		if conn, ok := srvr.JoinedConnections[response.UniqueAddr]; ok {
			srvr.removeConnection(&conn)
		} else {
			// TODO
		}
	case OUTPUT: // TODO or TO DROP
	case OUTPUT_WITH_ERROR: // TODO or TO DROP
	}
}

// Closes the connection and removes it from JoinedConnections
func (srvr *Server) removeConnection(connPtr *net.Conn) {
	uniqAddr := GetUniqueConnAddr(connPtr)
	(*connPtr).Close()

	delete(srvr.JoinedConnections, uniqAddr)
}
