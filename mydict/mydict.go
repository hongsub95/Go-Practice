package mydict

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errors.New("존재하지 않습니다")
}
