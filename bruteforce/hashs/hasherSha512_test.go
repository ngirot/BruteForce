package hashs

import (
	"encoding/hex"
	"testing"
)

func TestHasherSha512_Hash_WithSimpleWord(t *testing.T) {
	var hasher = NewHasherSha512()
	testHashSha512(t, hasher, "test", "ee26b0dd4af7e749aa1a8ee3c10ae9923f618980772e473f8819a5d4940e0db27ac185f8a0e1d5f84f88bc887fd67b143732c304cc5fa9ad8e6f57f50028a8ff")
}

func TestHasherSha512_Hash_WithUnicodeWord(t *testing.T) {
	var hasher = NewHasherSha512()
	testHashSha512(t, hasher, "ありがとう &!ç", "5b35cc2b8d75331db1ad1550a1e7b4f3538ca9bd9f1b4d9a6dff5b8c938b331a4844aa874639da3acd38c79d071d2a146916e215f12d9d562a2637a9d0927943")
}

func TestHasherSha512_Hash_ConsistencyWithSameHash(t *testing.T) {
	var hasher = NewHasherSha512()
	var testString = "test"

	var firstResult = sha512ToString(hasher.Hash(testString))

	for i:=0 ; i<10 ; i++ {
		var anotherResult = sha512ToString(hasher.Hash(testString))
		if anotherResult != firstResult {
			t.Errorf("Hasher is not consistent : the first value was '%s', but it another all returned '%s'", firstResult, anotherResult)
		}
	}
}

func TestHasherSha512_IsValid_WithAValidHash(t *testing.T) {
	var hasher = NewHasherSha512()
	var hash = "ee26b0dd4af7e749aa1a8ee3c10ae9923f618980772e473f8819a5d4940e0db27ac185f8a0e1d5f84f88bc887fd67b143732c304cc5fa9ad8e6f57f50028a8ff"

	if !hasher.IsValid(hash) {
		t.Errorf("The hash '%s' should be valid", hash)
	}
}

func TestHasherSha512_IsValid_WithAValueWithIncorrectSize(t *testing.T) {
	var hasher = NewHasherSha512()
	var hash = "ee26b0dd4af7e749aa1a8ee3c10ae9923f618980772e473f8819a5d4940e0db27ac185f8a0e1d5f84f88bc887fd67b143732c304cc5fa9ad8e6f57f50028"

	if hasher.IsValid(hash) {
		t.Errorf("The hash '%s' should not be valid", hash)
	}
}

func TestHasherSha512_IsValid_WithAValueWithNotvalidBase64Char(t *testing.T) {
	var hasher = NewHasherSha512()
	var hash = "ee2%b0dd4af7e749aa1a8ee3c10ae9923f618980772e473f8819a5d4940e0db27ac185f8a0e1d5f84f88bc887fd67b143732c304cc5fa9ad8e6f57f50028a8ff"

	if hasher.IsValid(hash) {
		t.Errorf("The hash '%s' should not be valid", hash)
	}
}

func testHashSha512(t *testing.T, hasher Hasher, value string, expectedHash string) {
	var actual = sha512ToString(hasher.Hash(value))
	if actual != expectedHash {
		t.Errorf("Hash value for string '%s' should be '%s' but was '%s'", value, expectedHash, actual)
	}
}

func sha512ToString(data [] byte) string {
	return hex.EncodeToString(data)
}