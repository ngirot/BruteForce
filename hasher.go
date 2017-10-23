package main

import (
	"encoding/hex"
	"crypto/sha256"
)

var cachedHasher = sha256.New()

func Hash(data string) string {
	return format(binaryHash(convert(data)))
}

func binaryHash(data []byte) []byte {
	cachedHasher.Reset()
	cachedHasher.Write([]byte(data))
	return cachedHasher.Sum(nil)
}

func convert(s string) []byte {
	return []byte(s)
}

func format(data []byte) string {
	return hex.EncodeToString(data)
}
