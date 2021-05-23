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

	var result = hasher.ProcessWithWildcard([]string{"e", "f"}, "", "", 1, "87c568e037a5fa50b1bc911e8ee19a77c4dd3c22bce9932f86fdd8a216afe1681c89737fada6859e91047eece711ec16da62d6ccb9fd0de2c51f132347350d8c")
	assertResultSha512(t, result, "e")
}

func TestHasherSha512_ProcessWithWildcard_WithSaltBefore(t *testing.T) {
	var hasher = NewHasherSha512()

	var result = hasher.ProcessWithWildcard([]string{"d", "e"}, "t", "", 1, "52b526411070a0a92075ea7c2575f759f480f2f4788d56300091696fc7eabb71a74a5fbca04b1934e215ca00bb6b977f6069a34588caa81f622616caacbc83bf")
	assertResultSha512(t, result, "e")
}

func TestHasherSha512_ProcessWithWildcard_WithSaltAfter(t *testing.T) {
	var hasher = NewHasherSha512()

	var result = hasher.ProcessWithWildcard([]string{"d", "e", "f"}, "", "t", 1, "3190862beb8c1d3628e48946f5d154798707a97e829be5c51db208f6baf46d98561a15c2811822eeedc028ed890cbd1d070ee794ceabea8b9ae848c21861b203")
	assertResultSha512(t, result, "f")
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

func assertResultSha512(t *testing.T, result string, expectedWord string) {
	if result != expectedWord {
		t.Errorf("Should have found '%s' but was '%s'", expectedWord, result)
	}
}

func sha512ToString(data []byte) string {
	return hex.EncodeToString(data)
}
