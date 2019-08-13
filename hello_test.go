package main

import (
	"log"
	"testing"
)

func TestHello(t *testing.T) {
	log.Printf("### Testing function hello(): ")

	assertCorrectMsg := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("Got %q, wanted %q", got, want)
		}
	}

	t.Run("Saying hello to people", func(t *testing.T) {
		got := hello("Sw")
		want := "Hello, Sw!"
		assertCorrectMsg(t, got, want)

	})

	t.Run("Say hello w/o name", func(t *testing.T) {
		got := hello("")
		want := "Hello, World@!"
		assertCorrectMsg(t, got, want)

	})

}
