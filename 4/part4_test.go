package main

import "testing"

func TestPuzzleWithTestInput(t *testing.T) {
	got := part1("test.txt")
	want := 18
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestPuzzleWithTest2Input(t *testing.T) {
	got := part1("test2.txt")
	want := 0
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestPart2(t *testing.T	){
	got := part2("test.txt")
	want := 9
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}