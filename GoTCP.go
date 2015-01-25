package main

import (
	"bytes"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
)

const (
	CONN_HOST = "localhost"
	CONN_TYPE = "tcp"
)

func main() {
	// Listen for incoming connections.
	port := flag.Int("port", 9100, "port to listen on")

	flag.Parse()

	fmt.Println("port:", *port)

	CONN_PORT := strconv.Itoa(*port)

	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	// echo message to console
	message := "Received ("
	message += strconv.Itoa(reqLen)
	message += " bytes): "
	n := bytes.Index(buf, []byte{0})
	message += string(buf[:n-1])
	// message += "\" ! Honestly I have no clue about what to do with your messages, so Bye Bye!\n"

	fmt.Println(message)
	// Send a response back to person contacting us.
	conn.Write([]byte("Message received."))
	// Close the connection when you're done with it.
	conn.Close()
}
