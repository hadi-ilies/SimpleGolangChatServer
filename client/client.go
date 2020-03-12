package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

//connect to server
func connect(ipAddr string, port string) net.Conn {
	client, err := net.Dial("tcp", ipAddr+":"+port)
	if err != nil {
		log.Fatal("Error during the connection :", err)
	}
	return client
}

//StartClient Start a simple client that will connect to the server and send text
func StartClient(ipAddr string, port string) {
	//connect client
	client := connect(ipAddr, port)
	// To create dynamic array
	messages := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		//buffer that store server's response
		buf := make([]byte, 1024)
		text := getLine(scanner)
		if len(text) != 0 {
			fmt.Println("message: \"", text, "\" sent !")
			client.Write([]byte(text))
			messages = append(messages, text)
		} else {
			break
		}
		// Read the incoming connection into the buffer.
		resLen, err := client.Read(buf)
		if err != nil {
			log.Fatalln("Error has been read")
		}
		println("Server: ", string(buf), " nb byte read = ", resLen)
	}
	println(messages)
}
