package test

import (
	"fmt"
)

func ExamplePrintfV() {
	fmt.Printf("%v", 1)
	// Output: 1
}

func ExamplePrintfT() {
	fmt.Printf("%T", 1)
	// Output: int
}

func ExamplePrintfSharpV() {
	fmt.Printf("%#v", "1")
	// Output: "1"
}

type Vertex struct {
	X, Y int
}

func ExamplePrintStruct() {
	v := Vertex{1, 2}

	fmt.Printf("%#v", v)
	// Output: test.Vertex{X:1, Y:2}
}
