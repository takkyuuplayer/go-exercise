package minna

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	if sum(1, 2) != 3 {
		t.Fatal(fmt.Sprintf("sum(1,2) should be %v(%T) but got %v(%T).", sum(1, 2), sum(1, 2), 3, 3))
	}
}
