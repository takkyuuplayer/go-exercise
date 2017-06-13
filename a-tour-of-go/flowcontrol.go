package main

import (
	"fmt"
	"math"
	"runtime"
)

func main() {
	forloop()
	forloop2()

	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(pow(3, 2, 10))
	fmt.Println(pow(3, 3, 20))

	newtownSqrt(2)

	switchStatement()

	deferStatement()

	deferStack()
}

func forloop() {
	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)
}

func forloop2() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func newtownSqrt(x float64) float64 {
	z := 1.0

	for i := 0; i < 10; i++ {
		z = z - (z*z-2)/(2*z)
		fmt.Printf("%g times ==>  %g\n", i, z)
	}

	return z
}

func switchStatement() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Printf("%s. ", os)
	}
}

func deferStatement() {
	defer fmt.Println("World")

	fmt.Println("Hello")
}

func deferStack() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

}
