package skey

import (
	"encoding/binary"
	"errors"
	"strconv"

	"github.com/cespare/xxhash"
)

type Key uint64

type Skey struct {
	passphrase string
	seed       int
	iterations int
	Keys       []Key
	counter    int
}

func New(passphrase string, seed int, iterations int) Skey {
	if iterations < 1 {
		panic("iterations must be at least 1")
	}

	sk := Skey{
		passphrase: passphrase,
		seed:       seed,
		iterations: iterations}

	sk.getKeys()

	return sk
}

func (sk *Skey) GetServerInit() Key {
	return hash64(sk.Keys[0])
}

func (sk *Skey) GetCurrent() Key {
	return sk.Keys[sk.counter]
}

func (sk *Skey) GetNext() (Key, error) {
	sk.counter++

	if sk.counter == sk.iterations {
		sk.counter--
		return 0, errors.New("out of keys")
	}

	return sk.Keys[sk.counter], nil
}

func Check(new Key, old Key) bool {
	return hash64(new) == old
}

func (sk *Skey) IsLast() bool {
	return (sk.counter + 1) == sk.iterations
}

func (sk *Skey) getKeys() {
	sk.Keys = make([]Key, sk.iterations)
	sk.Keys[0] = sk.getFirstKey()
	for i := 1; i < sk.iterations; i++ {
		sk.Keys[i] = hash64(sk.Keys[i-1])
	}
	reverse(sk.Keys)
}

func (sk *Skey) getFirstKey() Key {
	h := xxhash.New()
	h.Write([]byte(sk.passphrase))
	h.Write([]byte(strconv.Itoa(sk.seed)))
	return Key(h.Sum64())
}

func hash64(k Key) Key {
	byteKey := make([]byte, 8)
	binary.LittleEndian.PutUint64(byteKey, uint64(k))
	return Key(xxhash.Sum64(byteKey))
}

func reverse(xs []Key) []Key {
	for i, j := 0, len(xs)-1; i < j; i, j = i+1, j-1 {
		xs[i], xs[j] = xs[j], xs[i]
	}
	return xs
}

func (sk *Skey) GetInfo() (seed, iterations, keysLeft int, curKey uint64) {
	for i, key := range sk.Keys {
		if key == sk.GetCurrent() {
			keysLeft = sk.iterations - i
			break
		}
	}

	return sk.seed, sk.iterations, keysLeft, uint64(sk.GetCurrent())
}
