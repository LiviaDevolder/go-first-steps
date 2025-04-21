package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCount(t *testing.T) {
	t.Run("print 3 to Go!", func(t *testing.T) {
		// Arrange
		buffer := &bytes.Buffer{}
		expect := `3
2
1
Go!`

		// Act
		Count(buffer, &SpyCountOperations{})

		result := buffer.String()

		// Assert
		if result != expect {
			t.Errorf("result %s, expect %s", result, expect)
		}
	})

	t.Run("pause before each printing", func(t *testing.T) {
		spyPrintSleep := &SpyCountOperations{}
		Count(spyPrintSleep, spyPrintSleep)

		expect := []string{
			pause,
			write,
			pause,
			write,
			pause,
			write,
			pause,
			write,
		}

		if !reflect.DeepEqual(expect, spyPrintSleep.Calls) {
			t.Errorf("expect %v calls, but got %v", expect, spyPrintSleep.Calls)
		}
	})
}

func TestCustomSleep(t *testing.T) {
	// Arrange
	pauseTime := 5 * time.Second

	spyTime := &SpyTime{}

	sleeper := CustomSleeper{pauseTime, spyTime.Pause}

	// Act
	sleeper.Pause()

	// Assert
	if spyTime.durationPause != pauseTime {
		t.Errorf("should be paused for %v, but was paused for %v", pauseTime, spyTime.durationPause)
	}
}

type SpyCountOperations struct {
	Calls []string
}

func (s *SpyCountOperations) Pause() {
	s.Calls = append(s.Calls, pause)
}

func (s *SpyCountOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type SpyTime struct {
	durationPause time.Duration
}

func (t *SpyTime) Pause(duration time.Duration) {
	t.durationPause = duration
}
