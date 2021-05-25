package cpu

import (
	"github.com/ngirot/BruteForce/bruteforce/hashs/hashers"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

type hasherBcrypt struct {
}

func NewHasherBcrypt() hashers.Hasher {
	return &hasherBcrypt{}
}

func (h *hasherBcrypt) ProcessWithWildcard(charSet []string, saltBefore string, saltAfter string, numberOfWildCards int, expectedDigest string) string {
	return genericProcessWithWildCard(h, charSet, saltBefore, saltAfter, numberOfWildCards, expectedDigest)
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
