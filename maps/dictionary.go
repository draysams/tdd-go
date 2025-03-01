package maps

type Dictionary map[string]string
type DictionaryError string

const (
	ErrorNotFound         = DictionaryError("could not find the word")
	ErrorWordExists       = DictionaryError("word exists, cannot add word")
	ErrorWordDoesNotExist = DictionaryError("word does not exist, cannot update")
)

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
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
		return ErrorWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		d[word] = definition
	case ErrorNotFound:
		return ErrorWordDoesNotExist
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case ErrorNotFound:
		return ErrorWordDoesNotExist
	case nil:
		delete(d, word)
	default:
		return err
	}
	return nil
}

func (e DictionaryError) Error() string {
	return string(e)
}
