package main

import (
	"fmt"
	"net"
)

func main() {
	//Connect to Server on localhost:8080
	conn, err := net.Dial("tcp", "localhost:8080") //net.dial is a function in Go that is used to establish a network connection from client to server
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer conn.Close()

	//Receive message from Server
	buffer := make([]byte, 1024) //make allocates and initializes slices and maps.
	//[]byte creates a slice where ach element is a byte
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Server says:", string(buffer[:n]))
}
