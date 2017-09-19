package main

import (
	"bufio"
	"net"
)

type user struct {
	username	string
	con			net.Conn
	recv		*bufio.Reader
}



