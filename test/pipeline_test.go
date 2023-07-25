package test

import (
	"context"
	"sync"
	"testing"
	"time"
)

func TestPipeline(t *testing.T) {
	start := func(ctx context.Context) int {
		totalMessage := 0
		wg := &sync.WaitGroup{}

		wg.Add(1)
		messages := make(chan int, 5)
		go func() {
			defer wg.Done()
			defer close(messages)

			dequeue := time.Tick(37 * time.Millisecond)
			for {
				select {
				case <-ctx.Done():
					return
				case <-dequeue:
					totalMessage++
					t.Log(totalMessage)
					messages <- totalMessage
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
						t.Log(cached)
					}
					cached = make([]int, 0, 10)
				case result, ok := <-results:
					if !ok {
						if len(cached) > 0 {
							t.Log(cached)
						}
						return
					}
					cached = append(cached, result)
					if len(cached) == 10 {
						t.Log(cached)
						cached = make([]int, 0, 10)
					}
				}
			}
		}()

		wg.Wait()
		return totalMessage
	}

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	t.Log(start(ctx))
}
