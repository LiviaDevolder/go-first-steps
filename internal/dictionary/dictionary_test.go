package dictionary

import (
	"testing"
)

func TestSearch(t *testing.T) {
	// Arrange
	dictionary := Dictionary{"test": "just a test"}

	t.Run("known word", func(t *testing.T) {
		expect := "just a test"

		// Act
		result, _ := dictionary.Search("test")

		verifyStrings(t, result, expect)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, result := dictionary.Search("unkwnown")

		verifyError(t, result, ErrorNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		// Arrange
		dictionary := Dictionary{}
		definition := "just a test"
		word := "test"

		// Act & Assert
		err := dictionary.Add(word, definition)

		verifyError(t, err, nil)
		verifyDefinition(t, dictionary, word, definition)
	})

	t.Run("word already exists", func(t *testing.T) {
		// Arrange
		word := "test"
		definition := "just a test"
		dictionary := Dictionary{word: definition}

		// Act & Assert
		err := dictionary.Add(word, definition)

		verifyError(t, err, ErrorAlreadyExists)
		verifyDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("word exists", func(t *testing.T) {
		// Arrange
		word := "test"
		definition := "just a test"
		newDefinition := "new definition"

		dictionary := Dictionary{word: definition}

		// Act
		err := dictionary.Update(word, newDefinition)

		// Assert
		verifyError(t, err, nil)
		verifyDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		// Arrange
		word := "test"
		definition := "just a test"

		dictionary := Dictionary{}

		// Act
		err := dictionary.Update(word, definition)

		// Assert
		verifyError(t, err, ErrorDoesntExists)
	})
}

func TestDelete(t *testing.T) {
	// Arrange
	word := "test"
	definition := "just a test"

	dictionary := Dictionary{word: definition}

	// Act
	dictionary.Delete(word)

	// Assert
	_, err := dictionary.Search(word)
	if err != ErrorNotFound {
		t.Errorf("%s should be deleted", word)
	}
}

func verifyStrings(t *testing.T, result, expect string) {
	t.Helper()

	if result != expect {
		t.Errorf("result %s, expect %s, searched %s", result, expect, "test")
	}
}

func verifyError(t *testing.T, result, expect error) {
	t.Helper()

	if result != expect {
		t.Errorf("result %s, expect %s", result, expect)
	}
}

func verifyDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	// Act
	result, err := dictionary.Search(word)

	// Assert
	if err != nil {
		t.Fatal("failed to add the word to the dictionary", err)
	}

	if result != definition {
		t.Errorf("result %s, expect %s", result, definition)
	}
}
