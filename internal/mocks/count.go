package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	lastWord  = "Go!"
	initCount = 3
	write     = "write"
	pause     = "pause"
)

type Sleeper interface {
	Pause()
}

func Count(output io.Writer, sleeper Sleeper) {
	for i := initCount; i > 0; i-- {
		sleeper.Pause()
		fmt.Fprintln(output, i)
	}

	sleeper.Pause()

	fmt.Fprint(output, lastWord)
}

type CustomSleeper struct {
	duration time.Duration
	pause    func(time.Duration)
}

func (s *CustomSleeper) Pause() {
	s.pause(s.duration)
}

func main() {
	sleeper := &CustomSleeper{1 * time.Second, time.Sleep}
	Count(os.Stdout, sleeper)
}
