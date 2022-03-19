package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("there")
	want := "Hello, there"

	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}