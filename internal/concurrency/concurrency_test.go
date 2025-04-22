package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "waat://furhurterwe.geds"
}

func TestWebsiteCheckers(t *testing.T) {
	// Arrange
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	expect := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	// Act
	result := WebsiteCheckers(mockWebsiteChecker, websites)

	// Assert
	if !reflect.DeepEqual(expect, result) {
		t.Fatalf("expect %v, result %v", expect, result)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)

	return true
}

func BenchmarkWebsiteCheckers(b *testing.B) {
	urls := make([]string, 100)

	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		WebsiteCheckers(slowStubWebsiteChecker, urls)
	}
}
