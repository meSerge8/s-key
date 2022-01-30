package client

import (
	"fmt"
	"net"
	"server/skey"
)

var (
	passwd  string
	address string
	sk      *skey.Skey
	id      *uint32
	con     net.Conn
)

func connect() error {
	var err error
	con, err = net.Dial("tcp", address)
	if err != nil {
		return err
	}
	defer con.Close()

	if err := identify(); err != nil {
		return err
	}

	fmt.Println("Now you can start messaging!")
	return messaging()
}

func identify() error {
	if id == nil {
		fmt.Println("Init session")
		return create()
	} else {
		fmt.Println("Restore last session with id =", *id)
		return restore()
	}
}

func messaging() error {
	var msg string

	for {
		fmt.Print(">")
		fmt.Scan(&msg)
		if msg == "exit" {
			break
		}
		if err := compose(msg); err != nil {
			return err
		}
	}
	return close()
}
