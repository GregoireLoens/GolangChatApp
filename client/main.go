package main

import (
	"os"
	"fmt"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("USAGE:\n./ChatClient Adresse:Port")
		return
	}
	connect := os.Args[1]
	serv := conn_to_serv(connect)
	if serv == nil {
		return
	}
	client_loop(serv)
}