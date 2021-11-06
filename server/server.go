package main

import "fmt"

func main() {
	err := start()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Bye ...")
}
