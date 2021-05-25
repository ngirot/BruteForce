package cpu

import (
	"github.com/ngirot/BruteForce/bruteforce/hashs/hashers"
	"testing"
)

func TestHasherBcrypt_Hash_WithSimpleWord(t *testing.T) {
	var hasher = NewHasherBcrypt()
	testHashBcrypt(t, hasher, "test", "$2a$10$CHKhRsTMUlT2x8tOdkzJF.Gt4wp0dJk5qWRaumcfqazMMCAxxerGi")
}

func TestHasherBcrypt_Hash_WithUnicodeWord(t *testing.T) {
	var hasher = NewHasherBcrypt()
	testHashBcrypt(t, hasher, "ありがとう &!ç", "$2a$10$n6WTi0p4AFDqbKxQ/j2JuuXu.U1efAOxd9.DT5V8U5XQYKZfFc99C")
}

func TestHasherBcrypt_ProcessWithWildcard_WithSimpleWord(t *testing.T) {
	var hasher = NewHasherBcrypt()

	var result = hasher.ProcessWithWildcard([]string{"e", "f"}, "", "", 1, "$2y$10$uNf7HlBxEyqKQHg/vWARcOgAe1UVAlvXqC9vKDcxYbiUf9i7q37WK")

	assertResultBcrypt(t, result, "e")
}

func TestHasherBcrypt_ProcessWithWildcard_WithSaltBefore(t *testing.T) {
	var hasher = NewHasherBcrypt()

	var result = hasher.ProcessWithWildcard([]string{"d", "e"}, "t", "", 1, "$2y$10$.3rUbnuUCDMcPWmc99TgJOm5U7xO82M.BCp8oEmtL4m2wRxEFR8a2")

	assertResultBcrypt(t, result, "e")
}

func TestHasherBcrypt_ProcessWithWildcard_WithSaltAfter(t *testing.T) {
	var hasher = NewHasherBcrypt()

	var result = hasher.ProcessWithWildcard([]string{"d", "e", "f"}, "", "t", 1, "$2y$10$NYXNaJZKMbBzFQPEHa6n8.mB3lL7vyRcqy4vvOKnR8tVvYAVzBxE.")

	assertResultBcrypt(t, result, "f")
}

func TestHasherBcrypt_IsValid_WithAValidHash(t *testing.T) {
	var hasher = NewHasherBcrypt()
	var hash = "$2y$12$IG1fulw.SpW3meNG/Ht/hOypXW5Nv8gDfcx1Cx6SE.X7tIuQ0vzMm"

	if !hasher.IsValid(hash) {
		t.Errorf("The hash '%s' should be valid", hash)
	}
}

func TestHasherBcrypt_IsValid_WithAnOlderHeaderValidHash(t *testing.T) {
	var hasher = NewHasherBcrypt()
	var hash = "$2a$12$IG1fulw.SpW3meNG/Ht/hOypXW5Nv8gDfcx1Cx6SE.X7tIuQ0vzMm"

	if !hasher.IsValid(hash) {
		t.Errorf("The hash '%s' should be valid", hash)
	}
}

func TestHasherBcrypt_IsValid_WithAValueWithIncorrectSize(t *testing.T) {
	var hasher = NewHasherBcrypt()
	var hash = "$2y$12$IG1fulw.SpW3meNG/Ht/hOypXW5Nv8gDfcx1Cx6SE.X7tIuQ0vzMmxxx"

	if hasher.IsValid(hash) {
		t.Errorf("The hash '%s' should not be valid", hash)
	}
}

func TestHasherBcrypt_IsValid_WithAValueWithIncorrectHeader(t *testing.T) {
	var hasher = NewHasherBcrypt()
	var hash = "$1y$12$IG1fulw.SpW3meNG/Ht/hOypXW5Nv8gDfcx1Cx6SE.X7tIuQ0vzMmxxx"

	if hasher.IsValid(hash) {
		t.Errorf("The hash '%s' should not be valid", hash)
	}
}

func TestHasherBcrypt_IsValid_WithAValueWithIncorrectCost(t *testing.T) {
	var hasher = NewHasherBcrypt()
	var hash = "$1y$1x$IG1fulw.SpW3meNG/Ht/hOypXW5Nv8gDfcx1Cx6SE.X7tIuQ0vzMmxxx"

	if hasher.IsValid(hash) {
		t.Errorf("The hash '%s' should not be valid", hash)
	}
}

func testHashBcrypt(t *testing.T, hasher hashers.Hasher, value string, expectedHash string) {
	var result = hasher.Compare([]byte(value), []byte(expectedHash))
	if !result {
		t.Errorf("Hash value '%s' should be valid for string '%s'", expectedHash, value)
	}
}

func assertResultBcrypt(t *testing.T, result string, expectedWord string) {
	if result != expectedWord {
		t.Errorf("Should have found '%s' but was '%s'", expectedWord, result)
	}
}
