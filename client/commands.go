package client

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"server/skey"
	"server/utils"
)

func create() error {
	if err := sendByte(0); err != nil {
		return err
	}

	i, err := readUint32()
	if err != nil {
		return err
	}
	id = new(uint32)
	*id = i

	seed, err := readUint32()
	if err != nil {
		return err
	}

	iterations, err := readUint32()
	if err != nil {
		return err
	}

	sk = new(skey.Skey)
	*sk = skey.New(passwd, int(seed), int(iterations))

	return getStatus()
}

func restore() error {
	if err := sendByte(1); err != nil {
		return err
	}

	if _, err := con.Write(utils.ToBytes32(*id)); err != nil {
		return err
	}

	return getStatus()
}

func compose(msg string) error {
	log.Println("compose")

	if sk.IsLast() {
		if err := refresh(); err != nil {
			return err
		}
	}

	if err := sendByte(3); err != nil {
		return err
	}

	if err := checkKey(); err != nil {
		return err
	}

	if _, err := con.Write(append([]byte(msg), 0)); err != nil {
		return err
	}

	return getStatus()
}

func refresh() error {
	log.Println("refresh")
	if err := sendByte(2); err != nil {
		return err
	}

	fmt.Println("good1")
	if err := checkKey(); err != nil {
		// log.Fatal(err)
		fmt.Println(err)

		if err.Error() != "out of keys" {
			return err
		}
	}
	// bb := make([]byte, 5)
	// con.Read(bb)
	// fmt.Println(bb)
	// os.Exit(0)

	fmt.Println("good3")

	seed, err := readUint32()
	if err != nil {
		return err
	}

	iterations, err := readUint32()
	if err != nil {
		return err
	}

	fmt.Println("good4")
	fmt.Println("PSWD:", passwd)
	fmt.Println("SEED:", seed)
	fmt.Println("ITER:", iterations)
	*sk = skey.New(passwd, int(seed), int(iterations))
	log.Println("refreshed")
	fmt.Println("good5")

	return getStatus()
}

func checkKey() error {
	keyBs := utils.ToBytes64(uint64(sk.GetCurrent()))
	if _, err := con.Write(keyBs); err != nil {
		return err
	}

	if err := getStatus(); err != nil {
		return err
	}

	if _, err := sk.GetNext(); err != nil {
		return err
	}

	return nil
}

func close() error {
	if err := sendByte(4); err != nil {
		return err
	}
	return getStatus()
}

func getStatus() error {
	b := make([]byte, 1)
	if _, err := con.Read(b); err != nil {
		return err
	}

	switch b[0] {
	case 1:
		return nil
	case 0:
		reader := bufio.NewReader(con)
		bs, err := reader.ReadBytes(0)
		if err != nil {
			return err
		}
		return errors.New(string(bs))
	default:
		return errors.New("Unknown command from server")
	}
}

func sendByte(x byte) error {
	_, err := con.Write([]byte{x})
	return err
}

func readUint32() (uint32, error) {
	bs := make([]byte, 4)
	if _, err := con.Read(bs); err != nil {
		return 0, err
	}
	fmt.Println("Read32:", bs)
	return utils.FromBytes32(bs), nil
}
