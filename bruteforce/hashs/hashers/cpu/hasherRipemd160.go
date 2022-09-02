package cpu

import (
	"bytes"
	"encoding/hex"
	"github.com/ngirot/BruteForce/bruteforce/hashs/hashers"
	"golang.org/x/crypto/ripemd160"
	"hash"
)

type hasherRipemd160 struct {
	cache hash.Hash
}

func NewHasherRipemd160() hashers.Hasher {
	return &hasherRipemd160{ripemd160.New()}
}

func (h *hasherRipemd160) ProcessWithWildcard(charSet []string, saltBefore string, saltAfter string, numberOfWildCards int, expectedDigest string) string {
	return genericProcessWithWildCard(h, charSet, saltBefore, saltAfter, numberOfWildCards, expectedDigest)
}

func (h *hasherRipemd160) Example() string {
	return hex.EncodeToString(h.Hash([]string{"1234567890"})[0])
}

func (h *hasherRipemd160) DecodeInput(data string) []byte {
	var result, _ = hex.DecodeString(data)
	return result
}

func (h *hasherRipemd160) Hash(datas []string) [][]byte {
	var result = make([][]byte, len(datas))
	for i, value := range datas {
		result[i] = h.binaryHash(h.convert(value))
	}
	return result
}

func (h *hasherRipemd160) IsValid(data string) bool {
	return hashers.GenericBase64Validator(h, data)
}

func (h *hasherRipemd160) Compare(transformedData []byte, referenceData []byte) bool {
	return bytes.Equal(transformedData, referenceData)
}

func (h *hasherRipemd160) binaryHash(data []byte) []byte {
	h.cache.Reset()
	h.cache.Write([]byte(data))
	return h.cache.Sum(nil)
}

func (h *hasherRipemd160) convert(s string) []byte {
	return []byte(s)
}
