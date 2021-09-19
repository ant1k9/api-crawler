package errors

import "errors"

var (
	ErrPaginationEnd         = errors.New("end of pagination")
	ErrPaginatorDoesNotExist = errors.New("paginator does not exist")
)
