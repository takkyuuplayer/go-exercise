package test

import "testing"

func TestConstant(t *testing.T) {
	const (
		a = 1
		b
		c = 2
		d
	)

	if a != 1 {
		t.Errorf(`a = %#v, want %#v`, a, 1)
	}

	if b != 1 {
		t.Errorf(`b = %#v, want %#v`, b, 1)
	}

	if c != 2 {
		t.Errorf(`c = %#v, want %#v`, c, 2)
	}

	if d != 2 {
		t.Errorf(`d = %#v, want %#v`, d, 2)
	}
}

func TestAutoIncrement(t *testing.T) {
	const (
		zero = iota
		one
		two
		three
	)

	if zero != 0 {
		t.Errorf(`zero = %#v, want %#v`, zero, 0)
	}

	if three != 3 {
		t.Errorf(`three = %#v, want %#v`, three, 3)
	}
}

func TestAutoBitShift(t *testing.T) {
	const (
		one = 1 << iota
		two
		four
		eight
	)

	if one != 1 {
		t.Errorf(`one = %#v, want %#v`, one, 1)
	}

	if eight != 8 {
		t.Errorf(`eight = %#v, want %#v`, eight, 8)
	}
}
