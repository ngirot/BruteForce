package hashs

import (
	"encoding/base64"
	"encoding/hex"
)

type Hasher interface {
	Hash(data string) []byte
	IsValid(data string) bool
}

func genericBase64Validator(hasher Hasher, data string) bool {
	if !isBase64(data) {
		return false
	}

	var hexData = hasher.Hash("valid")
	return len(hexToString(hexData)) == len(data)
}

func hexToString(data []byte) string {
	return hex.EncodeToString(data)
}

func isBase64(data string) bool {
	var _,err = base64.StdEncoding.DecodeString(data)
	return err == nil
}