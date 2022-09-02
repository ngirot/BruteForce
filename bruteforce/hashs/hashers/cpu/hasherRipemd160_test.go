package cpu

import (
	"encoding/hex"
	"github.com/ngirot/BruteForce/bruteforce/hashs/hashers"
	"testing"
)

func TestHasherRipemd160_Hash_WithSimpleWord(t *testing.T) {
	var hasher = NewHasherRipemd160()
	testHashRipemd160(t, hasher, []string{"test"}, []string{"5e52fee47e6b070565f74372468cdc699de89107"})
}

func TestHasherRipemd160_Hash_WithMultipleWord(t *testing.T) {
	var hasher = NewHasherRipemd160()
	testHashRipemd160(t, hasher,
		[]string{"test1", "test2"},
		[]string{"9295fac879006ff44812e43b83b515a06c2950aa", "80b85ebf641abccdd26e327c5782353137a0a0af"})

}

func TestHasherRipemd160_Hash_WithUnicodeWord(t *testing.T) {
	var hasher = NewHasherRipemd160()
	testHashRipemd160(t, hasher, []string{"ありがとう &!ç"}, []string{"56e0854b5d8453b6cd4036b64e005e16354440e1"})
}

func TestHasherRipemd160_ProcessWithWildcard_WithSimpleWord(t *testing.T) {
	var hasher = NewHasherRipemd160()

	var result = hasher.ProcessWithWildcard([]string{"e", "f"}, "", "", 1, "0d42741db982eb2a3f615f46e41114bb64a1a476")
	assertResultRipemd160(t, result, "e")
}

func TestHasherRipemd160_ProcessWithWildcard_WithSaltBefore(t *testing.T) {
	var hasher = NewHasherRipemd160()

	var result = hasher.ProcessWithWildcard([]string{"d", "e"}, "t", "", 1, "02f40ba8bf5582afadb8ad6482efbde01cb94894")
	assertResultRipemd160(t, result, "e")
}

func TestHasherRipemd160_ProcessWithWildcard_WithSaltAfter(t *testing.T) {
	var hasher = NewHasherRipemd160()

	var result = hasher.ProcessWithWildcard([]string{"d", "e", "f"}, "", "t", 1, "c4092d57c59203eff022ec36275bf2b750995498")
	assertResultRipemd160(t, result, "f")
}

func TestHasherRipemd160_Hash_ConsistencyWithSameHash(t *testing.T) {
	var hasher = NewHasherRipemd160()
	var testString = "test"

	var firstResult = ripemd160ToString(hasher.Hash([]string{testString})[0])

	for i := 0; i < 10; i++ {
		var anotherResult = ripemd160ToString(hasher.Hash([]string{testString})[0])
		if anotherResult != firstResult {
			t.Errorf("Hasher is not consistent : the first value was '%s', but it another all returned '%s'", firstResult, anotherResult)
		}
	}
}

func TestHasherRipemd160_IsValid_WithAValidHash(t *testing.T) {
	var hasher = NewHasherRipemd160()
	var hash = "4ccd2e9a26c540facea42ec90f6988a75a64a325"

	if !hasher.IsValid(hash) {
		t.Errorf("The hash '%s' should be valid", hash)
	}
}

func TestHasherRipemd160_IsValid_WithAValueWithIncorrectSize(t *testing.T) {
	var hasher = NewHasherRipemd160()
	var hash = "098f6bcd4621d373cade4e83"

	if hasher.IsValid(hash) {
		t.Errorf("The hash '%s' not should be valid", hash)
	}
}

func TestHasherRipemd160_IsValid_WithAValueWithNotvalidBase64Char(t *testing.T) {
	var hasher = NewHasherRipemd160()
	var hash = "4ccd2e9a26c540fa$ea42ec90f6988a75a64a325"

	if hasher.IsValid(hash) {
		t.Errorf("The hash '%s' not should be valid", hash)
	}
}

func testHashRipemd160(t *testing.T, hasher hashers.Hasher, values []string, expectedHashs []string) {
	var actuals = hasher.Hash(values)

	for i, _ := range values {
		var actual = ripemd160ToString(actuals[i])
		if actual != expectedHashs[i] {
			t.Errorf("Hash value [position %d] for string '%s' should be '%s' but was '%s'", i, values[i], expectedHashs[i], actual)
		}
	}
}

func assertResultRipemd160(t *testing.T, result string, expectedWord string) {
	if result != expectedWord {
		t.Errorf("Should have found '%s' but was '%s'", expectedWord, result)
	}
}

func ripemd160ToString(data []byte) string {
	return hex.EncodeToString(data)
}
