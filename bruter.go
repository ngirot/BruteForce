package main

type tester func(data string) bool
type status func(data string)

var alphabet = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l"}

func TestAllStrings(testFn tester, statusFn status) string {
	var wordSize int

	for wordSize = 1; wordSize <= 3 ; wordSize++  {

		var word = make([]int, wordSize, wordSize)
		var allWordCompleted = false

		for !allWordCompleted {
			var x = convertWord(word[:])

			statusFn(x)
			if testFn(x) {
				return x
			}

			allWordCompleted = incrementWord(word)
		}
	}

	return ""
}

func incrementWord(word []int) bool{
	var overflow = true

	var i int

	for i = len(word) -1 ; i>=0 && overflow ; i--  {
		var newValue = word[i] + 1
		if newValue % len(alphabet) == 0 {
			newValue = 0
			overflow = true
		} else {
			overflow = false
		}

		word[i] = newValue
	}

	return overflow
}

func convertWord(word []int) string {
	var converted string

	for _,letter := range word {
		converted += alphabet[letter]
	}

	return converted
}