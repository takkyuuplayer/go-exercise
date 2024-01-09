package tour

import (
	"fmt"
	"math"
	"testing"
)

func TestPointer(t *testing.T) {
	t.Parallel()
	var p *int

	if p != nil {
		t.Fatal("point zero value is nil: " + fmt.Sprintf("%T(%v)", p, p))
	}

	i, j := 42, 2701

	p = &i
	*p = 21

	if i != 21 {
		t.Fatal("i is updated: " + fmt.Sprintf("%T(%v)", i, i))
	}

	p = &j
	*p = *p / 37

	if j != 73 {
		t.Fatal("j is updated: " + fmt.Sprintf("%T(%v)", j, j))
	}
}

func TestStruct(t *testing.T) {
	t.Parallel()
	v := Vertex{1, 2}

	v.X = 4

	if v.X != 4 {
		t.Fatal("v.X is updated: " + fmt.Sprintf("%T(%v)", v.X, v.X))
	}

	if v.Y != 2 {
		t.Fatal("v.Y is 2: " + fmt.Sprintf("%T(%v)", v.Y, v.Y))
	}
}

func TestPointerOfStruct(t *testing.T) {
	t.Parallel()
	v := Vertex{1, 2}

	p := &v

	if fmt.Sprintf("%T", p) != "*tour.Vertex" {
		t.Fatal("type is Vertex's pointer: " + fmt.Sprintf("%T(%v)", p, p))
	}

	if (*p).X != 1 {
		t.Fatal("(*pointer).X to show v.X : " + fmt.Sprintf("%T(%v)", (*p).X, (*p).X))
	}

	p.X = 1e9

	if p.X != 1e9 {
		t.Fatal("shorthand notation: " + fmt.Sprintf("%T(%v)", p.X, p.X))
	}
}

func TestStructLiterals(t *testing.T) {
	t.Parallel()
	v2 := Vertex{X: 1}

	if v2.X != 1 {
		t.Fatal("v2.X is initialized: " + fmt.Sprintf("%T(%v)", v2.X, v2.X))
	}

	if v2.Y != 0 {
		t.Fatal("v2.Y is zero value: " + fmt.Sprintf("%T(%v)", v2.Y, v2.Y))
	}
}

func TestArray(t *testing.T) {
	t.Parallel()
	var a [2]string

	a[0] = "Hello"
	a[1] = "World"

	if a[1] != "World" {
		t.Fatal("array definition: " + fmt.Sprintf("%T(%v)", a[1], a[1]))
	}
}

func TestSlices(t *testing.T) {
	t.Parallel()
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]

	if s[0] != 3 {
		t.Fatal("Slice is variable length arrray: " + fmt.Sprintf("%T(%v)", s, s))
	}

	if len(s) != 3 {
		t.Fatal("0,1,2 is available: " + fmt.Sprintf("%T(%v)", len(s), len(s)))
	}
}

func TestSliceLiterals(t *testing.T) {
	t.Parallel()
	primes := []int{2, 3, 5, 7, 11, 13}

	if primes[1] != 3 {
		t.Fatal("Slice can be defined with literal: " + fmt.Sprintf("%T(%v)", primes[1], primes[1]))
	}
}

func TestSliceNil(t *testing.T) {
	t.Parallel()

	var s []int

	if s != nil {
		t.Fatal("Slice's 0 value is nil: " + fmt.Sprintf("%T(%v)", s, s))
	}
}

func TestSliceWithMake(t *testing.T) {
	t.Parallel()
	a := make([]int, 5)

	if a[0] != 0 {
		t.Fatal("define dynamic size array: " + fmt.Sprintf("%T(%v)", a[0], a[0]))
	}
	if len(a) != 5 {
		t.Fatal("length = 5: " + fmt.Sprintf("%T(%v)", len(a), len(a)))
	}
	if cap(a) != 5 {
		t.Fatal("capasity = 5: " + fmt.Sprintf("%T(%v)", cap(a), cap(a)))
	}

	b := make([]int, 0, 5)

	if len(b) != 0 {
		t.Fatal("length = 0: " + fmt.Sprintf("%T(%v)", len(b), len(b)))
	}
	if cap(b) != 5 {
		t.Fatal("capasity = 5: " + fmt.Sprintf("%T(%v)", cap(b), cap(b)))
	}

}

func TestAppend(t *testing.T) {
	t.Parallel()
	var s []int

	if len(s) != 0 {
		t.Fatal("nil: " + fmt.Sprintf("%T(%v)", len(s), len(s)))
	}

	s = append(s, 0)

	if len(s) != 1 {
		t.Fatal("appended: " + fmt.Sprintf("%T(%v)", len(s), len(s)))
	}

	s = append(s, 1)

	if len(s) != 2 {
		t.Fatal("appended: " + fmt.Sprintf("%T(%v)", len(s), len(s)))
	}

	a := append(s, 2, 3, 4)

	if len(s) != 2 {
		t.Fatal("append is not bang function: " + fmt.Sprintf("%T(%v)", len(s), len(s)))
	}
	if len(a) != 5 {
		t.Fatal("appended: " + fmt.Sprintf("%T(%v)", len(s), len(s)))
	}
}

func TestRange(t *testing.T) {
	t.Parallel()
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}

	for idx, val := range pow {
		if idx != 0 {
			t.Fatal("1st index should be index: " + fmt.Sprintf("%T(%v)", idx, idx))
		}
		if val != 1 {
			t.Fatal("1st val should be value: " + fmt.Sprintf("%T(%v)", val, val))
		}
		if true {
			break
		}
	}

	for _, val := range pow {
		if val != 1 {
			t.Fatal("index is ignored: " + fmt.Sprintf("%T(%v)", val, val))
		}
		if true {
			break
		}
	}
}

func TestMap(t *testing.T) {
	t.Parallel()
	var m map[string]int

	if m != nil {
		t.Fatal("map's 0 value is nil: " + fmt.Sprintf("%T(%v)", m, m))
	}

	m = make(map[string]int)
	m["one"] = 1

	if m["one"] != 1 {
		t.Fatal("map can be created by make: " + fmt.Sprintf("%T(%v)", m["one"], m["one"]))
	}
}

func TestMapLiterals(t *testing.T) {
	t.Parallel()
	m := map[string]int{
		"one": 1,
		"two": 2,
	}

	if m["one"] != 1 {
		t.Fatal("map can be defiend with literal: " + fmt.Sprintf("%T(%v)", m["one"], m["one"]))
	}
}

func TestMapWithInferring(t *testing.T) {
	t.Parallel()
	m := map[string]Vertex{
		"O": {0, 0},
		"A": {1, 0},
		"B": {0, 1},
	}

	if m["A"].X != 1 {
		t.Fatal("ommit Vertex: " + fmt.Sprintf("%T(%v)", m["O"].X, m["O"].X))
	}
}

func TestMapOperation(t *testing.T) {
	t.Parallel()
	m := map[string]int{
		"one": 1,
		"two": 2,
	}

	elem, ok := m["one"]

	if elem != 1 {
		t.Fatal("one = 1: " + fmt.Sprintf("%T(%v)", elem, elem))
	}
	if ok != true {
		t.Fatal("defined: " + fmt.Sprintf("%T(%v)", ok, ok))
	}

	elem, ok = m["three"]

	if elem != 0 {
		t.Fatal("int's 0 value: " + fmt.Sprintf("%T(%v)", elem, elem))
	}
	if ok != false {
		t.Fatal("not defined: " + fmt.Sprintf("%T(%v)", ok, ok))
	}
}

func TestFunctionValues(t *testing.T) {
	t.Parallel()
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	if int(hypot(5, 12)) != 13 {
		t.Fatal("assiging function to variable: " + fmt.Sprintf("%T(%v)", int(hypot(5, 12)), int(hypot(5, 12))))
	}

}
