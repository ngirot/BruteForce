package hashs

import (
	"testing"
	"encoding/hex"
)

func TestHashSha1SimpleWord(t *testing.T) {
	var hasher = NewHasherSha1()
	if sha1ToString(hasher.Hash("test")) != "a94a8fe5ccb19ba61c4c0873d391e987982fbbd3" {
		t.Fail()
	}
}

func TestHashSha1ComplexWord(t *testing.T) {
	var hasher = NewHasherSha1()
	if sha1ToString(hasher.Hash("ありがとう &!ç")) != "cb789c4b10a21cd6fa614384436ac57b0e18c1cd" {
		t.Fail()
	}
}

func TestHashSha1MultipleTimes(t *testing.T) {
	var hasher = NewHasherSha1()
	var testString = "test"

	for i:=0 ; i<10 ; i++ {
		if sha1ToString(hasher.Hash(testString)) != "a94a8fe5ccb19ba61c4c0873d391e987982fbbd3" {
			t.Fail()
		}
	}
}

func TestIsValidSha1WithARealSha1(t *testing.T) {
	var hasher = NewHasherSha1()

	if !hasher.IsValid("cb789c4b10a21cd6fa614384436ac57b0e18c1cd") {
		t.Fail()
	}
}

func TestIsValidSha1WithAnInvalidSize(t *testing.T) {
	var hasher = NewHasherSha1()

	if hasher.IsValid("cb789c4b10a21cd6fa614384436ac57b0e") {
		t.Fail()
	}
}

func TestIsValidSha1WithNotABase64(t *testing.T) {
	var hasher = NewHasherSha1()

	if hasher.IsValid("cb789c4b10a21µd6fa614384436ac57b0e18c1cd") {
		t.Fail()
	}
}

func sha1ToString(data [] byte) string {
	return hex.EncodeToString(data)
}