package skey

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	sk := New("hello", 2, 10)
	fmt.Println(sk)
}

func TestCheck(t *testing.T) {
	sk := New("hello", 2, 10)
	old := sk.GetCurrent()
	new, err := sk.GetNext()

	if err != nil {
		t.Fatalf("Out of keys")
	}
	if !Check(new, old) {
		t.Error("Keys are not equal")
	}
}

func TestIterations(t *testing.T) {
	sk := New("hello", 2, 1000)

	old := sk.GetServerInit()
	new := sk.GetCurrent()
	err := error(nil)
	if !Check(new, old) {
		t.Fatalf("Keys are not equal: %d", 0)
	}
	old = new

	for i := 1; i < 1000; i++ {
		new, err = sk.GetNext()
		if err != nil {
			t.Fatalf("Out of keys: %d", i)
		}

		if !Check(new, old) {
			t.Fatalf("Keys are not equal: %d", i)
		}
		old = new
	}
}

func TestOutOfRange(t *testing.T) {
	sk := New("hello", 2, 1000)

	old := sk.GetServerInit()
	new := sk.GetCurrent()
	err := error(nil)
	if !Check(new, old) {
		t.Fatalf("Keys are not equal: %d", 0)
	}
	old = new

	for i := 1; i < 1000; i++ {
		new, err = sk.GetNext()
		if err != nil {
			t.Fatalf("Out of keys: %d", i)
		}

		if !Check(new, old) {
			t.Fatalf("Keys are not equal: %d", i)
		}
		old = new
	}

	fmt.Println(sk.counter)
	new, err = sk.GetNext()

	if err == nil {
		t.Fatal("Out of range happened but 'err' is null")
	}

	if new != 0 {
		t.Fatal("Out of range value is not null")
	}
}
