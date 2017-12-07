package bruteforce

import (
	"runtime"
	"github.com/ngirot/BruteForce/bruteforce/words"
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
	var numberOfChans = runtime.NumCPU()*2 + 1

	for i := 0; i < numberOfChans; i++ {
		var worder = words.NewWorder(alphabet, numberOfChans, i)
		go wordConsumer(worder, builder, resultChannel)
	}

	return <-resultChannel
}
