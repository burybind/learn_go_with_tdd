package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessasge := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Brendan", 1)
		want := "Hello, Brendan"
		assertCorrectMessasge(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", english)
		want := "Hello, World"
		assertCorrectMessasge(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Carlos", spanish)
		want := "Hola, Carlos"
		assertCorrectMessasge(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Juan", french)
		want := "Bonjour, Juan"
		assertCorrectMessasge(t, got, want)
	})

	t.Run("in Italian", func(t *testing.T) {
		got := Hello("Juan", 99)
		want := "Ciao, Juan"
		assertCorrectMessasge(t, got, want)
	})
}
