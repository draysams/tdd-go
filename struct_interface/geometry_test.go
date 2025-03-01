package struct_interface

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.00, 10.00}
	got := Perimeter(rectangle)
	expected := float64(40)

	if got != expected {
		t.Errorf("expected %.2f but got %.2f", expected, got)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{12, 6}
	got := Area(rectangle)
	expected := float64(72)

	if got != expected {
		t.Errorf("expected %.2f but got %.2f", expected, got)
	}
}
