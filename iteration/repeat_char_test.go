package iteration

import "testing"

func TestRepeatChar(t *testing.T) {
	got := repeatChar("a")
	expected := "aaaaa"

	if got != expected {
		t.Errorf("expected %q but got %q", expected, got)
	}
}

func BenchmarkRepeatChar(b *testing.B) {
	for range b.N {
		repeatChar("a")
	}
}
