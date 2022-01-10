package main

import (
	"fmt"
	"log"
	"server/client"
	"server/server"
)

func main() {

	var mode int
	fmt.Scanf("%d", &mode)

	if mode == 1 {
		s := server.New("localhost:8080", "passphrase")
		fmt.Println(s.Launch())
	}
	if mode == 2 {
		c := client.New("localhost:8080", "passphrase")
		err := c.Launch()
		if err != nil {
			log.Fatalln(err)
		}
	}
}
