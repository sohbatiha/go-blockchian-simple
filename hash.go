package main

import (
	"crypto/sha256"
	"fmt"
)

func GenerateHash(data ...interface{}) []byte {
	hash := sha256.New()
	fmt.Fprint(hash, data...)

	return hash.Sum(nil)
}
