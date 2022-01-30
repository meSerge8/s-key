package server

import (
	"errors"
	"fmt"
	"log"
	"net"
	"server/server/skeydb"
)

type command struct {
	id         int
	foo        func(net.Conn, *skeydb.Client) (*skeydb.Client, error)
	needClient bool
}

var commands []command

func initCommands() {
	commands = []command{
		{0, create, false},
		{1, find, false},
		{2, refresh, true},
		{3, receive, true},
		{4, closeConnection, false},
	}
}

func handle(con net.Conn) {
	var client *skeydb.Client

	defer con.Close()
	defer db.Save()

	for {
		cmd, err := getCommand(con)
		if err != nil {
			break
		}

		if !canAccess(cmd, client) {
			fmt.Println("Access denied")
			break
		}

		if client, err = cmd.foo(con, client); err != nil {
			if client == nil {
				break
			}
			log.Println(err.Error())
		}

		db.SaveClient(*client)
	}

	fmt.Print("Connection closed\n\n")
}

func getCommand(con net.Conn) (command, error) {
	bs := make([]byte, 1)
	if _, err := con.Read(bs); err != nil {
		return command{}, err
	}

	com := int(bs[0])

	for _, c := range commands {
		if c.id == com {
			return c, nil
		}
	}

	return command{}, errors.New("No such command")
}

func canAccess(c command, cl *skeydb.Client) bool {
	if (c.needClient == true) && (cl == nil) {
		return false
	}

	return true
}
