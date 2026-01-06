package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Hello World!")

	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Println("accepted a connection")
	fmt.Println(conn.LocalAddr().String())  // My address , i.e. address of the server created in this program by net.Listen()
	fmt.Println(conn.RemoteAddr().String()) // Address of who is connecting
}
