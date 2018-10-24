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
		go wordConsumer(worder, builder, wordConf.Salt, resultChannel)
	}

	return waitForResult(resultChannel, numberOfParallelRoutines)
}

func isHash(word string, salt string, test tester, notifyTesting status) string {
	notifyTesting(word)
	if test(word + salt) {
		return word
	} else {
		return ""
	}
}

func wordConsumer(worder words.Worder, builder TesterBuilder, salt string, r chan string) {
	var tester = builder.Build()

	for {
		var word = worder.Next()
		if word == "" {
			r <- ""
		}

		if isHash(word, salt, tester.Test, tester.Notify) != "" {
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
