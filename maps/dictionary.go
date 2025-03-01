package maps

type Dictionary map[string]string
type DictionaryError string

const (
	ErrorNotFound   = DictionaryError("could not find the word")
	ErrorWordExists = DictionaryError("word exists, cannot add word")
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

func (e DictionaryError) Error() string {
	return string(e)
}
