package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"net"
	"strings"
)

const (
	connHost = "localhost"
	connPort = "8999"
	connType = "tcp"
)

func main() {
	fmt.Println("Starting " + connType + " server on " + connHost + ":" + connPort)
	
	l, err := net.Listen(connType, connHost+":"+connPort)
	if err != nil {
		fmt.Print("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()
	
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error connecting:", err.Error())
			return
		}
		fmt.Println("Client connected.")
		
		fmt.Println("Client " + c.RemoteAddr().String() + " connected.")
		
		go handleConnection(c)
	}
}

func handleConnection(conn net.Conn) {
	buffer, err := bufio.NewReader(conn).ReadBytes('\n')
	// buffer, err := bufio.NewReader(conn).ReadString('\n')
	
	if err != nil {
		fmt.Println("Client left.")
		conn.Close()
		return
	}
	stringReceived := string(buffer[:len(buffer)-1])
	// convert CRLF to LF
	
	checkBool := strings.TrimRight(stringReceived, "\r\n") == "quit"
	log.Println("Client message:", stringReceived)

	if checkBool {
		os.Exit(1)
	}
	
	conn.Write(buffer)
	//conn.Write([]byte(buffer))
	
	handleConnection(conn)
}