package array_slice

import "testing"

func TestSum(t *testing.T) {
	sumArray := [5]int{1, 2, 3, 4, 5}

	got := SumArray(sumArray)
	expected := 15

	if got != expected {
		t.Errorf("expected %d but got %d, given %v", expected, got, sumArray)
	}
}
