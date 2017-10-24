package main

import (
	"./hashs"
	"fmt"
)

func Launch(hash string) string {
	return TestAllStrings(testValue(hash), displayValue)
}

var hasher = hashs.NewHasher()
var parsed = 0

func displayValue(data string) {
	parsed++
	if parsed%1000000 == 0 {
		fmt.Printf("Done: %s\n", data)
	}
}

func testValue(hash string) func(string) bool {
	return func(data string) bool {
		return hasher.Hash(data) == hash
	}
}
