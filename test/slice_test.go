package test

import (
	"reflect"
	"testing"
)

func TestSlice(t *testing.T) {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	b := a[2:5]

	if !reflect.DeepEqual(b, []int{3, 4, 5}) {
		t.Errorf(`b = %#v, want %#v`, b, []int{3, 4, 5})
	}

	if len(b) != 3 {
		t.Errorf(`len(b) = %#v, want %#v`, len(b), 3)
	}

	if cap(b) != 8 {
		t.Errorf(`cap(b) = %#v, want %#v`, cap(b), 8)
	}
}

func TestZeroValue(t *testing.T) {
	var a []int

	if a != nil {
		t.Errorf(`a = %#v, want %#v`, a, nil)
	}

	if len(a) != 0 {
		t.Errorf(`len(a) = %#v, want %#v`, len(a), 0)
	}

	a = nil

	if len(a) != 0 {
		t.Errorf(`len(a) = %#v, want %#v`, len(a), 0)
	}
}

func TestSliceIsReference(t *testing.T) {
	a := [...]int{1, 2, 3}

	b := a[1:]
	b[0] = 4

	if !reflect.DeepEqual(a, [...]int{1, 4, 3}) {
		t.Errorf(`a = %#v, want %#v`, a, [...]int{1, 4, 3})
	}

	if !reflect.DeepEqual(b, []int{4, 3}) {
		t.Errorf(`b = %#v, want %#v`, b, []int{4, 3})
	}
}

func TestAppend(t *testing.T) {
	origin := [...]int{0, 1, 2, 3}

	a := origin[:2]
	b := append(a, 1)

	if len(origin) != 4 {
		t.Errorf(`len(origin) = %#v, want %#v`, len(origin), 4)
	}

	if !reflect.DeepEqual(origin, [...]int{0, 1, 1, 3}) { // underlying array is overwritten
		t.Errorf(`origin = %#v, want %#v`, origin, [...]int{0, 1, 1, 3})
	}

	if len(a) != 2 {
		t.Errorf(`len(a) = %#v, want %#v`, len(a), 2)
	}

	if cap(a) != 4 {
		t.Errorf(`cap(a) = %#v, want %#v`, cap(a), 4)
	}

	if !reflect.DeepEqual(a, []int{0, 1}) {
		t.Errorf(`a = %#v, want %#v`, a, []int{0, 1})
	}

	if len(b) != 3 {
		t.Errorf(`len(b) = %#v, want %#v`, len(b), 3)
	}

	if cap(b) != 4 {
		t.Errorf(`cap(b) = %#v, want %#v`, cap(b), 4)
	}

	if !reflect.DeepEqual(b, []int{0, 1, 1}) {
		t.Errorf(`b = %#v, want %#v`, b, []int{0, 1, 1})
	}
}

func TestCopy(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}

	if !reflect.DeepEqual(a, []int{1, 2, 3}) {
		t.Errorf(`a = %#v, want %#v`, a, []int{1, 2, 3})
	}

	c := copy(a, b)

	if !reflect.DeepEqual(a, []int{4, 5, 6}) {
		t.Errorf(`a = %#v, want %#v`, a, []int{4, 5, 6})
	}

	if !reflect.DeepEqual(b, []int{4, 5, 6}) {
		t.Errorf(`b = %#v, want %#v`, b, []int{4, 5, 6})
	}

	if c != 3 {
		t.Errorf(`c = %#v, want %#v`, c, 3)
	}
}
