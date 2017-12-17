package hashs

import (
	"testing"
	"encoding/hex"
)

func TestHashMd5SimpleWord(t *testing.T) {
	var hasher = NewHasherMd5()
	if md5ToString(hasher.Hash("test")) != "098f6bcd4621d373cade4e832627b4f6" {
		t.Fail()
	}
}

func TestHashMd5ComplexWord(t *testing.T) {
	var hasher = NewHasherMd5()
	if md5ToString(hasher.Hash("ありがとう &!ç")) != "9ebcc60effbdef7c4d101a7ced1c6b01" {
		t.Fail()
	}
}

func TestHashMd5MultipleTimes(t *testing.T) {
	var hasher = NewHasherMd5()
	var testString = "test"

	for i:=0 ; i<10 ; i++ {
		if md5ToString(hasher.Hash(testString)) != "098f6bcd4621d373cade4e832627b4f6" {
			t.Fail()
		}
	}
}

func TestIsValidMd5WithARealMd5(t *testing.T) {
	var hasher = NewHasherMd5()

	if !hasher.IsValid("098f6bcd4621d373cade4e832627b4f6") {
		t.Fail()
	}
}

func TestIsValidMd5WithAnInvalidSize(t *testing.T) {
	var hasher = NewHasherMd5()

	if hasher.IsValid("098f6bcd4621d373cade4e83") {
		t.Fail()
	}
}

func TestIsValidMd5WithNotABase64(t *testing.T) {
	var hasher = NewHasherMd5()

	if hasher.IsValid("098f6bcd4621$373cade4e832627b4f6") {
		t.Fail()
	}
}

func md5ToString(data [] byte) string {
	return hex.EncodeToString(data)
}