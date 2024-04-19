package main

import (
	"fmt"

	"github.com/ShubhamVG/rshell-but-better/internal/server"
)

func main() {
	fmt.Println("Starting server...")
	server := server.NewServer("localhost", 8080)
	server.Start()
}
