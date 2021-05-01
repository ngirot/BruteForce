package hashers

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"hash"
)

type hasherSha256 struct {
	cache hash.Hash
}

func (h *hasherSha256) ProcessWithGpu(charSet []string, saltBefore string, saltAfter string, numberOfWildCards int, expectedDigest string) string {
	panic("implement me")
}

func NewHasherSha256() Hasher {
	return &hasherSha256{sha256.New()}
}

func (h *hasherSha256) Example() string {
	return hex.EncodeToString(h.Hash([]string{"1234567890"})[0])
}

func (h *hasherSha256) DecodeInput(data string) []byte {
	var result, _ = hex.DecodeString(data)
	return result
}

func (h *hasherSha256) Hash(datas []string) [][]byte {
	var result = make([][]byte, len(datas))
	for i, value := range datas {
		result[i] = h.binaryHash(h.convert(value))
	}
	return result
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
