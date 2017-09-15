package main

import (
	"fmt"
	"net"
)

func client_handling(con *net.Conn) {
	var (
		client user
	)
	if con == nil{
		return
	}
	init_client(&client, con)
	for {
			data, err := client.recv.ReadString('\n')
			if err != nil {
				println(client.username+": Disconnected.")
				return
				}
			print(client.username+": "+data)
	}
}

func server_loop(listen net.Listener) {
	for {
		con, err := listen.Accept()
		if err != nil {
			println(err)
		}
		go client_handling(&con)
	}
}

func init_server() net.Listener {
	fmt.Println("Launching Server...")

	listen, err := net.Listen("tcp", ":4242")
	if err != nil {
		fmt.Println(err)
	}
	return listen
}
