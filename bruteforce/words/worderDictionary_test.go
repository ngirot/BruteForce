package words

import (
	"testing"
)

func TestWorderDictionaryGoToNextWord(t *testing.T) {
	var worder = NewWorderDictionary([]string{"word1", "word2"}, 1, 0)

	if worder.Next() != "word1" {
		t.Fail()
	}

	if worder.Next() != "word2" {
		t.Fail()
	}
}

func TestWorderDictionaryGoToNextWordShouldStopWhenAllValuesHasBeenReturned(t *testing.T) {
	var worder = NewWorderDictionary([]string{"word1", "word2"}, 1, 0)

	worder.Next()
	worder.Next()

	if worder.Next() != "" {
		t.Fail()
	}
}

func TestWorderDictionaryGoToNextWordAndSkipSome(t *testing.T) {
	var worder = NewWorderDictionary([]string{"word0", "word1", "word2", "word3", "word4", "word5"}, 2, 0)

	if worder.Next() != "word0" {
		t.Fail()
	}

	if worder.Next() != "word2" {
		t.Fail()
	}
}

func TestWorderDictionarySkip(t *testing.T)  {
	var worder = NewWorderDictionary([]string{"word0", "word1", "word2", "word3"}, 1, 2)
	if worder.Next() != "word2" {
		t.Fail()
	}
}



func TestWorderDictionarySkipWithAStep(t *testing.T)  {
	var worder = NewWorderDictionary([]string{"word0", "word1", "word2", "word3"}, 5, 2)
	if worder.Next() != "word2" {
		t.Fail()
	}
}

func TestWorderDictionaryGoToWordReturnEmptyStringIfSkipIsTooHigh(t *testing.T) {
	var worder = NewWorderDictionary([]string{"word0", "word1", "word2"}, 1, 4)

	if worder.Next() != "" {
		t.Fail()
	}
}

