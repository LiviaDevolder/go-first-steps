package runner

import (
	"fmt"
	"net/http"
	"time"
)

const tenSecondsLimit = 10 * time.Second

func Runner(a, b string) (winner string, erro error) {
	return Custom(a, b, tenSecondsLimit)
}

func Custom(a, b string, timeLimit time.Duration) (winner string, erro error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeLimit):
		return "", fmt.Errorf("time limit exceeded to %s and %s", a, b)
	}
}

func ping(URL string) chan bool {
	ch := make(chan bool)

	go func() {
		http.Get(URL)
		ch <- true
	}()

	return ch
}
