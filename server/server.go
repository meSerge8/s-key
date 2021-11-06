package main

import (
	"fmt"
	"os"
)

var (
	address_port = "localhost:8080"
	passphrase   = "passphrase"
	iterations   = 1000
	seed         string
)

func launch() error {
	getAddresPort()

	fmt.Println(address_port)

	return nil
}

func getAddresPort() {
	if len(os.Args) != 2 {
		return
	}

	address_port = os.Args[1]
}
