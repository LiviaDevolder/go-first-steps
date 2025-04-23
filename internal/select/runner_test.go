package runner

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRunner(t *testing.T) {
	t.Run("get faster URL", func(t *testing.T) {
		// Arrange
		slowServer := createServerWithDelay(20 * time.Millisecond)
		fastServer := createServerWithDelay(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		expect := fastURL

		// Act
		result, err := Runner(slowURL, fastURL)

		// Assert
		if err != nil {
			t.Fatalf("didnt expected an error, but got %v", err)
		}
		
		if result != expect {
			t.Errorf("result %s, expect %s", result, expect)
		}
	})

	t.Run("throw an error if the server didnt return in 10s", func(t *testing.T) {
		server := createServerWithDelay(25 * time.Millisecond)

		defer server.Close()

		_, err := Custom(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("didnt got an error")
		}
	})
}

func createServerWithDelay(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
