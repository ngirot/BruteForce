package bruteforce

import (
	"fmt"
	"github.com/ngirot/BruteForce/bruteforce/conf"
	"strconv"
	"testing"
	"time"
)

type testFunc func(*testing.T)

var cpuConfiguration = conf.NewProcessingUnitConfiguration(false)

func Test_Launcher_ShouldFindSimpleSha256HashFromAlphabet_CPU_IntegrationTest(t *testing.T) {
	runAlphabetTest(t, "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad", "sha256", cpuConfiguration)
}

func Test_Launcher_ShouldFindSimpleSha256HashFromAlphabet_CPU_IntegrationTestInsideLimit(t *testing.T) {
	runAlphabetLimitTestInLimit(t, "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad", "sha256", cpuConfiguration)
}

func Test_Launcher_ShouldFindSimpleSha256HashFromAlphabet_CPU_IntegrationTestBeforeLimit(t *testing.T) {
	runAlphabetLimitTestBeforeLimit(t, "03ac674216f3e15c761ee1a5e255f067953623c8b388b4459e13f978d7c846f4", "sha256", cpuConfiguration)
}

func Test_Launcher_ShouldFindSimpleSha256HashFromAlphabet_CPU_IntegrationTestAfterLimit(t *testing.T) {
	runAlphabetLimitTestAfterLimit(t, "6b86b273ff34fce19d6b804eff5a3f5747ada4eaa22f1d49c01e52ddb7875b4b", "sha256", cpuConfiguration)
}

func Test_Launcher_ShouldFindSimpleSha256HashFromAlphabetWithSaltAfter_CPU_IntegrationTest(t *testing.T) {
	runAlphabetSaltAfter(t, "d43f1b69d87466016e70505bf91c0fd2f075a905b6d18b3bbe2f60ae2ca3dac6", "sha256", cpuConfiguration)
}

func Test_Launcher_ShouldFindSimpleSha256HashFromAlphabetWithSaltBefore_CPU_IntegrationTest(t *testing.T) {
	runAlphabetSaltBefore(t, "73963907b19c38e6cf8d4c9e4c08d966269abd5a17dca0171804be9121bd525c", "sha256", cpuConfiguration)
}

func Test_Launcher_ShouldFindSimpleSha256HashFromDictionary_CPU_IntegrationTest(t *testing.T) {
	runDictionary(t, "3eb95849b1228d1f28f57e60aa691e78f295245e79681e3bbb7a5807d4e01ed6", "sha256", cpuConfiguration)
}

func Test_Launcher_ShouldFindSimpleSha256HashFromDictionaryWithSaltAfter_CPU_IntegrationTest(t *testing.T) {
	runDictionartSaltAfter(t, "a6c624948b941bd972cf11d58af3cb5bb3ed326eb8107fe7bb5cb4f549706455", "sha256", cpuConfiguration)
}

func Test_Launcher_ShouldFindSimpleSha256HashFromDictionaryWithSaltBefore_CPU_IntegrationTest(t *testing.T) {
	runDictionarySaltBefore(t, "75e1aea225bdb80135b82dce07dfb302a2654bd680113eaf05807017e6a83fe4", "sha256", cpuConfiguration)
}

func runAlphabetTest(t *testing.T, valueParam string, typeParam string, processUnitConfiguration conf.ProcessingUnitConfiguration) {
	timeout(t, func(t *testing.T) {
		result, error := Launch(conf.HashConf{
			Value:    valueParam,
			HashType: typeParam,
		}, conf.WordConf{Alphabet: "testAlphabet.data"}, processUnitConfiguration, conf.WorSizeLimitConf{Min: 0, Max: 0})

		checkResult("abc", result, error, t)
	})
}

func runAlphabetSaltAfter(t *testing.T, valueParam string, typeParam string, processUnitConfiguration conf.ProcessingUnitConfiguration) {
	timeout(t, func(t *testing.T) {
		result, error := Launch(conf.HashConf{
			Value:    valueParam,
			HashType: typeParam,
		}, conf.WordConf{Alphabet: "testAlphabet.data", SaltAfter: "salty"}, processUnitConfiguration, conf.WorSizeLimitConf{Min: 0, Max: 0})

		checkResult("efg", result, error, t)
	})
}

func runAlphabetSaltBefore(t *testing.T, valueParam string, typeParam string, processUnitConfiguration conf.ProcessingUnitConfiguration) {
	timeout(t, func(t *testing.T) {
		result, error := Launch(conf.HashConf{
			Value:    valueParam,
			HashType: typeParam,
		}, conf.WordConf{Alphabet: "testAlphabet.data", SaltBefore: "salty"}, processUnitConfiguration, conf.WorSizeLimitConf{Min: 0, Max: 0})

		checkResult("efg", result, error, t)
	})
}

func runDictionary(t *testing.T, valueParam string, typeParam string, processUnitConfiguration conf.ProcessingUnitConfiguration) {
	result, error := Launch(conf.HashConf{
		Value:    valueParam,
		HashType: typeParam,
	}, conf.WordConf{Dictionary: "testDictionary.data"}, processUnitConfiguration, conf.WorSizeLimitConf{Min: 0, Max: 0})

	checkResult("fromDic", result, error, t)
}

func runDictionartSaltAfter(t *testing.T, valueParam string, typeParam string, processUnitConfiguration conf.ProcessingUnitConfiguration) {
	result, error := Launch(conf.HashConf{
		Value:    valueParam,
		HashType: typeParam,
	}, conf.WordConf{Dictionary: "testDictionary.data", SaltAfter: "salty"}, processUnitConfiguration, conf.WorSizeLimitConf{Min: 0, Max: 0})

	checkResult("maybe", result, error, t)
}

func runDictionarySaltBefore(t *testing.T, valueParam string, typeParam string, processUnitConfiguration conf.ProcessingUnitConfiguration) {
	result, error := Launch(conf.HashConf{
		Value:    valueParam,
		HashType: typeParam,
	}, conf.WordConf{Dictionary: "testDictionary.data", SaltBefore: "salty"}, processUnitConfiguration, conf.WorSizeLimitConf{Min: 0, Max: 0})

	checkResult("maybe", result, error, t)
}

func runAlphabetLimitTestInLimit(t *testing.T, valueParam string, typeParam string, processUnitConfiguration conf.ProcessingUnitConfiguration) {
	timeout(t, func(t *testing.T) {
		result, error := Launch(conf.HashConf{
			Value:    valueParam,
			HashType: typeParam,
		}, conf.WordConf{Alphabet: "testAlphabet.data"}, processUnitConfiguration, conf.WorSizeLimitConf{Min: 3, Max: 3})

		checkResult("abc", result, error, t)
	})
}

func runAlphabetLimitTestBeforeLimit(t *testing.T, valueParam string, typeParam string, processUnitConfiguration conf.ProcessingUnitConfiguration) {
	timeout(t, func(t *testing.T) {
		result, error := Launch(conf.HashConf{
			Value:    valueParam,
			HashType: typeParam,
		}, conf.WordConf{Alphabet: "testAlphabet.data"}, processUnitConfiguration, conf.WorSizeLimitConf{Min: 2, Max: 3})

		checkResult("", result, error, t)
	})
}

func runAlphabetLimitTestAfterLimit(t *testing.T, valueParam string, typeParam string, processUnitConfiguration conf.ProcessingUnitConfiguration) {
	timeout(t, func(t *testing.T) {
		result, error := Launch(conf.HashConf{
			Value:    valueParam,
			HashType: typeParam,
		}, conf.WordConf{Alphabet: "testAlphabet.data"}, processUnitConfiguration, conf.WorSizeLimitConf{Min: 2, Max: 3})

		checkResult("", result, error, t)
	})
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
	timeout := time.After(5000 * time.Millisecond)
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

func buildSizeLimit(t *testing.T, min int, max int) conf.WorSizeLimitConf {
	var sizeLimit = strconv.Itoa(min) + "-" + strconv.Itoa(max)
	var config, err = conf.NewWordSizeLimitConf(sizeLimit)
	if err != nil {
		fmt.Printf("%s", err)
		t.Errorf("Test invalid: word size limit '%s' is not valid", sizeLimit)
	}

	return config
}
