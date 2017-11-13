package test

import "testing"

var x = 0

func incrementX() {
	x++
}

func deferIncrementX() func() {
	return func() { x++ }
}

func TestDefer(t *testing.T) {
	defer incrementX() // 4

	if x != 0 {
		t.Errorf(`x = %#v, want %#v`, x, 0)
	}

	incrementX()

	if x != 1 {
		t.Errorf(`x = %#v, want %#v`, x, 1)
	}

	defer deferIncrementX()() // 3

	if x != 1 {
		t.Errorf(`x = %#v, want %#v`, x, 1)
	}

	defer func() { // 2
		if x != 2 {
			t.Errorf(`x = %#v, want %#v`, x, 2)
		}
	}()

	defer incrementX() // 1
}
