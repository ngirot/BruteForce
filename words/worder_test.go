package words

import "testing"

func TestGoToNextWord(t *testing.T) {
	var worder = NewWorder([]string{"a", "b"})

	if worder.Next() != "a" {
		t.Fail()
	}

	if worder.Next() != "b" {
		t.Fail()
	}
}

func TestGoToWordWithBiggerSize(t *testing.T) {
	var worder = NewWorder([]string{"a", "b"})

	worder.Next()
	worder.Next()
	if worder.Next() != "aa" {
		t.Fail()
	}
}
