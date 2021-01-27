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

func GenerateMask(difficulty int) []byte {
	full, half := difficulty/2, difficulty%2
	var mask []byte

	for i := 0; i < full; i++ {
		mask = append(mask, 0)
	}

	if half > 0 {
		mask = append(mask, 0xf)
	}

	return mask
}

func checkHashCondition(mask []byte, hash []byte) bool {
	for i := range mask {
		if hash[i] > mask[i] {
			return false
		}
	}

	return true
}

func GenerateHashWithDifficulty(mask []byte, data ...interface{}) ([]byte, int32) {
	len := len(data)

	data = append(data, nil)
	var i int32

	for {
		data[len] = i
		hash := GenerateHash(data...)

		if checkHashCondition(mask, hash) {
			return hash, i
		}
		i++
	}

}
