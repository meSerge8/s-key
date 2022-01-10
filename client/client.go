package client

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"server/skey"
)

type Client struct {
	Address_port string
	Passphrase   string
	iterations   int
	seed         int
	clientKeys   []skey.Key
}

func New(addr, pass string) Client {
	return Client{addr, pass, 1000, 10, make([]skey.Key, 10)}
}

func (s Client) Launch() (err error) {
	conn, err := net.Dial("tcp", "localhost:8080")
	defer conn.Close()
	if err != nil {
		return err
	}

	msg, err := ReadAll(conn)
	if err != nil {
		return err
	}

	fmt.Println(msg)

	return nil
}

func checkErr(err error) {

	if err != nil {

		log.Fatal(err)
	}
}

func ReadAll(c net.Conn) (string, error) {
	buffer := bytes.NewBuffer([]byte{})

	_, err := buffer.ReadFrom(c)
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
