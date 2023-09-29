package client

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func Client() {
	register()
}

func register() {
	fmt.Println("REGISTERING ...")
	signalAddress := "x03.gazer.cloud:8585"
	localAddress := ":8585" // default port

	remote, _ := net.ResolveUDPAddr("udp", signalAddress)
	local, _ := net.ResolveUDPAddr("udp", localAddress)
	conn, _ := net.ListenUDP("udp", local)
	bytesWritten, err := conn.WriteTo([]byte("HELLO"), remote)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	fmt.Println(bytesWritten, " bytes written")

	//listen(conn, local.String())
}

func listen(conn *net.UDPConn, local string) {
	for {
		fmt.Println("listening")
		buffer := make([]byte, 1024)
		bytesRead, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("[ERROR]", err)
			continue
		}

		fmt.Println("recevied UDP frame: ", string(buffer[0:bytesRead]))
		if string(buffer[0:bytesRead]) == "Hello!" {
			continue
		}

		for _, a := range strings.Split(string(buffer[0:bytesRead]), ",") {
			if a != local {
				go chatter(conn, a)
			}
		}
	}
}

func chatter(conn *net.UDPConn, remote string) {
	addr, _ := net.ResolveUDPAddr("udp", remote)
	for {
		conn.WriteTo([]byte("Hello!"), addr)
		fmt.Println("sent Hello! to ", remote)
		time.Sleep(5 * time.Second)
	}
}
