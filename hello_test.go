package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Brendan")
	want := "Hello, Brendan"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
