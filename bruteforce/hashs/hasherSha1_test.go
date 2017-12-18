package hashs

import (
	"testing"
	"encoding/hex"
)

func TestHasherSha1_Hash_WithSimpleWord(t *testing.T) {
	var hasher = NewHasherSha1()
	testHashSha1(t, hasher, "test", "a94a8fe5ccb19ba61c4c0873d391e987982fbbd3")
}

func TestHasherSha1_Hash_WithUnicodeWord(t *testing.T) {
	var hasher = NewHasherSha1()
	testHashSha1(t, hasher, "ありがとう &!ç", "cb789c4b10a21cd6fa614384436ac57b0e18c1cd")
}

func TestHasherSha1_Hash_ConsistencyWithSameHash(t *testing.T) {
	var hasher = NewHasherSha1()
	var testString = "test"

	var firstResult = sha1ToString(hasher.Hash(testString))

	for i:=0 ; i<10 ; i++ {
		var anotherResult = sha1ToString(hasher.Hash(testString))
		if anotherResult != firstResult {
			t.Errorf("Hasher is not consistent : the first value was '%s', but it another all returned '%s'", firstResult, anotherResult)
		}
	}
}

func TestHasherSha1_IsValid_WithAValidHash(t *testing.T) {
	var hasher = NewHasherSha1()
	var hash = "cb789c4b10a21cd6fa614384436ac57b0e18c1cd"

	if !hasher.IsValid(hash) {
		t.Errorf("The hash '%s' should be valid", hash)
	}
}

func TestHasherSha1_IsValid_WithAValueWithIncorrectSize(t *testing.T) {
	var hasher = NewHasherSha1()
	var hash = "cb789c4b10a21cd6fa614384436ac57b0e"

	if hasher.IsValid(hash) {
		t.Errorf("The hash '%s' should not be valid", hash)
	}
}

func TestHasherSha1_IsValid_WithAValueWithNotvalidBase64Char(t *testing.T) {
	var hasher = NewHasherSha1()
	var hash = "cb789c4b10a21µd6fa614384436ac57b0e18c1cd"

	if hasher.IsValid(hash) {
		t.Errorf("The hash '%s' should not be valid", hash)
	}
}

func testHashSha1(t *testing.T, hasher Hasher, value string, expectedHash string) {
	var actual = sha1ToString(hasher.Hash(value))
	if actual != expectedHash {
		t.Errorf("Hash value for string '%s' should be '%s' but was '%s'", value, expectedHash, actual)
	}
}

func sha1ToString(data [] byte) string {
	return hex.EncodeToString(data)
}