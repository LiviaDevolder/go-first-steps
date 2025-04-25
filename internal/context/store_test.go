package context

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	t        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <- ctx.Done():
				s.t.Log("spy store cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}

		data <- result
	} ()

	select {
	case <- ctx.Done():
		return "", ctx.Err()
	case res := <- data:
		return res, nil
	}
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true

	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func TestHandler(t *testing.T) {
	data := "Hello, world!"

	t.Run("return store data", func(t *testing.T) {
		// Arrange
		store := SpyStore{response: data, t: t}
		svr := Server(&store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		// Act
		svr.ServeHTTP(response, request)

		// Assert
		if response.Body.String() != data {
			t.Errorf("result %s expect %s", response.Body.String(), data)
		}
	})

	t.Run("cancel the job when cancel the request", func(t *testing.T) {
		// Arrange
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := &SpyResponseWriter{}

		// Act
		svr.ServeHTTP(response, request)

		// Assert
		if response.written {
			t.Errorf("a answer shouldnt be written")
		}
	})
}
