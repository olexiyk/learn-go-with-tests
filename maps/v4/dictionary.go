package main

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("not found")
var ErrWordExists = errors.New("word exists")

func (d Dictionary) Search(s string) (string, error) {
	v, ok := d[s]
	if !ok {
		return "", ErrNotFound
	}
	return v, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	if err == ErrNotFound {
		d[word] = definition
		return nil
	}
	return ErrWordExists
}
