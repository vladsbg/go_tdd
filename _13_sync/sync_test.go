package _13_sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementar o contador 3 vezes, resulta no valor 3", func(t *testing.T) {
		want := 3
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, want)
	})

	t.Run("roda concorrentemente em segurança", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()
		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func(w *sync.WaitGroup) {
				counter.Inc()
				w.Done()
			}(&wg)
		}
		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(t *testing.T, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Fatalf("Esperava %v, obteve %v", want, got)
	}
}
