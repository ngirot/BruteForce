package words

import (
	"testing"
)

func TestWorderAlphabet_Next_ShouldGoFromOneLetterToAnother(t *testing.T) {
	var worder = NewWorderAlphabet(BuildAlphabet([]string{"a", "b"}), 1, 0)

	var expectedFirstWord = "a"
	var firstWord = worder.Next()
	if firstWord != expectedFirstWord {
		t.Errorf("First word should be '%s' but was '%s'", expectedFirstWord, firstWord)
	}

	var expectedSecondWord = "b"
	var secondWord = worder.Next()
	if secondWord != expectedSecondWord {
		t.Errorf("Second word should be '%s' but was '%s'", expectedSecondWord, secondWord)
	}
}

func TestWorderAlphabet_Next_ShouldLoopIncreasingSizeWhenAllLettersWasReturned(t *testing.T) {
	var worder = NewWorderAlphabet(BuildAlphabet([]string{"a", "b"}), 1, 0)

	worder.Next()
	worder.Next()

	var expectedResult = "aa"
	var result = worder.Next()
	if result != expectedResult {
		t.Errorf("After all letters are consumed, the word '%s' was expected to be '%s'", result, expectedResult)
	}
}

func TestWorderAlphabet_Next_ShouldSkipWordsDuringLoop(t *testing.T) {
	var worder = NewWorderAlphabet(BuildAlphabet([]string{"0", "1", "2", "3", "4", "5"}), 2, 0)

	var expectedFirstWord = "0"
	var firstWord = worder.Next()
	if firstWord != expectedFirstWord {
		t.Errorf("First word should be '%s' but was '%s'", expectedFirstWord, firstWord)
	}

	var expectedSecondWord = "2"
	var secondWord = worder.Next()
	if secondWord != expectedSecondWord {
		t.Errorf("Second word should be '%s' but was '%s'", expectedSecondWord, secondWord)
	}
}

func TestWorderAlphabet_Next_ShouldSkipByIncresingWordSizeWhenAllLettersWasReturnedOrSkipped(t *testing.T) {
	var worder = NewWorderAlphabet(BuildAlphabet([]string{"0", "1", "2", "3"}), 2, 0)

	worder.Next()
	worder.Next()

	var expectedResult = "00"
	var result = worder.Next()
	if result != expectedResult {
		t.Errorf("After all letters are consumed (and somme skipped), the word '%s' was expected to be '%s'", result, expectedResult)
	}
}

func TestWorderAlphabet_Next_ShouldUseSkipAtInitialisation(t *testing.T) {
	var worder = NewWorderAlphabet(BuildAlphabet([]string{"0", "1", "2", "3"}), 1, 2)

	var expectedResult = "2"
	var result = worder.Next()
	if result != expectedResult {
		t.Errorf("The word '%s' was expected to be '%s' when some was skipped at initialisation", result, expectedResult)
	}
}

func TestWorderAlphabet_Next_ShouldMixSkipAndStep(t *testing.T) {
	var worder = NewWorderAlphabet(BuildAlphabet([]string{"0", "1", "2", "3"}), 5, 2)

	var expectedResult = "2"
	var result = worder.Next()
	if result != expectedResult {
		t.Errorf("The word '%s' was expected to be '%s' when some was skipped at initialisation and a some are skipped by step value", result, expectedResult)
	}
}

func TestWorderAlphabet_Next_ShouldSkipMoreWordsThanAlphabetSize(t *testing.T) {
	var worder = NewWorderAlphabet(BuildAlphabet([]string{"0", "1"}), 1, 3)

	var expectedResult = "01"
	var result = worder.Next()
	if result != expectedResult {
		t.Errorf("The word '%s' was expected to be '%s' when the initialisation skip is bigger than initialisation size", result, expectedResult)
	}
}

func TestWorderAlphabet_Next_ShouldSkipMoreWordsThanTwiceAlphabetSize(t *testing.T) {
	var worder = NewWorderAlphabet(BuildAlphabet([]string{"0", "1"}), 1, 12)

	var expectedResult = "110"
	var result = worder.Next()
	if result != expectedResult {
		t.Errorf("The word '%s' was expected to be '%s' when the initialisation skip is bigger than initialisation size", result, expectedResult)
	}
}
