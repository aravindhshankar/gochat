package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
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

	readbuf := make([]byte, 100)
	writebuf := []byte("You've connected to " + conn.LocalAddr().String() + "\n")
	n, err := conn.Write(writebuf)
	if err != nil {
		fmt.Println("Couldn't write to Remote address : ", conn.RemoteAddr().String())
	}
	log.Printf("%d bytes written to %s \n", n, conn.RemoteAddr().String())

	for {
		_, err = conn.Write([]byte("$ ")) // prompt asking the user to input
		if err != nil {
			log.Fatal("Connection unstable")
		}
		n, err = conn.Read(readbuf)
		if err != nil {
			fmt.Println(err)
		}
		ostring := string(readbuf[0 : n-1]) // strips the newline at the end
		log.Printf("%d bytes read from %s \n", n, conn.RemoteAddr().String())
		fmt.Println(ostring)
	}
}
