package hashs

import (
	"crypto/sha256"
	"hash"
)

type hasherSha256 struct {
	cache hash.Hash
}

func NewHasherSha256() Hasher {
	return &hasherSha256{sha256.New()}
}

func (h *hasherSha256) Hash(data string) []byte {
	return h.binaryHash(h.convert(data))
}

func (h *hasherSha256) IsValid(data string) bool {
	return genericBase64Validator(h, data)
}

func (h *hasherSha256) binaryHash(data []byte) []byte {
	h.cache.Reset()
	h.cache.Write([]byte(data))
	return h.cache.Sum(nil)
}

func (h *hasherSha256) convert(s string) []byte {
	return []byte(s)
}
