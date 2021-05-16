package hashers

import (
	"encoding/hex"
	"testing"
)

func TestHasherSha256_Hash_WithSimpleWord(t *testing.T) {
	var hasher = NewHasherSha256()
	testHashSha1(t, hasher, []string{"test"}, []string{"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"})
}

func TestHasherSha256_Hash_WithMultipleWord(t *testing.T) {
	var hasher = NewHasherSha256()
	testHashSha1(t, hasher,
		[]string{"test1", "test2"},
		[]string{"1b4f0e9851971998e732078544c96b36c3d01cedf7caa332359d6f1d83567014", "60303ae22b998861bce3b28f33eec1be758a213c86c93c076dbe9f558c11c752"})
}

func TestHasherSha256_Hash_WithUnicodeWord(t *testing.T) {
	var hasher = NewHasherSha256()
	testHashSha1(t, hasher, []string{"ありがとう &!ç"}, []string{"f89eddccb44ae418616060aefe3ca6604d49bc3d0e37e75167333d498532d7aa"})
}

func TestHasherSha256_ProcessWithWildcard_WithSimpleWord(t *testing.T) {
	var hasher = NewHasherSha256()

	hasher.ProcessWithWildcard([]string{"e", "f"}, "", "", 1, "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08.Gt4wp0dJk5qWRaumcfqazMMCAxxerGi")
}


func TestHasherSha256_ProcessWithWildcard_WithSaltBefore(t *testing.T) {
	var hasher = NewHasherSha256()

	hasher.ProcessWithWildcard([]string{"d", "e"}, "t", "", 1, "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08")
}

func TestHasherSha256_ProcessWithWildcard_WithSaltAfter(t *testing.T) {
	var hasher = NewHasherSha256()

	hasher.ProcessWithWildcard([]string{"d", "e", "f"}, "", "t", 1, "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08")
}

func TestHasherSha256_Hash_ConsistencyWithSameHash(t *testing.T) {
	var hasher = NewHasherSha256()
	var testString = "test"

	var firstResult = sha256ToString(hasher.Hash([]string{testString})[0])

	for i:=0 ; i<10 ; i++ {
		var anotherResult = sha256ToString(hasher.Hash([]string{testString})[0])
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

func testHashSha256(t *testing.T, hasher Hasher, values []string, expectedHashs []string) {
	var actuals = hasher.Hash(values)

	for i, _ := range values {
		var actual = sha256ToString(actuals[i])
		if actual != expectedHashs[i] {
			t.Errorf("Hash value [position %d] for string '%s' should be '%s' but was '%s'", i, values[i], expectedHashs[i], actual)
		}
	}
}

func sha256ToString(data [] byte) string {
	return hex.EncodeToString(data)
}