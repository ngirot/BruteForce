package words

var alphabet = []string{
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	"(", ")", "`", "~", "!", "@", "#", "$", "%", "^", "&", "*", "-", "+", "=", "|", "\\", "{", "}", "[", "]", ":", ";", "\"", "'", "<", ">", ",", ".", "?", "/",
}

type Worder struct {
	letters []int
	wordSize uint16
}

func NewWorder() Worder {
	return Worder{make([]int, 1, 1), 1}
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
		if isOverflow(newValue) {
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
		converted += alphabet[letter]
	}

	return converted
}

func isOverflow(newValue int) bool {
	return newValue%len(alphabet) == 0
}
