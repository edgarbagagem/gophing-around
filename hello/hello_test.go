package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to Leonardo", func(t *testing.T) {
		got := Hello("Leonardo Da Vinci")
		target := "Hello, Leonardo Da Vinci!"

		if got != target {
			t.Errorf("got %q want %q", got, target)
		}
	})

	t.Run("saying hello to the world", func(t *testing.T) {
		got := Hello("")
		target := "Hello, world!"

		if got != target {
			t.Errorf("got %q want %q", got, target)
		}
	})
}
