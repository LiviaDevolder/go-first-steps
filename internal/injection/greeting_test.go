package main

import (
	"bytes"
	"testing"
)

func TestGreeting(t *testing.T) {
	// Arrange
	buffer := bytes.Buffer{}
	Greeting(&buffer, "Ariadne")
	expect := "Hello, Ariadne"

	// Act
	result := buffer.String()

	// Assert
	if result != expect {
		t.Errorf("result %s, expect %s", result, expect)
	}
}
