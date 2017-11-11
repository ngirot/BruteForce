package words

type Worder struct {
	letters []int
	alphabet Alphabet
	wordSize int
	step int
}

func NewWorder(alphabet Alphabet, step int, skip int) Worder {
	var worder = Worder{make([]int, 1, 1), alphabet, 1, step}
	worder.updateToNextWord(skip)
	return worder
}

func (w *Worder) Next() string {
	var word = w.generateWord()

	w.updateToNextWord(w.step)

	return word
}

func (w *Worder) updateToNextWord(step int) {
	var overflow = step
	var position int

	for overflow != 0 {
		for position = w.wordSize-1; position >= 0 && overflow != 0; position-- {
			var newValue = w.letters[position] + overflow
			overflow = newValue / w.alphabet.length()

			w.letters[position] = newValue % w.alphabet.length()
		}

		if overflow > 0 {
			overflow--
			w.wordSize++
			w.letters = prepend(w.letters, 0)
		}
	}
}

func (w *Worder) generateWord() string {
	var converted string

	for _,position := range w.letters {
		converted += w.alphabet.letter(position)
	}

	return converted
}

func prepend(slice []int, value int) []int {
	var s1 = make([]int, len(slice) + 1)
	s1[0] = value
	copy(s1[1:], slice)
	return s1
}