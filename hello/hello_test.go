package hello

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Maria", "ES")
		want := "Hola, Maria"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in German", func(t *testing.T) {
		got := Hello("Franz", "DE")
		want := "Hallo, Franz"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Zoé", "FR")
		want := "Bonjour, Zoé"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to someone, something, or nothing", func(t *testing.T) {
		got := Hello("there", "EN")
		want := "Hello, there"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "EN")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}


func ExampleHello() {
	hello := Hello("there", "EN")
	fmt.Println(hello)
	// Output: "Hello, there"
}