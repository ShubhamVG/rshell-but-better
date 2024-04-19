package server

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ShubhamVG/rshell-but-better/internal/datastructs"
	"github.com/ShubhamVG/rshell-but-better/internal/network"
)

// ============================Exportables==========================

type Server struct {
	JoinedConnections map[network.UniqueConnAddr]net.Conn
	ReceivedResponses datastructs.Bucket[network.Response]
	Address           string
	Port              uint
}

func NewServer(addr string, port uint) Server {
	joinedConnections := map[network.UniqueConnAddr]net.Conn{}
	receivedResponses := datastructs.NewBucket[network.Response](100)

	return Server{
		Address:           addr,
		Port:              port,
		JoinedConnections: joinedConnections,
		ReceivedResponses: receivedResponses,
	}
}

func (srvr *Server) Send(reqPtr *network.Request) error {
	req := *reqPtr
	uniqAddr := req.UniqueAddr

	// Stitching the status code and payload and then sending it
	if conn, ok := srvr.JoinedConnections[uniqAddr]; ok {
		payloadBuffer := []byte{req.Status}
		reqBuffer := []byte(req.Payload)
		payloadBuffer = append(payloadBuffer, reqBuffer...)
		conn.SetWriteDeadline(time.Now().Add(writeTimeLimit))

		if _, err := conn.Write(payloadBuffer); err != nil {
			return err
		}

		return nil
	}

	return fmt.Errorf("connection not in JoinedConnection")
}

func (srvr *Server) Start() error {
	addr := fmt.Sprintf("%s:%d", srvr.Address, srvr.Port)
	listener, err := net.Listen("tcp", addr)

	if err != nil {
		return err
	}

	defer srvr.notifyAndCloseAllConnections()
	defer listener.Close()

	go srvr.acceptConnections(&listener)
	go srvr.handleReceives()

	// Prevents the process from exiting right away
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
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
			continue
		}

		uniqAddr := network.GetUniqueConnAddr(&conn)
		srvr.JoinedConnections[uniqAddr] = conn
	}
}

// Receive and call processResponse
func (srvr *Server) handleReceives() {
	buffer := make([]byte, 1024)

	for uniqAddr, conn := range srvr.JoinedConnections {
		conn.SetReadDeadline(time.Now().Add(readTimeLimit))
		n, err := conn.Read(buffer)

		switch err {
		case os.ErrDeadlineExceeded:
			continue
		case net.ErrClosed:
			srvr.removeConnection(&conn)
			continue
		}

		receivedResponse := network.ParseIntoResponse(uniqAddr, buffer[:n])
		srvr.ReceivedResponses.Append(receivedResponse)
		srvr.processResponse(receivedResponse)
	}
}

func (srvr *Server) notifyAndCloseAllConnections() {
	for _, conn := range srvr.JoinedConnections {
		conn.Write([]byte{network.REQUESTING_CLOSE})
		conn.Close()
	}
}

func (srvr *Server) processResponse(response network.Response) {
	switch response.Status {
	case network.PING:
		// TODO
	case network.REQUESTING_CLOSE:
		if conn, ok := srvr.JoinedConnections[response.UniqueAddr]; ok {
			srvr.removeConnection(&conn)
		}
	}
}

// Closes the connection and removes it from JoinedConnections
func (srvr *Server) removeConnection(connPtr *net.Conn) {
	uniqAddr := network.GetUniqueConnAddr(connPtr)
	(*connPtr).Close()

	delete(srvr.JoinedConnections, uniqAddr)
}
