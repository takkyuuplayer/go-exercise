package test

import (
	"math"
	"testing"
)

func TestNaN(t *testing.T) {
	t.Parallel()
	nan := math.NaN()

	if math.IsNaN(nan) == false {
		t.Errorf(`math.IsNaN(nan) = %#v, do NOT want %#v`, math.IsNaN(nan), false)
	}
}
