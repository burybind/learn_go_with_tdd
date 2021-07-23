package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, &counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := Counter{}

		var wg sync.WaitGroup // a WaitGroup waits for a collection of goroutines to finish
		wg.Add(wantedCount)   // defining how many goroutines to wait for

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done() // signaling to the WaitGroup that a goroutine has finished
			}()
		}
		wg.Wait() // used to block until the specified number of goroutines have finished

		assertCounter(t, &counter, wantedCount)
	})
}

func assertCounter(t testing.TB, counter *Counter, wantedCount int) {
	if counter.Value() != wantedCount {
		t.Errorf("got %d, want %d", counter.Value(), wantedCount)
	}
}
