package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Nihar")
	want := "Hello, Nihar"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
