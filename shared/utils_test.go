package utils

import (
	"testing"
)

func TestFindIndex(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}

	// Test case 1: Existing value
	index := FindIndex(slice, 3)
	if index != 2 {
		t.Errorf("Expected index %d, but got %d", 2, index)
	}

	// Test case 2: Non-existing value
	index = FindIndex(slice, 9)
	if index != -1 {
		t.Errorf("Expected index %d, but got %d", -1, index)
	}

	sliceStr := []string{"a", "b", "c", "d", "e"}

	// Test case 1: Existing value
	index = FindIndex(sliceStr, "c")
	if index != 2 {
		t.Errorf("Expected index %d, but got %d", 2, index)
	}

	// Test case 2: Non-existing value
	index = FindIndex(sliceStr, "f")
	if index != -1 {
		t.Errorf("Expected index %d, but got %d", -1, index)
	}

}

func TestRandomChance(t *testing.T) {

	// Test case 1: 0% chance
	if got := RandomChance(0); got != true {
		t.Errorf("Expected true but got %v", got)
	}

	// Test case 2: 100% chance
	if got := RandomChance(1); got != false {
		t.Errorf("Expected false but got %v", got)
	}
}
