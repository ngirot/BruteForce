package hashs

import (
	"encoding/hex"
	"testing"
)

func TestHasherMd5_Hash_WithSimpleWord(t *testing.T) {
	var hasher = NewHasherMd5()
	testHashMd5(t, hasher, "test", "098f6bcd4621d373cade4e832627b4f6")
}

func TestHasherMd5_Hash_WithUnicodeWord(t *testing.T) {
	var hasher = NewHasherMd5()
	testHashMd5(t, hasher, "ありがとう &!ç", "9ebcc60effbdef7c4d101a7ced1c6b01")
}

func TestHasherMd5_Hash_ConsistencyWithSameHash(t *testing.T) {
	var hasher = NewHasherMd5()
	var testString = "test"

	var firstResult = md5ToString(hasher.Hash(testString))

	for i:=0 ; i<10 ; i++ {
		var anotherResult = md5ToString(hasher.Hash(testString))
		if anotherResult != firstResult {
			t.Errorf("Hasher is not consistent : the first value was '%s', but it another all returned '%s'", firstResult, anotherResult)
		}
	}
}

func TestHasherMd5_IsValid_WithAValidHash(t *testing.T) {
	var hasher = NewHasherMd5()
	var hash = "098f6bcd4621d373cade4e832627b4f6"

	if !hasher.IsValid(hash) {
		t.Errorf("The hash '%s' should be valid", hash)
	}
}

func TestHasherMd5_IsValid_WithAValueWithIncorrectSize(t *testing.T) {
	var hasher = NewHasherMd5()
	var hash = "098f6bcd4621d373cade4e83"

	if hasher.IsValid(hash) {
		t.Errorf("The hash '%s' not should be valid", hash)
	}
}

func TestHasherMd5_IsValid_WithAValueWithNotvalidBase64Char(t *testing.T) {
	var hasher = NewHasherMd5()
	var hash = "098f6bcd4621$373cade4e832627b4f6"

	if hasher.IsValid(hash) {
		t.Errorf("The hash '%s' not should be valid", hash)
	}
}

func testHashMd5(t *testing.T, hasher Hasher, value string, expectedHash string) {
	var actual = md5ToString(hasher.Hash(value))
	if actual != expectedHash {
		t.Errorf("Hash value for string '%s' should be '%s' but was '%s'", value, expectedHash, actual)
	}
}

func md5ToString(data [] byte) string {
	return hex.EncodeToString(data)
}