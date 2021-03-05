package hashs

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"hash"
)

type hasherSha256 struct {
	cache hash.Hash
}

func NewHasherSha256() Hasher {
	return &hasherSha256{sha256.New()}
}

func (h *hasherSha256) Example() []byte {
	return h.Hash("1234567890")
}

func (h *hasherSha256) DecodeInput(data string) []byte {
	var result, _ = hex.DecodeString(data)
	return result
}

func (h *hasherSha256) Hash(data string) []byte {
	return h.binaryHash(h.convert(data))
}

func (h *hasherSha256) IsValid(data string) bool {
	return genericBase64Validator(h, data)
}

func (h *hasherSha256) Compare(transformedData []byte, referenceData []byte) bool {
	return bytes.Equal(transformedData, referenceData)
}

func (h *hasherSha256) binaryHash(data []byte) []byte {
	h.cache.Reset()
	h.cache.Write([]byte(data))
	return h.cache.Sum(nil)
}

func (h *hasherSha256) convert(s string) []byte {
	return []byte(s)
}
