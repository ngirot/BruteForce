package hashs

import (
	"testing"
	"encoding/hex"
)

func TestHashSha512SimpleWord(t *testing.T) {
	var hasher = NewHasherSha512()
	if sha512ToString(hasher.Hash("test")) != "ee26b0dd4af7e749aa1a8ee3c10ae9923f618980772e473f8819a5d4940e0db27ac185f8a0e1d5f84f88bc887fd67b143732c304cc5fa9ad8e6f57f50028a8ff" {
		t.Fail()
	}
}

func TestHashSha512ComplexWord(t *testing.T) {
	var hasher = NewHasherSha512()
	if sha512ToString(hasher.Hash("ありがとう &!ç")) != "5b35cc2b8d75331db1ad1550a1e7b4f3538ca9bd9f1b4d9a6dff5b8c938b331a4844aa874639da3acd38c79d071d2a146916e215f12d9d562a2637a9d0927943" {
		t.Fail()
	}
}

func TestHashSha512MultipleTimes(t *testing.T) {
	var hasher = NewHasherSha512()
	var testString = "test"

	for i:=0 ; i<10 ; i++ {
		if sha512ToString(hasher.Hash(testString)) != "ee26b0dd4af7e749aa1a8ee3c10ae9923f618980772e473f8819a5d4940e0db27ac185f8a0e1d5f84f88bc887fd67b143732c304cc5fa9ad8e6f57f50028a8ff" {
			t.Fail()
		}
	}
}

func sha512ToString(data [] byte) string {
	return hex.EncodeToString(data)
}