package server

import (
	"net"
	"server/skey"
)

func initSession(conn net.Conn) (skey.Key, error) {
	// for i, client := range s.clients {

	// }
	return 123, nil
}

func initClient(conn net.Conn) (err error) {

	// if err = recvHello(client); err != nil {
	// 	return
	// }

	// if err = sendId(client); err != nil {
	// 	return
	// }

	// if err = sendSeed(client); err != nil {
	// 	return
	// }

	return
}

// func recvHello(client Client) (err error) {
// 	var buf []byte
// 	n, err := client.conn.Read(buf)

// 	if err != nil {
// 		return
// 	}

// 	log.Println("recv:", buf)

// 	if string(buf[:n]) != "hello" {
// 		return errors.New("wrong hello pharse")
// 	}

// 	return
// }

// func sendId(client Client) (err error) {
// 	_, err = client.conn.Write([]byte(seed))

// 	if err != nil {
// 		return errors.New("send seed failed")
// 	}

// 	log.Println("send id:", client.id)

// 	return
// }

// func sendSeed(client Client) (err error) {

// 	_, err = client.conn.Write([]byte(seed))

// 	if err != nil {
// 		return errors.New("send seed failed")
// 	}

// 	log.Println("send seed:", seed)

// 	return
// }
// func checkPassword(i int, conn net.Conn) bool {
// 	buf := make([]byte, 128)
// 	_, err := conn.Read(buf)

// 	if err != nil {
// 		return false
// 	}

// 	passwd := string(buf)
// 	log.Println("Password:", passwd)

// 	return true
// }
