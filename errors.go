package passgen

import "errors"

var (
	ErrEmptyTemplate = errors.New("empty template")
	ErrZeroLength    = errors.New("zero password length")
)
