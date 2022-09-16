package mydict

import "errors"

type Dictionary map[string]string

var errNotFound = errors.New("존재하지 않습니다")
var errWordExist = errors.New("존재하는 단어입니다")

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExist
	}
	return nil
}
