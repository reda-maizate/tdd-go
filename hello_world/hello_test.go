package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Réda")
	want := "Hello, Réda"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
