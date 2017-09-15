package main

import "bufio"

type user struct {
	username	string
	userID		int
	recv		*bufio.Reader
	send		*bufio.Writer
}


type server struct {
	list *user
}
