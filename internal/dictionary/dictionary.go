package dictionary

type Dictionary map[string]string

type DictionaryError string

func (e DictionaryError) Error() string {
	return string(e)
}

const (
	ErrorNotFound      = DictionaryError("word not found")
	ErrorAlreadyExists = DictionaryError("word already exists")
	ErrorDoesntExists  = DictionaryError("word doesnt exists")
)

func (d Dictionary) Search(word string) (string, error) {
	definition, exists := d[word]

	if !exists {
		return "", ErrorNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrorNotFound:
		d[word] = definition
	case nil:
		return ErrorAlreadyExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrorNotFound:
		return ErrorDoesntExists
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
