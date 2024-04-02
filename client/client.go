package client

import (
	"net"
	"os/exec"
)

type Client struct {
	Address          string
	Port             string
	commandToExecute chan string
}

func NewClient(addr, port string) Client {
	commandBuffer := make(chan string, 1)
	return Client{Address: addr, Port: port, commandToExecute: commandBuffer}
}

func (client *Client) communicate() {
	addr := client.Address + ":" + client.Port
	conn, err := net.Dial("tcp", addr)

	if err != nil {
		// TODO
	}

	receivedBuffer := make([]byte, 1024)
	n, err := conn.Read(receivedBuffer)

	if err != nil {
		// TODO
	}

	bufferToSend := client.processRequest(receivedBuffer[:n])
	_, err = conn.Write(bufferToSend)
}

func (client *Client) executeCommand() {
	rawCommand := <-client.commandToExecute
	command, params := parseIntoCommandAndParams(rawCommand)
	output, err := exec.Command(command, params...).Output()

	if err != nil {
		// TODO
	}

	// TODO
}

// TODO
func (client *Client) processRequest(buffer []byte) []byte {
	return buffer
}
