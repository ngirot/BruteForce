package cpu

import (
	"encoding/hex"
	"github.com/ngirot/BruteForce/bruteforce/hashs/hashers"
	"testing"
)

func TestHasherMd5_Hash_WithSimpleWord(t *testing.T) {
	var hasher = NewHasherMd5()
	testHashMd5(t, hasher, []string{"test"}, []string{"098f6bcd4621d373cade4e832627b4f6"})
}

func TestHasherMd5_Hash_WithMultipleWord(t *testing.T) {
	var hasher = NewHasherMd5()
	testHashMd5(t, hasher,
		[]string{"test1", "test2"},
		[]string{"5a105e8b9d40e1329780d62ea2265d8a", "ad0234829205b9033196ba818f7a872b"})

}

func TestHasherMd5_Hash_WithUnicodeWord(t *testing.T) {
	var hasher = NewHasherMd5()
	testHashMd5(t, hasher, []string{"ありがとう &!ç"}, []string{"9ebcc60effbdef7c4d101a7ced1c6b01"})
}

func TestHasherMd5_ProcessWithWildcard_WithSimpleWord(t *testing.T) {
	var hasher = NewHasherMd5()

	var result = hasher.ProcessWithWildcard([]string{"e", "f"}, "", "", 1, "e1671797c52e15f763380b45e841ec32")
	assertResultMd5(t, result, "e")
}


func TestHasherMd5_ProcessWithWildcard_WithSaltBefore(t *testing.T) {
	var hasher = NewHasherMd5()

	var result = hasher.ProcessWithWildcard([]string{"d", "e"}, "t", "", 1, "569ef72642be0fadd711d6a468d68ee1")
	assertResultMd5(t, result, "e")
}

func TestHasherMd5_ProcessWithWildcard_WithSaltAfter(t *testing.T) {
	var hasher = NewHasherMd5()

	var result = hasher.ProcessWithWildcard([]string{"d", "e", "f"}, "", "t", 1, "49af3b640275c9b552a5f3f3d96a6062")
	assertResultMd5(t, result, "f")
}

func TestHasherMd5_Hash_ConsistencyWithSameHash(t *testing.T) {
	var hasher = NewHasherMd5()
	var testString = "test"

	var firstResult = md5ToString(hasher.Hash([]string{testString})[0])

	for i := 0; i < 10; i++ {
		var anotherResult = md5ToString(hasher.Hash([]string{testString})[0])
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

func testHashMd5(t *testing.T, hasher hashers.Hasher, values []string, expectedHashs []string) {
	var actuals = hasher.Hash(values)

	for i, _ := range values {
		var actual = md5ToString(actuals[i])
		if actual != expectedHashs[i] {
			t.Errorf("Hash value [position %d] for string '%s' should be '%s' but was '%s'", i, values[i], expectedHashs[i], actual)
		}
	}
}

func assertResultMd5(t *testing.T, result string, expectedWord string) {
	if result != expectedWord {
		t.Errorf("Should have found '%s' but was '%s'", expectedWord, result)
	}
}

func md5ToString(data []byte) string {
	return hex.EncodeToString(data)
}
