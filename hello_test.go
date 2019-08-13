package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := hello("Swetli")
	want := "Hello, Swetli!"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
