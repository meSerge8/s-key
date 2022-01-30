package main

import (
	"os"
	"server/client"
	"server/server"
)

func main() {

	switch os.Args[1] {
	case "server":
		server.Launch()
	case "client":
		client.Launch()
	}

}
