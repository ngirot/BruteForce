package hashs

import (
	"testing"
	"encoding/hex"
)


func TestHasherSha256_Hash_WithSimpleWord(t *testing.T) {
	var hasher = NewHasherSha256()
	testHashSha256(t, hasher, "test", "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08")
}

func TestHasherSha256_Hash_WithUnicodeWord(t *testing.T) {
	var hasher = NewHasherSha256()
	testHashSha256(t, hasher, "ありがとう &!ç", "f89eddccb44ae418616060aefe3ca6604d49bc3d0e37e75167333d498532d7aa")
}

func TestHasherSha256_Hash_ConsistencyWithSameHash(t *testing.T) {
	var hasher = NewHasherSha256()
	var testString = "test"

	var firstResult = sha256ToString(hasher.Hash(testString))

	for i:=0 ; i<10 ; i++ {
		var anotherResult = sha256ToString(hasher.Hash(testString))
		if anotherResult != firstResult {
			t.Errorf("Hasher is not consistent : the first value was '%s', but it another all returned '%s'", firstResult, anotherResult)
		}
	}
}

func TestHasherSha256_IsValid_WithAValidHash(t *testing.T) {
	var hasher = NewHasherSha256()
	var hash = "f89eddccb44ae418616060aefe3ca6604d49bc3d0e37e75167333d498532d7aa"

	if !hasher.IsValid(hash) {
		t.Errorf("The hash '%s' should be valid", hash)
	}
}

func TestHasherSha256_IsValid_WithAValueWithIncorrectSize(t *testing.T) {
	var hasher = NewHasherSha256()
	var hash = "f89eddccb44ae418616060aefe3ca6604d49bc3d0e37e7516733"

	if hasher.IsValid(hash) {
		t.Errorf("The hash '%s' should not be valid", hash)
	}
}

func TestHasherSha256_IsValid_WithAValueWithNotvalidBase64Char(t *testing.T) {
	var hasher = NewHasherSha256()
	var hash = "f89eddccb44ae41:616060aefe3ca6604d49bc3d0e37e75167333d498532d7aa"

	if hasher.IsValid(hash) {
		t.Errorf("The hash '%s' should not be valid", hash)
	}
}

func testHashSha256(t *testing.T, hasher Hasher, value string, expectedHash string) {
	var actual = sha256ToString(hasher.Hash(value))
	if actual != expectedHash {
		t.Errorf("Hash value for string '%s' should be '%s' but was '%s'", value, expectedHash, actual)
	}
}

func sha256ToString(data [] byte) string {
	return hex.EncodeToString(data)
}