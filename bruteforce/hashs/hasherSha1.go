package hashs

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"hash"
)

type hasherSha1 struct {
	cache hash.Hash
}

func NewHasherSha1() Hasher {
	return &hasherSha1{sha1.New()}
}

func (h *hasherSha1) Example() []byte {
	return h.Hash("1234567890")
}

func (h *hasherSha1) DecodeInput(data string) []byte {
	var result, _ = hex.DecodeString(data)
	return result
}

func (h *hasherSha1) Hash(data string) []byte {
	return h.binaryHash(h.convert(data))
}

func (h *hasherSha1) IsValid(data string) bool {
	return genericBase64Validator(h, data)
}

func (h *hasherSha1) Compare(transformedData []byte, referenceData []byte) bool {
	return bytes.Equal(transformedData, referenceData)
}

func (h *hasherSha1) binaryHash(data []byte) []byte {
	h.cache.Reset()
	h.cache.Write([]byte(data))
	return h.cache.Sum(nil)
}

func (h *hasherSha1) convert(s string) []byte {
	return []byte(s)
}
