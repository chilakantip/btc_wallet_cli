package utils

import (
	"crypto/sha256"

	"golang.org/x/crypto/ripemd160"
)

func ShaHash(b []byte, out []byte) {
	s := sha256.New()
	s.Write(b[:])
	tmp := s.Sum(nil)
	s.Reset()
	s.Write(tmp)
	copy(out[:], s.Sum(nil))
}

func RimpHash(in []byte, out []byte) {
	sha := sha256.New()
	sha.Write(in)
	rim := ripemd160.New()
	rim.Write(sha.Sum(nil)[:])
	copy(out, rim.Sum(nil))
}

func Sha2Sum(b []byte) (out []byte) {
	out = make([]byte, 32)
	ShaHash(b, out[:])
	return
}
