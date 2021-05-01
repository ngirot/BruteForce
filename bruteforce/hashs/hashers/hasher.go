package hashers

import (
	"encoding/base64"
	"encoding/hex"
)

type Hasher interface {
	Example() string
	DecodeInput(data string) []byte
	Hash(datas []string) [][]byte
	Compare(transformedData []byte, referenceData []byte) bool
	ProcessWithGpu(charSet []string, saltBefore string, saltAfter string, numberOfWildCards int, expectedDigest string) string
	IsValid(data string) bool
}

func genericBase64Validator(hasher Hasher, data string) bool {
	if !isBase64(data) {
		return false
	}

	var hexData = hasher.Hash([]string{"valid"})
	return len(hexToString(hexData[0])) == len(data)
}

func hexToString(data []byte) string {
	return hex.EncodeToString(data)
}

func isBase64(data string) bool {
	var _, err = base64.StdEncoding.DecodeString(data)
	return err == nil
}
