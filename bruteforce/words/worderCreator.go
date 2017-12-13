package words

func CreateWorder(alphabetFile string, dictionaryFile string, step int, skip int) Worder {
	if dictionaryFile != "" {
		return createDictionaryWorder(dictionaryFile, step, skip)
	} else {
		return createAlphabetWorder(alphabetFile, step, skip)
	}

}

func createDictionaryWorder(dictionatyFile string, step int, skip int) Worder {
	return NewWorderDictionaryFromFile(dictionatyFile, step, skip)
}

func createAlphabetWorder(alphabetFile string, step int, skip int) Worder {
	var alphabet = DefaultAlphabet()
	if alphabetFile != "" {
		alphabet = LoadAlphabet(alphabetFile)
	}

	return NewWorderAlphabet(alphabet, step, skip)
}
