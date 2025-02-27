package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Print("Listening on the port: 6379")

	// Crete a new server or TCP listener
	l, err := net.Listen("tcp", ":6379") // l: This stores the TCP listener object (server socket).
	if err != nil {
		fmt.Print(err) // err: Stores any error encountered in net.Listen() or l.Accept().
		return
	}

	// Listening for connection or Receiving request
	conn, err := l.Accept() // conn: Stores the active client connection (TCP) once a client connects.
	// l.Accept(): Waits for a client to connect to the server. When a client connects, it returns a connection object (conn).
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close() // Ensures the connection closes properly when the function exits.
	// Prevents memory leaks by cleaning up resources automatically.

	for {
		resp := NewResp(conn)
		value, err := resp.Read()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(value)

		// ignore request and send back a PONG
		conn.Write([]byte("+OK\r\n"))
	}
}
