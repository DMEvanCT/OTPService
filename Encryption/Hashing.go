package Encryption

import (
	"crypto/md5"
	"fmt"
	"io"
)

func Md5Hash(secret_to_hash, secret_salt string) []byte {
	h := md5.New()
	io.WriteString(h, secret_to_hash)
	io.WriteString(h, secret_salt)
	fmt.Printf("%x", h.Sum(nil))
	return h.Sum(nil)
}