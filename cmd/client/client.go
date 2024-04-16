package main

import (
	"fmt"
	"os"

	"github.com/ShubhamVG/rshell-but-better/internal/client"
)

func main() {
	server, err := client.NewClient("localhost", "8080")

	if err != nil {
		fmt.Println("Failed starting the server.")
		os.Exit(1)
	}

	server.Communicate()
}
