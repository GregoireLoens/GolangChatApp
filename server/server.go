package main

import (
	"fmt"
	"net"
	"bufio"
)

func server_loop(conn net.Conn) {
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print(message)
	}
}

func init_server() net.Conn {
	fmt.Println("Launching Server...")

	listen, err := net.Listen("tcp", ":4242")
	if err != nil {
		fmt.Println(err)
	}
	conn, _ := listen.Accept()
	return conn
}
