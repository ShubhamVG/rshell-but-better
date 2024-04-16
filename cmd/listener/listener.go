package main

import (
	"fmt"
	"os"

	"../../server"
)

func main() {
	listener := server.NewServer("localhost", 8080)

	if err := listener.Start(); err != nil {
		fmt.Println("Failed to start listener.")
		os.Exit(1)
	}
}
