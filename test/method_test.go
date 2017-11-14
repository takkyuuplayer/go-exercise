package test

import (
	"image/color"
	"math"
	"reflect"
	"testing"
)

type Point struct{ X, Y float64 }
type ColoredPoint struct {
	Point
	color.RGBA
}

func (p *Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(k float64) {
	p.X *= k
	p.Y *= k
}

func TestMethodWithCopiedValue(t *testing.T) {
	p1 := Point{1, 1}
	p2 := Point{4, 5}

	if p1.Distance(p2) != 5 {
		t.Errorf(`p1.Distance(p2) = %#v, want %#v`, p1.Distance(p2), 5)
	}
}

func TestMethodWithPointer(t *testing.T) {
	p := Point{1, 1}
	p.ScaleBy(3)

	if !reflect.DeepEqual(p, Point{3, 3}) {
		t.Errorf(`p = %#v, want %#v`, p, Point{3, 3})
	}
}

func TestHasARelation(t *testing.T) {
	p1 := ColoredPoint{Point{1, 1}, color.RGBA{255, 0, 0, 255}}
	p2 := ColoredPoint{Point{4, 5}, color.RGBA{255, 0, 0, 255}}

	if p1.Distance(p2.Point) != 5 {
		t.Errorf(`p1.Distance(p2) = %#v, want %#v`, p1.Distance(p2.Point), 5)
	}
}

func TestMethodExpression(t *testing.T) {
	p1 := Point{1, 1}
	p2 := Point{4, 5}

	scaleP1 := p1.ScaleBy

	scaleP1(4)

	if !reflect.DeepEqual(p1, Point{4, 4}) {
		t.Errorf(`p1 = %#v, want %#v`, p1, Point{4, 4})
	}

	distance := (*Point).Distance

	if distance(&p1, p2) != 1 {
		t.Errorf(`distance(&p1, p2) = %#v, want %#v`, distance(&p1, p2), 1)
	}
}
