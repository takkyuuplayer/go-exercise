package test

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPipeline(t *testing.T) {
	start := func(ctx context.Context) {
		wg := &sync.WaitGroup{}

		wg.Add(1)
		messages := make(chan int, 5)
		go func() {
			defer wg.Done()
			defer close(messages)

			counter := 0
			dequeue := time.Tick(5 * time.Millisecond)
			for {
				select {
				case <-ctx.Done():
					return
				case <-dequeue:
					counter++
					messages <- counter
				}
			}
		}()

		wg.Add(1)
		results := make(chan int)
		go func() {
			defer wg.Done()
			defer close(results)
			for i := range messages {
				results <- i * 2
			}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			interval := time.Tick(100 * time.Millisecond)
			cached := make([]int, 0, 10)
			for {
				select {
				case <-interval:
					if len(cached) > 0 {
						assert.Greater(t, len(cached), 0)
					}
					cached = make([]int, 0, 10)
				case result, ok := <-results:
					if !ok {
						if len(cached) > 0 {
							assert.Greater(t, len(cached), 0)
						}
						return
					}
					cached = append(cached, result)
					if len(results) == 10 {
						assert.Len(t, cached, 10)
						cached = make([]int, 0, 10)
					}
				}
			}
		}()

		wg.Wait()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	start(ctx)
}
