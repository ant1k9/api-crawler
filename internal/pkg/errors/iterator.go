package errors

import "errors"

var (
	ErrIteratorDoesNotExist = errors.New("iterator does not exist")
	ErrEmptyCollection      = errors.New("empty collection")
)
