package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Nihar", "")
		want := "Hello, Nihar"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when empty string is supplied.", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)

	})
	t.Run("In Spanish.", func(t *testing.T) {
		got := Hello("Nihar", "Spanish")
		want := "Hola, Nihar"
		assertCorrectMessage(t, got, want)
	})
	t.Run("In French.", func(t *testing.T) {
		got := Hello("Nihar", "French")
		want := "Bonjour, Nihar"
		assertCorrectMessage(t, got, want)
	})
	t.Run("In Hindi.", func(t *testing.T) {
		got := Hello("Nihar", "Hindi")
		want := "Namaste, Nihar"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
