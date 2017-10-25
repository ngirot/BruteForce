package words

type Worder struct {
	letters []int
	alphabet []string
	wordSize uint16
}

func NewWorder(alphabet []string) Worder {
	return Worder{make([]int, 1, 1), alphabet, 1}
}

func (w *Worder) Next() string {
	var word = w.generateWord()

	w.updateToNextWord()

	return word
}

func (w *Worder) updateToNextWord() {
	var overflow = true

	var position int

	for position = len(w.letters) - 1; position >= 0 && overflow; position-- {
		var newValue= w.letters[position] + 1	
		if w.isOverflow(newValue) {
			newValue = 0
			overflow = true
		} else {
			overflow = false
		}

		w.letters[position] = newValue
	}

	if overflow {
		w.wordSize++
		w.letters = make([]int, w.wordSize, w.wordSize)
	}
}

func (w *Worder) generateWord() string {
	var converted string

	for _,letter := range w.letters {
		converted += w.alphabet[letter]
	}

	return converted
}

func (w *Worder) isOverflow(newValue int) bool {
	return newValue%len(w.alphabet) == 0
}
