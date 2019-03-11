package main

import (
	"os"
	"testing"
)

func TestDefaultPort(t *testing.T) {
	got := port()
	want := "8080"

	if want != got {
		t.Errorf("got %s, but want %s", got, want)
	}
}

func TestCustomPort(t *testing.T) {
	os.Setenv("PORT", "9090")
	got := port()
	want := "9090"

	if want != got {
		t.Errorf("got %s, but want %s", got, want)
	}
}
