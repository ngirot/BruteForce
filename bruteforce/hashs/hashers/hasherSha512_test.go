package hashers

import (
	"encoding/hex"
	"testing"
)

func TestHasherSha512_Hash_WithSimpleWord(t *testing.T) {
	var hasher = NewHasherSha512()
	testHashSha512(t, hasher, []string{"test"}, []string{"ee26b0dd4af7e749aa1a8ee3c10ae9923f618980772e473f8819a5d4940e0db27ac185f8a0e1d5f84f88bc887fd67b143732c304cc5fa9ad8e6f57f50028a8ff"})
}

func TestHasherSha512_Hash_WithMultipleWord(t *testing.T) {
	var hasher = NewHasherSha512()
	testHashSha512(t, hasher,
		[]string{"test1", "test2"},
		[]string{"b16ed7d24b3ecbd4164dcdad374e08c0ab7518aa07f9d3683f34c2b3c67a15830268cb4a56c1ff6f54c8e54a795f5b87c08668b51f82d0093f7baee7d2981181", "6d201beeefb589b08ef0672dac82353d0cbd9ad99e1642c83a1601f3d647bcca003257b5e8f31bdc1d73fbec84fb085c79d6e2677b7ff927e823a54e789140d9"})
}

func TestHasherSha512_Hash_WithUnicodeWord(t *testing.T) {
	var hasher = NewHasherSha512()
	testHashSha512(t, hasher, []string{"ありがとう &!ç"}, []string{"5b35cc2b8d75331db1ad1550a1e7b4f3538ca9bd9f1b4d9a6dff5b8c938b331a4844aa874639da3acd38c79d071d2a146916e215f12d9d562a2637a9d0927943"})
}

func TestHasherSha512_ProcessWithWildcard_WithSimpleWord(t *testing.T) {
	var hasher = NewHasherSha512()

	hasher.ProcessWithWildcard([]string{"e", "f"}, "", "", 1, "ee26b0dd4af7e749aa1a8ee3c10ae9923f618980772e473f8819a5d4940e0db27ac185f8a0e1d5f84f88bc887fd67b143732c304cc5fa9ad8e6f57f50028a8ff.Gt4wp0dJk5qWRaumcfqazMMCAxxerGi")
}


func TestHasherSha512_ProcessWithWildcard_WithSaltBefore(t *testing.T) {
	var hasher = NewHasherSha512()

	hasher.ProcessWithWildcard([]string{"d", "e"}, "t", "", 1, "ee26b0dd4af7e749aa1a8ee3c10ae9923f618980772e473f8819a5d4940e0db27ac185f8a0e1d5f84f88bc887fd67b143732c304cc5fa9ad8e6f57f50028a8ff")
}

func TestHasherSha512_ProcessWithWildcard_WithSaltAfter(t *testing.T) {
	var hasher = NewHasherSha512()

	hasher.ProcessWithWildcard([]string{"d", "e", "f"}, "", "t", 1, "ee26b0dd4af7e749aa1a8ee3c10ae9923f618980772e473f8819a5d4940e0db27ac185f8a0e1d5f84f88bc887fd67b143732c304cc5fa9ad8e6f57f50028a8ff")
}

func TestHasherSha512_Hash_ConsistencyWithSameHash(t *testing.T) {
	var hasher = NewHasherSha512()
	var testString = "test"

	var firstResult = sha512ToString(hasher.Hash([]string{testString})[0])

	for i := 0; i < 10; i++ {
		var anotherResult = sha512ToString(hasher.Hash([]string{testString})[0])
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

func testHashSha512(t *testing.T, hasher Hasher, values []string, expectedHashs []string) {
	var actuals = hasher.Hash(values)

	for i, _ := range values {
		var actual = sha512ToString(actuals[i])
		if actual != expectedHashs[i] {
			t.Errorf("Hash value [position %d] for string '%s' should be '%s' but was '%s'", i, values[i], expectedHashs[i], actual)
		}
	}
}

func sha512ToString(data []byte) string {
	return hex.EncodeToString(data)
}
