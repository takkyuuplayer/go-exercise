package test_test

import (
	"fmt"
	"testing"
)

func TestFoo(t *testing.T) {
	t.Parallel()
	data := make([]int, 25)
	size := len(data)
	for i := 0; i < size; i++ {
		data[i] = i
	}

	perUnit := 20
	for start := 0; start < size; start += perUnit {
		end := start + perUnit
		if end > size {
			end = size
		}
		var units []int
		units = append(units, data[start:end]...)
		fmt.Println(units)
	}
}
