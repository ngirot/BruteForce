package main

type tester func(data string) bool
type status func(data string)

func TestAllStrings(test tester, notifyTesting status) string {
	var worder = NewWorder()

	for {
		var word = worder.Next()

		notifyTesting(word)
		if test(word) {
			return word
		}
	}
}