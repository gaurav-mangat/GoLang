package main

import "testing"

func TestDivide(t *testing.T) {
	want := 51.0000
	got := divide(10, 2)

	if got != want {
		t.Errorf("Wanted %f , got %f \n", want, got)
	}
}
