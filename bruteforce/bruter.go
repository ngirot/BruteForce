package bruteforce

import (
	"github.com/ngirot/BruteForce/bruteforce/conf"
	"github.com/ngirot/BruteForce/bruteforce/hashs/hashers"
	"github.com/ngirot/BruteForce/bruteforce/words"
	"math/rand"
)

type tester func(data []string) int
type status func(data string)

func TestAllStringsForAlphabet(builder TesterBuilder, wordConf conf.WordConf, processingUnitConfiguration conf.ProcessingUnitConfiguration) string {
	var charSet = words.CreateWorder(wordConf, 1, 0).GetCharsetIfAvailable().AsCharset()
	var tester = builder.Build()

	var maxWildCards = processingUnitConfiguration.NumberOfWildcardsForDeportedComputingUnit(len(charSet))

	var result = processWildcardsWithoutWorder(maxWildCards, tester, charSet, builder, wordConf)
	if result != "" {
		return result
	}

	return processWildcardsWithWorder(builder, wordConf, processingUnitConfiguration, tester, maxWildCards, charSet)
}

func processWildcardsWithWorder(builder TesterBuilder, wordConf conf.WordConf, processingUnitConfiguration conf.ProcessingUnitConfiguration, tester Tester, maxWildCards int, charSet []string) string {
	var resultChannel = make(chan string)
	var numberOfParallelRoutines = processingUnitConfiguration.NumberOfGoRoutines()

	for i := 0; i < numberOfParallelRoutines; i++ {
		var currentHasher = builder.Build().Hasher()
		var currentWorder = words.CreateWorder(wordConf, numberOfParallelRoutines, i)
		go processWithWildCards(currentWorder, tester, maxWildCards, charSet, currentHasher, wordConf.SaltBefore, wordConf.SaltAfter, resultChannel)
	}

	return waitForResult(resultChannel, numberOfParallelRoutines)
}

func processWildcardsWithoutWorder(maxWildCards int, tester Tester, charSet []string, builder TesterBuilder, wordConf conf.WordConf) string {
	var hasher = builder.Build().Hasher()

	var wildcards = 1
	for wildcards <= maxWildCards {
		tester.Notify(createRandomWord(wildcards, charSet))
		var result = hasher.ProcessWithWildcard(charSet, wordConf.SaltBefore, wordConf.SaltAfter, wildcards, tester.Target())
		if result != "" {
			return result
		}
		wildcards++
	}
	return ""
}

func processWithWildCards(currentWorder words.Worder, tester Tester, maxWildCards int, charSet []string, hasher hashers.Hasher, saltBefore string, saltAfter string, resultChannel chan string) {
	for {
		var currentWord = currentWorder.Next()
		tester.Notify(currentWord + createRandomWord(maxWildCards, charSet))
		var result = hasher.ProcessWithWildcard(charSet, saltBefore+currentWord, saltAfter, maxWildCards, tester.Target())

		if result != "" {
			resultChannel <- currentWord + result
		}
	}
}

func TestAllStringsForDictionary(builder TesterBuilder, wordConf conf.WordConf, processingUnitConfiguration conf.ProcessingUnitConfiguration) string {
	var resultChannel = make(chan string)
	var numberOfParallelRoutines = processingUnitConfiguration.NumberOfGoRoutines()

	for i := 0; i < numberOfParallelRoutines; i++ {
		var wordChannel = make(chan []string)

		var worder = words.CreateWorder(wordConf, numberOfParallelRoutines, i)
		go fillWords(worder, processingUnitConfiguration, wordChannel)
		go testWords(builder, wordConf, wordChannel, resultChannel)
	}

	return waitForResult(resultChannel, numberOfParallelRoutines)
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
			return
		}
		result := isHash(words, wordConf.SaltBefore, wordConf.SaltAfter, tester.Test, tester.Notify)
		if result != "" {
			resultChan <- result
		}
	}
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
