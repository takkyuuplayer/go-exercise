package tour

import (
	"fmt"
	"math"
	"testing"
)

func (v *Vertex) Abs() float64 {
	if v == nil {
		return 0
	}

	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func TestMethodsForStruct(t *testing.T) {
	v := Vertex{3, 4}

	if v.Abs() != 5 {
		t.Fatal("method is called: " + fmt.Sprintf("%T(%v)", v.Abs(), v.Abs()))
	}
}

func TestPointerReceivers(t *testing.T) {
	v := Vertex{3, 4}

	v.Scale(10)

	if v.X != 30 {
		t.Fatal("Changed: " + fmt.Sprintf("%T(%v)", v.X, v.X))
	}
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func TestMethodsForType(t *testing.T) {
	f := MyFloat(-2)

	if f.Abs() != 2 {
		t.Fatal("should calculate absolute value: " + fmt.Sprintf("%T(%v)", f.Abs(), f.Abs()))
	}

}

type Abser interface {
	Abs() float64
}

func TestInterface(t *testing.T) {
	var a Abser

	a = MyFloat(-2)

	if a.Abs() != 2 {
		t.Fatal("MyFloat has Abs method: " + fmt.Sprintf("%T(%v)", a.Abs(), a.Abs()))
	}

	a = &Vertex{3, 4}

	if a.Abs() != 5 {
		t.Fatal("*Vertex has Abs method: " + fmt.Sprintf("%T(%v)", a.Abs(), a.Abs()))
	}
}

func TestInterfaceWithNil(t *testing.T) {
	var i Abser

	if fmt.Sprintf("%T(%v)", i, i) != "<nil>(<nil>)" {
		t.Fatal("nil interface has no value and type: " + fmt.Sprintf("%T(%v)", fmt.Sprintf("%T(%v)", i, i), fmt.Sprintf("%T(%v)", i, i)))
	}

	var ty *Vertex

	i = ty

	if i.Abs() != 0 {
		t.Fatal("default value is nil: " + fmt.Sprintf("%T(%v)", i.Abs(), i.Abs()))
	}
}

func TestEmptyInterface(t *testing.T) {
	var i interface{}

	i = 42

	if fmt.Sprintf("%T(%v)", i, i) != "int(42)" {
		t.Fatal("can be any: " + fmt.Sprintf("%T(%v)", fmt.Sprintf("%T(%v)", i, i), fmt.Sprintf("%T(%v)", i, i)))
	}

	i = "hello"

	if fmt.Sprintf("%T(%v)", i, i) != "string(hello)" {
		t.Fatal("can be any: " + fmt.Sprintf("%T(%v)", fmt.Sprintf("%T(%v)", i, i), fmt.Sprintf("%T(%v)", i, i)))
	}
}

func TestTypeAssertions(t *testing.T) {
	var i interface{} = "hello"

	if fmt.Sprintf("%T(%v)", i, i) != "string(hello)" {
		t.Fatal("This is string: " + fmt.Sprintf("%T(%v)", fmt.Sprintf("%T(%v)", i, i), fmt.Sprintf("%T(%v)", i, i)))
	}

	s := i.(string)

	if fmt.Sprintf("%T(%v)", s, s) != "string(hello)" {
		t.Fatal("This is string: " + fmt.Sprintf("%T(%v)", fmt.Sprintf("%T(%v)", s, s), fmt.Sprintf("%T(%v)", s, s)))
	}

	_, ok := i.(string)

	if ok != true {
		t.Fatal("it is string: " + fmt.Sprintf("%T(%v)", ok, ok))
	}

	f, ok := i.(float64)

	if f != 0 {
		t.Fatal("0 value: " + fmt.Sprintf("%T(%v)", f, f))
	}
	if ok != false {
		t.Fatal("it isn't string: " + fmt.Sprintf("%T(%v)", ok, ok))
	}
}

func TestTypeSwitches(t *testing.T) {
	var i interface{} = 24

	switch v := i.(type) {
	case int:
	default:
		t.Fatal("It should be int: " + fmt.Sprintf("%T(%v)", v, v))
	}
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func TestStringerInterface(t *testing.T) {
	a := Person{"Arthur Dent", 42}

	if fmt.Sprintf("%v", a) != "Arthur Dent (42 years)" {
		t.Fatal("Person.String() is called: " + fmt.Sprintf("%T(%v)", fmt.Sprintf("%v", a), fmt.Sprintf("%v", a)))
	}
}
