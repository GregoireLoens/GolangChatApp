package main

import "os"

func main() {
	connect := os.Args[1]
	serv := conn_to_serv(connect)
	client_loop(serv)
}