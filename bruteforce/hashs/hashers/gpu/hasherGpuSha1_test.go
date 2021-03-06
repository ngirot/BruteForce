// +build opencl

package gpu

import (
	"encoding/hex"
	"github.com/ngirot/BruteForce/bruteforce/hashs/hashers"
	"testing"
)

var hasherSha1 = NewHasherGpuSha1()

func TestHasherGpuSha1_Hash_WithSimpleWord(t *testing.T) {
	testHashSha1(t, hasherSha1, []string{"test"}, []string{"a94a8fe5ccb19ba61c4c0873d391e987982fbbd3"})
}

func TestHasherGpuSha1_Hash_WithMultipleWord(t *testing.T) {
	testHashSha1(t, hasherSha1,
		[]string{"test1", "test2"},
		[]string{"b444ac06613fc8d63795be9ad0beaf55011936ac", "109f4b3c50d7b0df729d299bc6f8e9ef9066971f"})
}

func TestHasherGpuSha1_Hash_WithUnicodeWord(t *testing.T) {
	testHashSha1(t, hasherSha1, []string{"ありがとう &!ç"}, []string{"cb789c4b10a21cd6fa614384436ac57b0e18c1cd"})
}

func TestHasherGpuSha1_ProcessWithWildcard_WithSimpleWord(t *testing.T) {
	var result = hasherSha1.ProcessWithWildcard([]string{"e", "f"}, "", "", 1, "58e6b3a414a1e090dfc6029add0f3555ccba127f")
	assertResultSha1(t, result, "e")
}

func TestHasherGpuSha1_ProcessWithWildcard_WithSaltBefore(t *testing.T) {
	var result = hasherSha1.ProcessWithWildcard([]string{"d", "e"}, "t", "", 1, "33e9505d12942e8259a3c96fb6f88ed325b95797")
	assertResultSha1(t, result, "e")
}

func TestHasherGpuSha1_ProcessWithWildcard_WithSaltAfter(t *testing.T) {
	var result = hasherSha1.ProcessWithWildcard([]string{"d", "e", "f"}, "", "t", 1, "d352dbdf6170085acaf7ed62197a4de1452a0073")
	assertResultSha1(t, result, "f")
}

func TestHasherGpuSha1_Hash_ConsistencyWithSameHash(t *testing.T) {
	var testString = "test"

	var firstResult = sha1ToString(hasherSha1.Hash([]string{testString})[0])

	for i := 0; i < 10; i++ {
		var anotherResult = sha1ToString(hasherSha1.Hash([]string{testString})[0])
		if anotherResult != firstResult {
			t.Errorf("Hasher is not consistent : the first value was '%s', but it another all returned '%s'", firstResult, anotherResult)
		}
	}
}

func TestHasherGpuSha1_IsValid_WithAValidHash(t *testing.T) {
	var hash = "cb789c4b10a21cd6fa614384436ac57b0e18c1cd"

	if !hasherSha1.IsValid(hash) {
		t.Errorf("The hash '%s' should be valid", hash)
	}
}

func TestHasherGpuSha1_IsValid_WithAValueWithIncorrectSize(t *testing.T) {
	var hash = "cb789c4b10a21cd6fa614384436ac57b0e"

	if hasherSha1.IsValid(hash) {
		t.Errorf("The hash '%s' should not be valid", hash)
	}
}

func TestHasherGpuSha1_IsValid_WithAValueWithNotvalidBase64Char(t *testing.T) {
	var hash = "cb789c4b10a21µd6fa614384436ac57b0e18c1cd"

	if hasherSha1.IsValid(hash) {
		t.Errorf("The hash '%s' should not be valid", hash)
	}
}

func testHashSha1(t *testing.T, hasher hashers.Hasher, values []string, expectedHashs []string) {
	var actuals = hasher.Hash(values)

	for i, _ := range values {
		var actual = sha1ToString(actuals[i])
		if actual != expectedHashs[i] {
			t.Errorf("Hash value [position %d] for string '%s' should be '%s' but was '%s'", i, values[i], expectedHashs[i], actual)
		}
	}
}

func assertResultSha1(t *testing.T, result string, expectedWord string) {
	if result != expectedWord {
		t.Errorf("Should have found '%s' but was '%s'", expectedWord, result)
	}
}

func sha1ToString(data []byte) string {
	return hex.EncodeToString(data)
}
