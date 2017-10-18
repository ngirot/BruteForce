package main

import (
	"encoding/hex"
	"crypto/sha256"
)

func Hash(data string) string {
	return format(binaryHash(convert(data)))
}

func binaryHash(data []byte) []byte {
	var h = sha256.New()
	h.Write([]byte(data))
	return h.Sum(nil)
}

func convert(s string) []byte {
	return []byte(s)
}

func format(data []byte) string {
	return hex.EncodeToString(data)
}
