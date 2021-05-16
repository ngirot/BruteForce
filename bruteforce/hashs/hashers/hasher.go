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
	ProcessWithWildcard(charSet []string, saltBefore string, saltAfter string, numberOfWildCards int, expectedDigest string) string
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

func expand(charset []string, expansionFactor int) []string {
	if expansionFactor == 0 {
		return []string{}
	}

	var result []string

	for i := 0; i < len(charset); i++ {
		var array = expand(charset, expansionFactor-1)
		if len(array) > 0 {
			for j := 0; j < len(array); j++ {
				result = append(result, charset[i]+array[j])
			}
		} else {
			result = append(result, charset[i])
		}
	}

	return result
}

func genericProcessWithWildCard(hasher Hasher, charSet []string, saltBefore string, saltAfter string, numberOfWildCards int, expectedDigest string) string {
	var expanded = expand(charSet, numberOfWildCards)

	for i := 0; i < len(expanded); i++ {
		var word = saltBefore + expanded[i] + saltAfter
		var currentHash = hasher.Hash([]string{word})[0]

		if hasher.Compare(currentHash, hasher.DecodeInput(expectedDigest)) {
			return expanded[i]
		}
	}

	return ""
}
