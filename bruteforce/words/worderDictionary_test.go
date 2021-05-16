package words

import (
	"testing"
)

func TestWorderDictionary_Next_ShouldGoFromOneWordToAnother(t *testing.T) {
	var worder = NewWorderDictionary([]string{"word1", "word2"}, 1, 0)

	var expectedFirstWord = "word1"
	var firstWord = worder.Next()
	if firstWord != expectedFirstWord {
		t.Errorf("First word should be '%s' but was '%s'", expectedFirstWord, firstWord)
	}

	var expectedSecondWord = "word2"
	var secondWord = worder.Next()
	if secondWord != expectedSecondWord {
		t.Errorf("Second word should be '%s' but was '%s'", expectedFirstWord, firstWord)
	}
}

func TestWorderDictionary_Next_ShouldReturnEmptyStringWhenAllWOrdsWereReturned(t *testing.T) {
	var worder = NewWorderDictionary([]string{"word1", "word2"}, 1, 0)

	worder.Next()
	worder.Next()

	var result = worder.Next()
	if result != "" {
		t.Errorf("When the dictionary is empty, the result should be an empty string but was '%s'", result)
	}
}

func TestWorderDictionary_Next_SkipSomeWordsWithAStepValue(t *testing.T) {
	var worder = NewWorderDictionary([]string{"word0", "word1", "word2", "word3"}, 2, 0)

	var expectedFirstWord = "word0"
	var firstWord = worder.Next()
	if firstWord != expectedFirstWord {
		t.Errorf("First word should be '%s' but was '%s'", expectedFirstWord, firstWord)
	}

	var expectedSecondWord = "word2"
	var secondWord = worder.Next()
	if secondWord != expectedSecondWord {
		t.Errorf("Second word should be '%s' but was '%s'", expectedFirstWord, firstWord)
	}
}

func TestWorderDictionary_Next_ShouldSkipValuesAtInitialisation(t *testing.T) {
	var worder = NewWorderDictionary([]string{"word0", "word1", "word2", "word3"}, 1, 2)

	var expectedFirstWord = "word2"
	var firstWord = worder.Next()
	if firstWord != expectedFirstWord {
		t.Errorf("First word should be '%s' but was '%s'", expectedFirstWord, firstWord)
	}
}

func TestWorderDictionary_Next_ShouldMixSkipAndStep(t *testing.T) {
	var worder = NewWorderDictionary([]string{"word0", "word1", "word2", "word3"}, 5, 2)

	var expectedFirstWord = "word2"
	var firstWord = worder.Next()
	if firstWord != expectedFirstWord {
		t.Errorf("First word should be '%s' but was '%s'", expectedFirstWord, firstWord)
	}
}

func TestWorderDictionary_Next_ShouldRetundEmptyStringIfSkipValueIdBiggerThanDictionarySize(t *testing.T) {
	var worder = NewWorderDictionary([]string{"word0", "word1", "word2"}, 1, 4)

	var result = worder.Next()
	if result != "" {
		t.Errorf("When the dictionary is empty, the result should be an empty string but was '%s'", result)
	}
}
