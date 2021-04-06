package bruteforce

import (
	"github.com/ngirot/BruteForce/bruteforce/conf"
	"github.com/ngirot/BruteForce/bruteforce/words"
)

type tester func(data []string) int
type status func(data string)

func TestAllStrings(builder TesterBuilder, wordConf conf.WordConf, processingUnitConfiguration conf.ProcessingUnitConfiguration) string {

	var resultChannel = make(chan string)
	var numberOfParallelRoutines = processingUnitConfiguration.NumberOfGoRoutines()

	for i := 0; i < numberOfParallelRoutines; i++ {
		var worder = words.CreateWorder(wordConf.Alphabet, wordConf.Dictionary, numberOfParallelRoutines, i)
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

func wordConsumer(worder words.Worder, builder TesterBuilder, saltBefore string,saltAfter string, numberOfWordsPerIteration int, r chan string) {
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
