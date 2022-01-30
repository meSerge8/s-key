package skeydb

import (
	"io/ioutil"
	"math/rand"
	"os"
	"server/skey"
	"server/utils"
	"sync"
	"time"
)

type SkeyDB struct {
	pass       string
	iterations uint32

	path  string
	mutex sync.Mutex

	clients map[uint32]skey.Key
}

func New(filePath, pass string, iterations uint32) (SkeyDB, error) {

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			return SkeyDB{}, err
		}
		file.Close()
	}

	clients, err := upload(filePath)
	if err != nil {
		return SkeyDB{}, err
	}

	return SkeyDB{
		pass:       pass,
		iterations: iterations,
		path:       filePath,
		clients:    clients}, nil
}

func upload(path string) (map[uint32]skey.Key, error) {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	clients := make(map[uint32]skey.Key, len(bs)/12)
	for i := 0; i < len(bs)/12; i++ {
		pos := 12 * i
		id := utils.FromBytes32(bs[pos : pos+4])
		key := utils.FromBytes64(bs[pos+4 : pos+12])
		clients[id] = skey.Key(key)
	}

	return clients, nil
}

func (sk *SkeyDB) GetClients() []Client {
	cs := make([]Client, len(sk.clients))
	var i int

	for id, key := range sk.clients {
		cs[i] = Client{id, key}
		i++
	}

	return cs
}

func (sk *SkeyDB) Save() error {
	file, err := os.Create(sk.path)
	if err != nil {
		return err
	}

	for id, key := range sk.clients {
		id := utils.ToBytes32(id)
		key := utils.ToBytes64(uint64(key))

		_, err = file.Write(append(id, key...))
		if err != nil {
			return err
		}
	}

	return nil
}

func (sk *SkeyDB) RefreshKey(c Client) (cl Client, seed uint32) {
	seed = getRand()
	s := skey.New(sk.pass, int(seed), int(sk.iterations))
	sk.clients[c.id] = s.GetServerInit()
	return Client{c.id, s.GetServerInit()}, seed
}

func (sk *SkeyDB) Find(id uint32) (Client, bool) {
	if _, ok := sk.clients[id]; !ok {
		return Client{}, false
	}
	return Client{id, sk.clients[id]}, true
}

func (sk *SkeyDB) Create() (c Client, seed uint32) {
	id := getID(sk.clients)
	seed = getRand()
	s := skey.New(sk.pass, int(seed), int(sk.iterations))
	return Client{id, s.GetServerInit()}, seed
}

func getID(cs map[uint32]skey.Key) uint32 {
	var i uint32
	for {
		if _, ok := cs[i]; !ok {
			break
		}
		i++
	}
	return i
}

func (sk *SkeyDB) SaveClient(cl Client) {
	sk.clients[cl.id] = cl.key
}

func getRand() uint32 {
	rand.Seed(time.Now().UnixNano())
	return rand.Uint32()
}
