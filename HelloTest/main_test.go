package main

import "testing"

func TestMain(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Say Hello to people", func(t *testing.T) {
		got := Hello("Mervan", English)
		want := "Hello Mervan"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Say Hello Wold When Empty string supplied.", func(t *testing.T) {
		got := Hello("", English)
		want := "Hello World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Turkish Hello", func(t *testing.T) {
		got := Hello("Mervan", Turkish)
		want := "Merhaba Mervan"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Spanish Hello", func(t *testing.T) {
		got := Hello("Mervan", Spanish)
		want := "Hola Mervan"
		assertCorrectMessage(t, got, want)
	})

}
