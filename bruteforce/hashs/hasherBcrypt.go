package hashs

import (
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

type hasherBcrypt struct {
}

func NewHasherBcrypt() Hasher {
	return &hasherBcrypt{}
}

func (h *hasherBcrypt) Example() []byte {
	var result, _ = bcrypt.GenerateFromPassword([]byte("1234567890"), 10)
	return result
}

func (h *hasherBcrypt) DecodeInput(data string) []byte {
	return []byte(data)
}

func (h *hasherBcrypt) Hash(data string) []byte {
	return []byte(data)
}

func (h *hasherBcrypt) IsValid(data string) bool {
	var pattern = "^\\$2[ayb]\\$.{56}$"
	match, _ := regexp.MatchString(pattern, data)
	return match
}

func (h *hasherBcrypt) Compare(transformedData []byte, referenceData []byte) bool {
	return bcrypt.CompareHashAndPassword(referenceData, transformedData) == nil
}
