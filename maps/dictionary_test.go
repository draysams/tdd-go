package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}
	t.Run("known word", func(t *testing.T) {

		got, _ := dictionary.Search("test")
		want := "this is a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := ErrorNotFound

		assertError(t, err, want)
	})
}

func assertStrings(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q, given %q", got, want, "test")
	}

}

func assertError(t testing.TB, err, want error) {
	t.Helper()
	if err == nil {
		t.Fatalf("expected error %q but didn't receive it", want)
	}

	if err != want {
		t.Errorf("got %q want %q", err, want)
	}

}
