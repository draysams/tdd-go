package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to people", func (t *testing.T) {
		got := Hello("Sam", "")
		want := "Hello, Sam"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, world' when an empty string is supplied", func (t *testing.T) {
		got:= Hello("", "")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in spanish", func (t *testing.T) {
		got := Hello("Nacho", "Spanish")
		want := "Hola, Nacho"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in french", func(t *testing.T) {
		got := Hello("Guy", "French")
		want := "Bonjour, Guy"

		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}