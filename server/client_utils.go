package main

import (
	"bufio"
	"net"
)

func init_client(client *user, con *net.Conn){
	client.username = "Anonymous"
	client.recv = bufio.NewReader(*con)
	client.send = bufio.NewWriter(*con)
}
