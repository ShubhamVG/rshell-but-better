package main

import (
	"../../client"
)

func main() {
	receiver, err := client.NewClient("localhost", "8080")

	for err != nil {
		receiver, err = client.NewClient("localhost", "8080")
	}

	receiver.Communicate()
}
