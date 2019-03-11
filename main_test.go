package main

import "testing"

func TestPort(t *testing.T) {
	got := port()
	want := "8080"

	if want != got {
		t.Errorf("got %s, but want %s", got, want)
	}
}
