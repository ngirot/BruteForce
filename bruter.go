package main

import (
	"./words"
)

type tester func(data string) bool
type status func(data string)

var alphabet = []string{
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	"(", ")", "`", "~", "!", "@", "#", "$", "%", "^", "&", "*", "-", "+", "=", "|", "\\", "{", "}", "[", "]", ":", ";", "\"", "'", "<", ">", ",", ".", "?", "/",
}

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

func wordConsumer(c chan string, test tester, notifyTesting status, r chan string) {
	for word := range c {
		if isHash(word, test, notifyTesting) != "" {
			r <- word
		}
	}
}

func TestAllStrings(test tester, notifyTesting status) string {
	var wordChannel = make(chan string, 500)
	go wordProducer(words.NewWorder(alphabet), wordChannel)

	var resultChannel = make(chan string)

	for i:=0 ; i<25 ; i++ {
		go wordConsumer(wordChannel, test, notifyTesting, resultChannel)
	}

	return <- resultChannel
}
