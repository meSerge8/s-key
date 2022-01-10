package server

import (
	"log"
	"net"
	"server/skey"
)

type Server struct {
	Address_port string
	Passphrase   string
	iterations   int
	seed         int
	clientKeys   []skey.Key
}

func New(addr, pass string) Server {
	return Server{addr, pass, 1000, 10, make([]skey.Key, 10)}
}

func (s Server) Launch() (err error) {
	socket, err := net.Listen("tcp", s.Address_port)
	if err != nil {
		return err
	}
	log.Println("Listening at", s.Address_port)

	for id := 0; true; id++ {
		conn, err := socket.Accept()
		if err != nil {
			return err
		}
		go handle(conn)
	}

	return
}

func (s Server) Exit() error {
	return nil
}

func (s Server) Keyinit() error {
	return nil
}

func handle(conn net.Conn) {
	defer conn.Close()
	conn.Write([]byte("hello\naskdlasdasjdkdjaskdlasjdkasdjasdjasdjaskldjasdaskldjadjasldjsakdjaskjdakdaskjdlakdjaskdjaslkdjaskdjaskldjakdjaskldjasldjaslkdjadjaskldjaskdlakldjasdasdjaskldasjdklasjdaskldjaskldjasdklajdasdklajdlakdjasdakdjasdjaldkasjdaskldjdkasdjaslk"))
}
