package main

import (
	"words"
	"runtime"
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

func wordProducer(worder words.Worder, c chan string) {
	for {
		c <- worder.Next()
	}
}

func wordConsumer(c chan string, builder TesterBuilder, r chan string) {
	var tester = builder.Build()

	for word := range c {
		if isHash(word, tester.Test, tester.Notify) != "" {
			r <- word
		}
	}
}

func TestAllStrings(builder TesterBuilder, alphabetFile string) string {
	var alphabet = words.LoadAlphabet(alphabetFile)
	var resultChannel = make(chan string)
	var numberOfChans = runtime.NumCPU()*2 + 1

	for i := 0; i < numberOfChans; i++ {
		var wordChannel = make(chan string, 200)
		go wordProducer(words.NewWorder(alphabet, numberOfChans, i), wordChannel)
		go wordConsumer(wordChannel, builder, resultChannel)
	}

	return <-resultChannel
}
