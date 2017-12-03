package main

import (
	"./hashs"
	"bytes"
	"encoding/hex"
	"fmt"
)

func Launch(hash string, alphabetFile string, hashType string) (string, error) {
	var builder = new(TesterBuilder)

	if builderFunc, error := buildTester(hash, hashType); error == nil {
		builder.Build = builderFunc
		return TestAllStrings(*builder, alphabetFile), nil
	} else {
		return "", error
	}
}

var parsed = 0

func buildTester(hash string, hashType string) (func() Tester, error) {
	if hasherCreator, e := hashs.HasherCreator(hashType); e == nil {
		return func() Tester {
			var hasher = hasherCreator()
			var tester = new(Tester)
			tester.Notify = displayValue
			tester.Test = testValue(hash, hasher)
			return *tester
		}, nil
	} else {
		return nil, e
	}
}

func displayValue(data string) {
	parsed++
	if parsed%1000000 == 0 {
		fmt.Printf("Done: %s\n", data)
	}
}

func testValue(hash string, hasher hashs.Hasher) func(string) bool {
	var hashBytes, _ = hex.DecodeString(hash)
	return func(data string) bool {
		return bytes.Equal(hasher.Hash(data), hashBytes)
	}
}
