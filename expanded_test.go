package wotlib

import (
	"testing"
)

func TestTDSet(t *testing.T) {
	emptySet := NewExpandedThingDescriptionSet()
	if len(emptySet) != 0 {
		t.Fatalf("Set has invalid length")
	}

	expanded, err := FromBytes(testTDOne)
	if err != nil {
		t.Fatalf("Failed to build expanded td: %v", err)
	}

	set := NewExpandedThingDescriptionSet(expanded)
	if len(set) != 1 {
		t.Fatalf("Set has invalid length")
	}

	emptySet.Append(expanded)
	if len(emptySet) != 1 {
		t.Fatalf("Set now should have one more item")
	}

	emptySet.Remove(expanded.ID)
	if len(emptySet) != 0 {
		t.Fatalf("Set should be empty")
	}
}
