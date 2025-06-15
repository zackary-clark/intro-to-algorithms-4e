package ch2

import (
	"slices"
	"testing"
)

func Test_InsertionSort(t *testing.T) {
	input := []int{3, 2, 1, 4, 8, 6, 5, 7, 9}
	result := InsertionSort(input)
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if !slices.Equal(result, expected) {
		t.Error("input", input, "not sorted: expected ", expected, "got: ", result)
	}
}
