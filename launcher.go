package main

import (
	"./hashs"
	"fmt"
)

func Launch(hash string) string {
	var builder = new(TesterBuilder)
	builder.Build = buildTester(hash)

	return TestAllStrings(*builder)
}

var parsed = 0

func buildTester(hash string) func() Tester {
	return func() Tester {
		var hasher = hashs.NewHasher()
		var tester = new(Tester)
		tester.Notify = displayValue
		tester.Test = testValue(hash, hasher)
		return *tester
	}
}

func displayValue(data string) {
	parsed++
	if parsed%1000000 == 0 {
		fmt.Printf("Done: %s\n", data)
	}
}

func testValue(hash string, hasher hashs.Hasher) func(string) bool {
	return func(data string) bool {
		return hasher.Hash(data) == hash
	}
}
