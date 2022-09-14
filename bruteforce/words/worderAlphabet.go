package words

import (
	"github.com/ngirot/BruteForce/bruteforce/maths"
)

type worderAlphabet struct {
	letters  []int
	alphabet Alphabet
	wordSize int
	step     int
	maxSize  int
}

func NewWorderAlphabet(alphabet Alphabet, step int, skip int, minSize int, maxSize int) Worder {
	var wordSize = maths.MaxInt(1, minSize)
	var worder = worderAlphabet{make([]int, wordSize), alphabet, wordSize, step, maxSize}
	worder.updateToNextWord(skip)
	return &worder
}

func (w *worderAlphabet) Next() string {
	if w.isMaxSizeReached() {
		return ""
	}

	var word = w.generateWord()

	w.updateToNextWord(w.step)

	return word
}

func (w *worderAlphabet) isMaxSizeReached() bool {
	return w.maxSize != 0 && w.wordSize > w.maxSize
}

func (w *worderAlphabet) updateToNextWord(step int) {
	for i := 0; i < step; i++ {
		var overflow = 1
		var position int

		for overflow != 0 {
			for position = w.wordSize - 1; position >= 0 && overflow != 0; position-- {
				var newValue = w.letters[position] + overflow
				overflow = newValue / w.alphabet.Length()

				w.letters[position] = newValue % w.alphabet.Length()
			}

			if overflow > 0 {
				overflow = 0
				w.wordSize++
				w.letters = make([]int, w.wordSize)
			}
		}
	}
}

func (w *worderAlphabet) GetCharsetIfAvailable() Alphabet {
	return w.alphabet
}

func (w *worderAlphabet) generateWord() string {
	var converted string

	for _, position := range w.letters {
		converted += w.alphabet.Letter(position)
	}

	return converted
}
