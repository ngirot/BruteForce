package hashs

import (
	"testing"
)

func TestAllHasherTypesReturnAtLeastOneElement(t *testing.T) {
	var types = AllHasherTypes()

	if len(types)  <= 0 {
		t.Fail()
	}
}

func TestAllHasherContainsAtLeastSha256(t *testing.T) {
	var types = AllHasherTypes()

	var found = false
	for _,ty := range types {
		if ty == "sha256" {
			found = true
		}
	}

	if !found {
		t.Fail()
	}
}

func TestAllhasherNeverReturnEmptyType(t *testing.T) {
	var types = AllHasherTypes()

	for _,ty := range types {
		if ty == "" {
			t.Fail()
		}
	}

}
