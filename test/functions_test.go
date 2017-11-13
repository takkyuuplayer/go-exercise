package test

import "testing"

func calc(a, b int) (x, y int) {
	x, y = a+b, a-b
	return
}

func incrementAdd(a, b int) int {
	a++
	b++

	return a + b
}

func argumentMap(m map[string]int) map[string]int {
	m["test"] = 1

	return m
}

func TestWhenArgumentsAreValue(t *testing.T) {
	a, b := 1, 2
	x := incrementAdd(a, b)

	if a != 1 {
		t.Errorf(`a = %#v, want %#v`, a, 1)
	}

	if b != 2 {
		t.Errorf(`b = %#v, want %#v`, b, 2)
	}

	if x != 5 {
		t.Errorf(`x = %#v, want %#v`, x, 5)
	}
}

func TestWhenArgumentsAreReference(t *testing.T) {
	m := map[string]int{
		"test": 0,
		"foo":  1,
	}

	ret := argumentMap(m)

	if m["test"] != 1 {
		t.Errorf(`m["test"] = %#v, want %#v`, m["test"], 1)
	}

	if ret["test"] != 1 {
		t.Errorf(`ret["test"] = %#v, want %#v`, ret["test"], 1)
	}
}

func TestNamedReturnValue(t *testing.T) {
	x, y := calc(2, 1)

	if x != 3 {
		t.Errorf(`x = %#v, want %#v`, x, 3)
	}

	if y != 1 {
		t.Errorf(`y = %#v, want %#v`, y, 1)
	}
}

func multiValues() (int, int) {
	return 1, 2
}

func TestMultiValues(t *testing.T) {
	x := incrementAdd(multiValues())

	if x != 5 {
		t.Errorf(`x = %#v, want %#v`, x, 5)
	}
}

func TestFunctionValues(t *testing.T) {
	var x func(int, int) int

	if x != nil {
		t.Errorf(`function zero value is nil`)
	}

	x = incrementAdd

	if x(1, 2) != 5 {
		t.Errorf(`x(1,2) = %#v, want %#v`, x(1, 2), 5)
	}
}

func TestCapturingIterationVariables(t *testing.T) {
	var f []func() int

	for i := 0; i < 10; i++ {
		f = append(f, func() int { return i })
	}

	// i is updated
	if f[0]() != 10 {
		t.Errorf(`f[0]() = %#v, want %#v`, f[0](), 10)
	}

	if f[9]() != 10 {
		t.Errorf(`f[9]() = %#v, want %#v`, f[9](), 10)
	}
}
