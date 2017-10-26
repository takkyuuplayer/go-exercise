package test

import (
	"math"
	"testing"
)

func TestNaN(t *testing.T) {
	nan := math.NaN()

	if nan == nan {
		t.Errorf(`nan = %#v, do NOT want %#v`, nan, nan)
	}

	if nan >= nan {
		t.Errorf(`nan >= %#v, want %#v`, nan, nan)
	}

	if math.IsNaN(nan) == false {
		t.Errorf(`math.IsNaN(nan) = %#v, do NOT want %#v`, math.IsNaN(nan), false)
	}
}
