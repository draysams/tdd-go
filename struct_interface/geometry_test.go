package struct_interface

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10, 10)
	expected := float64(40)

	if got != expected {
		t.Errorf("expected %.2f but got %.2f", expected, got)
	}
}

func TestArea(t *testing.T) {
	got := Area(12, 6)
	expected := float64(72)

	if got != expected {
		t.Errorf("expected %.2f but got %.2f", expected, got)
	}
}
