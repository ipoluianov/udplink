package main

import (
	"fmt"
	"os"

	"github.com/ipoluianov/udplink/client"
	"github.com/ipoluianov/udplink/server"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Error: wrong arguments")
		return
	}
	cmd := os.Args[1]
	fmt.Println("UdpLink")
	fmt.Println("-------")
	switch cmd {
	case "c":
		client.Client()
	case "s":
		server.Server()
	}
	fmt.Println("Finished")
}
