package main

import (
	"errors"
	"fmt"
	"strings"
)

type menuAction func() (menuAction, error)

func start() (err error) {
	fmt.Print(helloStr)

	actionFoo := menuAction(doMenu)
	for {
		actionFoo, err = actionFoo()

		if err != nil {
			break
		}
	}

	if err.Error() == "Exit" {
		err = nil
	}

	return
}

func doMenu() (menuAction, error) {
	fmt.Print(menuStr)

	str := strings.TrimSpace(getInput())
	foo := menuAction(doMenu)

	switch str {
	case "1":
		foo = doLaunch

	case "2":
		foo = doKeyInit

	case "3":
		foo = doExit
	}

	return foo, nil
}

func doLaunch() (foo menuAction, err error) {
	fmt.Print(launchStr)
	return doMenu, launch()
}

func doKeyInit() (foo menuAction, err error) {
	fmt.Print(keyInitStr)
	return doMenu, keyInit()
}

func doExit() (foo menuAction, err error) {
	fmt.Print(exitStr)
	return nil, errors.New("Exit")
}

func getInput() (res string) {
	fmt.Print(">")

	fmt.Scanln(&res)

	return res
}

const (
	helloStr = `______________________________________
S/KEY by Melnikov Dorofeev 9317{23,24}

`
	menuStr = `_____________
MENU

1. Launch
2. KeyInit
3. Exit

`
	launchStr = `_____________
LAUNCH

`
	keyInitStr = `_____________
KEY INIT

Enter new init values:

`
	exitStr = `_____________
EXIT

`
)
