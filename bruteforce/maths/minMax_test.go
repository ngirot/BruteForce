package maths

import "testing"

func TestMaxInt_WithFirstParameterBigger(t *testing.T) {
	checkMax(5, 1, 5, t)
}

func TestMaxInt_WithFirstParameterLower(t *testing.T) {
	checkMax(3, 4, 4, t)
}

func TestMaxInt_WithEqualParameters(t *testing.T) {
	checkMax(2, 2, 2, t)
}

func TestMinInt_WithFirstParameterBigger(t *testing.T) {
	checkMin(5, 1, 1, t)
}

func TestMinInt_WithFirstParameterLower(t *testing.T) {
	checkMin(3, 4, 3, t)
}

func TestMinInt_WithEqualParameters(t *testing.T) {
	checkMin(2, 2, 2, t)
}

func checkMax(firstParameter int, secondParameter int, expectedResult int, t *testing.T) {
	var max = MaxInt(firstParameter, secondParameter)
	if max != expectedResult {
		t.Errorf("max(%d,%d) should be %d but was %d", firstParameter, secondParameter, expectedResult, max)
	}
}

func checkMin(firstParameter int, secondParameter int, expectedResult int, t *testing.T) {
	var max = MinInt(firstParameter, secondParameter)
	if max != expectedResult {
		t.Errorf("min(%d,%d) should be %d but was %d", firstParameter, secondParameter, expectedResult, max)
	}
}
