package utils

import "encoding/binary"

func ToBytes32(x uint32) []byte {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, x)
	return bs
}

func ToBytes64(x uint64) []byte {
	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, x)
	return bs
}

func FromBytes32(bs []byte) uint32 {
	return binary.BigEndian.Uint32(bs)
}

func FromBytes64(bs []byte) uint64 {
	return binary.BigEndian.Uint64(bs)

}
