package bruteforce

import (
	"github.com/ngirot/BruteForce/bruteforce/conf"
	"testing"
	"time"
)

type testFunc func(*testing.T)

func Test_Launcher_ShouldFindSimpleSha256HashFromAlphabet_IntegrationTest(t *testing.T) {
	timeout(t, func(t *testing.T) {
		result, error := Launch(conf.HashConf{
			Value:    "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad",
			HashType: "sha256",
		}, conf.WordConf{Alphabet: "testAlphabet.data"})

		checkResult("abc", result, error, t)
	})
}

func Test_Launcher_ShouldFindSimpleSha256HashFromAlphabetWithSalt_IntegrationTest(t *testing.T) {
	timeout(t, func(t *testing.T) {
		result, error := Launch(conf.HashConf{
			Value:    "d43f1b69d87466016e70505bf91c0fd2f075a905b6d18b3bbe2f60ae2ca3dac6",
			HashType: "sha256",
		}, conf.WordConf{Alphabet: "testAlphabet.data", Salt: "salty"})

		checkResult("efg", result, error, t)
	})
}

func Test_Launcher_ShouldFindSimpleSha256HashFromDictionary_IntegrationTest(t *testing.T) {
	result, error := Launch(conf.HashConf{
		Value:    "3eb95849b1228d1f28f57e60aa691e78f295245e79681e3bbb7a5807d4e01ed6",
		HashType: "sha256",
	}, conf.WordConf{Dictionary: "testDictionary.data"})

	checkResult("fromDic", result, error, t)
}

func Test_Launcher_ShouldFindSimpleSha256HashFromDictionaryWithSalt_IntegrationTest(t *testing.T) {
	result, error := Launch(conf.HashConf{
		Value:    "a6c624948b941bd972cf11d58af3cb5bb3ed326eb8107fe7bb5cb4f549706455",
		HashType: "sha256",
	}, conf.WordConf{Dictionary: "testDictionary.data", Salt: "salty"})

	checkResult("maybe", result, error, t)
}

func checkResult(expected string, result string, e error, t *testing.T) {
	if e != nil {
		t.Errorf("A simple hash should not generate an error")
	}
	if result != expected {
		t.Errorf("Expected '%s' but found '%s'", expected, result)
	}
}

func timeout(t *testing.T, testFunction testFunc) {
	timeout := time.After(1000 * time.Millisecond)
	completed := make(chan bool)

	go func() {
		testFunction(t)
		completed <- true
	}()

	select {
	case <-timeout:
		t.Errorf("Test timeout")
	case <-completed:
	}
}
