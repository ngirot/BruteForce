package hashs

import (
	"testing"
)

func TestAllHasherTypes_ReturnAtLeastOneElement(t *testing.T) {
	var types = AllHasherTypes()

	if len(types) <= 0 {
		t.Error("All hasher list should have at least one element")
	}
}

func TestAllHasher_ContainsAtLeastSha256(t *testing.T) {
	var types = AllHasherTypes()

	var found = false
	for _, ty := range types {
		if ty == "sha256" {
			found = true
		}
	}

	if !found {
		t.Error("The hasher list should contains at least SHA256")
	}
}

func TestAllhasher_NeverReturnEmptyType(t *testing.T) {
	var types = AllHasherTypes()

	for _, ty := range types {
		if ty == "" {
			t.Error("A hasher type should never be an empty string")
		}
	}

}
