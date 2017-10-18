package main

type tester func(data string) bool
type status func(data string)

var alphabet = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l"}

func TestAllStrings(test tester, notifyTesting status) string {
	var wordSize int

	for wordSize = 1; wordSize <= 3 ; wordSize++  {

		var letters = make([]int, wordSize, wordSize)
		var allWordsCompleted = false

		for !allWordsCompleted {
			var word = generateWord(letters[:])

			notifyTesting(word)
			if test(word) {
				return word
			}

			allWordsCompleted = updateToNextWord(letters)
		}
	}

	return ""
}

func updateToNextWord(word []int) bool {
	var overflow = true

	var position int

	for position = len(word) -1 ; position >=0 && overflow ; position--  {
		var newValue = word[position] + 1
		if isOverflow(newValue) {
			newValue = 0
			overflow = true
		} else {
			overflow = false
		}

		word[position] = newValue
	}

	return overflow
}
func isOverflow(newValue int) bool {
	return newValue%len(alphabet) == 0
}

func generateWord(word []int) string {
	var converted string

	for _,letter := range word {
		converted += alphabet[letter]
	}

	return converted
}