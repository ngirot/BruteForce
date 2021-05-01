package hashers

import (
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

type hasherBcrypt struct {
}

func (h *hasherBcrypt) ProcessWithGpu(charSet []string, saltBefore string, saltAfter string, numberOfWildCards int, expectedDigest string) string {
	panic("implement me")
}

func NewHasherBcrypt() Hasher {
	return &hasherBcrypt{}
}

func (h *hasherBcrypt) Example() string {
	var result, _ = bcrypt.GenerateFromPassword([]byte("1234567890"), 10)
	return string(result)
}

func (h *hasherBcrypt) DecodeInput(data string) []byte {
	return []byte(data)
}

func (h *hasherBcrypt) Hash(datas []string) [][]byte {
	return [][]byte{[]byte(datas[0])}
}

func (h *hasherBcrypt) IsValid(data string) bool {
	var pattern = "^\\$2[ayb]\\$.{56}$"
	match, _ := regexp.MatchString(pattern, data)
	return match
}

func (h *hasherBcrypt) Compare(transformedData []byte, referenceData []byte) bool {
	return bcrypt.CompareHashAndPassword(referenceData, transformedData) == nil
}
