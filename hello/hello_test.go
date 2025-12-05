package main

import "testing"

func TestHelloWorld(t *testing.T) {
	got := HelloWorld()
	target := "Hello, world!"

	if got != target {
		t.Errorf("got %q want %q", got, target)
	}
}

func TestHello(t *testing.T) {
	got := Hello("Leonardo Da Vinci")
	target := "Hello, Leonardo Da Vinci!"

	if got != target {
		t.Errorf("got %q want %q", got, target)
	}
}
