package maps

import "testing"

var dictionaryWord = "test"
var dictionaryDefinition = "this is a test"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{dictionaryWord: dictionaryDefinition}
	t.Run("known word", func(t *testing.T) {
		want := "this is a test"
		assertDefinition(t, dictionary, "test", want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := ErrorNotFound
		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add(dictionaryWord, dictionaryDefinition)
		want := dictionaryDefinition
		assertDefinition(t, dictionary, dictionaryWord, want)
	})

	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{dictionaryWord: dictionaryDefinition}
		err := dictionary.Add(dictionaryWord, "new word")
		assertError(t, err, ErrorWordExists)
		assertDefinition(t, dictionary, dictionaryWord, dictionaryDefinition)
	})

}

func assertDefinition(t testing.TB, dictionary Dictionary, word, want string) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("didn't expect an error but received", want)
	}

	if got != want {
		t.Fatal("should find added word ", err)
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
