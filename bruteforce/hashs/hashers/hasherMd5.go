package hashers

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
	return hex.EncodeToString(h.Hash([]string{"1234567890"})[0])
}

func (h *hasherMd5) DecodeInput(data string) []byte {
	var result, _ = hex.DecodeString(data)
	return result
}

func (h *hasherMd5) Hash(datas []string) [][]byte {
	var result = make([][]byte, len(datas))
	for i, value := range datas {
		result[i] = h.binaryHash(h.convert(value))
	}
	return result
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
