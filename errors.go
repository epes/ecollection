package ecollection

import "errors"

var (
	ErrEmpty = errors.New("collection is empty")
	ErrFull  = errors.New("collection is full")
)
