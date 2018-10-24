package words

import (
	"bufio"
	"os"
)

type worderDictionary struct {
	words []string
	step int
	position int
}


func NewWorderDictionaryFromFile(fileName string, step int, skip int) Worder {
	var file,_ = os.Open(fileName)
	defer file.Close()

	var words = make([]string,0)
	var scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	return &worderDictionary{words, step, skip}
}

func NewWorderDictionary(words []string, step int, skip int) Worder {
	return &worderDictionary{words, step, skip}
}

func (w *worderDictionary) Next() string {
	if w.position >= len(w.words) {
		return ""
	}

	var result = w.words[w.position]
	w.position+=w.step
	return result
}