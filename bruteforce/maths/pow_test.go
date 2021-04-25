package maths

import "testing"

func TestPowInt_withZeroExponent(t *testing.T) {
	check(5, 0, 1, t)
}

func TestPowInt_withOneExponent(t *testing.T) {
	check(5, 1, 5, t)
}

func TestPowInt_WithZeroBase(t *testing.T) {
	check(0, 100, 0, t)
}

func TestPowInt_StandardCase(t *testing.T) {
	check(2, 10, 1024, t)
}

func check(base int, exponent int, expectedResult int, t *testing.T) {
	var pow = PowInt(base, exponent)
	if pow != expectedResult {
		t.Errorf("%d^%d should be %d but was %d", base, exponent, expectedResult, pow)
	}
}
