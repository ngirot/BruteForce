package hashers

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

func (h *hasherSha1) Example() string {
	return hex.EncodeToString(h.Hash([]string{"1234567890"})[0])
}

func (h *hasherSha1) DecodeInput(data string) []byte {
	var result, _ = hex.DecodeString(data)
	return result
}

func (h *hasherSha1) Hash(datas []string) [][]byte {
	var result = make([][]byte, len(datas))
	for i, value := range datas {
		result[i] = h.binaryHash(h.convert(value))
	}
	return result
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
