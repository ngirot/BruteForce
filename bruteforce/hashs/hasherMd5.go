package hashs

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"hash"
)

type hasherMd5 struct {
	cache hash.Hash
}

func NewHasherMd5() Hasher {
	return &hasherMd5{md5.New()}
}

func (h *hasherMd5) Example() string {
	return hex.EncodeToString(h.Hash("1234567890"))
}

func (h *hasherMd5) DecodeInput(data string) []byte {
	var result, _ = hex.DecodeString(data)
	return result
}

func (h *hasherMd5) Hash(data string) []byte {
	return h.binaryHash(h.convert(data))
}

func (h *hasherMd5) IsValid(data string) bool {
	return genericBase64Validator(h, data)
}

func (h *hasherMd5) Compare(transformedData []byte, referenceData []byte) bool {
	return bytes.Equal(transformedData, referenceData)
}

func (h *hasherMd5) binaryHash(data []byte) []byte {
	h.cache.Reset()
	h.cache.Write([]byte(data))
	return h.cache.Sum(nil)
}

func (h *hasherMd5) convert(s string) []byte {
	return []byte(s)
}
