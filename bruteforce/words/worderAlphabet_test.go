package words

import (
	"testing"
)

func TestWorderAlphabetGoToNextWord(t *testing.T) {
	var worder = NewWorderAlphabet(BuildAlphabet([]string{"a", "b"}), 1, 0)

	if worder.Next() != "a" {
		t.Fail()
	}

	if worder.Next() != "b" {
		t.Fail()
	}
}

func TestWorderAlphabetGoToWordWithBiggerSize(t *testing.T) {
	var worder = NewWorderAlphabet(BuildAlphabet([]string{"a", "b"}), 1, 0)

	worder.Next()
	worder.Next()
	if worder.Next() != "aa" {
		t.Fail()
	}
}

func TestWorderAlphabetGoToNextWordAndSkipSome(t *testing.T) {
	var worder = NewWorderAlphabet(BuildAlphabet([]string{"0", "1", "2", "3", "4", "5"}), 2, 0)

	if worder.Next() != "0" {
		t.Fail()
	}

	if worder.Next() != "2" {
		t.Fail()
	}
}

func TestWorderAlphabetGoToWordWithBiggerSizeAndSkipSome(t *testing.T) {
	var worder = NewWorderAlphabet(BuildAlphabet([]string{"0", "1", "2", "3"}), 2, 0)

	worder.Next()
	worder.Next()
	if worder.Next() != "00" {
		t.Fail()
	}
}

func TestWorderAlphabetSkip(t *testing.T)  {
	var worder = NewWorderAlphabet(BuildAlphabet([]string{"0", "1", "2", "3"}), 1, 2)
	if worder.Next() != "2" {
		t.Fail()
	}
}

func TestWorderAlphabetSkipWithAStep(t *testing.T)  {
	var worder = NewWorderAlphabet(BuildAlphabet([]string{"0", "1", "2", "3"}), 5, 2)
	if worder.Next() != "2" {
		t.Fail()
	}
}

func TestWorderAlphabetSkipWithBiggerSize(t *testing.T)  {
	var worder = NewWorderAlphabet(BuildAlphabet([]string{"0", "1"}), 1, 3)
	if worder.Next() != "01" {
		t.Fail()
	}
}
