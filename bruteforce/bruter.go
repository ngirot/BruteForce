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
		if isHash(word, tester.Test, tester.Notify) != "" {
			r <- word
		}
	}
}

func TestAllStrings(builder TesterBuilder, alphabetFile string) string {
	var alphabet = words.DefaultAlphabet()
	if alphabetFile != "" {
		alphabet = words.LoadAlphabet(alphabetFile)
	}
	var resultChannel = make(chan string)
	var numberOfChans = conf.BestNumberOfGoRoutine()

	for i := 0; i < numberOfChans; i++ {
		var worder = words.NewWorder(alphabet, numberOfChans, i)
		go wordConsumer(worder, builder, resultChannel)
	}

	return <-resultChannel
}
