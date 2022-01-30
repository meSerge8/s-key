package client

import "fmt"

var info = `Id:%d
Seed:%d
Iterations:%d
Keys left:%d
Current key:%d
`

func Launch() {
	address = scanValue("Enter ip-address")
	passwd = scanValue("Enter password")

	for {
		runMenu()
		if err := connect(); err != nil {
			fmt.Println(err.Error())
		}

	}
}

func scanValue(ask string) (str string) {
	for {
		fmt.Printf("%s:", ask)
		fmt.Scanf("%s\n", &str)
		if str != "" {
			return
		}
	}
}

func runMenu() {
	fmt.Println("____________________")
	fmt.Println("Password:", passwd)
	if sk != nil {
		s, i, l, c := sk.GetInfo()
		fmt.Printf(info, *id, s, i, l, c)
	}
	fmt.Println("____________________")
	fmt.Println("\nPress enter to connect ...")
	fmt.Scanf("\n")
}
