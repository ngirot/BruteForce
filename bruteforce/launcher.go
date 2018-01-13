package bruteforce

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/ngirot/BruteForce/bruteforce/hashs"
	"time"
	"github.com/ngirot/BruteForce/bruteforce/display"
	"errors"
	"github.com/ngirot/BruteForce/bruteforce/conf"
)

func Launch(hash conf.HashConf, wordConf conf.WordConf) (string, error) {
	var builder = new(TesterBuilder)

	if !hashs.IsValidhash(hash) {
		return "", errors.New("hash value '" + hash.Value + "' is not valid for type '" + hash.HashType + "'")
	}

	if builderFunc, error := buildTester(hash); error == nil {
		builder.Build = builderFunc
		return TestAllStrings(*builder, wordConf), nil
	} else {
		return "", error
	}
}

func buildTester(hash conf.HashConf) (func() Tester, error) {
	if hasherCreator, e := hashs.HasherCreator(hash.HashType); e == nil {
		var heart = make(chan bool)
		go heartBeat(heart)

		var spinner = display.NewDefaultSpinner()

		return func() Tester {
			var hasher = hasherCreator()
			var tester = new(Tester)
			tester.Notify = displayValue(spinner, heart)
			tester.Test = testValue(hash.Value, hasher)
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
