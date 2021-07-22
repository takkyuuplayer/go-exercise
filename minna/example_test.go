package minna

import "fmt"

func ExamplePrintln_hello() {
	fmt.Println("Hello")
	// Output: Hello
}

func ExamplePrintln_unordered() {
	for _, v := range []int{1, 2, 3} {
		fmt.Println(v)
	}
	// Unordered output:
	// 2
	// 3
	// 1
}
