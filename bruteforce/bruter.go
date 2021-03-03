package bruteforce

import (
	"github.com/ngirot/BruteForce/bruteforce/conf"
	"github.com/ngirot/BruteForce/bruteforce/words"
)

type tester func(data string) bool
type status func(data string)

func TestAllStrings(builder TesterBuilder, wordConf conf.WordConf) string {

	var resultChannel = make(chan string)
	var numberOfParallelRoutines = conf.BestNumberOfGoRoutine()

	for i := 0; i < numberOfParallelRoutines; i++ {
		var worder = words.CreateWorder(wordConf.Alphabet, wordConf.Dictionary, numberOfParallelRoutines, i)
		go wordConsumer(worder, builder, wordConf.SaltBefore, wordConf.SaltAfter, resultChannel)
	}

	return waitForResult(resultChannel, numberOfParallelRoutines)
}

func isHash(word string, saltBefore string, saltAfter string, test tester, notifyTesting status) string {
	notifyTesting(word)
	if test(saltBefore + word + saltAfter) {
		return word
	} else {
		return ""
	}
}

func wordConsumer(worder words.Worder, builder TesterBuilder, saltBefore string, saltAfter string, r chan string) {
	var tester = builder.Build()

	for {
		var word = worder.Next()
		if word == "" {
			r <- ""
			return
		}

		if isHash(word, saltBefore, saltAfter, tester.Test, tester.Notify) != "" {
			r <- word
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
