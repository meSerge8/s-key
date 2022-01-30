package server

import (
	"fmt"
	"log"

	"net"
	"server/server/skeydb"
	"strconv"
)

var (
	passwd     string
	iterations uint32
	address    string
	db         skeydb.SkeyDB
	filename   = "clients.txt"
)

func Launch() {
	listener, err := initServer()
	if err != nil {
		fmt.Println("Initialization failed:", err.Error())
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			break
		}

		fmt.Println("Accept new connection")
		go handle(conn)
	}
}

func initValues() {
	address = scanValue("Enter ip-address")
	passwd = scanValue("Enter password")
	for {
		iterStr := scanValue("Enter number of iterations")
		iterInt, err := strconv.Atoi(iterStr)
		if err == nil {
			iterations = uint32(iterInt)
			break
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

func initDB() error {
	var err error
	db, err = skeydb.New(filename, passwd, iterations)
	if err != nil {
		return err
	}
	return nil
}

func getListener() (*net.TCPListener, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", address)
	if err != nil {
		return nil, err
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return nil, err
	}

	return listener, nil
}

func initServer() (listener *net.TCPListener, err error) {
	initValues()
	initCommands()
	if err := initDB(); err != nil {
		return nil, err
	}

	if listener, err = getListener(); err != nil {
		return nil, err
	}

	fmt.Println("____________________")
	fmt.Printf(start, address, passwd, iterations)
	for _, v := range db.GetClients() {
		fmt.Println(v)
	}
	fmt.Println("____________________")

	return listener, nil
}

var start = `Server started successfully
IP-address: %s
Password: %s
Iterations: %d

CurrentUsers:
`
