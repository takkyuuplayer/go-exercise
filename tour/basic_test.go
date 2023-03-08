package tour

import (
	"fmt"
	"math"
	"math/cmplx"
	"testing"
)

func TestExportedValue(t *testing.T) {
	if fmt.Sprintf("%v", math.Pi) != "3.141592653589793" {
		t.Fatal("To access exported value, use Capital character")
	}
}

func add(x, y int) int {
	return x + y
}

func TestFunction(t *testing.T) {
	if add(42, 13) != 55 {
		t.Fatal("function can be defined")
	}
}

func swap(x, y string) (string, string) {
	return y, x
}

func TestMultipleResults(t *testing.T) {
	a, b := swap("hello", "world")
	if a != "world" && b != "hello" {
		t.Fatal("function can return multiple values")
	}
}

func split(sum int) (ret1, ret2 int) {
	ret1 = sum * 4 / 9
	ret2 = sum - ret1
	return
}

func TestNamedReturn(t *testing.T) {
	a, b := split(17)
	if a != 7 && b != 10 {
		t.Fatal("function can return named values")
	}
}

var c bool

func TestVariable(t *testing.T) {
	var i int

	if i != 0 || c != false {
		t.Fatal("var defines variable with initial value")
	}
}

func TestVariableDeclaration(t *testing.T) {
	i := 3
	j, k := 4, 5

	if i != 3 || j != 4 || k != 5 {
		t.Fatal(":= defines variable")
	}
}

func TestBasicTypes(t *testing.T) {
	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)
	const f = "%T(%v)"

	if fmt.Sprintf(f, ToBe, ToBe) != "bool(false)" {
		t.Fatalf("%%T show type")
		t.Fatalf("%%V show value")
	}

	if fmt.Sprintf(f, MaxInt, MaxInt) != "uint64(18446744073709551615)" {
		t.Fatalf("%%T show type")
		t.Fatalf("%%V show value")
	}

	if fmt.Sprintf(f, z, z) != "complex128((2+3i))" {
		t.Fatalf("%%T show type")
		t.Fatalf("%%V show value")
	}
}

func TestZeroValues(t *testing.T) {
	var i int
	var f float64
	var b bool
	var s string

	if fmt.Sprintf("%v", i) != "0" {
		t.Fatal("int zero value is 0")
	}

	if fmt.Sprintf("%v", f) != "0" {
		t.Fatal("float zero value is 0.0")
	}

	if fmt.Sprintf("%v", b) != "false" {
		t.Fatal("bool zero value is false")
	}

	if fmt.Sprintf("%v", s) != "" {
		t.Fatal("string zero value is empty")
	}
}

func TestCasting(t *testing.T) {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)

	if fmt.Sprintf("%v", f) != "5" {
		t.Fatal("cast to fload to calculate Sqrt: " + fmt.Sprintf("%v", f))
	}

	if fmt.Sprintf("%v", z) != "5" {
		t.Fatal("cast to fload to calculate Sqrt: " + fmt.Sprintf("%v", z))
	}
}

func TestTypeInference(t *testing.T) {
	v := 42

	if fmt.Sprintf("%T", v) != "int" {
		t.Fatal("type should be infered from RHS: " + fmt.Sprintf("%T", v))
	}
}

func TestConstant(t *testing.T) {
	const World = "世界"

	if fmt.Sprintf("%v", World) != "世界" {
		t.Fatal("const can be defined: " + fmt.Sprintf("%v", World))
	}
}
