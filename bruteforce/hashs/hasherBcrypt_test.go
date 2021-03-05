package hashs

import (
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

func testHashBcrypt(t *testing.T, hasher Hasher, value string, expectedHash string) {
	var result = hasher.Compare([]byte(value), []byte(expectedHash))
	if !result {
		t.Errorf("Hash value '%s' should be valid for string '%s'", expectedHash, value)
	}
}