package minna

import (
	"testing"
)

func TestSum(t *testing.T) {
	if sum(1, 2) != 3 {
		t.Fatalf("sum(1,2) should be %v(%T) but got %v(%T).", sum(1, 2), sum(1, 2), 3, 3)
	}
}
