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
	checkArea := func(t testing.TB, shape Shape, expected float64) {
		t.Helper()
		got := shape.Area()
		if got != expected {
			t.Errorf("expected %g but got %g", expected, got)
		}
	}

	t.Run("area of a rectangle", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		expected := float64(72)

		checkArea(t, rectangle, expected)
	})

	t.Run("area of a circle", func(t *testing.T) {
		circle := Circle{10}
		expected := 314.1592653589793

		checkArea(t, circle, expected)
	})
}
