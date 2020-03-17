package client

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

//connect: connect to server
func connect(ipAddr string, port string) net.Conn {
	client, err := net.Dial("tcp", ipAddr+":"+port)
	if err != nil {
		log.Fatal("Error during the connection :", err)
	}
	return client
}

//sendText: get and send the text from the stdin
func sendText(client net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		//buffer that store server's response
		text := getLine(scanner)
		if len(text) != 0 {
			//the text is sent to the server here !
			client.Write([]byte(text))
		} else {
			break
		}
	}
}

//StartClient : Start a simple client that will connect to the server and send text
func StartClient(ipAddr string, port string) {
	//connect client
	client := connect(ipAddr, port)
	//async text input
	go sendText(client)
	for {
		// Read the incoming connection into the buffer.
		client.SetReadDeadline(time.Now().Add(30 * time.Microsecond))
		buf := make([]byte, 1024)
		_, err := client.Read(buf)
		if err != nil {
			if err == io.EOF {
				// io.EOF, etc
				println("EOF detected !!!")
				return
			} else if err.(*net.OpError).Timeout() {
				//println("TIMEOUT")
				// no status msgs
				// note: TCP keepalive failures also cause this; keepalive is on by default
				continue
			}
			fmt.Println("Error reading:", err.Error())
		}
		//print message sent by the server
		println(string(buf))
	}
}
