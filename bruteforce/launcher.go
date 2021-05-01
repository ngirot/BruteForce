package bruteforce

import (
	"errors"
	"fmt"
	"github.com/ngirot/BruteForce/bruteforce/conf"
	"github.com/ngirot/BruteForce/bruteforce/display"
	"github.com/ngirot/BruteForce/bruteforce/hashs"
	"github.com/ngirot/BruteForce/bruteforce/hashs/hashers"
	"time"
)

func Launch(hash conf.HashConf, wordConf conf.WordConf, processingUnitConfiguration conf.ProcessingUnitConfiguration) (string, error) {
	var builder = new(TesterBuilder)

	if !hashs.IsValidHash(hash) {
		return "", errors.New("Hash value '" + hash.Value + "' is not valid for type '" + hash.HashType + "'\nExample of a valid hash : '" + hashs.ExampleHash(hash) + "'")
	}

	if builderFunc, error := buildTester(hash, processingUnitConfiguration); error == nil {
		builder.Build = builderFunc
		if processingUnitConfiguration.Type() == conf.Gpu {
			if wordConf.IsForAlphabet() {
				return TestAllStringsGpuForAlphabet(*builder, wordConf, processingUnitConfiguration), nil
			} else {
				return TestAllStringsGpuForDictionary(*builder, wordConf, processingUnitConfiguration), nil
			}
		} else {
			return TestAllStringsCpu(*builder, wordConf, processingUnitConfiguration), nil
		}
	} else {
		return "", error
	}
}

func buildTester(hash conf.HashConf, processingUnitConfiguration conf.ProcessingUnitConfiguration) (func() Tester, error) {
	if hasherCreator, e := hashs.HasherCreator(hash.HashType, processingUnitConfiguration); e == nil {
		var heart = make(chan bool)
		go heartBeat(heart)

		var spinner = display.NewDefaultSpinner()

		return func() Tester {
			var hasher = hasherCreator()
			var tester = new(Tester)
			tester.Notify = displayValue(spinner, heart)
			tester.Test = testValues(hash.Value, hasher)
			tester.Target = func() string { return hash.Value }
			tester.Hasher = func() hashers.Hasher { return hasher }
			return *tester
		}, nil
	} else {
		return nil, e
	}
}

func displayValue(spinner display.Spinner, heart chan bool) func(string) {
	return func(data string) {
		select {
		case <-heart:
			fmt.Printf("\r%s %s...", spinner.Spin(), data)
		default:
		}
	}
}

func testValues(hash string, hasher hashers.Hasher) func([]string) int {
	var hashBytes = hasher.DecodeInput(hash)
	return func(datas []string) int {
		digests := hasher.Hash(datas)
		for i, digest := range digests {
			if hasher.Compare(digest, hashBytes) {
				return i
			}
		}
		return -1
	}
}

func heartBeat(heart chan bool) {
	for {
		time.Sleep(time.Millisecond * 500)
		heart <- true
	}
}
