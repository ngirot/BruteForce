package words

import "github.com/ngirot/BruteForce/bruteforce/conf"

func CreateWorder(wordConf conf.WordConf, step int, skip int, minSize int, maxSize int) Worder {
	if wordConf.IsForAlphabet() {
		return createAlphabetWorder(wordConf.Alphabet, step, skip, minSize, maxSize)
	} else {
		return createDictionaryWorder(wordConf.Dictionary, step, skip)
	}

}

func createDictionaryWorder(dictionatyFile string, step int, skip int) Worder {
	return NewWorderDictionaryFromFile(dictionatyFile, step, skip)
}

func createAlphabetWorder(alphabetFile string, step int, skip int, minSize int, maxSize int) Worder {
	var alphabet = DefaultAlphabet()
	if alphabetFile != "" {
		alphabet = LoadAlphabet(alphabetFile)
	}

	return NewWorderAlphabet(alphabet, step, skip, minSize, maxSize)
}
