package bruteforce

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/ngirot/BruteForce/bruteforce/hashs"
	"time"
	"github.com/ngirot/BruteForce/bruteforce/display"
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
		var heart = make(chan bool)
		go heartBeat(heart)

		var spinner = display.NewDefaultSpinner()

		return func() Tester {
			var hasher = hasherCreator()
			var tester = new(Tester)
			tester.Notify = displayValue(spinner, heart)
			tester.Test = testValue(hash, hasher)
			return *tester
		}, nil
	} else {
		return nil, e
	}
}

func displayValue(spinner display.Spinner, heart chan bool) func(string){

	return func(data string) {
		select {
		case <- heart:
			fmt.Printf("\r%s %s...", spinner.Spin(), data)
		default:
		}
	}
}

func testValue(hash string, hasher hashs.Hasher) func(string) bool {
	var hashBytes, _ = hex.DecodeString(hash)
	return func(data string) bool {
		return bytes.Equal(hasher.Hash(data), hashBytes)
	}
}

func heartBeat(heart chan bool) {
	for {
		time.Sleep(time.Millisecond * 500)
		heart <- true
	}
}
