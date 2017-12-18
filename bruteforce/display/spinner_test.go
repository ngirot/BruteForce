

package display

import (
	"testing"
)

func TestSpin_ShouldIncrementValue(t *testing.T) {
	var s = NewCustomSpinner([]string{"0", "1", "2", "3"})

	testSpin(t, s, "0")
	testSpin(t, s, "1")
	testSpin(t, s, "2")
	testSpin(t, s, "3")
}

func TestSpin_ShouldLoop(t *testing.T) {
	var s = NewCustomSpinner([]string{"0", "1"})

	testSpin(t, s, "0")
	testSpin(t, s, "1")
	testSpin(t, s, "0")
}

func testSpin(t *testing.T, s Spinner, expected string) {
	var actual = s.Spin()
	if actual != expected {
		t.Errorf("The value of the spinner should be %s but was %s", expected, actual)
	}
}