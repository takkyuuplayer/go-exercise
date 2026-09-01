package tour

import (
	"fmt"
	"testing"
)

func TestRangeAndClose(t *testing.T) {
	t.Parallel()
	fib := func(n int, c chan int) {
		current, next := 0, 1

		for range n {
			c <- current
			current, next = next, current+next
		}

		close(c)
	}

	c := make(chan int, 10)

	go fib(cap(c), c)

	sum := 0
	for i := range c {
		sum += i
	}

	if sum != 88 {
		t.Fatal("sum_1^10 fib(i): " + fmt.Sprintf("%T(%v)", sum, sum))
	}
}

func TestSelect(t *testing.T) {
	t.Parallel()
	fib := func(c, quit chan int) {
		current, next := 0, 1

		for {
			select {
			case c <- current:
				current, next = next, current+next
			case <-quit:
				return
			}
		}
	}

	c := make(chan int)
	quit := make(chan int)

	sum := 0

	go func() {
		for range 10 {
			sum += <-c
		}
		quit <- 0
	}()

	fib(c, quit)

	if sum != 88 {
		t.Fatal("sum_1^10 fib(i): " + fmt.Sprintf("%T(%v)", sum, sum))
	}
}
