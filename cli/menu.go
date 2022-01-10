package cli

import (
	"fmt"
	"strings"
)

type menuAction func() (menuAction, error)

var handler cliHandler

func Start(h cliHandler) (err error) {
	fmt.Print(helloStr)

	handler = h
	for action := showMenu; action != nil; action, err = action() {
		if err != nil {
			return err
		}
	}

	return
}

func showMenu() (m menuAction, err error) {
	fmt.Print(menuStr)

	switch strings.TrimSpace(getInput()) {
	case "1":
		m = launch

	case "2":
		m = keyinit

	case "3":
		m = exit
	}

	return showMenu, nil
}

func launch() (menuAction, error) {
	fmt.Print(launchStr)
	return showMenu, handler.Launch()
}

func keyinit() (menuAction, error) {
	fmt.Print(keyInitStr)
	return showMenu, handler.Keyinit()
}

func exit() (menuAction, error) {
	fmt.Print(exitStr)
	return nil, handler.Exit()
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
