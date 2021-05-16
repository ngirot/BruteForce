package conf

import (
	"testing"
)

func TestBestNumberOfGoRoutines_ShouldReturnAtLeastOne(t *testing.T) {
	var n = BestNumberOfGoRoutine()
	if n <= 0 {
		t.Errorf("The number of CPU should be at least 1, but was %d", n)
	}
}
