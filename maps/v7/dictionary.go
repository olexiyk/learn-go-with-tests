package main

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("not found")
var ErrWordExists = errors.New("word exists")
var ErrWordDoesNotExist = errors.New("word does not exists")

func (d Dictionary) Search(word string) (string, error) {
	v, ok := d[word]
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

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	if err == ErrNotFound {
		return ErrWordExists
	}
	d[word] = definition
	return nil
}

func (d Dictionary) Delete(word string) {
	_, ok := d[word]
	if !ok {
		delete(d, word)
	}
}
