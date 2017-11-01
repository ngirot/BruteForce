package hashs

import (
	"crypto/sha256"
	"hash"
)

type Hasher struct {
	cache hash.Hash
}

func NewHasher() Hasher {
	return Hasher{sha256.New()}
}

func (h *Hasher) Hash(data string) []byte {
	return h.binaryHash(h.convert(data))
}

func (h *Hasher) binaryHash(data []byte) []byte {
	h.cache.Reset()
	h.cache.Write([]byte(data))
	return h.cache.Sum(nil)
}

func (h *Hasher) convert(s string) []byte {
	return []byte(s)
}
