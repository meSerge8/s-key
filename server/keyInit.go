package main

import (
	"fmt"
	"strconv"
)

func keyInit() (err error) {

	updatePassphrase()
	updateIterations()
	updateSeed()
	showUpdates()

	return nil
}

func updatePassphrase() {
	fmt.Println("Pass phrase (default: 'passphrase')")

	str := getInput()
	if str == "" {
		return
	}

	passphrase = str
}

func updateIterations() {
	fmt.Println("\nIterations (default: '1000')")

	str := getInput()
	if str == "" {
		return
	}

	itr, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println("Not int. Stay default ...")
		return
	}

	iterations = itr
}

func updateSeed() {
	fmt.Println("\nSeed (default: random seed)")

	str := getInput()

	if str == "" {
		return
	}

	seed = str
}

func showUpdates() {
	fmt.Printf("\nUpdated values:\n- passphrase: %s\n- iterations: %d\n- seed: %s\n",
		passphrase,
		iterations,
		seed)
}
