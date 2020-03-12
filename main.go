package main

import (
	"os"

	"./client"
	"./server"
)

func main() {
	if len(os.Args) < 3 {
		os.Exit(1)
	} else {
		if os.Args[1] == "server" {
			println("start server here")
			if len(os.Args) == 3 {
				server.StartServer(os.Args[2])
			} else {
				server.StartServer(os.Args[2], os.Args[3])
			}
		}
		if os.Args[1] == "client" {
			println("Start client here")
			client.StartClient(os.Args[3])
		}
	}
}
