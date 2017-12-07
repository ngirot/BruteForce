package hashs

import (
	"hash"
	"crypto/sha1"
)

type hasherSha1 struct {
	cache hash.Hash
}

func NewHasherSha1() Hasher {
	return &hasherSha256{sha1.New()}
}

func (h *hasherSha1) Hash(data string) []byte {
	return h.binaryHash(h.convert(data))
}

func (h *hasherSha1) binaryHash(data []byte) []byte {
	h.cache.Reset()
	h.cache.Write([]byte(data))
	return h.cache.Sum(nil)
}

func (h *hasherSha1) convert(s string) []byte {
	return []byte(s)
}
