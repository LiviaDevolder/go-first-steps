package main

import "testing"

func TestHello(t *testing.T) {

	verifyCorrectMessage := func(t *testing.T, result, expected string) {
		t.Helper()
		if result != expected {
			t.Errorf("Result '%s', expected '%s'", result, expected)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		// Arrange
		expected := "Hello, Lívia!"

		// Act
		result := Hello("Lívia", "")

		// Assert
		verifyCorrectMessage(t, result, expected)
	})

	t.Run("say hello to the world when there is no name", func(t *testing.T) {
		// Arrange
		expected := "Hello, world!"

		// Act
		result := Hello("", "")

		// Assert
		verifyCorrectMessage(t, result, expected)
	})

	t.Run("in spanish", func(t *testing.T) {
		// Arrange
		expected := "Hola, Antonio!"

		// Act
		result := Hello("Antonio", "spanish")

		// Assert
		verifyCorrectMessage(t, result, expected)
	})

	t.Run("in russian", func(t *testing.T) {
		// Arrange
		expected := "Привет, Ариаднэ!"

		// Act
		result := Hello("Ариаднэ", "russian")

		// Assert
		verifyCorrectMessage(t, result, expected)
	})

	t.Run("in portuguese", func(t *testing.T) {
		// Arrange
		expected := "Olá, Fátima!"

		// Act
		result := Hello("Fátima", "portuguese")

		// Assert
		verifyCorrectMessage(t, result, expected)
	})
}
