package hashs

import (
	"hash"
	"crypto/sha512"
)

type hasherSha512 struct {
	cache hash.Hash
}

func NewHasherSha512() Hasher {
	return &hasherSha512{sha512.New()}
}

func (h *hasherSha512) Hash(data string) []byte {
	return h.binaryHash(h.convert(data))
}

func (h *hasherSha512) IsValid(data string) bool {
	return genericBase64Validator(h, data)
}

func (h *hasherSha512) binaryHash(data []byte) []byte {
	h.cache.Reset()
	h.cache.Write([]byte(data))
	return h.cache.Sum(nil)
}

func (h *hasherSha512) convert(s string) []byte {
	return []byte(s)
}
