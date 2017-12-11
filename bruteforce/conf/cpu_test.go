package conf

import (
	"testing"
)

func TestBestNumberOfGoRoutinesShouldReturnAtLeastOne(t *testing.T) {
	var n = BestNumberOfGoRoutine()
	if n <= 0 {
		t.Fail()
	}
}