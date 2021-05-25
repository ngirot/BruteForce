package cpu

import (
	"testing"
)

func TestExpand_EmptyCharset(t *testing.T) {
	result := expand([]string{}, 100)

	assertSize(t, result, 0)
}

func TestExpand_ExpansionFactor_1(t *testing.T) {
	result := expand([]string{"a", "b"}, 1)

	assertSize(t, result, 2)
	assertStringAtPosition(t, result, 0, "a")
	assertStringAtPosition(t, result, 1, "b")
}

func TestExpand_ExpansionFactor_2(t *testing.T) {
	result := expand([]string{"a", "b"}, 2)

	assertSize(t, result, 4)
	assertStringAtPosition(t, result, 0, "aa")
	assertStringAtPosition(t, result, 1, "ab")
	assertStringAtPosition(t, result, 2, "ba")
	assertStringAtPosition(t, result, 3, "bb")
}

func TestExpand_ExpansionFactor_3(t *testing.T) {
	result := expand([]string{"a", "b"}, 3)

	assertSize(t, result, 8)
	assertStringAtPosition(t, result, 0, "aaa")
	assertStringAtPosition(t, result, 1, "aab")
	assertStringAtPosition(t, result, 2, "aba")
	assertStringAtPosition(t, result, 3, "abb")
	assertStringAtPosition(t, result, 4, "baa")
	assertStringAtPosition(t, result, 5, "bab")
	assertStringAtPosition(t, result, 6, "bba")
	assertStringAtPosition(t, result, 7, "bbb")
}

func assertSize(t *testing.T, result []string, expectedSize int) {
	if len(result) != expectedSize {
		t.Errorf("Expanded charset should have a lenght of %d but has %d", expectedSize, len(result))
	}
}

func assertStringAtPosition(t *testing.T, result []string, position int, expected string) {
	if result[position] != expected {
		t.Errorf("Expanded charset should have %s at position %d, but was %s", expected, position, result[position])
	}
}
