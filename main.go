package main

import (
	"fmt"
	"io"
	"net"
	"os"
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

	for { // This creates an infinite loop, meaning the server will continuously read incoming messages from the client.
		buf := make([]byte, 1024) // Creates a byte slice (buffer) of size 1024 bytes. It will store incoming message from client.

		// read message from client
		_, err = conn.Read(buf) // The underscore (_) ignores the number of bytes read because it’s not used in the logic.
		if err != nil {
			if err == io.EOF { // EOF (End of File): io.EOF: The client closed the connection, so we break the loop.
				break
			}
			fmt.Println("error reading from client: ", err.Error())
			os.Exit(1)
		}

		// No matter what message the client sends, the server always responds with +GoBolt says OK\r\n.
		conn.Write([]byte("+GoBolt says OK\r\n"))
		// This follows the Redis Simple String format, where a + prefix indicates a plain-text response.
		//The \r\n (carriage return + newline) is required by the Redis protocol.
	}
}
