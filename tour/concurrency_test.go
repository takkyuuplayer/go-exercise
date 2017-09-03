package tour

import (
	"fmt"
	"testing"
)

func TestChannels(t *testing.T) {
	sum := func(s []int, c chan int) {
		ret := 0
		for _, v := range s {
			ret += v
		}
		c <- ret
	}

	s := []int{7, 2, 8. - 9, 4, 0}

	c := make(chan int, 2)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c

	if x != 3 {
		t.Fatal("1st channel: " + fmt.Sprintf("%T(%v)", x, x))
	}
	if y != 9 {
		t.Fatal("2nd channel: " + fmt.Sprintf("%T(%v)", y, y))
	}
}

func TestRangeAndClose(t *testing.T) {
	fib := func(n int, c chan int) {
		current, next := 0, 1

		for i := 0; i < n; i++ {
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
		for i := 0; i < 10; i++ {
			sum += <-c
		}
		quit <- 0
	}()

	fib(c, quit)

	if sum != 88 {
		t.Fatal("sum_1^10 fib(i): " + fmt.Sprintf("%T(%v)", sum, sum))
	}
}
