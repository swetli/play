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

	t.Run("Say_hello_in_bg", func(t *testing.T) {
		got := hello("Св", "bulgarian")
		want := "Здрасти, Св"
		assertCorrectMsg(t, got, want)

	})

	t.Run("Say_hello_in_en", func(t *testing.T) {
		got := hello("Sw", "en_EN")
		want := "Hello, Sw"
		assertCorrectMsg(t, got, want)

	})

	t.Run("Say_hello_no_name", func(t *testing.T) {
		got := hello("", "")
		want := "Hello, World"
		assertCorrectMsg(t, got, want)

	})

}
