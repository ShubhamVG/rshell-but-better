package server

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	. "../utils"
)

// ============================Exportables==========================

type Server struct {
	JoinedConnections map[UniqueConnAddr]net.Conn
	RequestsToSend    Queue[Request]
	ReceivedResponses Queue[Request] // this is stupid
	Address           string
	Port              uint
	// NextPort          uint        // to be implemented or dropped later
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
	go srvr.handleSends()
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

		uniqueAddr := getUniqueConnAddr(&conn)
		srvr.JoinedConnections[uniqueAddr] = conn
	}
}

func (srvr *Server) destructivelyCloseAllConnections() {
	for _, conn := range srvr.JoinedConnections {
		conn.Close()
	}
}

func (srvr *Server) handleReceives() {
	for uniqueAddr, conn := range srvr.JoinedConnections {
		buffer := make([]byte, 1024)
		conn.SetReadDeadline(time.Now().Add(readTimeLimit))
		n, err := conn.Read(buffer)

		switch err {
		case nil:
			// TODO
			receivedReq := parseIntoRequest(uniqueAddr, buffer[:n])
			srvr.ReceivedResponses.Enqueue(receivedReq)
		case os.ErrDeadlineExceeded:
			continue
		case net.ErrClosed:
			// TODO (maybe remove the connection?)
		}
	}
}

// Closes the connection and removes it from JoinedConnections
func (srvr *Server) removeConnection(connPtr *net.Conn) {
	uniqueAddr := getUniqueConnAddr(connPtr)
	(*connPtr).Close()

	delete(srvr.JoinedConnections, uniqueAddr)
}

// TODO
func (srvr *Server) handleSends() {
}
