package hashs

import (
	"crypto/sha256"
	"hash"
)

type Hasher interface {
	Hash(data string) []byte
}

type hasher struct {
	cache hash.Hash
}

func NewHasher() Hasher {
	return &hasher{sha256.New()}
}

func (h *hasher) Hash(data string) []byte {
	return h.binaryHash(h.convert(data))
}

func (h *hasher) binaryHash(data []byte) []byte {
	h.cache.Reset()
	h.cache.Write([]byte(data))
	return h.cache.Sum(nil)
}

func (h *hasher) convert(s string) []byte {
	return []byte(s)
}
