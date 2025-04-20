package iterator

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	// Arrange
	expected := "aaaaa"

	// Act
	result := Repeat("a", 5)

	// Assert
	if result != expected {
		t.Errorf("expected '%s' but got '%s'", expected, result)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("D", 3)

	fmt.Println(repeated)
	// Output: DDD
}
