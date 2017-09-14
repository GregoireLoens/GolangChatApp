package main

import (
	"net"
	"bufio"
	"os"
	"fmt"
)

func client_loop(conn net.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text + "\n")
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: "+message)
	}
}

func conn_to_serv(connect string) net.Conn {
	serv, err := net.Dial("tcp", connect)
	if err != nil {
		println(err)
	}
	return serv
}
