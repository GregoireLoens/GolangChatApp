package main

import (
	"net"
	"bufio"
	"os"
	"fmt"
)

func read_serv(conn net.Conn){
	for {
		readserv, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("\n"+readserv)
	}
}

func client_loop(conn net.Conn) {
	go read_serv(conn)
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text)
	}
}

func conn_to_serv(connect string) net.Conn {
	serv, err := net.Dial("tcp", connect)
	if err != nil {
		println("Error with conection")
		return nil
	}
	return serv
}
