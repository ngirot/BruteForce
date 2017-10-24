package main

import (
	"encoding/hex"
	"crypto/sha256"
	"hash"
)

type Hasher struct {
	cache hash.Hash
}

func NewHasher() Hasher {
	return Hasher{sha256.New()}
}

func (h *Hasher) Hash(data string) string {
	return h.format(h.binaryHash(h.convert(data)))
}

func (h *Hasher) binaryHash(data []byte) []byte {
	h.cache.Reset()
	h.cache.Write([]byte(data))
	return h.cache.Sum(nil)
}

func (h *Hasher) convert(s string) []byte {
	return []byte(s)
}

func (h *Hasher) format(data []byte) string {
	return hex.EncodeToString(data)
}
