package main

import (
	"fmt"

	"web_server"
)

func main() {
	server, err := web_server.CreateServer(8000)
	if err != nil {
		panic(fmt.Sprintf("Error creating server [%s]...", err.Error))
	}

	server.Start()
}
