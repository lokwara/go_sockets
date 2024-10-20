package main

import (
	"fmt"
	"net"
)

func main() {
	// Listen on TCP port 8080
	ln, err := net.Listen("tcp", ":8080") //err is error, ln is listener object
	//ln listens for connection requests from clients (like a phone line waiting for an incoming call)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer ln.Close()

	fmt.Println("Server is listening on port 8080...")

	//Accept a connection
	//conn is short for connection, between two end points, client and a server.
	conn, err := ln.Accept()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()

	// Send a message to the client
	message := "Hello from the Server!"
	conn.Write([]byte(message))
}
