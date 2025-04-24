package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("increase counter 3 times returns 3", func(t *testing.T) {
		// Arrange
		counter := NewCounter()

		// Act
		counter.Increase()
		counter.Increase()
		counter.Increase()

		// Assert
		verifyCounter(t, counter, 3)
	})

	t.Run("runs successfully and safely", func(t *testing.T) {
		// Arrange
		expectedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(expectedCount)

		// Act
		for range expectedCount {
			go func(w *sync.WaitGroup) {
				counter.Increase()
				w.Done()
			}(&wg)
		}

		wg.Wait()

		// Assert
		verifyCounter(t, counter, expectedCount)
	})
}

func verifyCounter(t *testing.T, result *Counter, expect int) {
	t.Helper()

	if result.Value() != expect {
		t.Errorf("result %d expect %d", result.Value(), expect)
	}
}
