package domain

import "errors"

var ErrNotFound = errors.New("resource was not found")

type ErrConflict string

func (e ErrConflict) Error() string {
	return "resource with conflicting property '" + string(e) + "' already exists"
}

type ErrValidation string

func (e ErrValidation) Error() string {
	return string(e)
}
