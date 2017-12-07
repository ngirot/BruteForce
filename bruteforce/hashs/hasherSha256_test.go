package hashs

import (
	"testing"
	"encoding/hex"
)

func TestHashSha256SimpleWord(t *testing.T) {
	var hasher = NewHasherSha256()
	if sha256ToString(hasher.Hash("test")) != "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08" {
		t.Fail()
	}
}

func TestHashSha256ComplexWord(t *testing.T) {
	var hasher = NewHasherSha256()
	if sha256ToString(hasher.Hash("ありがとう &!ç")) != "f89eddccb44ae418616060aefe3ca6604d49bc3d0e37e75167333d498532d7aa" {
		t.Fail()
	}
}

func TestHashSha256MultipleTimes(t *testing.T) {
	var hasher = NewHasherSha256()
	var testString = "test"

	for i:=0 ; i<10 ; i++ {
		if sha256ToString(hasher.Hash(testString)) != "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08" {
			t.Fail()
		}
	}
}

func sha256ToString(data [] byte) string {
	return hex.EncodeToString(data)
}