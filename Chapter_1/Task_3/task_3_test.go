package main

import (
	"testing"
)

func TestCalculateChildren(t *testing.T) {
	got := CalculateChildren(15, 20)
	want := 42

	if got != want {
		t.Errorf("got %q, \n wanted %q", got, want)
	}
}
