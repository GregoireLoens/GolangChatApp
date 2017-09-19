package main

import "net"

const (
	HOST = "127.0.0.1"
	PORT = ":4242"
)

var connections []net.Conn

func main() {
	listen := init_server()
	server_loop(listen)
}
