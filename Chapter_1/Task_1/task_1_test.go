package main

import (
	"testing"
)

func TestTriangleLeft(t *testing.T) {
	got := TriangleLeft(4, "*")
	want := "*\n**\n***\n****\n"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestTriangleCenterTopDown(t *testing.T) {
	got := TriangleCenterTopDown(4, "*")
	want := " * * * * \n  * * * \n   * * \n    * \n"

	if got != want {
		t.Errorf("got %q, \n wanted %q", got, want)
	}
}

func TestTriangleCenter(t *testing.T) {
	got := TriangleCenter(4, "*")
	want := "    * \n   * * \n  * * * \n * * * * \n"

	if got != want {
		t.Errorf("got %q, \n wanted %q", got, want)
	}
}
