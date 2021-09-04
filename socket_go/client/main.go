package main

import (
	"bufio"
	"log"
	"fmt"
	"os"
	"net"
)

const (
	connHost = "localhost"
	connPort = "8999"
	connType = "tcp"
)

func main() {
	fmt.Println("Connecting to " + connType + " server " + connHost + ":" + connPort)
	
	conn, err := net.Dial(connType, connHost+":"+connPort)
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}
	
	reader := bufio.NewReader(os.Stdin)
	
	for {
		fmt.Println("Text to send: ")
		
		input, _ := reader.ReadString('\n')
		
		// Send to socket connection.
		conn.Write([]byte(input))
		
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// message, _ := bufio.NewReader(conn).ReadBytes('\n')
		
		log.Print("Server relay:", message)
	}
}