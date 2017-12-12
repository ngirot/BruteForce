

package display

import (
	"testing"
	"fmt"
)


func TestSpinShouldLoop(t *testing.T) {
	var s = NewCustomSpinner([]string{"0", "1"})

	if s.Spin() != "0" {
		t.Fail()
	}
	if s.Spin() != "1" {
		t.Fail()
	}
	if s.Spin() != "0" {
		t.Fail()
	}
}

func TestSpinShouldLoop2(t *testing.T) {
	var s = NewDefaultSpinner()

	fmt.Printf("%s\n", s.Spin())
	fmt.Printf("%s\n", s.Spin())
	fmt.Printf("%s\n", s.Spin())
	fmt.Printf("%s\n", s.Spin())
	fmt.Printf("%s\n", s.Spin())
	fmt.Printf("%s\n", s.Spin())
}