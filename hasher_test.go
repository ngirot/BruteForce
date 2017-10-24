package main

import (
	"testing"
)

func TestHashSimpleWord(t *testing.T) {
	var hasher = NewHasher()
	if hasher.Hash("test") != "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08" {
		t.Fail()
	}
}

func TestHashComplexWord(t *testing.T) {
	var hasher = NewHasher()
	if hasher.Hash("ありがとう &!ç") != "f89eddccb44ae418616060aefe3ca6604d49bc3d0e37e75167333d498532d7aa" {
		t.Fail()
	}
}

func TestHashMultipleTimes(t *testing.T) {
	var hasher = NewHasher()
	var testString = "test"

	for i:=0 ; i<10 ; i++ {
		if hasher.Hash(testString) != "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08" {
			t.Fail()
		}
	}
}