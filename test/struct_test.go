package test

import (
	"testing"
)

type Employee struct {
	Name  string
	Phone string
}

func TestPointer(t *testing.T) {
	t.Parallel()
	emp := Employee{
		Name:  "Takkyuu",
		Phone: "11122223333",
	}

	if emp.Name != "Takkyuu" {
		t.Errorf(`emp.Name = %#v, want %#v`, emp.Name, "Takkyuu")
	}

	p := &emp

	if (*p).Name != "Takkyuu" {
		t.Errorf(`(*p).Name = %#v, want %#v`, (*p).Name, "Takkyuu")
	}

	if p.Name != "Takkyuu" {
		t.Errorf(`p.Name = %#v, want %#v`, p.Name, "Takkyuu")
	}
}

func TestCompare(t *testing.T) {
	t.Parallel()
	type point struct{ x, y int }

	p1 := point{1, 2}
	p2 := point{2, 1}

	if p1 == p2 {
		t.Errorf(`p1 = %#v, do NOT want %#v`, p1, p2)
	}

	p3 := point{1, 2}

	if p1 != p3 {
		t.Errorf(`p1 = %#v, want %#v`, p1, p3)
	}
}

func TestEmbedding(t *testing.T) {
	t.Parallel()
	type Point struct {
		X, Y int
	}
	type Circle struct {
		Point
		Radis int
	}
	type Wheel struct {
		Circle
		Spokes int
	}

	p := Point{1, 2}
	c := Circle{p, 5}
	w := Wheel{c, 20}

	if w.X != w.Circle.Point.X {
		t.Errorf(`w.X = %#v, want %#v`, w.X, w.Circle.Point.X)
	}

	w.X = 100

	if w.Circle.Point.X != 100 {
		t.Errorf(`w.Circle.Point.X = %#v, want %#v`, w.Circle.Point.X, 100)
	}
}
