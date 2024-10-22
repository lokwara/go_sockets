package main

import (
	"fmt"
	"net"
	"time"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		// Send a message to the client
		message := "Hello from the server! The time is" + time.Now().String()
		_, err := conn.Write([]byte(message + "\n"))
		if err != nil {
			fmt.Println("Error writing to client:", err)
			return
		}

		// Pause for 5 Seconds before sending another message
		time.Sleep(5 * time.Second)
	}
}

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

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting Connection:", err)
			continue
		}

		fmt.Println("New Client connected!")

		// Handle each connection in a new goroutine
		go handleConnection(conn)
	}
}
