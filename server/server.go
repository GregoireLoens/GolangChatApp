package main

import (
	"fmt"
	"net"
)

func server_loop(listen net.Listener) {
	for {
		con, _ := listen.Accept()
		serv.connections = append(serv.connections, con)
		client := create_client(con)
		go handling_client(client)
	}
}

func init_server() net.Listener {
	fmt.Println("Launching Server...")
	listen, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Server Launched")
	return listen
}
