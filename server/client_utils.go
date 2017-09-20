package main

import (
	"net"
	"bufio"
	"fmt"
	"io"
)

func send_msg(client *user, msg string) {
	for i := range serv.connections {
		fmt.Fprintf(serv.connections[i],
			client.username+" ("+client.curr_chan+"): "+msg)
	}
}

func handling_client(client *user) {
	fmt.Println("New User connected")
	for {
		msg, err := client.recv.ReadString('\n')
		if err == io.EOF {
			fmt.Println(client.username+": "+"disconnected.")
			return
		}
		fmt.Print(client.username + ": " + msg)
		send_msg(client, msg)
	}
}

func create_client(con net.Conn) *user {
	var client *user

	client = new(user)
	client.con = con
	client.recv = bufio.NewReader(con)
	client.username = "Anonymous"
	client.curr_chan = "Global"
	return client
}
