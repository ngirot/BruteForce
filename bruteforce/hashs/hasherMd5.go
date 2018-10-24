package hashs

import (
	"crypto/md5"
	"hash"
)

type hasherMd5 struct {
	cache hash.Hash
}

func NewHasherMd5() Hasher {
	return &hasherMd5{md5.New()}
}

func (h *hasherMd5) Hash(data string) []byte {
	return h.binaryHash(h.convert(data))
}

func (h *hasherMd5) IsValid(data string) bool {
	return genericBase64Validator(h, data)
}

func (h *hasherMd5) binaryHash(data []byte) []byte {
	h.cache.Reset()
	h.cache.Write([]byte(data))
	return h.cache.Sum(nil)
}

func (h *hasherMd5) convert(s string) []byte {
	return []byte(s)
}
