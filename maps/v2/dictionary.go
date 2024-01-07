package main

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("not found")

func (d Dictionary) Search(s string) (string, error) {
	v, ok := d[s]
	if !ok {
		return "", ErrNotFound
	}
	return v, nil
}
