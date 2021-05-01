package bruteforce

import (
	"github.com/ngirot/BruteForce/bruteforce/conf"
	"github.com/ngirot/BruteForce/bruteforce/words"
	"math/rand"
)

type tester func(data []string) int
type status func(data string)

func TestAllStringsGpuForAlphabet(builder TesterBuilder, wordConf conf.WordConf, processingUnitConfiguration conf.ProcessingUnitConfiguration) string {
	var worder = words.CreateWorder(wordConf, 1, 0)
	var hasher = builder.Build().Hasher()
	var tester = builder.Build()

	var charSet = worder.GetCharsetIfAvailable().AsCharset()
	var maxWildCards = processingUnitConfiguration.NumberOfWildcardsForDeportedComputingUnit(len(charSet))

	var wildcards = 1
	for wildcards <= maxWildCards {
		tester.Notify(createRandomWord(wildcards, charSet))
		var result = hasher.ProcessWithGpu(charSet, wordConf.SaltBefore, wordConf.SaltAfter, wildcards, tester.Target())
		if result != "" {
			return result
		}
		wildcards++
	}

	for {
		var currentWord = worder.Next()
		tester.Notify(currentWord + createRandomWord(maxWildCards, charSet))
		var result = hasher.ProcessWithGpu(charSet, wordConf.SaltBefore+currentWord, wordConf.SaltAfter, maxWildCards, tester.Target())

		if result != "" {
			return currentWord + result
		}
	}

}

func TestAllStringsGpuForDictionary(builder TesterBuilder, wordConf conf.WordConf, processingUnitConfiguration conf.ProcessingUnitConfiguration) string {
	var worder = words.CreateWorder(wordConf, 1, 0)

	var wordChannel = make(chan []string)
	var resultChannel = make(chan string)
	go fillWords(worder, processingUnitConfiguration, wordChannel)
	go testWords(builder, wordConf, wordChannel, resultChannel)

	return waitForResult(resultChannel, 1)
}

func fillWords(worder words.Worder, processingUnitConfiguration conf.ProcessingUnitConfiguration, wordChan chan []string) {
	for {
		var words = make([]string, processingUnitConfiguration.NumberOfWordsPerIteration())
		var size = 0
		for i, _ := range words {
			var word = worder.Next()
			if word != "" {
				words[i] = word
				size++
			} else {
				break
			}
		}

		if size == 0 {
			wordChan <- []string{}
		} else {
			wordChan <- words[0:size]
		}
	}
}

func testWords(builder TesterBuilder, wordConf conf.WordConf, wordChan chan []string, resultChan chan string) {
	var tester = builder.Build()
	for words := range wordChan {
		if len(words) == 0 {
			resultChan <- ""
		}
		result := isHash(words, wordConf.SaltBefore, wordConf.SaltAfter, tester.Test, tester.Notify)
		if result != "" {
			resultChan <- result
		}
	}
}

func TestAllStringsCpu(builder TesterBuilder, wordConf conf.WordConf, processingUnitConfiguration conf.ProcessingUnitConfiguration) string {

	var resultChannel = make(chan string)
	var numberOfParallelRoutines = processingUnitConfiguration.NumberOfGoRoutines()

	for i := 0; i < numberOfParallelRoutines; i++ {
		var worder = words.CreateWorder(wordConf, numberOfParallelRoutines, i)
		go wordConsumer(worder, builder, wordConf.SaltBefore, wordConf.SaltAfter, processingUnitConfiguration.NumberOfWordsPerIteration(), resultChannel)
	}

	return waitForResult(resultChannel, numberOfParallelRoutines)
}

func isHash(words []string, saltBefore string, saltAfter string, test tester, notifyTesting status) string {
	notifyTesting(words[0])
	var saltedWords = make([]string, len(words))
	copy(saltedWords, words)
	for i, _ := range words {
		saltedWords[i] = saltBefore + words[i] + saltAfter
	}
	var result = test(saltedWords)
	if result != -1 {
		return words[result]
	} else {
		return ""
	}
}

func wordConsumer(worder words.Worder, builder TesterBuilder, saltBefore string, saltAfter string, numberOfWordsPerIteration int, r chan string) {
	var tester = builder.Build()

	for {
		var words = make([]string, numberOfWordsPerIteration)
		for i, _ := range words {
			var word = worder.Next()
			if word == "" {
				r <- ""
				return
			}
			words[i] = word
		}

		result := isHash(words, saltBefore, saltAfter, tester.Test, tester.Notify)
		if result != "" {
			r <- result
		}
	}
}

func waitForResult(resultChannel chan string, numberOfChannels int) string {
	var returned = 0
	for v := range resultChannel {
		if v != "" {
			return v
		} else {
			returned++
		}

		if returned == numberOfChannels {
			return ""
		}
	}
	return ""
}

func createRandomWord(size int, charSet []string) string {
	var result = ""
	for i := 0; i < size; i++ {
		result += charSet[rand.Intn(len(charSet))]
	}

	return result
}
