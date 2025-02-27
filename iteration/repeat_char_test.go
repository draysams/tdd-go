package iteration

import (
	"fmt"
	"testing"
)

func TestRepeatChar(t *testing.T) {
	got := RepeatChar("a")
	expected := "aaaaa"

	if got != expected {
		t.Errorf("expected %q but got %q", expected, got)
	}
}
func ExampleRepeatChar() {
	repeatedString := RepeatChar("a")
	fmt.Println(repeatedString)
	// output aaaaa
}

func BenchmarkRepeatChar(b *testing.B) {
	for range b.N {
		RepeatChar("a")
	}
}
