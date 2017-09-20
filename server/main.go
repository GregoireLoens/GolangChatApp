package main

const (
	HOST = "127.0.0.1"
	PORT = ":4242"
)

var serv server

func main() {
	listen := init_server()
	server_loop(listen)
}
