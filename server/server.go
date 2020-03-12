package server

import (
	"fmt"
	"net"
	"os"
)

func createListenner(port string, ipAddr ...string) net.Listener {
	//init server
	const connType string = "tcp"
	ip := getOutboundIP().String()
	if len(ipAddr) > 0 {
		ip = ipAddr[0]
	}
	println("IP equal: ", ip)
	// Listen for incoming connections.
	l, err := net.Listen(connType, ip+":"+port)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	return l
}

//StartServer allow to start/run a server in tcp on the localhost:8080
func StartServer(port string, ipAddr ...string) {
	//init server
	l := createListenner(port, ipAddr...)
	//close server when server stop
	defer l.Close()

	//Listen for an incoming connection.
	for {
		conn, err := l.Accept()
		if err != nil {
			println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		println(conn)
		// Handle connections in a new goroutine.
		go createRoom(conn)
	}
}

func createRoom(conn net.Conn) {
	defer conn.Close()
	for {
		// Make a buffer to hold incoming data.
		buf := make([]byte, 1024)
		// Read the incoming connection into the buffer.
		resLen, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}
		//print message sent by the client
		println("Client: ", string(buf), " nb byte read = ", resLen)
		// Send a response back to person contacting us.
		conn.Write([]byte("Message received."))
		// Close the connection when you're done with it.
	}
}
