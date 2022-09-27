package bruteforce

import (
	"github.com/ngirot/BruteForce/bruteforce/conf"
	"github.com/ngirot/BruteForce/bruteforce/hashs/hashers"
	"github.com/ngirot/BruteForce/bruteforce/maths"
	"github.com/ngirot/BruteForce/bruteforce/words"
	"math/rand"
)

type tester func(data []string) int
type status func(data string, numberComputed int)

func TestAllStringsForAlphabet(builder TesterBuilder, wordConf conf.WordConf, processingUnitConfiguration conf.ProcessingUnitConfiguration, wordSizeConfiguration conf.WorSizeLimitConf) string {
	var charSet = words.CreateWorder(wordConf, 1, 0, 0, 0).GetCharsetIfAvailable().AsCharset()
	var tester = builder.Build()
	var maxWildCards = processingUnitConfiguration.NumberOfWildcardsForDeportedComputingUnit(len(charSet))

	var result = processWildcardsWithoutWorder(maxWildCards, tester, charSet, builder, wordConf, wordSizeConfiguration)
	if result != "" || sizeLimitIsReached(wordSizeConfiguration, maxWildCards) {
		return result
	}

	return processWildcardsWithWorder(builder, wordConf, processingUnitConfiguration, wordSizeConfiguration, tester, maxWildCards, charSet)
}

func sizeLimitIsReached(wordSizeConfiguration conf.WorSizeLimitConf, maxWildCards int) bool {
	return wordSizeConfiguration.Max != 0 && wordSizeConfiguration.Max <= maxWildCards
}

func processWildcardsWithWorder(builder TesterBuilder, wordConf conf.WordConf, processingUnitConfiguration conf.ProcessingUnitConfiguration, wordSizeConfiguration conf.WorSizeLimitConf, tester Tester, maxWildCards int, charSet []string) string {
	var resultChannel = make(chan string)
	var numberOfParallelRoutines = processingUnitConfiguration.NumberOfGoRoutines()

	for i := 0; i < numberOfParallelRoutines; i++ {
		var currentHasher = builder.Build().Hasher()
		maxSizeWithoutWildcards := computeSizeWithoutWildCards(wordSizeConfiguration.Max, maxWildCards)
		minSizeWithoutWildcards := computeSizeWithoutWildCards(wordSizeConfiguration.Min, maxWildCards)

		var currentWorder = words.CreateWorder(wordConf, numberOfParallelRoutines, i, minSizeWithoutWildcards, maxSizeWithoutWildcards)
		go processWithWildCards(currentWorder, tester, maxWildCards, charSet, currentHasher, wordConf.SaltBefore, wordConf.SaltAfter, resultChannel)
	}

	return waitForResult(resultChannel, numberOfParallelRoutines)
}

func computeSizeWithoutWildCards(value int, maxWildCards int) int {
	var maxSizeWithoutWildcards = value - maxWildCards
	if value == 0 {
		maxSizeWithoutWildcards = 0
	}
	return maxSizeWithoutWildcards
}

func processWildcardsWithoutWorder(maxWildCards int, tester Tester, charSet []string, builder TesterBuilder, wordConf conf.WordConf, wordSizeConfiguration conf.WorSizeLimitConf) string {
	var hasher = builder.Build().Hasher()

	var currentWildcards = maths.MaxInt(1, wordSizeConfiguration.Min)
	var maxNumberOfWildCards = maxWildCards
	if wordSizeConfiguration.Max != 0 {
		maxNumberOfWildCards = maths.MinInt(maxWildCards, wordSizeConfiguration.Max)
	}

	for currentWildcards <= maxNumberOfWildCards {
		tester.Notify(createRandomWord(currentWildcards, charSet), maths.PowInt(len(charSet), currentWildcards))
		var result = hasher.ProcessWithWildcard(charSet, wordConf.SaltBefore, wordConf.SaltAfter, currentWildcards, tester.Target())
		if result != "" {
			return result
		}
		currentWildcards++
	}
	return ""
}

func processWithWildCards(currentWorder words.Worder, tester Tester, maxWildCards int, charSet []string, hasher hashers.Hasher, saltBefore string, saltAfter string, resultChannel chan string) {
	var computedPerPass = maths.PowInt(len(charSet), maxWildCards)
	for {
		var currentWord = currentWorder.Next()
		if currentWord == "" {
			resultChannel <- ""
			return
		}
		tester.Notify(currentWord+createRandomWord(maxWildCards, charSet), computedPerPass)
		var result = hasher.ProcessWithWildcard(charSet, saltBefore+currentWord, saltAfter, maxWildCards, tester.Target())

		if result != "" {
			resultChannel <- currentWord + result
			return
		}
	}
}

func TestAllStringsForDictionary(builder TesterBuilder, wordConf conf.WordConf, processingUnitConfiguration conf.ProcessingUnitConfiguration) string {
	var resultChannel = make(chan string)
	var numberOfParallelRoutines = processingUnitConfiguration.NumberOfGoRoutines()

	for i := 0; i < numberOfParallelRoutines; i++ {
		var wordChannel = make(chan []string)

		var worder = words.CreateWorder(wordConf, numberOfParallelRoutines, i, 0, 0)
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
	notifyTesting(words[0], len(words))
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
