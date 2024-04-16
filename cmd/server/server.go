package main

import (
	"fmt"

	"github.com/ShubhamVG/rshell-but-better/internal/server"
)

func main() {
	server := server.NewServer("localhost", 8080)
	fmt.Println("Starting server...")
	server.Start()
}
