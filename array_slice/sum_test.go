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
	given := [][]int{{1, 2}, {3, 4}}

	checkSums(t, got, expected, given)
}

func TestSumAllTails(t *testing.T) {
	t.Run("sum a slice's tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{2, 3, 4})
		expected := []int{5, 7}
		given := [][]int{{1, 2, 3}, {2, 3, 4}}

		checkSums(t, got, expected, given)

	})

	t.Run("sum an empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		expected := []int{0, 9}
		given := [][]int{{}, {3, 4, 5}}

		checkSums(t, got, expected, given)
	})
}

func checkSums(t testing.TB, got []int, expected []int, given [][]int) {
	t.Helper()
	if !slices.Equal(got, expected) {
		t.Errorf("expected %v gpt %v given %v", expected, got, given)
	}
}
