package test

import (
	"math/cmplx"
	"testing"
)

func TestComplex(t *testing.T) {
	x := complex(1, 2)
	y := complex(3, 4)

	if x*y != complex(-5, 10) {
		t.Errorf(`x * y = %#v, want %#v`, x*y, complex(-5, 10))
	}

	if real(x) != 1 {
		t.Errorf(`real(x) = %#v, want %#v`, real(x), 1)
	}

	if imag(y) != 4 {
		t.Errorf(`imag(y) = %#v, want %#v`, imag(y), 4)
	}
}

func TestSqrt(t *testing.T) {
	if cmplx.Sqrt(-1) != complex(0, 1) {
		t.Errorf(`cmplx.Sqrt(-1) = %#v, want %#v`, cmplx.Sqrt(-1), complex(0, 1))
	}
}
