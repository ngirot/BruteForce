package main

import "./words"

type tester func(data string) bool
type status func(data string)

func TestAllStrings(test tester, notifyTesting status) string {
	var worder = words.NewWorder()

	for {
		var word = worder.Next()

		notifyTesting(word)
		if test(word) {
			return word
		}
	}
}
