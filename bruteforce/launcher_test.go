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

		if error != nil {
			t.Errorf("A simple hash should not generate an error")
		}

		if result != "abc" {
			t.Errorf("Expected 'abc' but found '%s'", result)
		}
	})
}

func Test_Launcher_ShouldFindSimpleSha256HashFromAlphabetWithSalt_IntegrationTest(t *testing.T) {
	timeout(t, func(t *testing.T) {
		result, error := Launch(conf.HashConf{
			Value:    "30dd1b89fa858b25136181c1eae57f4afa256328a8a6d5275d18c01648f0d121",
			HashType: "sha256",
		}, conf.WordConf{Alphabet: "testAlphabet.data", Salt: "salty"})

		if error != nil {
			t.Errorf("A simple hash should not generate an error")
		}

		if result != "abc" {
			t.Errorf("Expected 'abc' but found '%s'", result)
		}
	})
}

func Test_Launcher_ShouldFindSimpleSha256HashFromDictionary_IntegrationTest(t *testing.T) {
	result, error := Launch(conf.HashConf{
		Value:    "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad",
		HashType: "sha256",
	}, conf.WordConf{Dictionary: "testDictionary.data"})

	if error != nil {
		t.Errorf("A simple hash should not generate an error")
	}

	if result != "abc" {
		t.Errorf("Expected 'abc' but found '%s'", result)
	}
}

func Test_Launcher_ShouldFindSimpleSha256HashFromDictionaryWithSalt_IntegrationTest(t *testing.T) {
	result, error := Launch(conf.HashConf{
		Value:    "30dd1b89fa858b25136181c1eae57f4afa256328a8a6d5275d18c01648f0d121",
		HashType: "sha256",
	}, conf.WordConf{Dictionary: "testDictionary.data", Salt: "salty"})

	if error != nil {
		t.Errorf("A simple hash should not generate an error")
	}

	if result != "abc" {
		t.Errorf("Expected 'abc' but found '%s'", result)
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
