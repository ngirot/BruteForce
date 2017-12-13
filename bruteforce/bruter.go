package bruteforce

import (
	"github.com/ngirot/BruteForce/bruteforce/words"
	"github.com/ngirot/BruteForce/bruteforce/conf"
)

type tester func(data string) bool
type status func(data string)

func isHash(word string, test tester, notifyTesting status) string {
	notifyTesting(word)
	if test(word) {
		return word
	} else {
		return ""
	}
}

func wordConsumer(worder words.Worder, builder TesterBuilder, r chan string) {
	var tester = builder.Build()

	for {
		var word = worder.Next()
		if word == "" {
			r <- ""
		}

		if isHash(word, tester.Test, tester.Notify) != "" {
			r <- word
		}
	}
}

func TestAllStrings(builder TesterBuilder, alphabetFile string, dictionaryFile string) string {

	var resultChannel = make(chan string)
	var numberOfChans = conf.BestNumberOfGoRoutine()

	for i := 0; i < numberOfChans; i++ {
		var worder = words.CreateWorder(alphabetFile, dictionaryFile, numberOfChans, i)
		go wordConsumer(worder, builder, resultChannel)
	}

	return waitForResult(resultChannel, numberOfChans)
}

func waitForResult(resultChannel chan string, numberOfChans int) string {
	var returned = 0
	for v := range resultChannel {
		if v != "" {
			return v
		} else {
			returned++
		}

		if returned == numberOfChans {
			return ""
		}
	}
	return ""
}
