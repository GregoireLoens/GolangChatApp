package main

import (
	"bufio"
	"net"
)

type server struct {
	connections []net.Conn

}

type user struct {
	username	string
	con			net.Conn
	curr_chan	string
	recv		*bufio.Reader
}



