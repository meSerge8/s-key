package server

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"server/server/skeydb"
	"server/skey"
	"server/utils"
)

func create(con net.Conn, c *skeydb.Client) (*skeydb.Client, error) {
	c = new(skeydb.Client)
	var seed uint32
	*c, seed = db.Create()

	resp := utils.ToBytes32(c.GetId())
	resp = append(resp, utils.ToBytes32(seed)...)
	resp = append(resp, utils.ToBytes32(iterations)...)

	if _, err := con.Write(resp); err != nil {
		return nil, err
	}

	fmt.Printf("Create new client:\tid = %d\tseed = %d\n", c.GetId(), seed)

	return c, sendOk(con)
}

func find(con net.Conn, c *skeydb.Client) (*skeydb.Client, error) {
	idBs := make([]byte, 4)
	if _, err := con.Read(idBs); err != nil {
		return nil, err
	}
	id := utils.FromBytes32(idBs)

	if client, ok := db.Find(id); !ok {
		errStr := fmt.Sprintf("No client with id =%d", id)
		fmt.Println(errStr)
		return nil, sendFail(con, errStr)
	} else {
		c = new(skeydb.Client)
		*c = client
	}

	fmt.Printf("Find existing client:\tid = %d\n", id)

	return c, sendOk(con)
}

func refresh(con net.Conn, c *skeydb.Client) (*skeydb.Client, error) {
	if err := checkKey(con, c); err != nil {
		return nil, err

	}

	var seed uint32
	*c, seed = db.RefreshKey(*c)

	if _, err := con.Write(utils.ToBytes32(seed)); err != nil {
		return nil, err
	}

	if _, err := con.Write(utils.ToBytes32(iterations)); err != nil {
		return nil, err
	}

	fmt.Printf("Refresh keys:\tid = %d\tseed = %d\n", c.GetId(), seed)

	return c, sendOk(con)
}

func receive(con net.Conn, c *skeydb.Client) (*skeydb.Client, error) {
	if err := checkKey(con, c); err != nil {
		return nil, err
	}

	reader := bufio.NewReader(con)
	bs, err := reader.ReadBytes(0)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Receive message:\tid = %d\tmessage = %s\n", c.GetId(), string(bs))

	return c, sendOk(con)
}

func checkKey(con net.Conn, cl *skeydb.Client) error {
	keyBs := make([]byte, 8)
	if _, err := con.Read(keyBs); err != nil {
		return err
	}

	key := skey.Key(utils.FromBytes64(keyBs))

	if !cl.ConfirmKey(key) {
		return sendFail(con, fmt.Sprintf("Wronk key from %v\n", cl))
	}

	if err := sendOk(con); err != nil {
		return err
	}

	return nil
}

func closeConnection(con net.Conn, c *skeydb.Client) (*skeydb.Client, error) {
	if err := sendOk(con); err != nil {
		return nil, err
	}
	fmt.Printf("Client closed connection:\tid = %d\n", c.GetId())

	return nil, errors.New("Connection closed by client")
}

func sendOk(con net.Conn) error {
	_, err := con.Write([]byte{1})
	return err
}

func sendFail(con net.Conn, description string) error {
	fmt.Println("send fail")
	if _, err := con.Write([]byte{0}); err != nil {
		return err
	}

	bs := append([]byte(description), 0)
	if _, err := con.Write(bs); err != nil {
		return err
	}

	return errors.New(description)
}
