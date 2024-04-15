package server

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	. "../commons"
	. "../utils"
)

// ============================Exportables==========================

type Server struct {
	JoinedConnections map[UniqueConnAddr]net.Conn
	ReceivedResponses Queue[Response] // this is stupid
	Address           string
	Port              uint
	// NextPort          uint        // to be implemented or dropped later
}

func (srvr *Server) Send(reqPtr *Request) error {
	req := *reqPtr
	uniqAddr := req.UniqueAddr

	if conn, ok := srvr.JoinedConnections[uniqAddr]; ok {
		contentBuffer := []byte(req.Body)
		_, err := conn.Write(contentBuffer)

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
	for uniqAddr, conn := range srvr.JoinedConnections {
		buffer := make([]byte, 1024)
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
	}
}

// Closes the connection and removes it from JoinedConnections
func (srvr *Server) removeConnection(connPtr *net.Conn) {
	uniqAddr := GetUniqueConnAddr(connPtr)
	(*connPtr).Close()

	delete(srvr.JoinedConnections, uniqAddr)
}
