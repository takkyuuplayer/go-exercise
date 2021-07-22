package test

import (
	"fmt"
)

func ExamplePrintf_v() {
	fmt.Printf("%v", 1)
	// Output: 1
}

func ExamplePrintf_t() {
	fmt.Printf("%T", 1)
	// Output: int
}

func ExamplePrintf_sharpv() {
	fmt.Printf("%#v", "1")
	// Output: "1"
}

type Vertex struct {
	X, Y int
}

func ExamplePrintf_struct() {
	v := Vertex{1, 2}

	fmt.Printf("%#v", v)
	// Output: test.Vertex{X:1, Y:2}
}
