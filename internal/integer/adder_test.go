package integer

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	// Arrange
	expected := 4

	// Act
	result := Adder(2, 2)

	// Assert
	if result != expected {
		t.Errorf("expected '%d', result '%d'", expected, result)
	}
}

func ExampleAdder() {
	sum := Adder(1, 5)

	fmt.Println(sum)
	// Output: 6
}
