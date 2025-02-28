package array_slice

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := SumArray(numbers)
		expected := 6

		if got != expected {
			t.Errorf("expected %d but got %d given, %v", expected, got, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{3, 4})
	expected := []int{3, 7}

	if !slices.Equal(got, expected) {
		t.Errorf("expected %v but got %v given %v", expected, got, [][]int{{1, 2}, []int{3, 4}})
	}
}
